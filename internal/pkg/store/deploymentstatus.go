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

func (s *Store) CreateDeploymentStatus(ctx context.Context, ds *ent.DeploymentStatus) (*ent.DeploymentStatus, error) {
	ret, err := s.c.DeploymentStatus.
		Create().
		SetStatus(ds.Status).
		SetDescription(ds.Description).
		SetLogURL(ds.LogURL).
		SetDeploymentID(ds.DeploymentID).
		Save(ctx)
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

func (s *Store) SyncDeploymentStatus(ctx context.Context, ds *ent.DeploymentStatus) (*ent.DeploymentStatus, error) {
	ret, err := s.c.DeploymentStatus.
		Create().
		SetStatus(ds.Status).
		SetDescription(ds.Description).
		SetLogURL(ds.LogURL).
		SetDeploymentID(ds.DeploymentID).
		SetCreatedAt(ds.CreatedAt).
		SetUpdatedAt(ds.UpdatedAt).
		Save(ctx)
	if ent.IsConstraintError(err) {
		return nil, e.NewErrorWithMessage(
			e.ErrorCodeEntityUnprocessable,
			fmt.Sprintf("Failed to sync the deployment status. The value of \"%s\" field is invalid.", err.(*ent.ValidationError).Name),
			err)
	} else if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
	}

	return ret, nil
}
