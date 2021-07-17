package store

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/perm"
	"github.com/hanjunlee/gitploy/ent/repo"
	"github.com/hanjunlee/gitploy/ent/user"
)

func (s *Store) ListPermsOfRepo(ctx context.Context, r *ent.Repo, q string, page, perPage int) ([]*ent.Perm, error) {
	return s.c.Perm.
		Query().
		Where(
			perm.And(
				perm.HasRepoWith(repo.IDEQ(r.ID)),
				perm.HasUserWith(user.LoginContains(q)),
			),
		).
		WithRepo().
		WithUser().
		Limit(perPage).
		Offset(offset(page, perPage)).
		All(ctx)
}

func (s *Store) FindPermOfRepo(ctx context.Context, r *ent.Repo, u *ent.User) (*ent.Perm, error) {
	return s.c.Perm.
		Query().
		Where(
			perm.And(
				perm.HasUserWith(user.IDEQ(u.ID)),
				perm.HasRepoWith(repo.IDEQ(r.ID)),
			),
		).
		WithRepo().
		WithUser().
		Only(ctx)
}
