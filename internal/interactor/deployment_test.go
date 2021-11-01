package interactor

import (
	"context"
	"reflect"
	"testing"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/deployment"
	"github.com/gitploy-io/gitploy/ent/review"
	"github.com/gitploy-io/gitploy/internal/interactor/mock"
	"github.com/gitploy-io/gitploy/vo"
	"github.com/golang/mock/gomock"
	"go.uber.org/zap"
)

func newMockInteractor(store Store, scm SCM) *Interactor {
	return &Interactor{
		Store: store,
		SCM:   scm,
		log:   zap.L(),
	}
}

func TestInteractor_Deploy(t *testing.T) {
	ctx := gomock.Any()

	t.Run("Return a new deployment.", func(t *testing.T) {
		input := struct {
			d *ent.Deployment
			e *vo.Env
		}{
			d: &ent.Deployment{
				Type: deployment.TypeBranch,
				Ref:  "main",
				Env:  "production",
			},
			e: &vo.Env{},
		}

		ctrl := gomock.NewController(t)
		store := mock.NewMockStore(ctrl)
		scm := mock.NewMockSCM(ctrl)

		const (
			UID = 1000
		)

		store.
			EXPECT().
			HasLockOfRepoForEnv(ctx, gomock.AssignableToTypeOf(&ent.Repo{}), gomock.AssignableToTypeOf("")).
			Return(false, nil)

		store.
			EXPECT().
			GetNextDeploymentNumberOfRepo(ctx, gomock.AssignableToTypeOf(&ent.Repo{})).
			Return(1, nil)

		scm.
			EXPECT().
			CreateRemoteDeployment(ctx, gomock.Eq(&ent.User{}), gomock.Eq(&ent.Repo{}), gomock.AssignableToTypeOf(&ent.Deployment{}), gomock.Eq(&vo.Env{})).
			Return(&vo.RemoteDeployment{
				UID: UID,
			}, nil)

		t.Logf("MOCK - compare the deployment parameter.")
		store.
			EXPECT().
			CreateDeployment(ctx, gomock.Eq(&ent.Deployment{
				Number: 1, // The next deployment number.
				Type:   input.d.Type,
				Ref:    input.d.Ref,
				Env:    input.d.Env,
				UID:    UID,
				Status: deployment.StatusCreated,
			})).
			DoAndReturn(func(ctx context.Context, d *ent.Deployment) (interface{}, interface{}) {
				d.ID = 1
				return d, nil
			})

		store.
			EXPECT().
			CreateDeploymentStatus(ctx, gomock.AssignableToTypeOf(&ent.DeploymentStatus{}))

		i := newMockInteractor(store, scm)

		d, err := i.Deploy(context.Background(), &ent.User{}, &ent.Repo{}, input.d, input.e)
		if err != nil {
			t.Errorf("Deploy returns a error: %s", err)
			t.FailNow()
		}

		expected := &ent.Deployment{
			ID:     1,
			Number: 1,
			Type:   input.d.Type,
			Ref:    input.d.Ref,
			Env:    input.d.Env,
			UID:    UID,
			Status: deployment.StatusCreated,
		}
		if !reflect.DeepEqual(d, expected) {
			t.Errorf("Deploy = %v, wanted %v", d, expected)
		}
	})

	t.Run("Return the waiting deployment and reviews.", func(t *testing.T) {
		input := struct {
			d *ent.Deployment
			e *vo.Env
		}{
			d: &ent.Deployment{
				Number: 3,
				Type:   deployment.TypeBranch,
				Ref:    "main",
				Env:    "production",
			},
			e: &vo.Env{
				Approval: &vo.Approval{
					Enabled:       true,
					RequiredCount: 1,
				},
				Review: &vo.Review{
					Enabled:   true,
					Reviewers: []string{"octocat"},
				},
			},
		}

		ctrl := gomock.NewController(t)
		store := mock.NewMockStore(ctrl)
		scm := mock.NewMockSCM(ctrl)

		store.
			EXPECT().
			HasLockOfRepoForEnv(ctx, gomock.AssignableToTypeOf(&ent.Repo{}), gomock.AssignableToTypeOf("")).
			Return(false, nil)

		store.
			EXPECT().
			GetNextDeploymentNumberOfRepo(ctx, gomock.AssignableToTypeOf(&ent.Repo{})).
			Return(1, nil)

		t.Logf("MOCK - compare the deployment parameter.")
		store.
			EXPECT().
			CreateDeployment(ctx, gomock.Eq(&ent.Deployment{
				Number: 1,
				Type:   input.d.Type,
				Ref:    input.d.Ref,
				Env:    input.d.Env,
				Status: deployment.StatusWaiting,
			})).
			DoAndReturn(func(ctx context.Context, d *ent.Deployment) (interface{}, interface{}) {
				d.ID = 1
				return d, nil
			})

		store.
			EXPECT().
			FindUserByLogin(ctx, gomock.AssignableToTypeOf("")).
			Return(&ent.User{}, nil)

		store.
			EXPECT().
			CreateReview(ctx, gomock.AssignableToTypeOf(&ent.Review{})).
			Return(&ent.Review{}, nil)

		store.
			EXPECT().
			CreateEvent(ctx, gomock.AssignableToTypeOf(&ent.Event{})).
			Return(&ent.Event{}, nil)

		i := newMockInteractor(store, scm)

		d, err := i.Deploy(context.Background(), &ent.User{}, &ent.Repo{}, input.d, input.e)
		if err != nil {
			t.Errorf("Deploy returns a error: %s", err)
			t.FailNow()
		}

		expected := &ent.Deployment{
			ID:     1,
			Number: 1,
			Type:   input.d.Type,
			Ref:    input.d.Ref,
			Env:    input.d.Env,
			Status: deployment.StatusWaiting,
		}
		if !reflect.DeepEqual(d, expected) {
			t.Errorf("Deploy = %v, wanted %v", d, expected)
		}
	})
}

