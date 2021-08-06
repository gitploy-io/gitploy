package interactor

import (
	"context"
	"fmt"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/approval"
	"github.com/hanjunlee/gitploy/ent/deployment"
	"github.com/hanjunlee/gitploy/vo"
)

func (i *Interactor) Deploy(ctx context.Context, u *ent.User, r *ent.Repo, d *ent.Deployment, env *vo.Env) (*ent.Deployment, error) {
	return i.deploy(ctx, u, r, d, env)
}

func (i *Interactor) Rollback(ctx context.Context, u *ent.User, r *ent.Repo, d *ent.Deployment, env *vo.Env) (*ent.Deployment, error) {
	d.IsRollback = true

	return i.deploy(ctx, u, r, d, env)
}

func (i *Interactor) deploy(ctx context.Context, u *ent.User, r *ent.Repo, d *ent.Deployment, env *vo.Env) (*ent.Deployment, error) {
	d.UserID = u.ID
	d.RepoID = r.ID

	d, err := i.Store.CreateDeployment(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("It failed to save a new deployment to the store: %w", err)
	}

	if !env.IsApprovalEabled() {
		return i.createDeploymentToSCM(ctx, u, r, d, env)
	}

	d.IsApprovalEnabled = true
	d.RequiredApprovalCount = env.Approval.RequiredCount
	d, _ = i.Store.UpdateDeployment(ctx, d)

	return d, nil
}

func (i *Interactor) IsApproved(ctx context.Context, d *ent.Deployment) bool {
	as, _ := i.ListApprovals(ctx, d)

	approved := 0
	for _, a := range as {
		if a.Status == approval.StatusApproved {
			approved = approved + 1
		}
	}

	return approved >= d.RequiredApprovalCount
}

func (i *Interactor) CreateDeploymentToSCM(ctx context.Context, u *ent.User, re *ent.Repo, d *ent.Deployment, env *vo.Env) (*ent.Deployment, error) {
	return i.createDeploymentToSCM(ctx, u, re, d, env)
}

func (i *Interactor) createDeploymentToSCM(ctx context.Context, u *ent.User, re *ent.Repo, d *ent.Deployment, e *vo.Env) (*ent.Deployment, error) {
	if d.IsRollback {
		// Rollback configures it can deploy the ref without any constraints.
		// 1) Set auto_merge false to avoid the merge conflict.
		// 2) Set required_contexts empty to skip the verfication.
		e.Task = "rollback"
		e.AutoMerge = false
		e.RequiredContexts = []string{}
	}

	rd, err := i.SCM.CreateDeployment(ctx, u, re, d, e)
	if err != nil {
		d.Status = deployment.StatusFailure
		if _, err := i.UpdateDeployment(ctx, d); err != nil {
			return nil, err
		}

		i.CreateDeploymentStatus(ctx, &ent.DeploymentStatus{
			Status:       string(deployment.StatusFailure),
			Description:  "Gitploy failed to create a deployment.",
			DeploymentID: d.ID,
		})
		return nil, err
	}

	// Save the state of the remote deployment.
	d.UID = rd.UID
	d.Sha = rd.SHA
	d.HTMLURL = rd.HTLMURL
	d.Status = deployment.StatusCreated

	if _, err := i.UpdateDeployment(ctx, d); err != nil {
		return nil, err
	}

	i.CreateDeploymentStatus(ctx, &ent.DeploymentStatus{
		Status:       string(deployment.StatusCreated),
		Description:  "Gitploy creates a new deployment.",
		DeploymentID: d.ID,
	})
	return d, nil
}
