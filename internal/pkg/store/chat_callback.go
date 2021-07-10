package store

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/chatcallback"
)

func (s *Store) CreateDeployChatCallback(ctx context.Context, cu *ent.ChatUser, repo *ent.Repo, cb *ent.ChatCallback) (*ent.ChatCallback, error) {
	return s.c.ChatCallback.
		Create().
		SetState(cb.State).
		SetType(cb.Type).
		SetChatUser(cu).
		SetRepo(repo).
		Save(ctx)
}

func (s *Store) FindChatCallbackByState(ctx context.Context, state string) (*ent.ChatCallback, error) {
	return s.c.ChatCallback.
		Query().
		Where(
			chatcallback.StateEQ(state),
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
