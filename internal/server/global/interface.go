package global

import (
	"context"

	"github.com/gitploy-io/gitploy/ent"
)

type (
	Interactor interface {
		FindUserByHash(ctx context.Context, hash string) (*ent.User, error)
	}
)
