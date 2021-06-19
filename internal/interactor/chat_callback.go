package interactor

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
)

func (i *Interactor) CreateDeployChatCallback(ctx context.Context, cu *ent.ChatUser, repo *ent.Repo, cb *ent.ChatCallback) (*ent.ChatCallback, error) {
	return i.store.CreateDeployChatCallback(ctx, cu, repo, cb)
}

func (i *Interactor) FindChatCallbackByID(ctx context.Context, id string) (*ent.ChatCallback, error) {
	return i.store.FindChatCallbackByID(ctx, id)
}

func (i *Interactor) FindChatCallbackWithEdgesByID(ctx context.Context, id string) (*ent.ChatCallback, error) {
	return i.store.FindChatCallbackWithEdgesByID(ctx, id)
}

func (i *Interactor) CloseChatCallback(ctx context.Context, cb *ent.ChatCallback) (*ent.ChatCallback, error) {
	return i.store.CloseChatCallback(ctx, cb)
}
