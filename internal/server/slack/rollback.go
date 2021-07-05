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

func (s *Slack) handleRollbackCmd(c *gin.Context, cmd slack.SlashCommand) {
	ctx := c.Request.Context()

	// user have to be exist if chat user is found.
	cu, err := s.i.FindChatUserWithUserByID(ctx, cmd.UserID)
	if ent.IsNotFound(err) {
		responseMessage(cmd, fmt.Sprint("Slack is not connected with Gitploy."))
		c.Status(http.StatusOK)
		return
	} else if err != nil {
		s.log.Error("failed to find the user.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	u := cu.Edges.User
	client := slack.New(cu.BotToken)

	fullname := trimRollbackCommandPrefix(cmd.Text)
	ns, n, err := parseFullName(fullname)
	if err != nil {
		responseMessage(cmd, fmt.Sprintf("`%s` is invalid format.", fullname))
		c.Status(http.StatusOK)
		return
	}

	r, err := s.i.FindRepoByNamespaceName(ctx, u, ns, n)
	if ent.IsNotFound(err) {
		responseMessage(cmd, fmt.Sprintf("The `%s` is not found.", fullname))
		c.Status(http.StatusOK)
		return
	} else if err != nil {
		s.log.Error("failed to find the repo.", zap.String("repo", fullname), zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	config, err := s.i.GetConfig(ctx, u, r)
	if vo.IsConfigNotFoundError(err) {
		responseMessage(cmd, "The config file is not found")
		c.Status(http.StatusOK)
		return
	} else if vo.IsConfigParseError(err) {
		responseMessage(cmd, "The config file is invliad format.")
		c.Status(http.StatusOK)
		return
	} else if err != nil {
		s.log.Error("failed to get the config file.", zap.String("repo", fullname), zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	cb := randstr()
	state := randstr()
	as := s.getSucceedDeploymentAggregation(ctx, r, config)

	err = client.OpenDialogContext(ctx, cmd.TriggerID, slack.Dialog{
		CallbackID:     cb,
		State:          state,
		Title:          "Rollback",
		SubmitLabel:    "Submit",
		NotifyOnCancel: true,
		Elements:       createRollbackDialogElement(as),
	})
	if err != nil {
		s.log.Error("failed to open the dialog.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	_, err = s.i.CreateDeployChatCallback(ctx, cu, r, &ent.ChatCallback{
		ID:    cb,
		Type:  chatcallback.TypeRollback,
		State: state,
	})
	if err != nil {
		s.log.Error("failed to create a new rollback.", zap.Error(err))
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

func trimRollbackCommandPrefix(txt string) string {
	return strings.TrimPrefix(txt, "rollback ")
}
