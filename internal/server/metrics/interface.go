package metrics

import (
	"context"
	"time"

	"github.com/gitploy-io/gitploy/ent"
)

type (
	Interactor interface {
		ListAllDeploymentCounts(ctx context.Context) ([]*ent.DeploymentCount, error)
		ListDeploymentCountsGreaterThanTime(ctx context.Context, updated time.Time) ([]*ent.DeploymentCount, error)
	}
)
