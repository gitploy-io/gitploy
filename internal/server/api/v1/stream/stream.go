package stream

import (
	"io"
	"net/http"

	"github.com/gin-contrib/sse"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/hanjunlee/gitploy/ent"
	gb "github.com/hanjunlee/gitploy/internal/server/global"
)

type (
	Stream struct {
		i   Interactor
		log *zap.Logger
	}
)

func NewStream(i Interactor) *Stream {
	return &Stream{
		i:   i,
		log: zap.L().Named("stream"),
	}
}

func (s *Stream) GetNotification(c *gin.Context) {
	v, _ := c.Get(gb.KeyUser)
	me, _ := v.(*ent.User)

	notifications := make(chan *ent.Notification, 100)

	// subscribe notification events
	// it'll unsubscribe after the connection is closed.
	sub := func(u *ent.User, n *ent.Notification) {
		if me.ID == u.ID {
			notifications <- n
			s.log.Debug("receive a new notification event.", zap.Int("id", n.ID))
		}
	}
	if err := s.i.Subscribe(sub); err != nil {
		s.log.Error("failed to subscribe notification events", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to connect.")
		return
	}

	defer func() {
		if err := s.i.Unsubscribe(sub); err != nil {
			s.log.Error("failed to unsubscribe notification events.")
		}

		close(notifications)
		s.log.Debug("connect is closed.")
	}()

	c.Stream(func(w io.Writer) bool {
		n, ok := <-notifications
		if !ok {
			return false
		}

		sse.Encode(w, sse.Event{
			Event: "notification",
			Data:  n,
		})
		return true
	})
}
