//go:generate mockgen -source ./interface.go -destination ./mock/interactor.go -package mock

package hooks

import (
	"context"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/extent"
)

type (
	Interactor interface {
		FindRepoByID(ctx context.Context, id int64) (*ent.Repo, error)
		FindDeploymentByUID(ctx context.Context, uid int64) (*ent.Deployment, error)
		SyncDeploymentStatus(ctx context.Context, ds *ent.DeploymentStatus) (*ent.DeploymentStatus, error)
		Deploy(ctx context.Context, u *ent.User, r *ent.Repo, d *ent.Deployment, env *extent.Env) (*ent.Deployment, error)
		UpdateDeployment(ctx context.Context, d *ent.Deployment) (*ent.Deployment, error)
		ProduceDeploymentStatisticsOfRepo(ctx context.Context, r *ent.Repo, d *ent.Deployment) (*ent.DeploymentStatistics, error)
		CreateEvent(ctx context.Context, e *ent.Event) (*ent.Event, error)
		GetConfig(ctx context.Context, u *ent.User, r *ent.Repo) (*extent.Config, error)
	}
)
