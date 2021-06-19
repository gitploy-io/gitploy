package store

import (
	"context"
	"time"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/notification"
)

func (s *Store) ListNotificationsFromTime(ctx context.Context, t time.Time) ([]*ent.Notification, error) {
	return s.c.Notification.
		Query().
		Where(
			notification.CreatedAtGTE(t),
		).
		All(ctx)
}

func (s *Store) SetNotificationDone(ctx context.Context, n *ent.Notification) (*ent.Notification, error) {
	return s.c.Notification.
		UpdateOne(n).
		SetNotified(true).
		Save(ctx)
}
