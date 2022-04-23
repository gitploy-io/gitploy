package store

import (
	"context"
	"fmt"

	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/ent/deploymentstatus"
	"github.com/gitploy-io/gitploy/pkg/e"
)

func (s *Store) ListDeploymentStatuses(ctx context.Context, d *ent.Deployment) ([]*ent.DeploymentStatus, error) {
	dss, err := s.c.DeploymentStatus.
		Query().
		Where(deploymentstatus.DeploymentIDEQ(d.ID)).
		All(ctx)
	if err != nil {
		return nil, e.NewErrorWithMessage(e.ErrorCodeInternalError, "Failed to list deployment statuses.", err)
	}

	return dss, nil
}

func (s *Store) FindDeploymentStatusByID(ctx context.Context, id int) (*ent.DeploymentStatus, error) {
	ds, err := s.c.DeploymentStatus.Query().
		Where(deploymentstatus.IDEQ(id)).
		WithDeployment().
		WithRepo().
		Only(ctx)
	if ent.IsNotFound(err) {
		return nil, e.NewError(e.ErrorCodeEntityNotFound, err)
	}

	return ds, nil
}

func (s *Store) CreateEntDeploymentStatus(ctx context.Context, ds *ent.DeploymentStatus) (*ent.DeploymentStatus, error) {
	// Build the query creating a deployment status.
	qry := s.c.DeploymentStatus.Create().
		SetStatus(ds.Status).
		SetDescription(ds.Description).
		SetLogURL(ds.LogURL).
		SetDeploymentID(ds.DeploymentID).
		SetRepoID(ds.RepoID)

	if !ds.CreatedAt.IsZero() {
		qry.SetCreatedAt(ds.CreatedAt.UTC())
	}

	if !ds.UpdatedAt.IsZero() {
		qry.SetUpdatedAt(ds.UpdatedAt.UTC())
	}

	ret, err := qry.Save(ctx)
	if ent.IsConstraintError(err) {
		return nil, e.NewErrorWithMessage(
			e.ErrorCodeEntityUnprocessable,
			fmt.Sprintf("Failed to create a deployment status. The value of \"%s\" field is invalid.", err.(*ent.ValidationError).Name),
			err)
	} else if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
	}

	return ret, nil
}