func TestInteractor_DeployToRemote(t *testing.T) {
	ctx := gomock.Any()

	t.Run("create a new remote deployment and update the deployment.", func(t *testing.T) {
		input := struct {
			d *ent.Deployment
			e *vo.Env
		}{
			d: &ent.Deployment{},
			e: &vo.Env{},
		}

		ctrl := gomock.NewController(t)
		store := mock.NewMockStore(ctrl)
		scm := mock.NewMockSCM(ctrl)

		const (
			UID = 1000
		)

		store.
			EXPECT().
			HasLockOfRepoForEnv(ctx, gomock.AssignableToTypeOf(&ent.Repo{}), gomock.AssignableToTypeOf("")).
			Return(false, nil)

		// Return a approved review.
		store.
			EXPECT().
			ListReviews(ctx, gomock.AssignableToTypeOf(&ent.Deployment{})).
			Return([]*ent.Review{
				{
					Status: review.StatusApproved,
				},
			}, nil)

		scm.
			EXPECT().
			CreateRemoteDeployment(ctx, gomock.AssignableToTypeOf(&ent.User{}), gomock.AssignableToTypeOf(&ent.Repo{}), gomock.AssignableToTypeOf(&ent.Deployment{}), gomock.AssignableToTypeOf(&vo.Env{})).
			Return(&vo.RemoteDeployment{
				UID: UID,
			}, nil)

		t.Log("MOCK - Compare the deployment parameter.")
		store.
			EXPECT().
			UpdateDeployment(ctx, gomock.Eq(&ent.Deployment{
				UID:    UID,
				Status: deployment.StatusCreated,
			})).
			DoAndReturn(func(ctx context.Context, d *ent.Deployment) (interface{}, interface{}) {
				return d, nil
			})

		store.
			EXPECT().
			CreateDeploymentStatus(ctx, gomock.AssignableToTypeOf(&ent.DeploymentStatus{}))

		i := newMockInteractor(store, scm)

		d, err := i.DeployToRemote(context.Background(), &ent.User{}, &ent.Repo{}, input.d, input.e)
		if err != nil {
			t.Errorf("CreateRemoteDeployment returns a error: %s", err)
			t.FailNow()
		}

		expected := &ent.Deployment{
			ID:     input.d.ID,
			UID:    UID,
			Status: deployment.StatusCreated,
		}
		if !reflect.DeepEqual(d, expected) {
			t.Errorf("CreateRemoteDeployment = %v, wanted %v", d, expected)
		}
	})
}
