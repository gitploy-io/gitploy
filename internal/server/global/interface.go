package global

import (
	"context"

	"github.com/gitploy-io/gitploy/model/ent"
)

type (
	Interactor interface {
		FindUserByHash(ctx context.Context, hash string) (*ent.User, error)
	}
)
