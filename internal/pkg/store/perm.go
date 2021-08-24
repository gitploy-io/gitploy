package store

import (
	"context"
	"time"

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

func (s *Store) CreatePerm(ctx context.Context, p *ent.Perm) (*ent.Perm, error) {
	return s.c.Perm.
		Create().
		SetRepoPerm(p.RepoPerm).
		SetUserID(p.UserID).
		SetRepoID(p.RepoID).
		Save(ctx)
}

func (s *Store) UpdatePerm(ctx context.Context, p *ent.Perm) (*ent.Perm, error) {
	return s.c.Perm.
		UpdateOne(p).
		SetRepoPerm(p.RepoPerm).
		Save(ctx)
}

func (s *Store) DeletePermsOfUserLessThanUpdatedAt(ctx context.Context, u *ent.User, t time.Time) error {
	return s.WithTx(ctx, func(tx *ent.Tx) error {
		_, err := tx.Perm.
			Delete().
			Where(
				perm.And(
					perm.UserIDEQ(u.ID),
					perm.UpdatedAtLT(t),
				),
			).
			Exec(ctx)
		return err
	})
}
