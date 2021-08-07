package store

import (
	"context"
	"time"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/deployment"
	"github.com/hanjunlee/gitploy/ent/repo"
)

func (s *Store) ListInactiveDeploymentsLessThanTime(ctx context.Context, t time.Time, page, perPage int) ([]*ent.Deployment, error) {
	return s.c.Deployment.
		Query().
		Where(
			deployment.And(
				deployment.StatusIn(deployment.StatusWaiting, deployment.StatusCreated),
				deployment.CreatedAtLT(t),
			),
		).
		WithRepo().
		WithUser().
		Limit(perPage).
		Offset(offset(page, perPage)).
		All(ctx)
}

func (s *Store) ListDeploymentsOfRepo(ctx context.Context, r *ent.Repo, env string, status string, page, perPage int) ([]*ent.Deployment, error) {
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
		WithRepo().
		Limit(perPage).
		Offset(offset(page, perPage)).
		All(ctx)
}

func (s *Store) FindDeploymentByID(ctx context.Context, id int) (*ent.Deployment, error) {
	return s.c.Deployment.
		Query().
		Where(
			deployment.IDEQ(id),
		).
		WithRepo().
		WithUser().
		WithDeploymentStatuses().
		First(ctx)
}

func (s *Store) FindDeploymentOfRepoByNumber(ctx context.Context, r *ent.Repo, number int) (*ent.Deployment, error) {
	return s.c.Deployment.
		Query().
		Where(
			deployment.And(
				deployment.RepoID(r.ID),
				deployment.NumberEQ(number),
			),
		).
		WithRepo().
		WithUser().
		WithDeploymentStatuses().
		First(ctx)
}

func (s *Store) FindDeploymentByUID(ctx context.Context, uid int64) (*ent.Deployment, error) {
	return s.c.Deployment.
		Query().
		Where(
			deployment.UIDEQ(uid),
		).
		WithRepo().
		WithUser().
		WithDeploymentStatuses().
		First(ctx)
}

// TODO: Lock deployments for the index which has repo_id.
func (s *Store) GetNextDeploymentNumberOfRepo(ctx context.Context, r *ent.Repo) (int, error) {
	var v []struct {
		RepoID string `json:"repo_id"`
		Max    int    `json:"max"`
	}
	err := s.c.Deployment.Query().
		Where(
			deployment.RepoID(r.ID),
		).
		GroupBy(deployment.FieldRepoID).
		Aggregate(ent.Max(deployment.FieldNumber)).
		Scan(ctx, &v)
	if err != nil {
		return 0, err
	}

	max := 0
	if len(v) != 0 {
		max = v[0].Max + 1
	}

	return max, nil
}

// CreateDeployment always set the next number of deployment
// when it creates.
func (s *Store) CreateDeployment(ctx context.Context, d *ent.Deployment) (*ent.Deployment, error) {
	return s.c.Deployment.Create().
		SetNumber(d.Number).
		SetType(d.Type).
		SetRef(d.Ref).
		SetEnv(d.Env).
		SetIsRollback(d.IsRollback).
		SetIsApprovalEnabled(d.IsApprovalEnabled).
		SetRequiredApprovalCount(d.RequiredApprovalCount).
		SetUserID(d.UserID).
		SetRepoID(d.RepoID).
		Save(ctx)
}

func (s *Store) UpdateDeployment(ctx context.Context, d *ent.Deployment) (*ent.Deployment, error) {
	return s.c.Deployment.UpdateOne(d).
		SetType(d.Type).
		SetRef(d.Ref).
		SetEnv(d.Env).
		SetUID(d.UID).
		SetSha(d.Sha).
		SetHTMLURL(d.HTMLURL).
		SetIsRollback(d.IsRollback).
		SetIsApprovalEnabled(d.IsApprovalEnabled).
		SetRequiredApprovalCount(d.RequiredApprovalCount).
		SetStatus(d.Status).
		Save(ctx)
}
