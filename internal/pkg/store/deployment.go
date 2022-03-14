package store

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	i "github.com/gitploy-io/gitploy/internal/interactor"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/ent/deployment"
	"github.com/gitploy-io/gitploy/model/ent/perm"
	"github.com/gitploy-io/gitploy/model/ent/predicate"
	"github.com/gitploy-io/gitploy/pkg/e"
)

func (s *Store) CountDeployments(ctx context.Context) (int, error) {
	cnt, err := s.c.Deployment.
		Query().
		Count(ctx)
	if err != nil {
		return 0, e.NewError(e.ErrorCodeInternalError, err)
	}

	return cnt, nil
}

func (s *Store) SearchDeploymentsOfUser(ctx context.Context, u *ent.User, opt *i.SearchDeploymentsOfUserOptions) ([]*ent.Deployment, error) {
	statusIn := func(ss []deployment.Status) predicate.Deployment {
		if len(ss) == 0 {
			// if not status were provided,
			// it always make this predicate truly.
			return func(s *sql.Selector) {}
		}

		return deployment.StatusIn(ss...)
	}

	// Search deployments that were triggered by the user.
	if opt.Owned {
		return s.c.Deployment.
			Query().
			Where(
				deployment.And(
					deployment.UserIDEQ(u.ID),
					statusIn(opt.Statuses),
					deployment.CreatedAtGTE(opt.From.UTC()),
					deployment.CreatedAtLT(opt.To.UTC()),
				),
			).
			Order(ent.Desc(deployment.FieldCreatedAt)).
			Offset(offset(opt.Page, opt.PerPage)).
			Limit(opt.PerPage).
			WithRepo().
			WithUser().
			All(ctx)
	}

	return s.c.Deployment.
		Query().
		Where(func(s *sql.Selector) {
			p := sql.Table(perm.Table)

			s.Join(p).OnP(
				sql.And(
					sql.ColumnsEQ(p.C(perm.FieldRepoID), s.C(deployment.FieldRepoID)),
					sql.EQ(p.C(perm.FieldUserID), u.ID),
				),
			)
		}).
		Where(
			deployment.And(
				statusIn(opt.Statuses),
				deployment.CreatedAtGTE(opt.From.UTC()),
				deployment.CreatedAtLT(opt.To.UTC()),
			),
		).
		Order(ent.Desc(deployment.FieldCreatedAt)).
		Offset(offset(opt.Page, opt.PerPage)).
		Limit(opt.PerPage).
		WithRepo().
		WithUser().
		All(ctx)
}

func (s *Store) ListInactiveDeploymentsLessThanTime(ctx context.Context, opt *i.ListInactiveDeploymentsLessThanTimeOptions) ([]*ent.Deployment, error) {
	return s.c.Deployment.
		Query().
		Where(
			deployment.And(
				deployment.StatusIn(deployment.StatusWaiting, deployment.StatusCreated),
				deployment.CreatedAtLT(opt.Less),
			),
		).
		WithRepo().
		WithUser().
		Limit(opt.PerPage).
		Offset(offset(opt.Page, opt.PerPage)).
		All(ctx)
}

func (s *Store) ListDeploymentsOfRepo(ctx context.Context, r *ent.Repo, opt *i.ListDeploymentsOfRepoOptions) ([]*ent.Deployment, error) {
	q := s.c.Deployment.
		Query()

	q = q.Where(
		deployment.RepoIDEQ(r.ID),
	)

	if env := opt.Env; env != "" {
		q = q.Where(
			deployment.EnvEQ(env),
		)
	}

	if status := opt.Status; status != "" {
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
		Limit(opt.PerPage).
		Offset(offset(opt.Page, opt.PerPage)).
		All(ctx)
}

func (s *Store) FindDeploymentByID(ctx context.Context, id int) (*ent.Deployment, error) {
	d, err := s.c.Deployment.
		Query().
		Where(
			deployment.IDEQ(id),
		).
		WithRepo().
		WithUser().
		WithDeploymentStatuses().
		Only(ctx)
	if ent.IsNotFound(err) {
		return nil, e.NewErrorWithMessage(e.ErrorCodeEntityNotFound, "The deployment is not found.", err)
	} else if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
	}

	return d, nil
}

func (s *Store) FindDeploymentOfRepoByNumber(ctx context.Context, r *ent.Repo, number int) (*ent.Deployment, error) {
	d, err := s.c.Deployment.
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
		Only(ctx)
	if ent.IsNotFound(err) {
		return nil, e.NewErrorWithMessage(e.ErrorCodeEntityNotFound, "The deployment is not found.", err)
	} else if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
	}

	return d, nil
}

