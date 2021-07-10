package store

import (
	"context"
	"time"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/notification"
)

func (s *Store) ListNotifications(ctx context.Context, u *ent.User, page, perPage int) ([]*ent.Notification, error) {
	return s.c.Notification.
		Query().
		Where(
			notification.UserID(u.ID),
		).
		WithUser().
		WithRepo().
		WithDeployment().
		Limit(perPage).
		Offset(offset(page, perPage)).
		Order(ent.Desc(notification.FieldCreatedAt)).
		All(ctx)
}

func (s *Store) ListNotificationsFromTime(ctx context.Context, t time.Time) ([]*ent.Notification, error) {
	return s.c.Notification.
		Query().
		Where(
			notification.CreatedAtGTE(t),
		).
		WithUser().
		WithRepo().
		WithDeployment().
		All(ctx)
}

func (s *Store) FindNotificationByID(ctx context.Context, id int) (*ent.Notification, error) {
	return s.c.Notification.Get(ctx, id)
}

func (s *Store) CreateNotification(ctx context.Context, n *ent.Notification) (*ent.Notification, error) {
	nc := s.c.Notification.
		Create().
		SetType(n.Type).
		SetNotified(n.Notified).
		SetUserID(n.UserID).
		SetRepoID(n.RepoID)

	if n.Type == notification.TypeDeployment {
		nc = nc.SetDeploymentID(n.DeploymentID)
	}

	return nc.Save(ctx)
}

func (s *Store) UpdateNotification(ctx context.Context, n *ent.Notification) (*ent.Notification, error) {
	return s.c.Notification.
		UpdateOne(n).
		SetNotified(n.Notified).
		SetChecked(n.Checked).
		Save(ctx)
}
