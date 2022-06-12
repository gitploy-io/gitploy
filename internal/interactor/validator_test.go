package interactor_test

import (
	"testing"

	"github.com/AlekSi/pointer"
	"github.com/golang/mock/gomock"

	i "github.com/gitploy-io/gitploy/internal/interactor"
	"github.com/gitploy-io/gitploy/internal/interactor/mock"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/ent/deployment"
	"github.com/gitploy-io/gitploy/model/ent/review"
	"github.com/gitploy-io/gitploy/model/extent"
	"github.com/gitploy-io/gitploy/pkg/e"
)

func TestRefValidator_Validate(t *testing.T) {
	t.Run("Returns an entity_unprocessable error when the ref is invalid", func(t *testing.T) {
		v := &i.RefValidator{
			Env: &extent.Env{DeployableRef: pointer.ToString("main")},
		}

		if err := v.Validate(&ent.Deployment{Ref: "branch"}); !e.HasErrorCode(err, e.ErrorCodeEntityUnprocessable) {
			t.Fatalf("Error is not entity_unprocessable: %v", err)
		}
	})
}

func TestStatusValidator_Validate(t *testing.T) {
	t.Run("Returns an deployment_status_invalid error when the status is not equal.", func(t *testing.T) {
		v := &i.StatusValidator{
			Status: deployment.StatusWaiting,
		}

		if err := v.Validate(&ent.Deployment{Status: deployment.StatusCreated}); !e.HasErrorCode(err, e.ErrorCodeDeploymentStatusInvalid) {
			t.Fatalf("Error is not deployment_status_invalid: %v", err)
		}
	})
}

func TestLockValidator_Validate(t *testing.T) {
	t.Run("Returns an deployment_locked error when the env is locked.", func(t *testing.T) {
		t.Log("Start mocking:")
		ctrl := gomock.NewController(t)
		store := mock.NewMockStore(ctrl)

		store.EXPECT().
			HasLockOfRepoForEnv(gomock.Any(), gomock.AssignableToTypeOf(&ent.Repo{}), gomock.AssignableToTypeOf("environment")).
			Return(true, nil)

		v := &i.LockValidator{
			Store: store,
		}

		if err := v.Validate(&ent.Deployment{Status: deployment.StatusCreated}); !e.HasErrorCode(err, e.ErrorCodeDeploymentLocked) {
			t.Fatalf("Error is not deployment_status_invalid: %v", err)
		}
	})
}

func TestReviewValidator_Validate(t *testing.T) {
	t.Run("Returns an deployment_not_approved error when a review is pending.", func(t *testing.T) {
		t.Log("Start mocking:")
		ctrl := gomock.NewController(t)
		store := mock.NewMockStore(ctrl)

		t.Log("Return a pending review.")
		store.EXPECT().
			ListReviews(gomock.Any(), gomock.AssignableToTypeOf(&ent.Deployment{})).
			Return([]*ent.Review{
				{Status: review.StatusPending},
			}, nil)

		v := &i.ReviewValidator{
			Store: store,
		}

		if err := v.Validate(&ent.Deployment{}); !e.HasErrorCode(err, e.ErrorCodeDeploymentNotApproved) {
			t.Fatalf("Error is not deployment_status_invalid: %v", err)
		}
	})

	t.Run("Returns an deployment_not_approved error when a review is rejected.", func(t *testing.T) {
		t.Log("Start mocking:")
		ctrl := gomock.NewController(t)
		store := mock.NewMockStore(ctrl)

		t.Log("Return a pending review.")
		store.EXPECT().
			ListReviews(gomock.Any(), gomock.AssignableToTypeOf(&ent.Deployment{})).
			Return([]*ent.Review{
				{Status: review.StatusRejected},
			}, nil)

		v := &i.ReviewValidator{
			Store: store,
		}

		if err := v.Validate(&ent.Deployment{}); !e.HasErrorCode(err, e.ErrorCodeDeploymentNotApproved) {
			t.Fatalf("Error is not deployment_status_invalid: %v", err)
		}
	})
}

func TestSerializationValidator_Validate(t *testing.T) {
	t.Run("Returns nil if the serialization is empty.", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		store := mock.NewMockStore(ctrl)

		v := &i.SerializationValidator{
			Env:   &extent.Env{},
			Store: store,
		}

		if err := v.Validate(&ent.Deployment{}); err != nil {
			t.Fatalf("Validate returns an error: %v", err)
		}
	})

	t.Run("Returns an deployment_serialization error if there is a running deployment.", func(t *testing.T) {
		t.Log("Start mocking:")
		ctrl := gomock.NewController(t)
		store := mock.NewMockStore(ctrl)

		t.Log("Return a running deployment.")
		store.EXPECT().
			FindPrevRunningDeployment(gomock.Any(), gomock.AssignableToTypeOf(&ent.Deployment{})).
			Return(&ent.Deployment{}, nil)

		v := &i.SerializationValidator{
			Env:   &extent.Env{Serialization: pointer.ToBool(true)},
			Store: store,
		}

		if err := v.Validate(&ent.Deployment{}); !e.HasErrorCode(err, e.ErrorCodeDeploymentSerialization) {
			t.Fatalf("Error is not deployment_serialization: %v", err)
		}
	})
}

func TestDynamicPayloadValidator_Validate(t *testing.T) {
	t.Run("Returns nil if it is rollback.", func(t *testing.T) {
		v := &i.DynamicPayloadValidator{Env: &extent.Env{}}

		if err := v.Validate(&ent.Deployment{IsRollback: true}); err != nil {
			t.Fatalf("Validate returns an error: %s", err)
		}
	})
}
