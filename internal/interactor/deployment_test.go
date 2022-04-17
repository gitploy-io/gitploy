package interactor_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"

	i "github.com/gitploy-io/gitploy/internal/interactor"
	"github.com/gitploy-io/gitploy/internal/interactor/mock"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/ent/deployment"
	"github.com/gitploy-io/gitploy/model/ent/review"
	"github.com/gitploy-io/gitploy/model/extent"
)

var (
	ctx = gomock.Any()
)

func TestInteractor_Deploy(t *testing.T) {

	t.Run("Return a new deployment.", func(t *testing.T) {
		t.Log("Start mocking:")
		ctrl := gomock.NewController(t)
		store := mock.NewMockStore(ctrl)
		scm := mock.NewMockSCM(ctrl)

		t.Log("\tValidate the request, and get the next deployment number.")
		store.
			EXPECT().
			HasLockOfRepoForEnv(ctx, gomock.AssignableToTypeOf(&ent.Repo{}), gomock.AssignableToTypeOf("environment")).
			Return(false, nil)

		store.
			EXPECT().
			GetNextDeploymentNumberOfRepo(ctx, gomock.AssignableToTypeOf(&ent.Repo{})).
			Return(1, nil)

		t.Log("\tCreate a new deployment with the response, and create a new deployment status.")
		scm.
			EXPECT().
			CreateRemoteDeployment(ctx, gomock.AssignableToTypeOf(&ent.User{}), gomock.AssignableToTypeOf(&ent.Repo{}), gomock.AssignableToTypeOf(&ent.Deployment{}), gomock.Eq(&extent.Env{})).
			Return(&extent.RemoteDeployment{}, nil)

		store.
			EXPECT().
			CreateDeployment(ctx, gomock.AssignableToTypeOf(&ent.Deployment{})).
			DoAndReturn(func(ctx context.Context, d *ent.Deployment) (*ent.Deployment, error) {
				return d, nil
			})

		store.
			EXPECT().
			CreateEvent(gomock.Any(), gomock.AssignableToTypeOf(&ent.Event{})).
			Return(&ent.Event{}, nil)

		store.
			EXPECT().
			CreateEntDeploymentStatus(ctx, gomock.AssignableToTypeOf(&ent.DeploymentStatus{}))

		it := i.NewInteractor(&i.InteractorConfig{
			Store: store,
			SCM:   scm,
		})

		d, err := it.Deploy(context.Background(), &ent.User{}, &ent.Repo{}, &ent.Deployment{}, &extent.Env{})
		if err != nil {
			t.Fatalf("Deploy returns a error: %s", err)
		}

		if d.Status != deployment.StatusCreated {
			t.Errorf("Deploy.Status = %v, wanted %v", d, deployment.StatusCreated)
		}
	})

	t.Run("Return the waiting deployment and reviews.", func(t *testing.T) {
		t.Log("Start mocking:")
		ctrl := gomock.NewController(t)
		store := mock.NewMockStore(ctrl)
		scm := mock.NewMockSCM(ctrl)

		t.Log("\tValidate the request, and get the next deployment number.")
		store.
			EXPECT().
			HasLockOfRepoForEnv(ctx, gomock.AssignableToTypeOf(&ent.Repo{}), gomock.AssignableToTypeOf("environment")).
			Return(false, nil)

		store.
			EXPECT().
			GetNextDeploymentNumberOfRepo(ctx, gomock.AssignableToTypeOf(&ent.Repo{})).
			Return(1, nil)

		t.Log("\tCreate a new deployment and request reviews.")
		store.
			EXPECT().
			CreateDeployment(ctx, gomock.AssignableToTypeOf(&ent.Deployment{})).
			DoAndReturn(func(ctx context.Context, d *ent.Deployment) (interface{}, interface{}) {
				return d, nil
			})

		store.
			EXPECT().
			CreateEvent(gomock.Any(), gomock.AssignableToTypeOf(&ent.Event{})).
			Return(&ent.Event{}, nil)

		it := i.NewInteractor(&i.InteractorConfig{
			Store: store,
			SCM:   scm,
		})

		d, err := it.Deploy(context.Background(), &ent.User{}, &ent.Repo{}, &ent.Deployment{}, &extent.Env{
			Review: &extent.Review{
				Enabled: true,
			},
		})
		if err != nil {
			t.Errorf("Deploy returns a error: %s", err)
			t.FailNow()
		}

		if d.Status != deployment.StatusWaiting {
			t.Errorf("Deploy.Status = %v, wanted %v", d, deployment.StatusCreated)
		}
	})
}

func TestInteractor_DeployToRemote(t *testing.T) {
	t.Run("Create a new remote deployment and update the deployment.", func(t *testing.T) {
		t.Log("Start mocking")
		ctrl := gomock.NewController(t)
		store := mock.NewMockStore(ctrl)
		scm := mock.NewMockSCM(ctrl)

		t.Log("\tValidate the request.")
		store.
			EXPECT().
			HasLockOfRepoForEnv(ctx, gomock.AssignableToTypeOf(&ent.Repo{}), gomock.AssignableToTypeOf("environment")).
			Return(false, nil)

		store.
			EXPECT().
			ListReviews(ctx, gomock.AssignableToTypeOf(&ent.Deployment{})).
			Return([]*ent.Review{
				{Status: review.StatusApproved},
			}, nil)

		t.Log("\tUpdate the deployment with the response.")
		scm.
			EXPECT().
			CreateRemoteDeployment(ctx, gomock.AssignableToTypeOf(&ent.User{}), gomock.AssignableToTypeOf(&ent.Repo{}), gomock.AssignableToTypeOf(&ent.Deployment{}), gomock.AssignableToTypeOf(&extent.Env{})).
			Return(&extent.RemoteDeployment{}, nil)

		store.
			EXPECT().
			UpdateDeployment(ctx, gomock.AssignableToTypeOf(&ent.Deployment{})).
			DoAndReturn(func(ctx context.Context, d *ent.Deployment) (interface{}, interface{}) {
				return d, nil
			})

		store.
			EXPECT().
			CreateEntDeploymentStatus(ctx, gomock.AssignableToTypeOf(&ent.DeploymentStatus{}))

		it := i.NewInteractor(&i.InteractorConfig{
			Store: store,
			SCM:   scm,
		})

		d, err := it.DeployToRemote(context.Background(), &ent.User{}, &ent.Repo{}, &ent.Deployment{
			Status: deployment.StatusWaiting,
		}, &extent.Env{})
		if err != nil {
			t.Errorf("DeployToRemote returns a error: %s", err)
			t.FailNow()
		}

		if d.Status != deployment.StatusCreated {
			t.Errorf("DeployToRemote = %v, wanted %v", d, deployment.StatusCreated)
		}
	})
}
