package interactor

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
)

func (i *Interactor) SaveChatUser(ctx context.Context, u *ent.User, cu *ent.ChatUser) (*ent.ChatUser, error) {
	_, err := i.FindChatUserByID(ctx, cu.ID)
	if ent.IsNotFound(err) {
		return i.CreateChatUser(ctx, u, cu)
	}

	return i.UpdateChatUser(ctx, u, cu)
}
