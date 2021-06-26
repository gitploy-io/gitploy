package store

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/chatuser"
)

func (s *Store) FindChatUserByID(ctx context.Context, id string) (*ent.ChatUser, error) {
	return s.c.ChatUser.Get(ctx, id)
}

func (s *Store) FindChatUserWithUserByID(ctx context.Context, id string) (*ent.ChatUser, error) {
	return s.c.ChatUser.
		Query().
		Where(
			chatuser.IDEQ(id),
		).
		WithUser().
		First(ctx)
}
