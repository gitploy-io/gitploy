package store

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
)

func (s *Store) FindUser() (*ent.User, error) {
	return s.c.User.Get(context.Background(), "17633736")
}

func (s *Store) FindChatUserByID(ctx context.Context, id string) (*ent.ChatUser, error) {
	return s.c.ChatUser.Get(ctx, id)
}

func (s *Store) CreateChatUser(ctx context.Context, u *ent.User, cu *ent.ChatUser) (*ent.ChatUser, error) {
	return s.c.ChatUser.
		Create().
		SetID(cu.ID).
		SetToken(cu.Token).
		SetBotToken(cu.BotToken).
		SetRefresh(cu.Refresh).
		SetExpiry(cu.Expiry).
		Save(ctx)
}

func (s *Store) UpdateChatUser(ctx context.Context, cu *ent.ChatUser) (*ent.ChatUser, error) {
	return s.c.ChatUser.
		UpdateOneID(cu.ID).
		SetToken(cu.Token).
		SetBotToken(cu.BotToken).
		SetRefresh(cu.Refresh).
		SetExpiry(cu.Expiry).
		Save(ctx)
}
