//go:generate mockgen -source ./interface.go -destination ./mock/interactor.go -package mock

package users

import (
	"context"

	i "github.com/gitploy-io/gitploy/internal/interactor"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/extent"
)

type (
	Interactor interface {
		SearchUsers(ctx context.Context, opts *i.SearchUsersOptions) ([]*ent.User, error)
		FindUserByID(ctx context.Context, id int64) (*ent.User, error)
		GetRateLimit(ctx context.Context, u *ent.User) (*extent.RateLimit, error)
		UpdateUser(ctx context.Context, u *ent.User) (*ent.User, error)
		DeleteUser(ctx context.Context, u *ent.User) error
	}
)
