package interactor

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/deployment"
	"github.com/hanjunlee/gitploy/vo"
)

func (i *Interactor) Deploy(ctx context.Context, u *ent.User, re *ent.Repo, d *ent.Deployment, env *vo.Env) (*ent.Deployment, error) {
	d.UserID = u.ID
	d.RepoID = re.ID

	num, err := i.Store.GetNextDeploymentNumberOfRepo(ctx, d)
	d.Number = num

	d, err = i.Store.CreateDeployment(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("failed to save a new deployment to the store: %w", err)
	}

	if !env.HasApproval() {
		return i.deployToSCM(ctx, u, re, d, env)
	}

	_, err = i.requestApprovals(ctx, d, env.Approval.Approvers)
	if err != nil {
		return nil, err
	}

	d.RequiredApprovalCount = env.Approval.RequiredCount
	d, _ = i.Store.UpdateDeployment(ctx, d)

	// Automatically approved for deployment timeout,
	// it deploys to SCM if a status of the deployment is waiting.
	time.AfterFunc(time.Duration(env.Approval.WaitMinute)*time.Minute, func() {
		ctx := context.Background()
		log := i.log.Named("waiting")

		d, err := i.FindDeploymentWithEdgesByID(ctx, d.ID)
		if err != nil {
			log.Error("failed to find the deployment.", zap.Int("deployment_id", d.ID))
			return
		}

		if d.Status != deployment.StatusWaiting {
			log.Debug("the deployment already processed.")
			return
		}

		if _, err = i.deployToSCM(ctx, u, re, d, env); err != nil {
			log.Error("failed to deploy to SCM.", zap.Error(err))
			return
		}
	})

	return d, nil
}

// TODO: Rollback command.

func (i *Interactor) deployToSCM(ctx context.Context, u *ent.User, re *ent.Repo, od *ent.Deployment, e *vo.Env) (*ent.Deployment, error) {
	nd, err := i.SCM.CreateDeployment(ctx, u, re, od, e)
	if err != nil {
		od.Status = deployment.StatusFailure
		i.UpdateDeployment(ctx, od)
		return nil, err
	}

	nd.Status = deployment.StatusCreated
	return i.UpdateDeployment(ctx, nd)
}

func (i *Interactor) requestApprovals(ctx context.Context, d *ent.Deployment, approvers []string) ([]*ent.Approval, error) {
	approvals := []*ent.Approval{}

	for _, ar := range approvers {
		u, err := i.FindUserByLogin(ctx, ar)
		if ent.IsNotFound(err) {
			continue
		} else if err != nil {
			d.Status = deployment.StatusFailure
			i.UpdateDeployment(ctx, d)
			return nil, fmt.Errorf("failed to get the user: %w", err)
		}

		a, err := i.CreateApproval(ctx, &ent.Approval{
			UserID:       u.ID,
			DeploymentID: d.ID,
		})
		if err != nil {
			d.Status = deployment.StatusFailure
			i.UpdateDeployment(ctx, d)
			return nil, fmt.Errorf("failed to request a approval: %w", err)
		}

		approvals = append(approvals, a)
	}

	return approvals, nil
}
