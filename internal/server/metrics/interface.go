package metrics

import (
	"context"
	"time"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/vo"
)

type (
	Interactor interface {
		ListAllDeploymentStatisticss(ctx context.Context) ([]*ent.DeploymentStatistics, error)
		ListDeploymentStatisticssGreaterThanTime(ctx context.Context, updated time.Time) ([]*ent.DeploymentStatistics, error)
		GetLicense(ctx context.Context) (*vo.License, error)
	}
)
