package sync

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
)

type (
	Interactor interface {
		FindUserByHash(ctx context.Context, hash string) (*ent.User, error)
		Sync(ctx context.Context, user *ent.User) error
	}
)
