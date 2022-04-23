// Copyright 2021 Gitploy.IO Inc. All rights reserved.
// Use of this source code is governed by the Gitploy Non-Commercial License
// that can be found in the LICENSE file.

//go:build !oss

package stream

import (
	"time"

	"github.com/gin-contrib/sse"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/ent/event"
)

// GetEvents streams events of deployment, or review.
func (s *Stream) GetEvents(c *gin.Context) {
	ctx := c.Request.Context()

	v, _ := c.Get(gb.KeyUser)
	u, _ := v.(*ent.User)

	events := make(chan *sse.Event, 10)

	// Subscribe events
	// it'll unsubscribe after the connection is closed.
	sub := func(e *ent.Event) {
		switch e.Kind {
		case event.KindDeployment:
			d, err := s.i.FindDeploymentByID(ctx, e.DeploymentID)
			if err != nil {
				s.log.Error("Failed to find the deployment.", zap.Error(err))
				return
			}

			if _, err := s.i.FindPermOfRepo(ctx, d.Edges.Repo, u); err != nil {
				s.log.Debug("Skip the event. The permission is denied.", zap.Error(err))
				return
			}

			s.log.Debug("Dispatch a deployment event.", zap.Int("id", d.ID))
			events <- &sse.Event{
				Event: "deployment",
				Data:  d,
			}
		case event.KindDeploymentStatus:
			ds, err := s.i.FindDeploymentStatusByID(ctx, e.DeploymentStatusID)
			if err != nil {
				s.log.Error("Failed to find the deployment status.", zap.Error(err))
				return
			}

			// Ensure that a user has access to the repository of the deployment.
			d, err := s.i.FindDeploymentByID(ctx, ds.DeploymentID)
			if err != nil {
				s.log.Error("Failed to find the deployment.", zap.Error(err))
				return
			}

			if _, err := s.i.FindPermOfRepo(ctx, d.Edges.Repo, u); err != nil {
				s.log.Debug("Skip the event. The permission is denied.", zap.Error(err))
				return
			}

			s.log.Debug("Dispatch a deployment_status event.", zap.Int("id", d.ID))
			events <- &sse.Event{
				Event: "deployment_status",
				Data:  ds,
			}

		case event.KindReview:
			r, err := s.i.FindReviewByID(ctx, e.ReviewID)
			if err != nil {
				s.log.Error("Failed to find the review.", zap.Error(err))
				return
			}

			d, err := s.i.FindDeploymentByID(ctx, r.DeploymentID)
			if err != nil {
				s.log.Error("Failed to find the deployment.", zap.Error(err))
				return
			}

			if _, err := s.i.FindPermOfRepo(ctx, d.Edges.Repo, u); err != nil {
				s.log.Debug("Skip the event. The permission is denied.")
				return
			}

			s.log.Debug("Dispatch a review event.", zap.Int("id", r.ID))
			events <- &sse.Event{
				Event: "review",
				Data:  r,
			}
		}
	}

	if err := s.i.SubscribeEvent(sub); err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to subscribe notification events").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
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
			c.Render(-1, e)
			w.Flush()
		}
	}
}