func (s *Store) FindDeploymentByUID(ctx context.Context, uid int64) (*ent.Deployment, error) {
	d, err := s.c.Deployment.
		Query().
		Where(
			deployment.UIDEQ(uid),
		).
		WithRepo().
		WithUser().
		WithDeploymentStatuses().
		Only(ctx)
	if ent.IsNotFound(err) {
		return nil, e.NewErrorWithMessage(e.ErrorCodeEntityNotFound, "The deployment is not found.", err)
	} else if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
	}

	return d, nil
}

func (s *Store) GetNextDeploymentNumberOfRepo(ctx context.Context, r *ent.Repo) (int, error) {
	var v []struct {
		RepoID int64 `json:"repo_id"`
		Max    int   `json:"max"`
	}
	err := s.c.Deployment.Query().
		Where(deployment.RepoID(r.ID)).
		GroupBy(deployment.FieldRepoID).
		Aggregate(ent.Max(deployment.FieldNumber)).
		Scan(ctx, &v)
	if err != nil {
		return 0, e.NewError(e.ErrorCodeInternalError, err)
	}

	if len(v) >= 2 {
		return 0, e.NewErrorWithMessage(e.ErrorCodeInternalError, "The result format of query is invalid.", err)
	}

	// The return value must be one when there is no deployment record.
	if len(v) == 0 {
		return 1, nil
	}

	return v[0].Max + 1, nil
}

// FindPrevRunningDeployment find a deployment of which the status is created, queued, or running.
func (s *Store) FindPrevRunningDeployment(ctx context.Context, d *ent.Deployment) (*ent.Deployment, error) {
	d, err := s.c.Deployment.
		Query().
		Where(
			deployment.And(
				deployment.RepoIDEQ(d.RepoID),
				deployment.EnvEQ(d.Env),
				deployment.StatusIn(
					deployment.StatusCreated,
					deployment.StatusQueued,
					deployment.StatusRunning,
				),
			),
		).
		Order(ent.Desc(deployment.FieldCreatedAt)).
		First(ctx)
	if ent.IsNotFound(err) {
		return nil, e.NewErrorWithMessage(e.ErrorCodeEntityNotFound, "The deployment is not found.", err)
	} else if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
	}

	return d, nil
}

func (s *Store) FindPrevSuccessDeployment(ctx context.Context, d *ent.Deployment) (*ent.Deployment, error) {
	d, err := s.c.Deployment.
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
	if ent.IsNotFound(err) {
		return nil, e.NewErrorWithMessage(e.ErrorCodeEntityNotFound, "The deployment is not found.", err)
	} else if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
	}

	return d, nil
}

// CreateDeployment create a new deployment, and
// it updates the 'latest_deployed_at' field of the repository.
func (s *Store) CreateDeployment(ctx context.Context, d *ent.Deployment) (*ent.Deployment, error) {
	// TODO: Group by a transaction
	d, err := s.c.Deployment.Create().
		SetNumber(d.Number).
		SetType(d.Type).
		SetRef(d.Ref).
		SetEnv(d.Env).
		SetDynamicPayload(d.DynamicPayload).
		SetUID(d.UID).
		SetSha(d.Sha).
		SetHTMLURL(d.HTMLURL).
		SetProductionEnvironment(d.ProductionEnvironment).
		SetIsRollback(d.IsRollback).
		SetStatus(d.Status).
		SetUserID(d.UserID).
		SetRepoID(d.RepoID).
		Save(ctx)
	if ent.IsConstraintError(err) {
		return nil, e.NewError(e.ErrorCodeDeploymentConflict, err)
	} else if ent.IsValidationError(err) {
		return nil, e.NewErrorWithMessage(e.ErrorCodeEntityUnprocessable, fmt.Sprintf("Failed to create a deployment. The value of \"%s\" field is invalid.", err.(*ent.ValidationError).Name), err)
	} else if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
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
		SetStatus(d.Status).
		Save(ctx)
	if ent.IsValidationError(err) {
		return nil, e.NewErrorWithMessage(
			e.ErrorCodeEntityUnprocessable,
			fmt.Sprintf("Failed to update a deployment. The value of \"%s\" field is invalid.", err.(*ent.ValidationError).Name),
			err)
	} else if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
	}

	return d, nil
}
