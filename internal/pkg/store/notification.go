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
		All(ctx)
}

func (s *Store) FindNotificationByID(ctx context.Context, id int) (*ent.Notification, error) {
	return s.c.Notification.Get(ctx, id)
}

func (s *Store) CreateNotification(ctx context.Context, n *ent.Notification) (*ent.Notification, error) {
	return s.c.Notification.
		Create().
		SetType(n.Type).
		SetRepoNamespace(n.RepoNamespace).
		SetRepoName(n.RepoName).
		SetDeploymentNumber(n.DeploymentNumber).
		SetDeploymentType(n.DeploymentType).
		SetDeploymentRef(n.DeploymentRef).
		SetDeploymentEnv(n.DeploymentEnv).
		SetDeploymentStatus(n.DeploymentStatus).
		SetDeploymentLogin(n.DeploymentLogin).
		SetNotified(n.Notified).
		SetUserID(n.UserID).
		Save(ctx)
}

func (s *Store) UpdateNotification(ctx context.Context, n *ent.Notification) (*ent.Notification, error) {
	return s.c.Notification.
		UpdateOne(n).
		SetRepoNamespace(n.RepoNamespace).
		SetRepoName(n.RepoName).
		SetDeploymentNumber(n.DeploymentNumber).
		SetDeploymentType(n.DeploymentType).
		SetDeploymentRef(n.DeploymentRef).
		SetDeploymentEnv(n.DeploymentEnv).
		SetDeploymentStatus(n.DeploymentStatus).
		SetDeploymentLogin(n.DeploymentLogin).
		SetNotified(n.Notified).
		SetChecked(n.Checked).
		Save(ctx)
}
