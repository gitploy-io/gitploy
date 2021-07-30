package stream

import (
	"math/rand"
	"net/http"
	"time"

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

	debugID := randstr()
	s.log.Debug("create a stream.", zap.String("debug_id", debugID))

	notifications := make(chan ent.Notification, 10)

	// Subscribe notification events
	// it'll unsubscribe after the connection is closed.
	sub := func(u *ent.User, n *ent.Notification) {
		var vn ent.Notification = *n

		// It is notified by Chat if user has connected with chat.
		if u.Edges.ChatUser != nil {
			vn.Notified = true
		}

		if me.ID == u.ID {
			notifications <- vn
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
	}()

	w := c.Writer

L:
	for {
		select {
		case <-w.CloseNotify():
			s.log.Debug("stream canceled.", zap.String("debug_id", debugID))
			break L
		case <-time.After(time.Hour):
			s.log.Debug("stream canceled.", zap.String("debug_id", debugID))
			break L
		case <-time.After(time.Second * 30):
			c.Render(-1, sse.Event{
				Event: "ping",
				Data:  "ping",
			})
			w.Flush()
		case n := <-notifications:
			c.Render(-1, sse.Event{
				Event: "notification",
				Data:  n,
			})
			w.Flush()
			s.log.Debug("server sent event.", zap.Int("notification_id", n.ID), zap.String("debug_id", debugID))
		}
	}
}

func randstr() string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, 8)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
