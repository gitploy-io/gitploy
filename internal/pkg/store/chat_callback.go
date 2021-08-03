package store

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/chatcallback"
)

func (s *Store) CreateChatCallback(ctx context.Context, cu *ent.ChatUser, repo *ent.Repo, cb *ent.ChatCallback) (*ent.ChatCallback, error) {
	return s.c.ChatCallback.
		Create().
		SetType(cb.Type).
		SetChatUser(cu).
		SetRepo(repo).
		Save(ctx)
}

func (s *Store) FindChatCallbackByHash(ctx context.Context, hash string) (*ent.ChatCallback, error) {
	return s.c.ChatCallback.
		Query().
		Where(
			chatcallback.HashEQ(hash),
		).
		WithChatUser().
		WithRepo().
		First(ctx)
}

func (s *Store) CloseChatCallback(ctx context.Context, cb *ent.ChatCallback) (*ent.ChatCallback, error) {
	return s.c.ChatCallback.
		UpdateOne(cb).
		SetIsOpened(false).
		Save(ctx)
}
