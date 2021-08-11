package store

import (
	"context"
	"time"

	"entgo.io/ent/dialect"
	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/event"
	"github.com/hanjunlee/gitploy/ent/notificationrecord"
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
		Limit(limit).
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

func (s *Store) CheckNotificationRecordOfEvent(ctx context.Context, e *ent.Event) bool {
	var hasRecord bool

	s.WithTx(ctx, func(tx *ent.Tx) error {
		query := tx.NotificationRecord.
			Query().
			Where(
				notificationrecord.EventIDEQ(e.ID),
			)

		// Use "SELECT ... FOR UPDATE" for MySQL and Postgres.
		if tx.GetDriverDialect() == dialect.MySQL || tx.GetDriverDialect() == dialect.Postgres {
			query = query.
				ForUpdate()
		}

		if cnt, _ := query.Count(ctx); cnt != 0 {
			hasRecord = true
			return nil
		}

		if _, err := tx.NotificationRecord.
			Create().
			SetEventID(e.ID).
			Save(ctx); err != nil {
			return err
		}

		return nil
	})

	return hasRecord
}
