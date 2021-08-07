package interactor

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/deployment"
	"github.com/hanjunlee/gitploy/internal/interactor/mock"
	"github.com/hanjunlee/gitploy/vo"
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
	any := gomock.Any()

	t.Run("create a new deployment and update with the remote deployment.", func(t *testing.T) {
		input := struct {
			u   *ent.User
			r   *ent.Repo
			d   *ent.Deployment
			env *vo.Env
		}{
			u: &ent.User{
				ID: "1",
			},
			r: &ent.Repo{
				ID: "1",
			},
			d: &ent.Deployment{
				ID: 1,
			},
			env: &vo.Env{},
		}
		rd := &vo.RemoteDeployment{
			UID:     1000,
			SHA:     "087135f",
			HTLMURL: "https://github.com/commit/087135f",
		}

		ctrl := gomock.NewController(t)
		store := mock.NewMockStore(ctrl)
		scm := mock.NewMockSCM(ctrl)

		store.
			EXPECT().
			CreateDeployment(ctx, gomock.AssignableToTypeOf(&ent.Deployment{})).
			DoAndReturn(func(ctx context.Context, d *ent.Deployment) (interface{}, interface{}) {
				return d, nil
			})

		scm.
			EXPECT().
			CreateDeployment(ctx, any, any, any, any).
			Return(rd, nil)

		store.
			EXPECT().
			UpdateDeployment(ctx, gomock.AssignableToTypeOf(&ent.Deployment{})).
			DoAndReturn(func(ctx context.Context, d *ent.Deployment) (interface{}, interface{}) {
				return d, nil
			})

		store.
			EXPECT().
			CreateDeploymentStatus(ctx, gomock.AssignableToTypeOf(&ent.DeploymentStatus{}))

		i := newMockInteractor(store, scm)

		d, err := i.Deploy(context.Background(), input.u, input.r, input.d, input.env)
		if err != nil {
			t.Errorf("Deploy returns a error: %s", err)
			t.FailNow()
		}

		expected := &ent.Deployment{
			ID:      input.d.ID,
			UID:     rd.UID,
			Sha:     rd.SHA,
			HTMLURL: rd.HTLMURL,
			Status:  deployment.StatusCreated,
			UserID:  input.u.ID,
			RepoID:  input.r.ID,
		}
		if !reflect.DeepEqual(d, expected) {
			t.Errorf("Deploy = %v, wanted %v", d, expected)
		}
	})
}

func TestInteractor_CreateDeploymentToSCM(t *testing.T) {
	ctx := gomock.Any()
	any := gomock.Any()

	t.Run("update deployment with the remote deployment.", func(t *testing.T) {
		input := struct {
			u   *ent.User
			r   *ent.Repo
			d   *ent.Deployment
			env *vo.Env
		}{
			u: &ent.User{},
			r: &ent.Repo{},
			d: &ent.Deployment{
				ID: 1,
			},
			env: &vo.Env{},
		}
		rd := &vo.RemoteDeployment{
			UID:     1000,
			SHA:     "087135f",
			HTLMURL: "https://github.com/commit/087135f",
		}

		ctrl := gomock.NewController(t)
		store := mock.NewMockStore(ctrl)
		scm := mock.NewMockSCM(ctrl)

		scm.
			EXPECT().
			CreateDeployment(ctx, any, any, any, any).
			Return(rd, nil)

		store.
			EXPECT().
			UpdateDeployment(ctx, gomock.AssignableToTypeOf(&ent.Deployment{})).
			DoAndReturn(func(ctx context.Context, d *ent.Deployment) (interface{}, interface{}) {
				return d, nil
			})

		store.
			EXPECT().
			CreateDeploymentStatus(ctx, gomock.AssignableToTypeOf(&ent.DeploymentStatus{}))

		i := newMockInteractor(store, scm)

		d, err := i.CreateDeploymentToSCM(context.Background(), input.u, input.r, input.d, input.env)
		if err != nil {
			t.Errorf("CreateDeploymentToSCM returns a error: %s", err)
			t.FailNow()
		}

		expected := &ent.Deployment{
			ID:      input.d.ID,
			UID:     rd.UID,
			Sha:     rd.SHA,
			HTMLURL: rd.HTLMURL,
			Status:  deployment.StatusCreated,
		}
		if !reflect.DeepEqual(d, expected) {
			t.Errorf("CreateDeploymentToSCM = %v, wanted %v", d, expected)
		}
	})
}
