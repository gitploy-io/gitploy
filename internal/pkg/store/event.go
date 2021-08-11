package store

import (
	"context"
	"time"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/event"
)

func (s *Store) ListEventsGreaterThanTime(ctx context.Context, t time.Time) ([]*ent.Event, error) {
	const limit = 100

	return s.c.Event.
		Query().
		Where(
			event.CreatedAtGT(t),
		).
		WithApproval().
		WithDeployment().
		Limit(100).
		All(ctx)
}

func (s *Store) CreateEvent(ctx context.Context, e *ent.Event) (*ent.Event, error) {
	create := s.c.Event.
		Create().
		SetKind(e.Kind).
		SetType(e.Type)

	if e.Kind == event.KindDeployment {
		create = create.SetDeploymentID(e.DeploymentID)
	} else if e.Kind == event.KindApproval {
		create = create.SetApprovalID(e.ApprovalID)
	}

	return create.Save(ctx)
}
