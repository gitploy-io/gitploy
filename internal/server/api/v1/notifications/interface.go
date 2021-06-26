package notifications

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
)

type (
	Interactor interface {
		ListNotifications(ctx context.Context, u *ent.User, page, perPage int) ([]*ent.Notification, error)
		FindNotificationByID(ctx context.Context, id int) (*ent.Notification, error)
		SetNotificationChecked(ctx context.Context, n *ent.Notification) (*ent.Notification, error)
	}
)
