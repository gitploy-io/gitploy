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

func (s *Store) CreateNotification(ctx context.Context, n *ent.Notification) (*ent.Notification, error) {
	return s.c.Notification.
		Create().
		SetType(n.Type).
		SetResourceID(n.ResourceID).
		SetNotified(n.Notified).
		SetUserID(n.UserID).
		Save(ctx)
}

func (s *Store) SetNotificationDone(ctx context.Context, n *ent.Notification) (*ent.Notification, error) {
	return s.c.Notification.
		UpdateOne(n).
		SetNotified(true).
		Save(ctx)
}

func (s *Store) ListNotifications(ctx context.Context, u *ent.User, page, perPage int) ([]*ent.Notification, error) {
	return s.c.Notification.
		Query().
		Where(
			notification.UserID(u.ID),
		).
		Limit(perPage).
		Offset(offset(page, perPage)).
		Order(ent.Desc(notification.FieldCreatedAt)).
		All(ctx)
}

func (s *Store) FindNotificationByID(ctx context.Context, id int) (*ent.Notification, error) {
	return s.c.Notification.Get(ctx, id)
}

func (s *Store) SetNotificationChecked(ctx context.Context, n *ent.Notification) (*ent.Notification, error) {
	return s.c.Notification.
		UpdateOne(n).
		SetChecked(true).
		Save(ctx)
}
