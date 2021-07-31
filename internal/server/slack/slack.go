package slack

import (
	"context"
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

	s.i.Subscribe(func(u *ent.User, n *ent.Notification) {
		if cu := u.Edges.ChatUser; cu != nil {
			ctx := context.Background()
			s.Notify(ctx, cu, n)
		}
	})

	return s
}

// Cmd handles Slash command of Slack.
// https://api.slack.com/interactivity/slash-commands
func (s *Slack) Cmd(c *gin.Context) {
	ctx := c.Request.Context()

	cmd, err := slack.SlashCommandParse(c.Request)
	if err != nil {
		s.log.Error("failed to parse the command.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	t := strings.TrimSpace(cmd.Text)

	var handleErr error
	if strings.HasPrefix(t, "deploy ") {
		handleErr = s.handleDeployCmd(ctx, cmd)
	} else if strings.HasPrefix(t, "rollback ") {
		handleErr = s.handleRollbackCmd(ctx, cmd)
	} else if strings.HasPrefix(t, "help") {
		s.handleHelpCmd(ctx, cmd)
	} else {
		s.handleHelpCmd(ctx, cmd)
	}

	if handleErr != nil {
		s.log.Error("It has failed to handle the command: %s", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	s.log.Debug("It has handled the command successfully.", zap.String("command", cmd.Command))
	c.Status(http.StatusOK)
}

func (s *Slack) handleHelpCmd(ctx context.Context, cmd slack.SlashCommand) {
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
	ctx := c.Request.Context()

	c.Request.ParseForm()
	payload := c.Request.PostForm.Get("payload")

	scb := slack.InteractionCallback{}
	err := scb.UnmarshalJSON([]byte(payload))
	if err != nil {
		s.log.Error("failed to unmarshal.", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	if scb.Type == slack.InteractionTypeDialogCancellation {
		c.Status(http.StatusOK)
		return
	}

	// Trim backticked double quote for string type.
	// https://github.com/slack-go/slack/issues/816
	state := strings.Trim(scb.State, "\"")

	cb, err := s.i.FindChatCallbackByState(ctx, state)
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
	if cb.Type == chatcallback.TypeDeploy {
		interactErr = s.interactDeploy(c, scb, cb)
	} else if cb.Type == chatcallback.TypeRollback {
		interactErr = s.interactRollback(c, scb, cb)
	}

	if interactErr != nil {
		s.log.Error("It has failed to interact the component: %s", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	s.log.Debug("interact the component successfully.", zap.String("type", string(cb.Type)))
	c.Status(http.StatusOK)
}
