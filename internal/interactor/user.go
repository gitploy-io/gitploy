package interactor

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/vo"
)

func (i *Interactor) GetRemoteUserByToken(ctx context.Context, token string) (*vo.RemoteUser, error) {
	return i.SCM.GetUser(ctx, token)
}

func (i *Interactor) SaveUser(ctx context.Context, u *ent.User) (*ent.User, error) {
	_, err := i.FindUserByID(ctx, u.ID)
	if ent.IsNotFound(err) {
		u, _ = i.CreateUser(ctx, u)
	} else if err != nil {
		return nil, err
	}

	return i.UpdateUser(ctx, u)
}
