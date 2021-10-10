package store

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/deployment"
	"github.com/gitploy-io/gitploy/ent/deploymentcount"
	"github.com/gitploy-io/gitploy/ent/perm"
	"github.com/gitploy-io/gitploy/ent/predicate"
	"go.uber.org/zap"
)

func (s *Store) SearchDeployments(ctx context.Context, u *ent.User, ss []deployment.Status, owned bool, from time.Time, to time.Time, page, perPage int) ([]*ent.Deployment, error) {
	statusIn := func(ss []deployment.Status) predicate.Deployment {
		if len(ss) == 0 {
			// if not status were provided,
			// it always make this predicate truly.
			return func(s *sql.Selector) {}
		}

		return deployment.StatusIn(ss...)
	}

	if owned {
		return s.c.Deployment.
			Query().
			Where(
				deployment.And(
					deployment.UserIDEQ(u.ID),
					statusIn(ss),
					deployment.CreatedAtGTE(from),
					deployment.CreatedAtLT(to),
				),
			).
			WithRepo().
			WithUser().
			Order(ent.Desc(deployment.FieldCreatedAt)).
			Offset(offset(page, perPage)).
			Limit(perPage).
			All(ctx)
	}

	return s.c.Deployment.
		Query().
		Where(func(s *sql.Selector) {
			t := sql.Table(perm.Table)

			// Join with Perm for Repo.ID
			s.Join(t).
				On(
					s.C(deployment.FieldRepoID),
					s.C(perm.FieldRepoID),
				).
				Where(
					sql.EQ(
						t.C(perm.FieldUserID),
						u.ID,
					),
				)
		}).
		Where(
			deployment.And(
				statusIn(ss),
				deployment.CreatedAtGTE(from),
				deployment.CreatedAtLT(to),
			),
		).
		WithRepo().
		WithUser().
		Order(ent.Desc(deployment.FieldCreatedAt)).
		Offset(offset(page, perPage)).
		Limit(perPage).
		All(ctx)
}

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
		deployment.RepoIDEQ(r.ID),
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

func (s *Store) GetNextDeploymentNumberOfRepo(ctx context.Context, r *ent.Repo) (int, error) {
	cnt, err := s.c.Deployment.Query().
		Where(
			deployment.RepoID(r.ID),
		).
		Count(ctx)
	if err != nil {
		return 0, err
	}

	return cnt + 1, nil
}

func (s *Store) FindPrevSuccessDeployment(ctx context.Context, d *ent.Deployment) (*ent.Deployment, error) {
	return s.c.Deployment.
		Query().
		Where(
			deployment.And(
				deployment.RepoIDEQ(d.RepoID),
				deployment.EnvEQ(d.Env),
				deployment.StatusEQ(deployment.StatusSuccess),
				deployment.CreatedAtLT(d.CreatedAt),
			),
		).
		Order(ent.Desc(deployment.FieldCreatedAt)).
		First(ctx)
}

// CreateDeployment create a new deployment, and
// it updates the 'latest_deployed_at' field of the repository.
func (s *Store) CreateDeployment(ctx context.Context, d *ent.Deployment) (*ent.Deployment, error) {
	d, err := s.c.Deployment.Create().
		SetNumber(d.Number).
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
		SetUserID(d.UserID).
		SetRepoID(d.RepoID).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	s.c.Repo.
		UpdateOneID(d.RepoID).
		SetLatestDeployedAt(d.CreatedAt).
		Save(ctx)

	return d, nil
}

func (s *Store) UpdateDeployment(ctx context.Context, d *ent.Deployment) (*ent.Deployment, error) {
	d, err := s.c.Deployment.
		UpdateOne(d).
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
	if err != nil {
		return nil, err
	}

	if d.Status != deployment.StatusSuccess {
		return d, nil
	}

	// Update the statistics of deployment by
	// increasing the count.
	err = s.WithTx(ctx, func(tx *ent.Tx) error {
		r, err := tx.Repo.Get(ctx, d.RepoID)
		if err != nil {
			return err
		}

		dc, err := s.c.DeploymentCount.
			Query().
			Where(
				deploymentcount.NamespaceEQ(r.Namespace),
				deploymentcount.NameEQ(r.Name),
				deploymentcount.EnvEQ(d.Env),
			).
			Only(ctx)
		if ent.IsNotFound(err) {
			s.c.DeploymentCount.
				Create().
				SetNamespace(r.Namespace).
				SetName(r.Name).
				SetEnv(d.Env).
				Save(ctx)
			return nil
		} else if err != nil {
			return err
		}

		s.c.DeploymentCount.
			UpdateOne(dc).
			SetCount(dc.Count + 1).
			Save(ctx)
		return nil
	})
	if err != nil {
		zap.L().Error("It has failed to increase the deployment count.", zap.Error(err))
	}

	return d, nil
}
