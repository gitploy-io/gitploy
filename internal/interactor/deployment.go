package interactor

import (
	"context"
	"fmt"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/ent/deployment"
	"github.com/gitploy-io/gitploy/model/extent"
	"github.com/gitploy-io/gitploy/pkg/e"
	"go.uber.org/zap"
)

type (
	// DeploymentInteractor provides application logic for interacting with deployments.
	DeploymentInteractor service

	// DeploymentStore defines operations for working with deployments.
	DeploymentStore interface {
		CountDeployments(ctx context.Context) (int, error)
		SearchDeploymentsOfUser(ctx context.Context, u *ent.User, opt *SearchDeploymentsOfUserOptions) ([]*ent.Deployment, error)
		ListInactiveDeploymentsLessThanTime(ctx context.Context, opt *ListInactiveDeploymentsLessThanTimeOptions) ([]*ent.Deployment, error)
		ListDeploymentsOfRepo(ctx context.Context, r *ent.Repo, opt *ListDeploymentsOfRepoOptions) ([]*ent.Deployment, error)
		FindDeploymentByID(ctx context.Context, id int) (*ent.Deployment, error)
		FindDeploymentByUID(ctx context.Context, uid int64) (*ent.Deployment, error)
		FindDeploymentOfRepoByNumber(ctx context.Context, r *ent.Repo, number int) (*ent.Deployment, error)
		FindPrevSuccessDeployment(ctx context.Context, d *ent.Deployment) (*ent.Deployment, error)
		GetNextDeploymentNumberOfRepo(ctx context.Context, r *ent.Repo) (int, error)
		CreateDeployment(ctx context.Context, d *ent.Deployment) (*ent.Deployment, error)
		UpdateDeployment(ctx context.Context, d *ent.Deployment) (*ent.Deployment, error)
	}

	// SearchDeploymentsOfUserOptions specifies the optional parameters that
	// search deployments.
	SearchDeploymentsOfUserOptions struct {
		ListOptions

		Statuses []deployment.Status
		Owned    bool
		From     time.Time
		To       time.Time
	}

	// ListInactiveDeploymentsLessThanTimeOptions specifies the optional parameters that
	// get inactive deployments.
	ListInactiveDeploymentsLessThanTimeOptions struct {
		ListOptions

		Less time.Time
	}

	ListDeploymentsOfRepoOptions struct {
		ListOptions

		Env    string
		Status string
	}

	// DeploymentSCM defines operations for working with remote users.
	DeploymentSCM interface {
		// SCM returns the deployment with UID and SHA.
		CreateRemoteDeployment(ctx context.Context, u *ent.User, r *ent.Repo, d *ent.Deployment, e *extent.Env) (*extent.RemoteDeployment, error)
		CancelDeployment(ctx context.Context, u *ent.User, r *ent.Repo, d *ent.Deployment, s *ent.DeploymentStatus) error

		GetConfig(ctx context.Context, u *ent.User, r *ent.Repo) (*extent.Config, error)
		GetConfigRedirectURL(ctx context.Context, u *ent.User, r *ent.Repo) (string, error)
		GetNewConfigRedirectURL(ctx context.Context, u *ent.User, r *ent.Repo) (string, error)
	}
)

