package store

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
)

func (s *Store) CreateChatCallback(ctx context.Context, cb *ent.ChatCallback) (*ent.ChatCallback, error) {
	return s.c.ChatCallback.
		Create().
		SetID(cb.ID).
		SetState(cb.State).
		SetType(cb.Type).
		Save(ctx)
}
