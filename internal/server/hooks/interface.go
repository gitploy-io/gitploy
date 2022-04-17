//go:generate mockgen -source ./interface.go -destination ./mock/interactor.go -package mock

package hooks

import (
	"context"

	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/extent"
)

type (
	Interactor interface {
		FindRepoByID(ctx context.Context, id int64) (*ent.Repo, error)
		FindDeploymentByUID(ctx context.Context, uid int64) (*ent.Deployment, error)
		Deploy(ctx context.Context, u *ent.User, r *ent.Repo, d *ent.Deployment, env *extent.Env) (*ent.Deployment, error)
		UpdateDeployment(ctx context.Context, d *ent.Deployment) (*ent.Deployment, error)
		CreateDeploymentStatus(ctx context.Context, ds *ent.DeploymentStatus) (*ent.DeploymentStatus, error)
		ProduceDeploymentStatisticsOfRepo(ctx context.Context, r *ent.Repo, d *ent.Deployment) (*ent.DeploymentStatistics, error)
		GetEvaluatedConfig(ctx context.Context, u *ent.User, r *ent.Repo, v *extent.EvalValues) (*extent.Config, error)
	}
)
