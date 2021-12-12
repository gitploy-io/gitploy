//go:generate mockgen -source ./interface.go -destination ./mock/interactor.go -package mock

package users

import (
	"context"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/extent"
)

type (
	Interactor interface {
		ListUsers(ctx context.Context, login string, page, perPage int) ([]*ent.User, error)
		FindUserByID(ctx context.Context, id int64) (*ent.User, error)
		GetRateLimit(ctx context.Context, u *ent.User) (*extent.RateLimit, error)
		UpdateUser(ctx context.Context, u *ent.User) (*ent.User, error)
		DeleteUser(ctx context.Context, u *ent.User) error
	}
)
