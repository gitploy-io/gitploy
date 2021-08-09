package store

import (
	"context"
	"time"

	"entgo.io/ent/dialect"
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

func (s *Store) ListPublishingNotificaitonsGreaterThanTime(ctx context.Context, t time.Time) ([]*ent.Notification, error) {
	var ns []*ent.Notification

	if err := s.WithTx(ctx, func(tx *ent.Tx) error {
		var (
			err error
			now = time.Now()
		)

		query := tx.Notification.
			Query().
			Where(
				notification.And(
					notification.NotifiedEQ(false),
					notification.CreatedAtGTE(t),
					notification.CreatedAtLT(now),
				),
			)

		// Use "SELECT ... FOR UPDATE" for MySQL and Postgres.
		if tx.GetDriverDialect() == dialect.MySQL || tx.GetDriverDialect() == dialect.Postgres {
			query = query.
				ForUpdate()
		}

		ns, err = query.
			WithUser().
			All(ctx)
		if err != nil {
			return err
		}

		_, err = tx.Notification.
			Update().
			Where(
				notification.And(
					notification.CreatedAtGTE(t),
					notification.NotifiedEQ(false),
					notification.CreatedAtLT(now),
				),
			).
			SetNotified(true).
			Save(ctx)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return ns, nil
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
		SetApprovalStatus(n.ApprovalStatus).
		SetApprovalLogin(n.ApprovalLogin).
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
