package interactor

import (
	"context"
	"fmt"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/deployment"
	"github.com/hanjunlee/gitploy/vo"
)

func (i *Interactor) Deploy(ctx context.Context, u *ent.User, re *ent.Repo, d *ent.Deployment, env *vo.Env) (*ent.Deployment, error) {
	d, err := i.Store.CreateDeployment(ctx, u, re, d)
	if err != nil {
		return nil, fmt.Errorf("failed to save a new deployment to the store: %w", err)
	}

	return i.deployToSCM(ctx, u, re, d, env)
}

func (i *Interactor) deployToSCM(ctx context.Context, u *ent.User, re *ent.Repo, od *ent.Deployment, e *vo.Env) (*ent.Deployment, error) {
	if !e.HasApproval() {
		nd, err := i.SCM.CreateDeployment(ctx, u, re, od, e)
		if err != nil {
			od.Status = deployment.StatusFailure
			i.UpdateDeployment(ctx, od)
			return nil, err
		}

		nd.Status = deployment.StatusCreated
		return i.UpdateDeployment(ctx, nd)
	}

	// TODO: handling approval.
	return nil, fmt.Errorf("Not implemented yet.")
}
