//go:generate mockgen -source ./interface.go -destination ./mock/interactor.go -package mock

package users

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/vo"
)

type (
	Interactor interface {
		ListUsers(ctx context.Context, login string, page, perPage int) ([]*ent.User, error)
		FindUserByID(ctx context.Context, id string) (*ent.User, error)
		GetRateLimit(ctx context.Context, u *ent.User) (*vo.RateLimit, error)
		UpdateUser(ctx context.Context, u *ent.User) (*ent.User, error)
		DeleteUser(ctx context.Context, u *ent.User) error
	}
)
