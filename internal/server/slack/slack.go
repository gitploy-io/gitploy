package slack

import (
	"net/http"

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
