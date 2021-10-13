package interactor

import (
	"context"

	"github.com/gitploy-io/gitploy/ent"
)

func (i *Interactor) ProduceDeploymentStatisticsOfRepo(ctx context.Context, r *ent.Repo, d *ent.Deployment) (*ent.DeploymentStatistics, error) {
	if d.IsRollback {
		return i.incrRollbackCount(ctx, r, d)
	}

	s, err := i.Store.FindDeploymentStatisticsOfRepoByEnv(ctx, r, d.Env)

	if ent.IsNotFound(err) {
		return i.Store.CreateDeploymentStatistics(ctx, &ent.DeploymentStatistics{
			Env:           d.Env,
			Count:         1,
			RollbackCount: 0,
			RepoID:        r.ID,
		})
	}

	s.Count = s.Count + 1

	return i.Store.UpdateDeploymentStatistics(ctx, s)
}

func (i *Interactor) incrRollbackCount(ctx context.Context, r *ent.Repo, d *ent.Deployment) (*ent.DeploymentStatistics, error) {
	s, err := i.Store.FindDeploymentStatisticsOfRepoByEnv(ctx, r, d.Env)

	if ent.IsNotFound(err) {
		return i.Store.CreateDeploymentStatistics(ctx, &ent.DeploymentStatistics{
			Env:           d.Env,
			Count:         0,
			RollbackCount: 1,
			RepoID:        r.ID,
		})
	}

	s.RollbackCount = s.RollbackCount + 1

	return i.Store.UpdateDeploymentStatistics(ctx, s)
}
