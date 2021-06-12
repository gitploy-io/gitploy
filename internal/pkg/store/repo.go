package store

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/deployment"
	"github.com/hanjunlee/gitploy/ent/perm"
	"github.com/hanjunlee/gitploy/ent/repo"
	"github.com/hanjunlee/gitploy/ent/user"
)

func (s *Store) ListRepos(ctx context.Context, u *ent.User, q string, page, perPage int) ([]*ent.Repo, error) {
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

func (s *Store) ListSortedRepos(ctx context.Context, u *ent.User, q string, page, perPage int) ([]*ent.Repo, error) {
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

func (s *Store) ListDeployments(ctx context.Context, r *ent.Repo, env string, status string, page, perPage int) ([]*ent.Deployment, error) {
	q := s.c.Deployment.
		Query()

	q = q.Where(
		deployment.HasRepoWith(repo.IDEQ(r.ID)),
	)

	if env != "" {
		q = q.Where(
			deployment.EnvEQ(env),
		)
	}

	if status != "" {
		q = q.Where(
			deployment.StatusEQ(deployment.Status(status)),
		)
	}

	return q.Order(
		ent.Desc(deployment.FieldCreatedAt),
	).
		WithUser(func(uq *ent.UserQuery) {
			uq.Select("id", "login", "avatar", "created_at", "updated_at")
		}).
		Limit(perPage).
		Offset(offset(page, perPage)).
		All(ctx)
}

func (s *Store) FindLatestDeployment(ctx context.Context, r *ent.Repo, env string) (*ent.Deployment, error) {
	return s.c.Deployment.
		Query().
		Where(
			deployment.EnvEQ(env),
			deployment.HasRepoWith(repo.IDEQ(r.ID)),
		).
		Order(
			ent.Desc(deployment.FieldCreatedAt),
		).
		WithUser(func(uq *ent.UserQuery) {
			uq.Select("id", "login", "avatar")
		}).
		First(ctx)
}

func (s *Store) CreateDeployment(ctx context.Context, u *ent.User, r *ent.Repo, d *ent.Deployment) (*ent.Deployment, error) {
	d, err := s.c.Deployment.Create().
		SetType(d.Type).
		SetRef(d.Ref).
		SetEnv(d.Env).
		SetUserID(u.ID).
		SetRepoID(r.ID).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	s.c.Repo.UpdateOneID(r.ID).
		SetLatestDeployedAt(d.CreatedAt).
		Save(ctx)

	return d, nil
}

func (s *Store) UpdateDeployment(ctx context.Context, d *ent.Deployment) (*ent.Deployment, error) {
	return s.c.Deployment.UpdateOne(d).
		SetUID(d.UID).
		SetType(d.Type).
		SetRef(d.Ref).
		SetSha(d.Sha).
		SetEnv(d.Env).
		SetStatus(d.Status).
		Save(ctx)
}

func (s *Store) FindPerm(ctx context.Context, u *ent.User, repoID string) (*ent.Perm, error) {
	return s.c.Perm.
		Query().
		Where(
			perm.And(
				perm.HasUserWith(user.IDEQ(u.ID)),
				perm.HasRepoWith(repo.IDEQ(repoID)),
			),
		).
		Only(ctx)
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
