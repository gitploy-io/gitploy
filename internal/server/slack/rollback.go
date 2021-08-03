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
	"github.com/hanjunlee/gitploy/ent/notification"
	"github.com/hanjunlee/gitploy/vo"
)

const (
	blockDeployment  = "block_deployment"
	actionDeployment = "aciton_deployment"
)

type (
	rollbackViewSubmission struct {
		DeploymentID int
	}

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
		responseMessage(cmd.ChannelID, cmd.ResponseURL, fmt.Sprintf("`%s` is invalid repository format.", args[1]))
		c.Status(http.StatusOK)
		return
	}

	r, err := s.i.FindRepoOfUserByNamespaceName(ctx, cu.Edges.User, ns, n)
	if ent.IsNotFound(err) {
		responseMessage(cmd.ChannelID, cmd.ResponseURL, fmt.Sprintf("The `%s` repository is not found.", args[1]))
		c.Status(http.StatusOK)
		return
	} else if err != nil {
		s.log.Error("It has failed to get the repo.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	config, err := s.i.GetConfig(ctx, cu.Edges.User, r)
	if vo.IsConfigNotFoundError(err) {
		responseMessage(cmd.ChannelID, cmd.ResponseURL, "The config file is not found")
		c.Status(http.StatusOK)
		return
	} else if vo.IsConfigParseError(err) {
		responseMessage(cmd.ChannelID, cmd.ResponseURL, "The config file is invliad format.")
		c.Status(http.StatusOK)
		return
	} else if err != nil {
		s.log.Error("It has failed to get the config.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	cb, err := s.i.CreateChatCallback(ctx, cu, r, &ent.ChatCallback{
		Type: chatcallback.TypeRollback,
	})
	if err != nil {
		s.log.Error("It has failed to create a new callback.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	as := s.getSucceedDeploymentAggregation(ctx, r, config)

	_, err = slack.New(cu.BotToken).
		OpenViewContext(ctx, cmd.TriggerID, buildRollbackView(cb.Hash, as))
	if err != nil {
		s.log.Error("It has failed to open a dialog.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func buildRollbackView(callbackID string, as []*deploymentAggregation) slack.ModalViewRequest {
	groups := []*slack.OptionGroupBlockObject{}

	for _, a := range as {
		options := []*slack.OptionBlockObject{}

		for _, d := range a.deployments {
			created, _ := goment.New(d.CreatedAt)

			options = append(options, &slack.OptionBlockObject{
				Text: &slack.TextBlockObject{
					Type: slack.PlainTextType,
					Text: fmt.Sprintf("#%d - %s deployed at %s", d.ID, d.GetShortRef(), created.FromNow()),
				},
				Value: strconv.Itoa(d.ID),
			})
		}

		groups = append(groups, &slack.OptionGroupBlockObject{
			Label: &slack.TextBlockObject{
				Type: slack.PlainTextType,
				Text: string(a.envName),
			},
			Options: options,
		})
	}

	return slack.ModalViewRequest{
		Type:       slack.VTModal,
		CallbackID: callbackID,
		Title: &slack.TextBlockObject{
			Type: slack.PlainTextType,
			Text: "Rollback",
		},
		Submit: &slack.TextBlockObject{
			Type: slack.PlainTextType,
			Text: "Submit",
		},
		Close: &slack.TextBlockObject{
			Type: slack.PlainTextType,
			Text: "Close",
		},
		Blocks: slack.Blocks{
			BlockSet: []slack.Block{
				slack.InputBlock{
					Type:    slack.MBTInput,
					BlockID: blockDeployment,
					Label: &slack.TextBlockObject{
						Type: slack.PlainTextType,
						Text: "Deployments",
					},
					Element: slack.SelectBlockElement{
						Type:     slack.OptTypeStatic,
						ActionID: actionDeployment,
						Placeholder: &slack.TextBlockObject{
							Type: slack.PlainTextType,
							Text: "Select target deployment",
						},
						OptionGroups: groups,
					},
				},
			},
		},
	}
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

	// InteractionCallbackParse always to be parsed successfully because
	// it was called in the Interact method.
	itr, _ := s.InteractionCallbackParse(c.Request)
	cb, _ := s.i.FindChatCallbackByHash(ctx, itr.View.CallbackID)

	cu, _ := s.i.FindChatUserByID(ctx, cb.Edges.ChatUser.ID)

	sm := parseRollbackSubmissions(itr)

	d, err := s.i.FindDeploymentByID(ctx, sm.DeploymentID)
	if err != nil {
		s.log.Error("It has failed to find the deployment.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	cf, err := s.i.GetConfig(ctx, cu.Edges.User, cb.Edges.Repo)
	if err != nil {
		s.log.Error("It has failed to get the config file.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	if !cf.HasEnv(d.Env) {
		responseMessage(itr.Channel.ID, itr.ResponseURL, fmt.Sprintf("The `%s` environment is not found.", d.Env))
		c.Status(http.StatusOK)
		return
	}

	next, err := s.i.GetNextDeploymentNumberOfRepo(ctx, cb.Edges.Repo)
	if err != nil {
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

	if err = s.i.Publish(ctx, notification.TypeDeploymentCreated, cb.Edges.Repo, d, nil); err != nil {
		s.log.Warn("failed to notify the deployment.", zap.Error(err))
	}

	c.Status(http.StatusOK)
}

func parseRollbackSubmissions(itr slack.InteractionCallback) *rollbackViewSubmission {
	sm := &rollbackViewSubmission{}

	values := itr.View.State.Values
	if v, ok := values[blockDeployment][actionDeployment]; ok {
		sm.DeploymentID = atoi(v.SelectedOption.Value)
	}

	return sm
}
