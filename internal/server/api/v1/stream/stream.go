package stream

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-contrib/sse"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/event"
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

func (s *Stream) GetEvents(c *gin.Context) {
	ctx := c.Request.Context()

	v, _ := c.Get(gb.KeyUser)
	u, _ := v.(*ent.User)

	debugID := randstr()
	s.log.Debug("create a stream.", zap.String("debug_id", debugID))

	events := make(chan *ent.Event, 10)

	// Subscribe events
	// it'll unsubscribe after the connection is closed.
	sub := func(event *ent.Event) {
		if ok, err := s.hasPermForEvent(ctx, u, event); !ok {
			return
		} else if err != nil {
			s.log.Error("It has failed to check the perm.", zap.Error(err))
			return
		}

		events <- event
	}
	if err := s.i.SubscribeEvent(sub); err != nil {
		s.log.Error("failed to subscribe notification events", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to connect.")
		return
	}

	defer func() {
		if err := s.i.UnsubscribeEvent(sub); err != nil {
			s.log.Error("failed to unsubscribe notification events.")
		}

		close(events)
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
		case e := <-events:
			var data interface{}
			if e.Type == event.TypeDeployment {
				data = e.Edges.Deployment
			} else if e.Type == event.TypeApproval {
				data = e.Edges.Approval
			}

			c.Render(-1, sse.Event{
				Event: e.Type.String(),
				Data:  data,
			})
			w.Flush()
			s.log.Debug("server sent event.", zap.Int("event_id", e.ID), zap.String("debug_id", debugID))
		}
	}
}

// hasPermForEvent checks the user has permission for the event.
func (s *Stream) hasPermForEvent(ctx context.Context, u *ent.User, e *ent.Event) (bool, error) {
	if e.Type == event.TypeDeployment {
		d, err := s.i.FindDeploymentByID(ctx, e.DeploymentID)
		if err != nil {
			return false, err
		}

		if _, err = s.i.FindPermOfRepo(ctx, d.Edges.Repo, u); ent.IsNotFound(err) {
			return false, nil
		} else if err != nil {
			return false, err
		}

		return true, nil
	}

	if e.Type == event.TypeApproval {
		a, err := s.i.FindApprovalByID(ctx, e.ApprovalID)
		if err != nil {
			return false, err
		}

		d, err := s.i.FindDeploymentByID(ctx, a.DeploymentID)
		if err != nil {
			return false, err
		}

		if _, err = s.i.FindPermOfRepo(ctx, d.Edges.Repo, u); ent.IsNotFound(err) {
			return false, nil
		} else if err != nil {
			return false, err
		}

		return true, nil
	}

	return false, fmt.Errorf("The type of event is not \"deployment\" or \"approval\".")
}

func randstr() string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, 8)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
