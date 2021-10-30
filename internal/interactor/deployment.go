package interactor

import (
	"context"
	"fmt"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/approval"
	"github.com/gitploy-io/gitploy/ent/deployment"
	"github.com/gitploy-io/gitploy/pkg/e"
	"github.com/gitploy-io/gitploy/vo"
	"go.uber.org/zap"
)

// Deploy creates a new remote deployment basically.
// But if the approval flag is enabled, it saves the deployment only.
func (i *Interactor) Deploy(ctx context.Context, u *ent.User, r *ent.Repo, d *ent.Deployment, env *vo.Env) (*ent.Deployment, error) {
	if locked, err := i.Store.HasLockOfRepoForEnv(ctx, r, d.Env); locked {
		return nil, e.NewError(
			e.ErrorCodeDeploymentLocked,
			err,
		)
	} else if err != nil {
		return nil, err
	}

	number, err := i.Store.GetNextDeploymentNumberOfRepo(ctx, r)
	if err != nil {
		return nil, e.NewError(
			e.ErrorCodeInternalError,
			err,
		)
	}

	if env.IsApprovalEabled() {
		d := &ent.Deployment{
			Number:                number,
			Type:                  d.Type,
			Env:                   d.Env,
			Ref:                   d.Ref,
			Status:                deployment.StatusWaiting,
			ProductionEnvironment: env.IsProductionEnvironment(),
			IsRollback:            d.IsRollback,
			IsApprovalEnabled:     true,
			RequiredApprovalCount: env.Approval.RequiredCount,
			UserID:                u.ID,
			RepoID:                r.ID,
		}

		i.log.Debug("Save a new deployment to wait approvals.", zap.Any("deployment", d))
		return i.Store.CreateDeployment(ctx, d)
	}

	i.log.Debug("Create a new remote deployment.")
	rd, err := i.createRemoteDeployment(ctx, u, r, d, env)
	if err != nil {
		return nil, err
	}

	d = &ent.Deployment{
		Number:                number,
		Type:                  d.Type,
		Env:                   d.Env,
		Ref:                   d.Ref,
		Status:                deployment.StatusCreated,
		UID:                   rd.UID,
		Sha:                   rd.SHA,
		HTMLURL:               rd.HTLMURL,
		ProductionEnvironment: env.IsProductionEnvironment(),
		IsRollback:            d.IsRollback,
		IsApprovalEnabled:     false,
		RequiredApprovalCount: 0,
		UserID:                u.ID,
		RepoID:                r.ID,
	}

	i.log.Debug("Create a new deployment with the payload.", zap.Any("deployment", d))
	d, err = i.Store.CreateDeployment(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("It failed to save a new deployment.: %w", err)
	}

	i.CreateDeploymentStatus(ctx, &ent.DeploymentStatus{
		Status:       string(deployment.StatusCreated),
		Description:  "Gitploy starts to deploy.",
		DeploymentID: d.ID,
	})

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

// DeployToRemote create a new remote deployment after the deployment was approved.
func (i *Interactor) DeployToRemote(ctx context.Context, u *ent.User, r *ent.Repo, d *ent.Deployment, env *vo.Env) (*ent.Deployment, error) {
	if locked, err := i.Store.HasLockOfRepoForEnv(ctx, r, d.Env); locked {
		return nil, e.NewError(
			e.ErrorCodeDeploymentLocked,
			err,
		)
	} else if err != nil {
		return nil, err
	}

	if d.IsApprovalEnabled && !i.IsApproved(ctx, d) {
		return nil, e.NewError(
			e.ErrorCodeDeploymentUnapproved,
			nil,
		)
	}

	rd, err := i.createRemoteDeployment(ctx, u, r, d, env)
	if err != nil {
		return nil, err
	}

	// Save the state of the remote deployment.
	d.UID = rd.UID
	d.Sha = rd.SHA
	d.HTMLURL = rd.HTLMURL
	d.Status = deployment.StatusCreated

	if d, err = i.UpdateDeployment(ctx, d); err != nil {
		return nil, err
	}

	i.CreateDeploymentStatus(ctx, &ent.DeploymentStatus{
		Status:       string(deployment.StatusCreated),
		Description:  "Gitploy creates a new deployment.",
		DeploymentID: d.ID,
	})

	return d, nil
}

func (i *Interactor) createRemoteDeployment(ctx context.Context, u *ent.User, r *ent.Repo, d *ent.Deployment, env *vo.Env) (*vo.RemoteDeployment, error) {
	// Rollback configures it can deploy the ref without any constraints.
	// 1) Set auto_merge false to avoid the merge conflict.
	// 2) Set required_contexts empty to skip the verfication.
	if d.IsRollback {
		env.AutoMerge = pointer.ToBool(false)
		env.RequiredContexts = &[]string{}
	}

	return i.SCM.CreateRemoteDeployment(ctx, u, r, d, env)
}

func (i *Interactor) runClosingInactiveDeployment(stop <-chan struct{}) {
	ctx := context.Background()

	ticker := time.NewTicker(time.Minute)
L:
	for {
		select {
		case _, ok := <-stop:
			if !ok {
				ticker.Stop()
				break L
			}
		case t := <-ticker.C:
			ds, err := i.ListInactiveDeploymentsLessThanTime(ctx, t.Add(-30*time.Minute).UTC(), 1, 30)
			if err != nil {
				i.log.Error("It has failed to read inactive deployments.", zap.Error(err))
				continue
			}

			for _, d := range ds {
				// Change the status of the deployment canceled, and also
				// cancel the remote deployment if it has.
				if d.Status == deployment.StatusCreated {
					if d.Edges.User != nil && d.Edges.Repo != nil {
						s := &ent.DeploymentStatus{
							Status:       "canceled",
							Description:  "Gitploy cancels the inactive deployment.",
							DeploymentID: d.ID,
						}
						if err := i.SCM.CancelDeployment(ctx, d.Edges.User, d.Edges.Repo, d, s); err != nil {
							i.log.Error("It has failed to cancel the remote deployment.", zap.Error(err))
							continue
						}

						if _, err := i.Store.CreateDeploymentStatus(ctx, s); err != nil {
							i.log.Error("It has failed to create a new deployment status.", zap.Error(err))
							continue
						}
					}
				}

				d.Status = deployment.StatusCanceled
				if _, err := i.UpdateDeployment(ctx, d); err != nil {
					i.log.Error("It has failed to update the deployment canceled.", zap.Error(err))
				}

				i.log.Debug("Cancel the inactive deployment.", zap.Int("deployment_id", d.ID))
			}
		}
	}
}
