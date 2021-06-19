package slack

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
	"github.com/slack-go/slack"
	"go.uber.org/zap"
)

type (
	Slack struct {
		i   Interactor
		log *zap.Logger
	}
)

func NewSlack(i Interactor) *Slack {
	return &Slack{
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

	if strings.HasPrefix(t, "deploy") {

	} else if strings.HasPrefix(t, "help") {
		s.handleHelpCmd(c, cmd)
	} else {
		s.handleHelpCmd(c, cmd)
	}
}

func (s *Slack) handleHelpCmd(c *gin.Context, cmd slack.SlashCommand) {
	msg := strings.Join([]string{
		"Below are the commands you can use: \n",
		"*Deploy* - Create a new deployment.",
		"`/gitploy deploy OWNER/REPO` - Create a new deployment to OWNER/REPO",
	}, "\n")
	responseMessage(cmd, msg)
}

func responseMessage(cmd slack.SlashCommand, message string) {
	client := slack.New("")
	client.SendMessage(cmd.ChannelID, slack.MsgOptionText(message, true), slack.MsgOptionResponseURL(cmd.ResponseURL, "ephemeral"))
}

func (s *Slack) Interact(c *gin.Context) {
	c.Request.ParseForm()
	payload := c.Request.PostForm.Get("payload")

	callback := &slack.InteractionCallback{}
	err := callback.UnmarshalJSON([]byte(payload))
	if err != nil {
		s.log.Error("failed to unmarshal.", zap.Error(err))
	}

	log.Info(callback.Submission)
	c.Status(http.StatusOK)
}
