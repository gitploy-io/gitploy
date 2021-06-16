package slack

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/slack-go/slack"
	"go.uber.org/zap"

	"github.com/hanjunlee/gitploy/ent"
	errs "github.com/hanjunlee/gitploy/internal/errors"
	"github.com/hanjunlee/gitploy/vo"
)

const (
	token = ""
)

func (s *Slack) Deploy(c *gin.Context) {
	cmd, err := slack.SlashCommandParse(c.Request)
	if err != nil {
		s.log.Error("failed to parse the command.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	// TODO: Authenticate a user
	sc := slack.New(token)
	u, _ := s.i.FindUser()

	ns, n, err := parseFullName(cmd.Text)
	if err != nil {
		sendResponse(sc, cmd, fmt.Sprintf("`%s` is invalid format.", cmd.Text))
		c.Status(http.StatusOK)
		return
	}

	ctx := c.Request.Context()
	r, err := s.i.FindRepoByNamespaceName(ctx, u, ns, n)
	if ent.IsNotFound(err) {
		sendResponse(sc, cmd, fmt.Sprintf("The `%s` is not found", cmd.Text))
		c.Status(http.StatusOK)
		return
	} else if err != nil {
		s.log.Error("failed to find the repo.", zap.String("namespace", ns), zap.String("name", n), zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	config, err := s.i.GetConfig(ctx, u, r)
	if errs.IsConfigNotFoundError(err) {
		sendResponse(sc, cmd, "The config file is not found")
		c.Status(http.StatusOK)
		return
	} else if errs.IsConfigParseError(err) {
		sendResponse(sc, cmd, "The config file is invliad format.")
		c.Status(http.StatusOK)
		return
	} else if err != nil {
		s.log.Error("failed to get the config file.", zap.String("namespace", ns), zap.String("name", n), zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	err = sc.OpenDialogContext(ctx, cmd.TriggerID, slack.Dialog{
		CallbackID:     "deploy.1",
		Title:          "Deploy",
		State:          "state",
		SubmitLabel:    "Submit",
		NotifyOnCancel: true,
		Elements:       createDialogElement(config),
	})
	if err != nil {
		s.log.Error("failed to open the dialog.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func sendResponse(c *slack.Client, cmd slack.SlashCommand, message string) error {
	_, _, _, err := c.SendMessage(cmd.ChannelID, slack.MsgOptionText(message, true), slack.MsgOptionResponseURL(cmd.ResponseURL, "ephemeral"))
	return err
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
				Name:  "environment",
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
