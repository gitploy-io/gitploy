package middlewares

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
)

type (
	Store interface {
		FindUserByHash(ctx context.Context, hash string) (*ent.User, error)
	}
)
