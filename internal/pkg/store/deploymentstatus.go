package store

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
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
