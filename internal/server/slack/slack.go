package slack

import (
	"net/http"
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

func (s *Slack) Cmd(c *gin.Context) {
	cmd, err := slack.SlashCommandParse(c.Request)
	if err != nil {
		s.log.Error("failed to parse the command.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	t := strings.TrimSpace(cmd.Text)
	s.log.Debug("Parse Slack command.", zap.String("text", t))

	if strings.HasPrefix(t, "deploy ") {
		s.handleDeployCmd(c, cmd)
	} else if strings.HasPrefix(t, "help") {
		s.handleHelpCmd(c, cmd)
	} else {
		s.handleHelpCmd(c, cmd)
	}
}

func (s *Slack) handleHelpCmd(c *gin.Context, cmd slack.SlashCommand) {
	msg := strings.Join([]string{
		"Below are the commands you can use: \n",
		"*Deploy*",
		"`/gitploy deploy OWNER/REPO` - Create a new deployment to OWNER/REPO",
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

	if state := strings.Trim(scb.State, "\""); state != cb.State {
		responseMessage(scb, "The state is invalid. You can interact with Slack by only `/gitploy`.")
		c.Status(http.StatusOK)
		return
	}

	defer s.i.CloseChatCallback(ctx, cb)

	switch cb.Type {
	case chatcallback.TypeDeploy:
		s.log.Debug("interact with the deploy command.")
		s.interactDeploy(c, scb)
	}

	c.Status(http.StatusOK)
}
