package store

import (
	"context"
	"time"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/deploymentcount"
)

func (s *Store) ListAllDeploymentCounts(ctx context.Context) ([]*ent.DeploymentCount, error) {
	return s.c.DeploymentCount.
		Query().
		All(ctx)
}

func (s *Store) ListDeploymentCountsGreaterThanTime(ctx context.Context, updated time.Time) ([]*ent.DeploymentCount, error) {
	return s.c.DeploymentCount.
		Query().
		Where(
			deploymentcount.UpdatedAtGT(updated),
		).
		All(ctx)
}
