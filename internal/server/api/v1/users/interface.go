package users

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
)

type (
	Interactor interface {
		FindUserByID(ctx context.Context, id string) (*ent.User, error)
	}
)
