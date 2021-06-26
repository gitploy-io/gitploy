package interactor

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
)

func (i *Interactor) FindChatUserWithUserByID(ctx context.Context, id string) (*ent.ChatUser, error) {
	return i.store.FindChatUserWithUserByID(ctx, id)
}

func (i *Interactor) SaveChatUser(ctx context.Context, u *ent.User, cu *ent.ChatUser) (*ent.ChatUser, error) {
	_, err := i.store.FindChatUserByID(ctx, cu.ID)
	if ent.IsNotFound(err) {
		return i.store.CreateChatUser(ctx, u, cu)
	}

	return i.store.UpdateChatUser(ctx, u, cu)
}
