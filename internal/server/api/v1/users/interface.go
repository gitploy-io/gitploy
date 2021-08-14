package users

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/vo"
)

type (
	Interactor interface {
		FindUserByID(ctx context.Context, id string) (*ent.User, error)
		GetRateLimit(ctx context.Context, u *ent.User) (*vo.RateLimit, error)
	}
)
