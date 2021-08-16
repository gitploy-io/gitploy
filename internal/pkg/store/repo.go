package store

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/deployment"
	"github.com/hanjunlee/gitploy/ent/perm"
	"github.com/hanjunlee/gitploy/ent/repo"
)

func (s *Store) ListReposOfUser(ctx context.Context, u *ent.User, q string, page, perPage int) ([]*ent.Repo, error) {
	return s.c.Repo.
		Query().
		Where(func(s *sql.Selector) {
			t := sql.Table(perm.Table)
			s.
				Join(t).
				On(t.C(perm.FieldRepoID), s.C(repo.FieldID)).
				Where(sql.EQ(t.C(perm.FieldUserID), u.ID))
		}).
		Where(
			repo.NameContains(q),
		).
		WithDeployments(func(dq *ent.DeploymentQuery) {
			dq.
				WithUser().
				Order(ent.Desc(deployment.FieldCreatedAt)).
				Limit(3)
		}).
		Limit(perPage).
		Offset(offset(page, perPage)).
		All(ctx)
}

func (s *Store) ListSortedReposOfUser(ctx context.Context, u *ent.User, q string, page, perPage int) ([]*ent.Repo, error) {
	return s.c.Repo.
		Query().
		Where(func(s *sql.Selector) {
			t := sql.Table(perm.Table)
			s.
				Join(t).
				On(t.C(perm.FieldRepoID), s.C(repo.FieldID)).
				Where(sql.EQ(t.C(perm.FieldUserID), u.ID))
		}).
		Where(
			repo.And(
				repo.NameContains(q),
			),
		).
		WithDeployments(func(dq *ent.DeploymentQuery) {
			dq.
				WithUser().
				Order(ent.Desc(deployment.FieldCreatedAt)).
				Limit(3)
		}).
		Order(ent.Desc(repo.FieldLatestDeployedAt)).
		Limit(perPage).
		Offset(offset(page, perPage)).
		All(ctx)
}

func (s *Store) UpdateRepo(ctx context.Context, r *ent.Repo) (*ent.Repo, error) {
	return s.c.Repo.
		UpdateOne(r).
		SetConfigPath(r.ConfigPath).
		Save(ctx)
}

func (s *Store) FindRepoOfUserByID(ctx context.Context, u *ent.User, id string) (*ent.Repo, error) {
	return s.c.Repo.
		Query().
		Where(func(s *sql.Selector) {
			t := sql.Table(perm.Table)
			s.
				Join(t).
				On(t.C(perm.FieldRepoID), s.C(repo.FieldID)).
				Where(sql.EQ(t.C(perm.FieldUserID), u.ID))
		}).
		Where(
			repo.IDEQ(id),
		).
		Only(ctx)
}

func (s *Store) FindRepoOfUserByNamespaceName(ctx context.Context, u *ent.User, namespace, name string) (*ent.Repo, error) {
	r, err := s.c.Repo.
		Query().
		Where(func(s *sql.Selector) {
			t := sql.Table(perm.Table)
			s.
				Join(t).
				On(t.C(perm.FieldRepoID), s.C(repo.FieldID)).
				Where(sql.EQ(t.C(perm.FieldUserID), u.ID))
		}).
		Where(
			repo.And(
				repo.NamespaceEQ(namespace),
				repo.NameEQ(name),
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
