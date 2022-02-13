package interactor

import (
	"context"
	"fmt"
	"time"

	"github.com/gitploy-io/gitploy/model/ent"
)

type (
	// DeploymentStatisticsInteractor provides application logic for interacting with deployment_statistics.
	DeploymentStatisticsInteractor service

	// DeploymentStatisticsStore defines operations for working with deployment_statistics.
	DeploymentStatisticsStore interface {
		ListAllDeploymentStatistics(ctx context.Context) ([]*ent.DeploymentStatistics, error)
		ListDeploymentStatisticsGreaterThanTime(ctx context.Context, updated time.Time) ([]*ent.DeploymentStatistics, error)
		FindDeploymentStatisticsOfRepoByEnv(ctx context.Context, r *ent.Repo, env string) (*ent.DeploymentStatistics, error)
		CreateDeploymentStatistics(ctx context.Context, s *ent.DeploymentStatistics) (*ent.DeploymentStatistics, error)
		UpdateDeploymentStatistics(ctx context.Context, s *ent.DeploymentStatistics) (*ent.DeploymentStatistics, error)
	}
)

func (i *DeploymentStatisticsInteractor) ProduceDeploymentStatisticsOfRepo(ctx context.Context, r *ent.Repo, d *ent.Deployment) (*ent.DeploymentStatistics, error) {
	s, err := i.store.FindDeploymentStatisticsOfRepoByEnv(ctx, r, d.Env)

	if ent.IsNotFound(err) {
		if s, err = i.store.CreateDeploymentStatistics(ctx, &ent.DeploymentStatistics{
			Env:    d.Env,
			RepoID: r.ID,
		}); err != nil {
			return nil, err
		}
	}

	if s, err = i.produceDeploymentStatisticsOfRepo(ctx, r, d, s); err != nil {
		return nil, err
	}

	return i.store.UpdateDeploymentStatistics(ctx, s)
}

func (i *DeploymentStatisticsInteractor) produceDeploymentStatisticsOfRepo(ctx context.Context, r *ent.Repo, d *ent.Deployment, s *ent.DeploymentStatistics) (*ent.DeploymentStatistics, error) {
	{
		if d.IsRollback {
			s.RollbackCount = s.RollbackCount + 1
		} else {
			s.Count = s.Count + 1
		}
	}

	{
		ld, err := i.store.FindPrevSuccessDeployment(ctx, d)
		if ent.IsNotFound(err) {
			return s, nil
		} else if err != nil {
			return nil, err
		}

		if d.Edges.User == nil {
			if d, err = i.store.FindDeploymentByID(ctx, d.ID); err != nil {
				return nil, err
			} else if d.Edges.User == nil {
				return nil, fmt.Errorf("The deployer is not found.")
			}
		}

		cms, fs, err := i.scm.CompareCommits(ctx, d.Edges.User, r, ld.Sha, d.Sha, 1, 100)
		if err != nil {
			return nil, err
		}

		s.CommitCount = s.CommitCount + len(cms)

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
