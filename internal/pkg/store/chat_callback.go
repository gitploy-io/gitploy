package store

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/chatcallback"
)

func (s *Store) CreateDeployChatCallback(ctx context.Context, cu *ent.ChatUser, repo *ent.Repo, cb *ent.ChatCallback) (*ent.ChatCallback, error) {
	return s.c.ChatCallback.
		Create().
		SetID(cb.ID).
		SetState(cb.State).
		SetType(cb.Type).
		SetChatUser(cu).
		SetRepo(repo).
		Save(ctx)
}

func (s *Store) FindChatCallbackByID(ctx context.Context, id string) (*ent.ChatCallback, error) {
	return s.c.ChatCallback.
		Get(ctx, id)
}

func (s *Store) FindChatCallbackWithEdgesByID(ctx context.Context, id string) (*ent.ChatCallback, error) {
	return s.c.ChatCallback.
		Query().
		Where(
			chatcallback.IDEQ(id),
		).
		WithChatUser().
		WithRepo().
		First(ctx)
}

func (s *Store) CloseChatCallback(ctx context.Context, cb *ent.ChatCallback) (*ent.ChatCallback, error) {
	return s.c.ChatCallback.
		UpdateOne(cb).
		Save(ctx)
}
