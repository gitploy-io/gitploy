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
		WithRepo().
		All(ctx)
}

func (s *Store) ListDeploymentStatisticsGreaterThanTime(ctx context.Context, updated time.Time) ([]*ent.DeploymentStatistics, error) {
	return s.c.DeploymentStatistics.
		Query().
		Where(
			deploymentstatistics.UpdatedAtGT(updated),
		).
		WithRepo().
		All(ctx)
}

func (s *Store) FindDeploymentStatisticsOfRepoByEnv(ctx context.Context, r *ent.Repo, env string) (*ent.DeploymentStatistics, error) {
	return s.c.DeploymentStatistics.
		Query().
		Where(
			deploymentstatistics.RepoIDEQ(r.ID),
			deploymentstatistics.EnvEQ(env),
		).
		WithRepo().
		Only(ctx)
}

func (s *Store) CreateDeploymentStatistics(ctx context.Context, ds *ent.DeploymentStatistics) (*ent.DeploymentStatistics, error) {
	return s.c.DeploymentStatistics.
		Create().
		SetEnv(ds.Env).
		SetCount(ds.Count).
		SetRepoID(ds.RepoID).
		Save(ctx)
}

func (s *Store) UpdateDeploymentStatistics(ctx context.Context, ds *ent.DeploymentStatistics) (*ent.DeploymentStatistics, error) {
	return s.c.DeploymentStatistics.
		UpdateOne(ds).
		SetCount(ds.Count).
		Save(ctx)
}
