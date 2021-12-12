package metrics

import (
	"context"
	"time"

	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/extent"
)

type (
	Interactor interface {
		CountActiveRepos(ctx context.Context) (int, error)
		CountRepos(ctx context.Context) (int, error)
		ListAllDeploymentStatistics(ctx context.Context) ([]*ent.DeploymentStatistics, error)
		ListDeploymentStatisticsGreaterThanTime(ctx context.Context, updated time.Time) ([]*ent.DeploymentStatistics, error)
		GetLicense(ctx context.Context) (*extent.License, error)
	}
)
