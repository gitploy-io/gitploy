package slack

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/slack-go/slack"
	"go.uber.org/zap"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/chatcallback"
	"github.com/hanjunlee/gitploy/ent/deployment"
	"github.com/hanjunlee/gitploy/vo"
)

type (
	ErrorsPayload struct {
		Errors []ErrorPayload `json:"errors"`
	}

	ErrorPayload struct {
		Name  string `json:"name"`
		Error string `json:"error"`
	}
)

// handleDeployCmd handles deploy command: "/gitploy deploy OWNER/REPO".
// It opens a dialog with fields to submit the payload for deployment.
func (s *Slack) handleDeployCmd(c *gin.Context) {
	ctx := c.Request.Context()

	// SlashCommandParse hvae to be success because
	// it was called in the Cmd method.
	cmd, _ := slack.SlashCommandParse(c.Request)

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

	// The length of args is always equal to two.
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
		responseMessage(cmd.ChannelID, cmd.ResponseURL, "The config file is not found.")
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

	state := randstr()

	// Create a new callback to interact with submissions.
	cb, err := s.i.CreateDeployChatCallback(ctx, cu, r, &ent.ChatCallback{
		Type:  chatcallback.TypeDeploy,
		State: state,
	})
	if err != nil {
		s.log.Error("It has failed to create a new callback.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	err = slack.New(cu.BotToken).
		OpenDialogContext(ctx, cmd.TriggerID, slack.Dialog{
			CallbackID:     itoa(cb.ID),
			State:          state,
			Title:          "Deploy",
			SubmitLabel:    "Submit",
			NotifyOnCancel: true,
			Elements:       createDeployDialogElement(config),
		})
	if err != nil {
		s.log.Error("It has failed to create a new callback.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

// interactDeploy deploy with the submitted payload.
func (s *Slack) interactDeploy(c *gin.Context) {
	ctx := c.Request.Context()

	// InteractionCallbackParse hvae to be success because
	// it was called in the Interact method.
	itr, _ := s.InteractionCallbackParse(c.Request)
	cb, _ := s.i.FindChatCallbackByState(ctx, itr.State)

	var (
		typ = itr.Submission["type"]
		ref = itr.Submission["ref"]
		env = itr.Submission["env"]
	)

	// Get the chat user with edges.
	cu, _ := s.i.FindChatUserByID(ctx, cb.Edges.ChatUser.ID)

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

	if !cf.HasEnv(env) {
		responseMessage(itr.Channel.ID, itr.ResponseURL, fmt.Sprintf("The `%s` environment is not found.", env))
		c.Status(http.StatusOK)
		return
	}

	// Prepare to deploy:
	// 1) commit SHA.
	// 2) next deployment number.
	var (
		sha    string
		number int
	)
	sha, err = s.getCommitSha(ctx, cu.Edges.User, cb.Edges.Repo, typ, ref)
	if vo.IsRefNotFoundError(err) {
		c.JSON(http.StatusOK, ErrorsPayload{
			Errors: []ErrorPayload{
				{
					Name:  "ref",
					Error: "The reference is not found.",
				},
			},
		})
		return
	} else if err != nil {
		s.log.Error("It has failed to get the SHA of commit.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	if number, err = s.i.GetNextDeploymentNumberOfRepo(ctx, cb.Edges.Repo); err != nil {
		s.log.Error("It has failed to get the next deployment number.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	d, err := s.i.Deploy(ctx, cu.Edges.User, cb.Edges.Repo, &ent.Deployment{
		Number: number,
		Type:   deployment.Type(typ),
		Ref:    ref,
		Sha:    sha,
		Env:    env,
	}, cf.GetEnv(env))
	if err != nil {
		s.log.Error("It has failed to deploy.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	if err = s.i.PublishDeployment(ctx, cb.Edges.Repo, d); err != nil {
		s.log.Warn("It has failed to publish the deployment.", zap.Error(err))
	}

	c.Status(http.StatusOK)
}

func parseFullName(fullname string) (string, string, error) {
	namespaceName := strings.Split(fullname, "/")
	if len(namespaceName) != 2 {
		return "", "", fmt.Errorf("It is a invalid formatted command.")
	}

	return namespaceName[0], namespaceName[1], nil
}

func createDeployDialogElement(c *vo.Config) []slack.DialogElement {
	options := []slack.DialogSelectOption{}
	for _, env := range c.Envs {
		options = append(options, slack.DialogSelectOption{
			Label: strings.Title(env.Name),
			Value: env.Name,
		})
	}

	return []slack.DialogElement{
		slack.DialogInputSelect{
			DialogInput: slack.DialogInput{
				Type:  slack.InputTypeSelect,
				Label: "Environment",
				Name:  "env",
			},
			Options: options,
		},
		slack.DialogInputSelect{
			DialogInput: slack.DialogInput{
				Type:  slack.InputTypeSelect,
				Label: "Type",
				Name:  "type",
			},
			Options: []slack.DialogSelectOption{
				{
					Label: "Commit",
					Value: "commit",
				},
				{
					Label: "Branch",
					Value: "branch",
				},
				{
					Label: "Tag",
					Value: "tag",
				},
			},
		},
		slack.DialogInput{
			Type:  "text",
			Label: "Ref",
			Name:  "ref",
			Hint:  "E.g. Commit - 25a667d6, Branch - main, Tag - v0.1.2",
		},
	}
}

func (s *Slack) getCommitSha(ctx context.Context, u *ent.User, re *ent.Repo, typ, ref string) (string, error) {
	switch typ {
	case "commit":
		c, err := s.i.GetCommit(ctx, u, re, ref)
		if err != nil {
			return "", err
		}

		return c.Sha, nil
	case "branch":
		b, err := s.i.GetBranch(ctx, u, re, ref)
		if err != nil {
			return "", err
		}

		return b.CommitSha, nil
	case "tag":
		t, err := s.i.GetTag(ctx, u, re, ref)
		if err != nil {
			return "", err
		}

		return t.CommitSha, nil
	default:
		return "", fmt.Errorf("Type must be one of commit, branch, tag.")
	}
}
