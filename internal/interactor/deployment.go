package interactor

import (
	"context"
	"fmt"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/approval"
	"github.com/gitploy-io/gitploy/ent/deployment"
	"github.com/gitploy-io/gitploy/vo"
	"go.uber.org/zap"
)

func (i *Interactor) Deploy(ctx context.Context, u *ent.User, r *ent.Repo, d *ent.Deployment, e *vo.Env) (*ent.Deployment, error) {
	d.ProductionEnvironment = e.IsProductionEnvironment()
	d.UserID = u.ID
	d.RepoID = r.ID

	if !e.IsApprovalEabled() {
		rd, err := i.SCM.CreateRemoteDeployment(ctx, u, r, d, e)
		if err != nil {
			return nil, err
		}

		// Save the state of the remote deployment.
		d.UID = rd.UID
		d.Sha = rd.SHA
		d.HTMLURL = rd.HTLMURL
		d.Status = deployment.StatusCreated

		d, err := i.Store.CreateDeployment(ctx, d)
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

	d.IsApprovalEnabled = true
	d.RequiredApprovalCount = e.Approval.RequiredCount
	d.Status = deployment.StatusWaiting

	return i.Store.CreateDeployment(ctx, d)
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

func (i *Interactor) CreateRemoteDeployment(ctx context.Context, u *ent.User, re *ent.Repo, d *ent.Deployment, e *vo.Env) (*ent.Deployment, error) {
	// Rollback configures it can deploy the ref without any constraints.
	// 1) Set auto_merge false to avoid the merge conflict.
	// 2) Set required_contexts empty to skip the verfication.
	if d.IsRollback {
		e.AutoMerge = pointer.ToBool(false)
		e.RequiredContexts = &[]string{}
	}

	rd, err := i.SCM.CreateRemoteDeployment(ctx, u, re, d, e)
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
