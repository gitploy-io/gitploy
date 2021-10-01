package interactor

import (
	"context"
	"reflect"
	"testing"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/deployment"
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

	t.Run("create a new deployment.", func(t *testing.T) {
		input := struct {
			u *ent.User
			r *ent.Repo
			d *ent.Deployment
			e *vo.Env
		}{
			u: &ent.User{
				ID: 1,
			},
			r: &ent.Repo{
				ID: 1,
			},
			d: &ent.Deployment{
				Number: 3,
				Type:   deployment.TypeCommit,
				Ref:    "3ee3221",
				Env:    "local",
			},
			e: &vo.Env{},
		}

		ctrl := gomock.NewController(t)
		store := mock.NewMockStore(ctrl)
		scm := mock.NewMockSCM(ctrl)

		const (
			ID  = 1
			UID = 1000
		)

		t.Logf("Returns a new remote deployment with UID = %d", UID)
		scm.
			EXPECT().
			CreateRemoteDeployment(ctx, gomock.Eq(input.u), gomock.Eq(input.r), gomock.Eq(input.d), gomock.Eq(input.e)).
			Return(&vo.RemoteDeployment{
				UID: UID,
			}, nil)

		t.Logf("Check the deployment input has UID")
		store.
			EXPECT().
			CreateDeployment(ctx, gomock.Eq(&ent.Deployment{
				Number: input.d.Number,
				Type:   input.d.Type,
				Ref:    input.d.Ref,
				Env:    input.d.Env,
				UID:    UID,
				Status: deployment.StatusCreated,
				UserID: input.u.ID,
				RepoID: input.r.ID,
			})).
			DoAndReturn(func(ctx context.Context, d *ent.Deployment) (interface{}, interface{}) {
				d.ID = ID
				return d, nil
			})

		store.
			EXPECT().
			CreateDeploymentStatus(ctx, gomock.AssignableToTypeOf(&ent.DeploymentStatus{}))

		i := newMockInteractor(store, scm)

		d, err := i.Deploy(context.Background(), input.u, input.r, input.d, input.e)
		if err != nil {
			t.Errorf("Deploy returns a error: %s", err)
			t.FailNow()
		}

		expected := &ent.Deployment{
			ID:     ID,
			Number: input.d.Number,
			Type:   input.d.Type,
			Ref:    input.d.Ref,
			Env:    input.d.Env,
			UID:    UID,
			Status: deployment.StatusCreated,
			UserID: input.u.ID,
			RepoID: input.r.ID,
		}
		if !reflect.DeepEqual(d, expected) {
			t.Errorf("Deploy = %v, wanted %v", d, expected)
		}
	})

	t.Run("create a new deployment with the approval configuration.", func(t *testing.T) {
		input := struct {
			u *ent.User
			r *ent.Repo
			d *ent.Deployment
			e *vo.Env
		}{
			u: &ent.User{
				ID: 1,
			},
			r: &ent.Repo{
				ID: 1,
			},
			d: &ent.Deployment{
				Number: 3,
				Type:   deployment.TypeCommit,
				Ref:    "3ee3221",
				Env:    "local",
			},
			e: &vo.Env{
				Approval: &vo.Approval{
					Enabled:       true,
					RequiredCount: 1,
				},
			},
		}

		ctrl := gomock.NewController(t)
		store := mock.NewMockStore(ctrl)
		scm := mock.NewMockSCM(ctrl)

		const (
			ID = 1
		)

		t.Logf("Check the deployment has configurations of approval.")
		store.
			EXPECT().
			CreateDeployment(ctx, gomock.Eq(&ent.Deployment{
				Number:                input.d.Number,
				Type:                  input.d.Type,
				Ref:                   input.d.Ref,
				Env:                   input.d.Env,
				IsApprovalEnabled:     true,
				RequiredApprovalCount: input.e.Approval.RequiredCount,
				Status:                deployment.StatusWaiting,
				UserID:                input.u.ID,
				RepoID:                input.r.ID,
			})).
			DoAndReturn(func(ctx context.Context, d *ent.Deployment) (interface{}, interface{}) {
				d.ID = ID
				return d, nil
			})

		i := newMockInteractor(store, scm)

		d, err := i.Deploy(context.Background(), input.u, input.r, input.d, input.e)
		if err != nil {
			t.Errorf("Deploy returns a error: %s", err)
			t.FailNow()
		}

		expected := &ent.Deployment{
			ID:                    ID,
			Number:                input.d.Number,
			Type:                  input.d.Type,
			Ref:                   input.d.Ref,
			Env:                   input.d.Env,
			IsApprovalEnabled:     true,
			RequiredApprovalCount: input.e.Approval.RequiredCount,
			Status:                deployment.StatusWaiting,
			UserID:                input.u.ID,
			RepoID:                input.r.ID,
		}
		if !reflect.DeepEqual(d, expected) {
			t.Errorf("Deploy = %v, wanted %v", d, expected)
		}
	})
}

func TestInteractor_CreateRemoteDeployment(t *testing.T) {
	ctx := gomock.Any()

	t.Run("create a new remote deployment and update the deployment.", func(t *testing.T) {
		input := struct {
			u *ent.User
			r *ent.Repo
			d *ent.Deployment
			e *vo.Env
		}{
			u: &ent.User{},
			r: &ent.Repo{},
			d: &ent.Deployment{
				ID: 1,
			},
			e: &vo.Env{},
		}

		ctrl := gomock.NewController(t)
		store := mock.NewMockStore(ctrl)
		scm := mock.NewMockSCM(ctrl)

		const (
			UID = 1000
		)

		t.Logf("Returns a new remote deployment with UID = %d", UID)
		scm.
			EXPECT().
			CreateRemoteDeployment(ctx, gomock.Eq(input.u), gomock.Eq(input.r), gomock.Eq(input.d), gomock.Eq(input.e)).
			Return(&vo.RemoteDeployment{
				UID: UID,
			}, nil)

		t.Logf("Check the deployment input has UID")
		store.
			EXPECT().
			UpdateDeployment(ctx, gomock.Eq(&ent.Deployment{
				ID:     input.d.ID,
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

		d, err := i.CreateRemoteDeployment(context.Background(), input.u, input.r, input.d, input.e)
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
