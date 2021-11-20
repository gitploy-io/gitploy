package interactor

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/internal/interactor/mock"
	"github.com/gitploy-io/gitploy/pkg/e"
	"github.com/gitploy-io/gitploy/vo"
)

func TestInteractor_GetEnv(t *testing.T) {
	t.Run("Return an error when the environment is not defined.", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		scm := mock.NewMockSCM(ctrl)

		scm.
			EXPECT().
			GetConfig(gomock.Any(), gomock.AssignableToTypeOf(&ent.User{}), gomock.AssignableToTypeOf(&ent.Repo{})).
			Return(&vo.Config{
				Envs: []*vo.Env{},
			}, nil)

		i := &Interactor{SCM: scm}

		_, err := i.GetEnv(context.Background(), &ent.User{}, &ent.Repo{}, "production")
		if !e.HasErrorCode(err, e.ErrorCodeConfigUndefinedEnv) {
			t.Fatalf("GetEnv error = %v, wanted ErrorCodeConfigUndefinedEnv", err)
		}
	})

	t.Run("Return the environment.", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		scm := mock.NewMockSCM(ctrl)

		scm.
			EXPECT().
			GetConfig(gomock.Any(), gomock.AssignableToTypeOf(&ent.User{}), gomock.AssignableToTypeOf(&ent.Repo{})).
			Return(&vo.Config{
				Envs: []*vo.Env{
					{
						Name: "production",
					},
				},
			}, nil)

		i := &Interactor{SCM: scm}

		_, err := i.GetEnv(context.Background(), &ent.User{}, &ent.Repo{}, "production")
		if err != nil {
			t.Fatalf("GetEnv returns an error: %v", err)
		}
	})
}
