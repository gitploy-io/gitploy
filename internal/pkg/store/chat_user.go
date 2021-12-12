package store

import (
	"context"

	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/ent/chatuser"
)

func (s *Store) FindChatUserByID(ctx context.Context, id string) (*ent.ChatUser, error) {
	return s.c.ChatUser.
		Query().
		Where(
			chatuser.IDEQ(id),
		).
		WithUser().
		First(ctx)
}

func (s *Store) CreateChatUser(ctx context.Context, cu *ent.ChatUser) (*ent.ChatUser, error) {
	return s.c.ChatUser.
		Create().
		SetID(cu.ID).
		SetToken(cu.Token).
		SetBotToken(cu.BotToken).
		SetRefresh(cu.Refresh).
		SetExpiry(cu.Expiry).
		SetUserID(cu.UserID).
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

func (s *Store) DeleteChatUser(ctx context.Context, cu *ent.ChatUser) error {
	return s.c.ChatUser.
		DeleteOneID(cu.ID).
		Exec(ctx)
}
