package store

import (
	"context"

	"github.com/gitploy-io/gitploy/ent"
)

func (s *Store) CreateDeploymentStatus(ctx context.Context, ds *ent.DeploymentStatus) (*ent.DeploymentStatus, error) {
	return s.c.DeploymentStatus.
		Create().
		SetStatus(ds.Status).
		SetDescription(ds.Description).
		SetLogURL(ds.LogURL).
		SetDeploymentID(ds.DeploymentID).
		Save(ctx)
}

func (s *Store) SyncDeploymentStatus(ctx context.Context, ds *ent.DeploymentStatus) (*ent.DeploymentStatus, error) {
	return s.c.DeploymentStatus.
		Create().
		SetStatus(ds.Status).
		SetDescription(ds.Description).
		SetLogURL(ds.LogURL).
		SetDeploymentID(ds.DeploymentID).
		SetCreatedAt(ds.CreatedAt).
		SetUpdatedAt(ds.UpdatedAt).
		Save(ctx)
}
