package interactor

import (
	"context"
	"fmt"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/deployment"
	errs "github.com/hanjunlee/gitploy/internal/errors"
	"github.com/hanjunlee/gitploy/vo"
)

func (i *Interactor) Deploy(ctx context.Context, u *ent.User, re *ent.Repo, d *ent.Deployment) (*ent.Deployment, error) {
	c, err := i.GetConfig(ctx, u, re)
	if err != nil {
		return nil, err
	}

	if !c.HasEnv(d.Env) {
		return nil, &errs.EnvNotFoundError{
			RepoName: re.Name,
		}
	}

	env := c.GetEnv(d.Env)

	d, err = i.Store.CreateDeployment(ctx, u, re, d)
	if err != nil {
		return nil, fmt.Errorf("failed to create a new deployment on the store: %w", err)
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

func (i *Interactor) GetConfig(ctx context.Context, u *ent.User, r *ent.Repo) (*vo.Config, error) {
	return i.GetConfig(ctx, u, r)
}
