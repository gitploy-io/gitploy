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

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/event"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
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

// GetEvents streams events of deployment, or approval.
func (s *Stream) GetEvents(c *gin.Context) {
	ctx := c.Request.Context()

	v, _ := c.Get(gb.KeyUser)
	u, _ := v.(*ent.User)

	debugID := randstr()

	events := make(chan *ent.Event, 10)

	// Subscribe events
	// it'll unsubscribe after the connection is closed.
	sub := func(e *ent.Event) {

		// Deleted type is always propagated to all.
		if e.Type == event.TypeDeleted {
			events <- e
			return
		}

		if ok, err := s.hasPermForEvent(ctx, u, e); !ok {
			s.log.Debug("Skip the event. The user has not the perm.")
			return
		} else if err != nil {
			s.log.Error("It has failed to check the perm.", zap.Error(err))
			return
		}

		events <- e
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
			break L
		case <-time.After(time.Hour):
			break L
		case <-time.After(time.Second * 30):
			c.Render(-1, sse.Event{
				Event: "ping",
				Data:  "ping",
			})
			w.Flush()
		case e := <-events:
			c.Render(-1, sse.Event{
				Event: "event",
				Data:  e,
			})
			w.Flush()
			s.log.Debug("server sent event.", zap.Int("event_id", e.ID), zap.String("debug_id", debugID))
		}
	}
}

// hasPermForEvent checks the user has permission for the event.
func (s *Stream) hasPermForEvent(ctx context.Context, u *ent.User, e *ent.Event) (bool, error) {
	if e.Kind == event.KindDeployment {
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

	if e.Kind == event.KindApproval {
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

	b := make([]rune, 4)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
