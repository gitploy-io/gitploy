package store

import (
	"context"
	"time"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/deploymentstatistics"
)

func (s *Store) ListAllDeploymentStatistics(ctx context.Context) ([]*ent.DeploymentStatistics, error) {
	return s.c.DeploymentStatistics.
		Query().
		All(ctx)
}

func (s *Store) ListDeploymentStatisticsGreaterThanTime(ctx context.Context, updated time.Time) ([]*ent.DeploymentStatistics, error) {
	return s.c.DeploymentStatistics.
		Query().
		Where(
			deploymentstatistics.UpdatedAtGT(updated),
		).
		All(ctx)
}
