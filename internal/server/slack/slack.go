package slack

import (
	"context"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/callback"
	"github.com/slack-go/slack"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

type (
	Slack struct {
		host  string
		proto string

		c *oauth2.Config
		i Interactor

		log *zap.Logger
	}

	SlackConfig struct {
		ServerHost  string
		ServerProto string
		*oauth2.Config
		Interactor
	}
)

func NewSlack(c *SlackConfig) *Slack {
	s := &Slack{
		host:  c.ServerHost,
		proto: c.ServerProto,
		c:     c.Config,
		i:     c.Interactor,
		log:   zap.L().Named("slack"),
	}

	s.i.SubscribeEvent(func(e *ent.Event) {
		s.Notify(context.Background(), e)
	})

	return s
}

// Cmd handles Slash command of Slack.
// https://api.slack.com/interactivity/slash-commands
func (s *Slack) Cmd(c *gin.Context) {
	av, _ := c.Get(KeyCmd)
	cmd := av.(slack.SlashCommand)

	if matched, _ := regexp.MatchString("^deploy[[:blank:]]+[0-9A-Za-z._-]*/[0-9A-Za-z._-]*$", cmd.Text); matched {
		s.handleDeployCmd(c)
	} else if matched, _ := regexp.MatchString("^rollback[[:blank:]]+[0-9A-Za-z._-]*/[0-9A-Za-z._-]*$", cmd.Text); matched {
		s.handleRollbackCmd(c)
	} else {
		s.handleHelpCmd(cmd.ChannelID, cmd.ResponseURL)
	}
}

func (s *Slack) handleHelpCmd(channelID, responseURL string) {
	msg := strings.Join([]string{
		"Below are the commands you can use:\n",
		"*Deploy*",
		"`/gitploy deploy OWNER/REPO` - Create a new deployment for OWNER/REPO.\n",
		"*Rollback*",
		"`/gitploy rollback OWNER/REPO` - Rollback by the deployment for OWNER/REPO.\n",
		"*Lock/Unlock*",
		"`/gitploy lock OWNER/REPO` - Lock the repository to disable deploying.",
		"`/gitploy unlock OWNER/REPO` - Unlock the repository to enable deploying.\n",
	}, "\n")

	postResponseMessage(channelID, responseURL, msg)
}

func postResponseMessage(channelID, responseURL, message string) error {
	_, _, _, err := slack.
		New("").
		SendMessage(
			channelID,
			slack.MsgOptionResponseURL(responseURL, "ephemeral"),
			slack.MsgOptionText(message, false),
		)
	return err
}

func postBotMessage(cu *ent.ChatUser, message string) error {
	_, _, _, err := slack.
		New(cu.BotToken).
		SendMessage(
			cu.ID,
			slack.MsgOptionText(message, false),
		)
	return err
}

// Interact interacts interactive components (dialog, button).
func (s *Slack) Interact(c *gin.Context) {
	ctx := c.Request.Context()

	v, _ := c.Get(KeyIntr)
	itr := v.(slack.InteractionCallback)

	cb, err := s.i.FindCallbackByHash(ctx, itr.View.CallbackID)
	if err != nil {
		s.log.Error("It has failed to find the callback.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	if cb.Type == callback.TypeDeploy {
		s.interactDeploy(c)
	} else if cb.Type == callback.TypeRollback {
		s.interactRollback(c)
	}
}

func (s *Slack) InteractionCallbackParse(r *http.Request) (slack.InteractionCallback, error) {
	r.ParseForm()
	payload := r.PostForm.Get("payload")

	scb := slack.InteractionCallback{}
	err := scb.UnmarshalJSON([]byte(payload))

	// Trim backticked double quote for string type.
	// https://github.com/slack-go/slack/issues/816
	state := strings.Trim(scb.State, "\"")
	scb.State = state

	return scb, err
}
