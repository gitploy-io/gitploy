package web

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
)

type (
	Store interface {
		FindUserByID(ctx context.Context, id string) (*ent.User, error)
		CreateUser(ctx context.Context, u *ent.User) (*ent.User, error)
		UpdateUser(ctx context.Context, u *ent.User) (*ent.User, error)
	}

	SCM interface {
		GetUser(ctx context.Context, token string) (*ent.User, error)
	}
)
