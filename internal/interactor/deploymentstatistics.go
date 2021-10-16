package interactor

import (
	"context"
	"fmt"

	"github.com/gitploy-io/gitploy/ent"
)

func (i *Interactor) ProduceDeploymentStatisticsOfRepo(ctx context.Context, r *ent.Repo, d *ent.Deployment) (*ent.DeploymentStatistics, error) {
	s, err := i.Store.FindDeploymentStatisticsOfRepoByEnv(ctx, r, d.Env)

	if ent.IsNotFound(err) {
		if s, err = i.Store.CreateDeploymentStatistics(ctx, &ent.DeploymentStatistics{
			Env:    d.Env,
			RepoID: r.ID,
		}); err != nil {
			return nil, err
		}
	}

	if s, err = i.produceDeploymentStatisticsOfRepo(ctx, r, d, s); err != nil {
		return nil, err
	}

	return i.Store.UpdateDeploymentStatistics(ctx, s)
}

func (i *Interactor) produceDeploymentStatisticsOfRepo(ctx context.Context, r *ent.Repo, d *ent.Deployment, s *ent.DeploymentStatistics) (*ent.DeploymentStatistics, error) {
	{
		if d.IsRollback {
			s.RollbackCount = s.RollbackCount + 1
		} else {
			s.Count = s.Count + 1
		}
	}

	{
		ld, err := i.Store.FindPrevSuccessDeployment(ctx, d)
		if ent.IsNotFound(err) {
			return s, nil
		} else if err != nil {
			return nil, err
		}

		if d.Edges.User == nil {
			if d, err = i.Store.FindDeploymentByID(ctx, d.ID); err != nil {
				return nil, err
			} else if d.Edges.User == nil {
				return nil, fmt.Errorf("The deployer is not found.")
			}
		}

		cms, fs, err := i.SCM.CompareCommits(ctx, d.Edges.User, r, ld.Sha, d.Sha, 1, 100)
		if err != nil {
			return nil, err
		}

		for _, cm := range cms {
			leadTime := d.UpdatedAt.Sub(cm.Author.Date)
			s.LeadTimeSeconds = s.LeadTimeSeconds + int(leadTime.Seconds())
		}

		// Changes from the latest deployment.
		for _, f := range fs {
			s.Additions = s.Additions + f.Additions
			s.Deletions = s.Deletions + f.Deletions
			s.Changes = s.Changes + f.Changes
		}
	}

	return s, nil
}
