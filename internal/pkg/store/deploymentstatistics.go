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

func (s *Store) FindDeploymentStatisticsOfRepoByEnv(ctx context.Context, r *ent.Repo, env string) (*ent.DeploymentStatistics, error) {
	return s.c.DeploymentStatistics.
		Query().
		Where(
			deploymentstatistics.NamespaceEQ(r.Namespace),
			deploymentstatistics.NameEQ(r.Name),
			deploymentstatistics.EnvEQ(env),
		).
		Only(ctx)
}

func (s *Store) CreateDeploymentStatistics(ctx context.Context, ds *ent.DeploymentStatistics) (*ent.DeploymentStatistics, error) {
	return s.c.DeploymentStatistics.
		Create().
		SetNamespace(ds.Namespace).
		SetName(ds.Name).
		SetEnv(ds.Env).
		SetCount(ds.Count).
		Save(ctx)
}

func (s *Store) UpdateDeploymentStatistics(ctx context.Context, ds *ent.DeploymentStatistics) (*ent.DeploymentStatistics, error) {
	return s.c.DeploymentStatistics.
		UpdateOne(ds).
		SetCount(ds.Count).
		Save(ctx)
}
