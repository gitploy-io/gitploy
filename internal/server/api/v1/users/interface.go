package users

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
)

type (
	Interactor interface {
		FindUserByID(ctx context.Context, id string) (*ent.User, error)
		ListNotifications(ctx context.Context, u *ent.User, page, perPage int) ([]*ent.Notification, error)
		FindNotificationByID(ctx context.Context, id int) (*ent.Notification, error)
		UpdateNotification(ctx context.Context, n *ent.Notification) (*ent.Notification, error)
	}
)
