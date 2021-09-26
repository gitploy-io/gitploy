package slack

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gitploy-io/gitploy/ent"
	"github.com/slack-go/slack"
	"go.uber.org/zap"
)

const (
	KeyCmd      = "gitploy.slack.command"
	KeyIntr     = "gitploy.slack.interaction"
	KeyChatUser = "gitploy.slack.user"
)

type (
	SlackMiddleware struct {
		i      Interactor
		secret string
		log    *zap.Logger
	}

	SlackMiddlewareConfig struct {
		Interactor Interactor
		Secret     string
	}
)

func NewSlackMiddleware(c *SlackMiddlewareConfig) *SlackMiddleware {
	return &SlackMiddleware{
		i:      c.Interactor,
		secret: c.Secret,
		log:    zap.L().Named("slack-middleware"),
	}
}

func (m *SlackMiddleware) Verify() gin.HandlerFunc {
	return func(c *gin.Context) {
		v, err := slack.NewSecretsVerifier(c.Request.Header, m.secret)
		if err != nil {
			m.log.Error("failed to generate the verifier.")
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		d := copyRawData(c.Request)
		v.Write(d)

		if err := v.Ensure(); err != nil {
			m.log.Error("invalid request.", zap.Error(err))
			c.AbortWithStatus(http.StatusBadRequest)
		}
	}
}

func (m *SlackMiddleware) ParseCmd() gin.HandlerFunc {
	return func(c *gin.Context) {
		cmd, err := slack.SlashCommandParse(c.Request)
		if err != nil {
			m.log.Error("It has failed to parse the command.", zap.Error(err))
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.Set(KeyCmd, cmd)
	}
}

func (m *SlackMiddleware) ParseIntr() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.ParseForm()
		payload := c.Request.PostForm.Get("payload")

		intr := slack.InteractionCallback{}
		if err := intr.UnmarshalJSON([]byte(payload)); err != nil {
			m.log.Error("It has failed to parse the interaction callback.", zap.Error(err))
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.Set(KeyIntr, intr)
	}
}

func (m *SlackMiddleware) SetChatUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		if v, ok := c.Get(KeyCmd); ok {
			cmd := v.(slack.SlashCommand)

			cu, err := m.i.FindChatUserByID(ctx, cmd.UserID)
			if ent.IsNotFound(err) {
				postResponseMessage(cmd.ChannelID, cmd.ResponseURL, "Slack is not connected with Gitploy.")
				c.Status(http.StatusOK)
				return
			} else if err != nil {
				m.log.Error("It has failed to get chat-user.", zap.Error(err))
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}

			c.Set(KeyChatUser, cu)
			return
		}

		if v, ok := c.Get(KeyIntr); ok {
			intr := v.(slack.InteractionCallback)

			cu, err := m.i.FindChatUserByID(ctx, intr.User.ID)
			// InteractionCallback doesn't have the response URL.
			if err != nil {
				m.log.Error("It has failed to get chat-user.", zap.Error(err))
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}

			c.Set(KeyChatUser, cu)
			return
		}
	}
}

func copyRawData(req *http.Request) []byte {
	bodyBytes, _ := ioutil.ReadAll(req.Body)
	req.Body.Close()

	req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	return bodyBytes
}
