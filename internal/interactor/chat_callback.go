package interactor

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
)

func (i *Interactor) CreateChatCallback(ctx context.Context, cb *ent.ChatCallback) (*ent.ChatCallback, error) {
	return i.store.CreateChatCallback(ctx, cb)
}
