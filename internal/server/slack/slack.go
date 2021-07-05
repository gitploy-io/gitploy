package slack

import (
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/chatcallback"
	"github.com/slack-go/slack"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

type (
	Slack struct {
		i   Interactor
		c   *oauth2.Config
		log *zap.Logger
	}
)

func NewSlack(c *oauth2.Config, i Interactor) *Slack {
	return &Slack{
		c:   c,
		i:   i,
		log: zap.L().Named("slack"),
	}
}

// Cmd handles Slash command of Slack.
// https://api.slack.com/interactivity/slash-commands
func (s *Slack) Cmd(c *gin.Context) {
	cmd, err := slack.SlashCommandParse(c.Request)
	if err != nil {
		s.log.Error("failed to parse the command.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	t := strings.TrimSpace(cmd.Text)

	var handleErr error
	if strings.HasPrefix(t, "deploy ") {
		handleErr = s.handleDeployCmd(c, cmd)
	} else if strings.HasPrefix(t, "rollback ") {
		handleErr = s.handleRollbackCmd(c, cmd)
	} else if strings.HasPrefix(t, "help") {
		s.handleHelpCmd(c, cmd)
	} else {
		s.handleHelpCmd(c, cmd)
	}

	if handleErr != nil {
		s.log.Error("failed to handle the command: %s", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	s.log.Debug("handling the command successfully.", zap.String("command", cmd.Command))
	c.Status(http.StatusOK)
}

func (s *Slack) handleHelpCmd(c *gin.Context, cmd slack.SlashCommand) {
	msg := strings.Join([]string{
		"Below are the commands you can use:",
		"",
		"*Deploy*",
		"`/gitploy deploy OWNER/REPO` - Create a new deployment for OWNER/REPO.",
		"",
		"*Rollback*",
		"`/gitploy rollback OWNER/REPO` - Rollback by the deployment for OWNER/REPO.",
	}, "\n")
	responseMessage(cmd, msg)
}

func responseMessage(obj interface{}, message string) {
	var (
		channelID   string
		responseURL string
	)
	switch i := obj.(type) {
	case slack.SlashCommand:
		channelID = i.ChannelID
		responseURL = i.ResponseURL
	case slack.InteractionCallback:
		channelID = i.Channel.ID
		responseURL = i.ResponseURL
	}

	client := slack.New("")
	client.SendMessage(channelID, slack.MsgOptionText(message, true), slack.MsgOptionResponseURL(responseURL, "ephemeral"))
}

// Interact interacts interactive components (dialog, button).
func (s *Slack) Interact(c *gin.Context) {
	c.Request.ParseForm()
	payload := c.Request.PostForm.Get("payload")

	scb := slack.InteractionCallback{}
	err := scb.UnmarshalJSON([]byte(payload))
	if err != nil {
		s.log.Error("failed to unmarshal.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	ctx := c.Request.Context()

	cb, err := s.i.FindChatCallbackByID(ctx, scb.CallbackID)
	if ent.IsNotFound(err) {
		responseMessage(scb, "The callback is not found. You can interact with Slack by only `/gitploy`.")
		c.Status(http.StatusOK)
		return
	} else if err != nil {
		s.log.Error("failed to find the callback.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	defer s.i.CloseChatCallback(ctx, cb)

	var interactErr error
	if scb.Type == slack.InteractionTypeDialogSubmission && cb.Type == chatcallback.TypeDeploy {
		interactErr = s.interactDeploy(c, scb)
	} else if scb.Type == slack.InteractionTypeDialogSubmission && cb.Type == chatcallback.TypeDeploy {
		interactErr = s.interactRollback(c, scb)
	}

	if interactErr != nil {
		s.log.Error("failed to interact the component: %s", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

// Check checks Slack is enabled.
func (s *Slack) Check(c *gin.Context) {
	c.Status(http.StatusOK)
}

func randstr() string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, 16)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}
