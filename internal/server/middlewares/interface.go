package middlewares

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
)

type (
	Interactor interface {
		FindUserByHash(ctx context.Context, hash string) (*ent.User, error)
	}
)
