package slack

import (
	"context"
	"net/http"
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
	cmd, err := slack.SlashCommandParse(c.Request)
	if err != nil {
		s.log.Error("It has failed to parse the command.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	args := strings.Split(cmd.Text, " ")
	if args[0] == "deploy" && len(args) == 2 {
		s.handleDeployCmd(c)
	} else if args[0] == "rollback" && len(args) == 2 {
		s.handleRollbackCmd(c)
	} else if args[0] == "lock" && len(args) == 2 {
		s.handleLockCmd(c)
	} else if args[0] == "unlock" && len(args) == 2 {
		s.handleUnlockCmd(c)
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

	itr, err := s.InteractionCallbackParse(c.Request)
	if err != nil {
		s.log.Error("It has failed to parse the interaction callback.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

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
