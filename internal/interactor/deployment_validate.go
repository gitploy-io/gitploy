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

// DeploymentValidator validate that it is deployable.
type DeploymentValidator struct {
	validators []Validator
}

func NewDeploymentValidator(validators []Validator) *DeploymentValidator {
	return &DeploymentValidator{
		validators: validators,
	}
}

func (v *DeploymentValidator) Validate(d *ent.Deployment) error {
	for _, v := range v.validators {
		if err := v.Validate(d); err != nil {
			return err
		}
	}

	return nil
}

// Validator defines the method that validate a deployment.
type Validator interface {
	Validate(d *ent.Deployment) error
}

// RefValidator validate that the 'ref' is matched with the 'deployable_ref' pattern.
type RefValidator struct {
	Env *extent.Env
}

func (v *RefValidator) Validate(d *ent.Deployment) error {
	ok, err := v.Env.IsDeployableRef(d.Ref)
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

// FrozenWindowValidator validate that the time is in the frozen window.
type FrozenWindowValidator struct {
	Env *extent.Env
}

func (v *FrozenWindowValidator) Validate(d *ent.Deployment) error {
	ok, err := v.Env.IsFreezed(time.Now().UTC())
	if err != nil {
		return err
	}

	if ok {
		return e.NewError(e.ErrorCodeDeploymentFrozen, nil)
	}

	return nil
}

// StatusValidator validate the deployment status is valid.
type StatusValidator struct {
	Status deployment.Status
}

func (v *StatusValidator) Validate(d *ent.Deployment) error {
	if d.Status != v.Status {
		return e.NewErrorWithMessage(
			e.ErrorCodeDeploymentStatusInvalid,
			"The deployment status is not waiting.",
			nil,
		)
	}

	return nil
}

// LockValidator validate that the environment of the repository is locked.
type LockValidator struct {
	Repo  *ent.Repo
	Store LockStore
}

func (v *LockValidator) Validate(d *ent.Deployment) error {
	locked, err := v.Store.HasLockOfRepoForEnv(context.Background(), v.Repo, d.Env)
	if err != nil {
		return err
	}

	if locked {
		return e.NewError(e.ErrorCodeDeploymentLocked, err)
	}

	return nil
}

// ReviewValidator verifies the request is approved or not.
// If one of the  reviews has approve the status is approved.
type ReviewValidator struct {
	Store ReviewStore
}

func (v *ReviewValidator) Validate(d *ent.Deployment) error {
	reviews, err := v.Store.ListReviews(context.Background(), d)
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
