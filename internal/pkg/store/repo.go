package store

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/deployment"
	"github.com/hanjunlee/gitploy/ent/perm"
	"github.com/hanjunlee/gitploy/ent/repo"
	"github.com/hanjunlee/gitploy/ent/user"
)

func (s *Store) ListReposOfUser(ctx context.Context, u *ent.User, q string, page, perPage int) ([]*ent.Repo, error) {
	return s.c.Repo.
		Query().
		Where(
			repo.And(
				repo.HasPermsWith(perm.HasUserWith(user.IDEQ(u.ID))),
				repo.NameContains(q),
			),
		).
		Limit(perPage).
		Offset(offset(page, perPage)).
		WithDeployments(func(dq *ent.DeploymentQuery) {
			dq.Order(ent.Desc(deployment.FieldCreatedAt)).
				Limit(5)
		}).
		All(ctx)
}

func (s *Store) ListSortedReposOfUser(ctx context.Context, u *ent.User, q string, page, perPage int) ([]*ent.Repo, error) {
	return s.c.Repo.
		Query().
		Where(
			repo.And(
				repo.HasPermsWith(perm.HasUserWith(user.IDEQ(u.ID))),
				repo.NameContains(q),
			),
		).
		Order(ent.Desc(repo.FieldLatestDeployedAt)).
		Limit(perPage).
		Offset(offset(page, perPage)).
		WithDeployments(func(dq *ent.DeploymentQuery) {
			dq.Order(ent.Desc(deployment.FieldCreatedAt)).
				Limit(5)
		}).
		All(ctx)
}

func (s *Store) UpdateRepo(ctx context.Context, r *ent.Repo) (*ent.Repo, error) {
	return s.c.Repo.
		UpdateOne(r).
		SetConfigPath(r.ConfigPath).
		Save(ctx)
}

func (s *Store) FindRepoByID(ctx context.Context, id string) (*ent.Repo, error) {
	return s.c.Repo.Get(ctx, id)
}

func (s *Store) FindRepoByNamespaceName(ctx context.Context, u *ent.User, namespace, name string) (*ent.Repo, error) {
	r, err := s.c.Repo.
		Query().
		Where(
			repo.And(
				repo.NamespaceEQ(namespace),
				repo.NameEQ(name),
				repo.HasPermsWith(perm.HasUserWith(user.IDEQ(u.ID))),
			),
		).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *Store) Activate(ctx context.Context, r *ent.Repo) (*ent.Repo, error) {
	return s.c.Repo.
		UpdateOne(r).
		SetActive(true).
		SetWebhookID(r.WebhookID).
		Save(ctx)
}

func (s *Store) Deactivate(ctx context.Context, r *ent.Repo) (*ent.Repo, error) {
	return s.c.Repo.
		UpdateOne(r).
		SetActive(false).
		SetWebhookID(0).
		Save(ctx)
}