// Deploy posts a new deployment to SCM with the payload.
// But if it requires a review, it saves the payload on the store and waits until reviewed.
// It returns an error for a undeployable payload.
func (i *DeploymentInteractor) Deploy(ctx context.Context, u *ent.User, r *ent.Repo, d *ent.Deployment, env *extent.Env) (*ent.Deployment, error) {
	v := newDeploymentValidator([]validator{
		&refValidator{env: env},
		&frozenWindowValidator{env: env},
		&lockValidator{repo: r, store: i.store},
	})
	if err := v.Validate(d); err != nil {
		return nil, err
	}

	number, err := i.store.GetNextDeploymentNumberOfRepo(ctx, r)
	if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
	}

	d = &ent.Deployment{
		Number:                number,
		Type:                  d.Type,
		Env:                   d.Env,
		Ref:                   d.Ref,
		ProductionEnvironment: env.IsProductionEnvironment(),
		IsRollback:            d.IsRollback,
		UserID:                u.ID,
		RepoID:                r.ID,
	}

	if env.HasReview() {
		d.Status = deployment.StatusWaiting

		i.log.Debug("Save the deployment to wait reviews.")
		d, err = i.store.CreateDeployment(ctx, d)
		if err != nil {
			return nil, err
		}

		for _, rvr := range env.Review.Reviewers {
			if rvr == u.Login {
				continue
			}

			i.log.Debug(fmt.Sprintf("Request a review to %s.", rvr))
			if _, err := i.requestReviewByLogin(ctx, d, rvr); err != nil {
				i.log.Error("Failed to request the review.", zap.Error(err))
			}
		}

		return d, nil
	}

	i.log.Debug("Create a new remote deployment.")
	rd, err := i.createRemoteDeployment(ctx, u, r, d, env)
	if err != nil {
		return nil, err
	}

	d.UID = rd.UID
	d.Sha = rd.SHA
	d.HTMLURL = rd.HTLMURL
	d.Status = deployment.StatusCreated

	i.log.Debug("Create a new deployment with the payload.", zap.Any("deployment", d))
	d, err = i.store.CreateDeployment(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("It failed to save a new deployment.: %w", err)
	}

	i.store.CreateDeploymentStatus(ctx, &ent.DeploymentStatus{
		Status:       string(deployment.StatusCreated),
		Description:  "Gitploy starts to deploy.",
		DeploymentID: d.ID,
	})

	return d, nil
}

// DeployToRemote posts a new deployment to SCM with the saved payload
// after review has finished.
// It returns an error for a undeployable payload.
func (i *DeploymentInteractor) DeployToRemote(ctx context.Context, u *ent.User, r *ent.Repo, d *ent.Deployment, env *extent.Env) (*ent.Deployment, error) {
	v := newDeploymentValidator([]validator{
		&statusValidator{status: deployment.StatusWaiting},
		&refValidator{env: env},
		&frozenWindowValidator{env: env},
		&lockValidator{repo: r, store: i.store},
		&reviewValidator{store: i.store},
	})
	if err := v.Validate(d); err != nil {
		return nil, err
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

	if d, err = i.store.UpdateDeployment(ctx, d); err != nil {
		return nil, err
	}

	i.store.CreateDeploymentStatus(ctx, &ent.DeploymentStatus{
		Status:       string(deployment.StatusCreated),
		Description:  "Gitploy creates a new deployment.",
		DeploymentID: d.ID,
	})

	return d, nil
}

func (i *DeploymentInteractor) createRemoteDeployment(ctx context.Context, u *ent.User, r *ent.Repo, d *ent.Deployment, env *extent.Env) (*extent.RemoteDeployment, error) {
	// Rollback configures it can deploy the ref without any constraints.
	// 1) Set auto_merge false to avoid the merge conflict.
	// 2) Set required_contexts empty to skip the verfication.
	if d.IsRollback {
		env.AutoMerge = pointer.ToBool(false)
		env.RequiredContexts = &[]string{}
	}

	return i.scm.CreateRemoteDeployment(ctx, u, r, d, env)
}

func (i *DeploymentInteractor) runClosingInactiveDeployment(stop <-chan struct{}) {
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
			ds, err := i.store.ListInactiveDeploymentsLessThanTime(ctx, &ListInactiveDeploymentsLessThanTimeOptions{
				ListOptions: ListOptions{Page: 1, PerPage: 30},
				Less:        t.Add(-30 * time.Minute).UTC(),
			})
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
						if err := i.scm.CancelDeployment(ctx, d.Edges.User, d.Edges.Repo, d, s); err != nil {
							i.log.Error("It has failed to cancel the remote deployment.", zap.Error(err))
							continue
						}

						if _, err := i.store.CreateDeploymentStatus(ctx, s); err != nil {
							i.log.Error("It has failed to create a new deployment status.", zap.Error(err))
							continue
						}
					}
				}

				d.Status = deployment.StatusCanceled
				if _, err := i.store.UpdateDeployment(ctx, d); err != nil {
					i.log.Error("It has failed to update the deployment canceled.", zap.Error(err))
				}

				i.log.Debug("Cancel the inactive deployment.", zap.Int("deployment_id", d.ID))
			}
		}
	}
}
