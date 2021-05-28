package store

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/perm"
	"github.com/hanjunlee/gitploy/ent/repo"
	"github.com/hanjunlee/gitploy/ent/user"
)

func (s *Store) ListRepos(ctx context.Context, u *ent.User, page, perPage int) ([]*ent.Repo, error) {
	// TODO: support sort by
	ps, err := s.c.Perm.
		Query().
		Where(
			perm.HasUserWith(user.IDEQ(u.ID)),
		).
		Limit(perPage).
		Offset(offset(page, perPage)).
		WithRepo().
		All(ctx)
	if err != nil {
		return nil, err
	}

	rs := []*ent.Repo{}

	for _, p := range ps {
		rs = append(rs, p.Edges.Repo)
	}
	return rs, nil
}

func (s *Store) FindRepo(ctx context.Context, u *ent.User, id string) (*ent.Repo, error) {
	p, err := s.c.Perm.
		Query().
		Where(
			perm.And(
				perm.HasUserWith(user.IDEQ(u.ID)),
				perm.HasRepoWith(repo.IDEQ(id)),
			),
		).
		WithRepo().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return p.Edges.Repo, nil
}
