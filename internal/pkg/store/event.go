package store

import (
	"context"
	"time"

	"entgo.io/ent/dialect"
	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/event"
	"github.com/gitploy-io/gitploy/ent/notificationrecord"
)

// ListEventsGreaterThanTime returns all events for deployment and approval
// that are greater than the time.
//
// It processes eager loading, especially, Approval loads a repository of deployment.
func (s *Store) ListEventsGreaterThanTime(ctx context.Context, t time.Time) ([]*ent.Event, error) {
	const limit = 100

	return s.c.Event.
		Query().
		Where(
			event.CreatedAtGT(t),
		).
		WithApproval(func(aq *ent.ApprovalQuery) {
			aq.
				WithUser().
				WithDeployment(func(dq *ent.DeploymentQuery) {
					dq.
						WithUser().
						WithRepo()
				})
		}).
		WithDeployment(func(dq *ent.DeploymentQuery) {
			dq.
				WithUser().
				WithRepo().
				WithDeploymentStatuses()
		}).
		Limit(limit).
		All(ctx)
}

func (s *Store) CreateEvent(ctx context.Context, e *ent.Event) (*ent.Event, error) {
	qry := s.c.Event.
		Create().
		SetKind(e.Kind).
		SetType(e.Type)

	if e.Type == event.TypeDeleted {
		qry = qry.SetDeletedID(e.DeletedID)
	} else if e.Kind == event.KindDeployment {
		qry = qry.SetDeploymentID(e.DeploymentID)
	} else if e.Kind == event.KindApproval {
		qry = qry.SetApprovalID(e.ApprovalID)
	}

	return qry.Save(ctx)
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
