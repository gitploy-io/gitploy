package slack

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nleeper/goment"
	"github.com/slack-go/slack"
	"go.uber.org/zap"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/chatcallback"
	"github.com/hanjunlee/gitploy/ent/deployment"
	"github.com/hanjunlee/gitploy/vo"
)

type (
	deploymentAggregation struct {
		envName     string
		deployments []*ent.Deployment
	}
)

// handleRollbackCmd handles rollback command: "/gitploy rollback OWNER/REPO".
func (s *Slack) handleRollbackCmd(c *gin.Context) {
	ctx := c.Request.Context()

	cmd, _ := slack.SlashCommandParse(c.Request)

	// user have to be exist if chat user is found.
	cu, err := s.i.FindChatUserByID(ctx, cmd.UserID)
	if ent.IsNotFound(err) {
		responseMessage(cmd.ChannelID, cmd.ResponseURL, "Slack is not connected with Gitploy.")
		c.Status(http.StatusOK)
		return
	} else if err != nil {
		s.log.Error("It has failed to get chat user.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	args := strings.Split(cmd.Text, " ")

	ns, n, err := parseFullName(args[1])
	if err != nil {
		responseMessage(cmd.ChannelID, cmd.ResponseURL, fmt.Sprintf("`%s` is invalid format.", args[1]))
		c.Status(http.StatusOK)
		return
	}

	r, err := s.i.FindRepoOfUserByNamespaceName(ctx, cu.Edges.User, ns, n)
	if ent.IsNotFound(err) {
		responseMessage(cmd.ChannelID, cmd.ResponseURL, fmt.Sprintf("The `%s` is not found.", args[1]))
		c.Status(http.StatusOK)
		return
	} else if err != nil {
		s.log.Error("It has failed to get the repo.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	config, err := s.i.GetConfig(ctx, cu.Edges.User, r)
	if vo.IsConfigNotFoundError(err) {
		responseMessage(cmd.ChannelID, cmd.ResponseURL, "The configuration file is not found")
		c.Status(http.StatusOK)
		return
	} else if vo.IsConfigParseError(err) {
		responseMessage(cmd.ChannelID, cmd.ResponseURL, "The configuration file is invliad format.")
		c.Status(http.StatusOK)
		return
	} else if err != nil {
		s.log.Error("It has failed to get the config.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	state := randstr()

	cb, err := s.i.CreateDeployChatCallback(ctx, cu, r, &ent.ChatCallback{
		Type:  chatcallback.TypeRollback,
		State: state,
	})
	if err != nil {
		s.log.Error("It has failed to create a new callback.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	as := s.getSucceedDeploymentAggregation(ctx, r, config)
	err = slack.New(cu.BotToken).
		OpenDialogContext(ctx, cmd.TriggerID, slack.Dialog{
			CallbackID:     itoa(cb.ID),
			State:          state,
			Title:          "Rollback",
			SubmitLabel:    "Submit",
			NotifyOnCancel: true,
			Elements:       createRollbackDialogElement(as),
		})
	if err != nil {
		s.log.Error("It has failed to open a dialog.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (s *Slack) getSucceedDeploymentAggregation(ctx context.Context, r *ent.Repo, cf *vo.Config) []*deploymentAggregation {
	a := []*deploymentAggregation{}

	for _, env := range cf.Envs {
		ds, _ := s.i.ListDeployments(ctx, r, env.Name, string(deployment.StatusSuccess), 1, 5)
		if len(ds) == 0 {
			continue
		}

		a = append(a, &deploymentAggregation{
			envName:     env.Name,
			deployments: ds,
		})
	}

	return a
}

func (s *Slack) interactRollback(c *gin.Context) {
	ctx := c.Request.Context()

	// InteractionCallbackParse hvae to be success because
	// it was called in the Interact method.
	itr, _ := s.InteractionCallbackParse(c.Request)
	cb, _ := s.i.FindChatCallbackByState(ctx, itr.State)

	var (
		id = itr.Submission["deployment_id"]
	)

	cu, _ := s.i.FindChatUserByID(ctx, cb.Edges.ChatUser.ID)

	d, err := s.i.FindDeploymentByID(ctx, atoi(id))
	if err != nil {
		s.log.Error("It has failed to find the deployment.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	cf, err := s.i.GetConfig(ctx, cu.Edges.User, cb.Edges.Repo)
	if vo.IsConfigNotFoundError(err) {
		responseMessage(itr.Channel.ID, itr.ResponseURL, "The config file is not found.")
		c.Status(http.StatusOK)
		return
	} else if vo.IsConfigParseError(err) {
		responseMessage(itr.Channel.ID, itr.ResponseURL, "The config file is invalid format.")
		c.Status(http.StatusOK)
		return
	} else if err != nil {
		s.log.Error("It has failed to get the config file.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	if !cf.HasEnv(d.Env) {
		responseMessage(itr.Channel.ID, itr.ResponseURL, fmt.Sprintf("The `%s` environment is not found.", d.Env))
		c.Status(http.StatusOK)
		return
	}

	// Prepare to rollback:
	// 1) next deployment number.
	var (
		next int
	)

	if next, err = s.i.GetNextDeploymentNumberOfRepo(ctx, cb.Edges.Repo); err != nil {
		s.log.Error("It has failed to get the next deployment number.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	d, err = s.i.Rollback(ctx, cu.Edges.User, cb.Edges.Repo, &ent.Deployment{
		Number: next,
		Type:   deployment.Type(d.Type),
		Ref:    d.Ref,
		Sha:    d.Sha,
		Env:    d.Env,
	}, cf.GetEnv(d.Env))
	if err != nil {
		s.log.Error("It has failed to deploy.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	if err = s.i.PublishDeployment(ctx, cb.Edges.Repo, d); err != nil {
		s.log.Warn("failed to notify the deployment.", zap.Error(err))
	}

	c.Status(http.StatusOK)
}

func createRollbackDialogElement(as []*deploymentAggregation) []slack.DialogElement {
	groups := []slack.DialogOptionGroup{}
	for _, a := range as {
		options := []slack.DialogSelectOption{}

		for _, d := range a.deployments {
			created, _ := goment.New(d.CreatedAt)
			options = append(options, slack.DialogSelectOption{
				Label: fmt.Sprintf("#%d - %s deployed at %s", d.ID, strRef(d), created.FromNow()),
				Value: strconv.Itoa(d.ID),
			})
		}

		groups = append(groups, slack.DialogOptionGroup{
			Label:   strings.Title(a.envName),
			Options: options,
		})
	}

	return []slack.DialogElement{
		slack.DialogInputSelect{
			DialogInput: slack.DialogInput{
				Type:  slack.InputTypeSelect,
				Label: "Deployment",
				Name:  "deployment_id",
			},
			OptionGroups: groups,
		},
	}
}

func strRef(d *ent.Deployment) string {
	ref := d.Ref

	if d.Type == deployment.TypeCommit {
		ref = d.Ref[:7]
	}

	return ref
}
