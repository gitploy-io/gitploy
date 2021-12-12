package store

import (
	"context"

	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/ent/callback"
)

func (s *Store) CreateCallback(ctx context.Context, cb *ent.Callback) (*ent.Callback, error) {
	return s.c.Callback.
		Create().
		SetType(cb.Type).
		SetRepoID(cb.RepoID).
		Save(ctx)
}

func (s *Store) FindCallbackByHash(ctx context.Context, hash string) (*ent.Callback, error) {
	return s.c.Callback.
		Query().
		Where(
			callback.HashEQ(hash),
		).
		WithRepo().
		First(ctx)
}
