//go:generate mockgen -source ./interface.go -destination ./mock/interactor.go -package mock

package hooks

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/notification"
)

type (
	Interactor interface {
		FindDeploymentByUID(ctx context.Context, uid int64) (*ent.Deployment, error)
		CreateDeploymentStatus(ctx context.Context, s *ent.DeploymentStatus) (*ent.DeploymentStatus, error)
		UpdateDeployment(ctx context.Context, d *ent.Deployment) (*ent.Deployment, error)
		Publish(ctx context.Context, typ notification.Type, r *ent.Repo, d *ent.Deployment, a *ent.Approval) error
	}
)
