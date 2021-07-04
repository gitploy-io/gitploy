package store

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/deployment"
	"github.com/hanjunlee/gitploy/ent/repo"
)

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

func (s *Store) FindDeploymentWithEdgesByID(ctx context.Context, id int) (*ent.Deployment, error) {
	return s.c.Deployment.
		Query().
		Where(
			deployment.IDEQ(id),
		).
		WithRepo().
		WithUser().
		First(ctx)
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

// TODO: Lock deployments for the index which has repo_id.
func (s *Store) GetNextDeploymentNumberOfRepo(ctx context.Context, d *ent.Deployment) (int, error) {
	cnt, err := s.c.Deployment.Query().
		Where(
			deployment.RepoID(d.RepoID),
		).
		Count(ctx)
	if err != nil {
		return 0, err
	}

	return cnt + 1, nil
}

// CreateDeployment always set the next number of deployment
// when it creates.
func (s *Store) CreateDeployment(ctx context.Context, d *ent.Deployment) (*ent.Deployment, error) {
	return s.c.Deployment.Create().
		SetNumber(d.Number).
		SetType(d.Type).
		SetRef(d.Ref).
		SetEnv(d.Env).
		SetUserID(d.UserID).
		SetRepoID(d.RepoID).
		Save(ctx)
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
