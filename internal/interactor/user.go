package interactor

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
)

func (i *Interactor) FindUser() (*ent.User, error) {
	return i.store.FindUser()
}

func (i *Interactor) FindUserByHash(ctx context.Context, hash string) (*ent.User, error) {
	return i.store.FindUserByHash(ctx, hash)
}

func (i *Interactor) GetSCMUserByToken(ctx context.Context, token string) (*ent.User, error) {
	return i.scm.GetUser(ctx, token)
}

func (i *Interactor) SaveSCMUser(ctx context.Context, u *ent.User) (*ent.User, error) {
	_, err := i.store.FindUserByID(ctx, u.ID)
	if ent.IsNotFound(err) {
		u, _ = i.store.CreateUser(ctx, u)
	} else if err != nil {
		return nil, err
	}

	return i.store.UpdateUser(ctx, u)
}

func (i *Interactor) SaveChatUser(ctx context.Context, u *ent.User, cu *ent.ChatUser) (*ent.ChatUser, error) {
	_, err := i.store.FindChatUserByID(ctx, cu.ID)
	if ent.IsNotFound(err) {
		return i.store.CreateChatUser(ctx, u, cu)
	}

	return i.store.UpdateChatUser(ctx, cu)
}
