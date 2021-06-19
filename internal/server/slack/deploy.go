package slack

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/slack-go/slack"
	"go.uber.org/zap"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/chatcallback"
	"github.com/hanjunlee/gitploy/ent/deployment"
	errs "github.com/hanjunlee/gitploy/internal/errors"
	"github.com/hanjunlee/gitploy/vo"
)

const (
	token = ""
)

func init() {
	// Seed for randstr
	rand.Seed(time.Now().UnixNano())
}

func (s *Slack) handleDeployCmd(c *gin.Context, cmd slack.SlashCommand) {
	ctx := c.Request.Context()

	u, err := s.i.FindUserWithChatUserByChatUserID(ctx, cmd.UserID)
	if u.Edges.ChatUser == nil {
		responseMessage(cmd, fmt.Sprint("Slack is not connected with Gitploy."))
		c.Status(http.StatusOK)
		return
	} else if err != nil {
		s.log.Error("failed to find the user.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	cu := u.Edges.ChatUser
	client := slack.New(cu.BotToken)

	fullname := trimDeployCommandPrefix(cmd.Text)
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
	if errs.IsConfigNotFoundError(err) {
		responseMessage(cmd, "The config file is not found")
		c.Status(http.StatusOK)
		return
	} else if errs.IsConfigParseError(err) {
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

	err = client.OpenDialogContext(ctx, cmd.TriggerID, slack.Dialog{
		CallbackID:     cb,
		State:          state,
		Title:          "Deploy",
		SubmitLabel:    "Submit",
		NotifyOnCancel: true,
		Elements:       createDialogElement(config),
	})
	if err != nil {
		s.log.Error("failed to open the dialog.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	_, err = s.i.CreateDeployChatCallback(ctx, cu, r, &ent.ChatCallback{
		ID:    cb,
		Type:  chatcallback.TypeDeploy,
		State: state,
	})
	if err != nil {
		s.log.Error("failed to create a new callback.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (s *Slack) interactDeploy(c *gin.Context, scb slack.InteractionCallback) {
	ctx := c.Request.Context()

	cb, _ := s.i.FindChatCallbackWithEdgesByID(ctx, scb.CallbackID)

	cu := cb.Edges.ChatUser
	u, _ := s.i.FindUserWithChatUserByChatUserID(ctx, cu.ID)

	d, err := s.i.Deploy(ctx, u, cb.Edges.Repo, &ent.Deployment{
		Type: deployment.Type(scb.Submission["type"]),
		Ref:  scb.Submission["ref"],
		Env:  scb.Submission["env"],
	})
	if err != nil {
		s.log.Error("failed to deploy.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	if _, err = s.i.Notify(ctx, d); err != nil {
		s.log.Warn("failed to notify the deployment.", zap.Error(err))
	}

	s.log.Debug("deployed successfully.")
	c.Status(http.StatusOK)
}

func sendResponse(c *slack.Client, cmd slack.SlashCommand, message string) error {
	_, _, _, err := c.SendMessage(cmd.ChannelID, slack.MsgOptionText(message, true), slack.MsgOptionResponseURL(cmd.ResponseURL, "ephemeral"))
	return err
}

func trimDeployCommandPrefix(txt string) string {
	return strings.TrimPrefix(txt, "deploy ")
}

func parseFullName(n string) (string, string, error) {
	namespaceName := strings.Split(n, "/")
	if len(namespaceName) != 2 {
		return "", "", fmt.Errorf("invalid format")
	}

	return namespaceName[0], namespaceName[1], nil
}

func createDialogElement(c *vo.Config) []slack.DialogElement {
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
				Type:  "select",
				Label: "Environment",
				Name:  "env",
			},
			Options: options,
		},
		slack.DialogInputSelect{
			DialogInput: slack.DialogInput{
				Type:  "select",
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

func randstr() string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, 16)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
