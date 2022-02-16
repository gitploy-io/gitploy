package interactor

import (
	"context"
	"time"

	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/ent/deployment"
	"github.com/gitploy-io/gitploy/model/ent/review"
	"github.com/gitploy-io/gitploy/model/extent"
	"github.com/gitploy-io/gitploy/pkg/e"
)

// deploymentValidator validate that it is deployable.
type deploymentValidator struct {
	validators []validator
}

func newDeploymentValidator(validators []validator) *deploymentValidator {
	return &deploymentValidator{
		validators: validators,
	}
}

func (v *deploymentValidator) Validate(d *ent.Deployment) error {
	for _, v := range v.validators {
		if err := v.Validate(d); err != nil {
			return err
		}
	}

	return nil
}

// validator defines the method that validate a deployment.
type validator interface {
	Validate(d *ent.Deployment) error
}

// refValidator validate that the 'ref' is matched with the 'deployable_ref' pattern.
type refValidator struct {
	env *extent.Env
}

func (v *refValidator) Validate(d *ent.Deployment) error {
	ok, err := v.env.IsDeployableRef(d.Ref)
	if err != nil {
		return err
	}
	if !ok {
		return e.NewErrorWithMessage(
			e.ErrorCodeEntityUnprocessable,
			"The ref is not matched with the 'deployable_ref' pattern.",
			nil)
	}

	return nil
}

// frozenWindowValidator validate that the time is in the frozen window.
type frozenWindowValidator struct {
	env *extent.Env
}

func (v *frozenWindowValidator) Validate(d *ent.Deployment) error {
	ok, err := v.env.IsFreezed(time.Now().UTC())
	if err != nil {
		return err
	}

	if ok {
		return e.NewError(e.ErrorCodeDeploymentFrozen, nil)
	}

	return nil
}

// statusValidator validate the deployment status is valid.
type statusValidator struct {
	status deployment.Status
}

func (v *statusValidator) Validate(d *ent.Deployment) error {
	if d.Status != v.status {
		return e.NewErrorWithMessage(
			e.ErrorCodeDeploymentStatusInvalid,
			"The deployment status is not waiting.",
			nil,
		)
	}

	return nil
}

// lockValidator validate that the environment of the repository is locked.
type lockValidator struct {
	repo  *ent.Repo
	store LockStore
}

func (v *lockValidator) Validate(d *ent.Deployment) error {
	locked, err := v.store.HasLockOfRepoForEnv(context.Background(), v.repo, d.Env)
	if err != nil {
		return err
	}

	if locked {
		return e.NewError(e.ErrorCodeDeploymentLocked, err)
	}

	return nil
}

// reviewValidator verifies the request is approved or not.
// If one of the  reviews has approve the status is approved.
type reviewValidator struct {
	store ReviewStore
}

func (v *reviewValidator) Validate(d *ent.Deployment) error {
	reviews, err := v.store.ListReviews(context.Background(), d)
	if err != nil {
		return err
	}

	for _, r := range reviews {
		if r.Status == review.StatusRejected {
			return e.NewError(e.ErrorCodeDeploymentNotApproved, nil)
		}
	}

	for _, r := range reviews {
		if r.Status == review.StatusApproved {
			return nil
		}
	}

	return e.NewError(e.ErrorCodeDeploymentNotApproved, nil)
}
