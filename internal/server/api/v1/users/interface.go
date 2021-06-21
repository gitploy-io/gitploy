package users

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
)

type (
	Interactor interface {
		FindUserWithChatUserByID(ctx context.Context, id string) (*ent.User, error)
	}
)
