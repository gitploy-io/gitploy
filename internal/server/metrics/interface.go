package metrics

import (
	"context"
	"time"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/vo"
)

type (
	Interactor interface {
		ListAllDeploymentCounts(ctx context.Context) ([]*ent.DeploymentCount, error)
		ListDeploymentCountsGreaterThanTime(ctx context.Context, updated time.Time) ([]*ent.DeploymentCount, error)
		GetLicense(ctx context.Context) (*vo.License, error)
	}
)
