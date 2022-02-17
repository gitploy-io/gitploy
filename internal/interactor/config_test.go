package interactor_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"

	i "github.com/gitploy-io/gitploy/internal/interactor"
	"github.com/gitploy-io/gitploy/internal/interactor/mock"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/extent"
)

func TestConfigInteractor_GetEvaluatedConfig(t *testing.T) {
	t.Run("Return the evaluated config", func(t *testing.T) {
		t.Log("Start mocking: ")
		ctrl := gomock.NewController(t)
		scm := mock.NewMockSCM(ctrl)

		t.Log("\tGet the config.")
		scm.EXPECT().
			GetConfig(gomock.Any(), gomock.AssignableToTypeOf(&ent.User{}), gomock.AssignableToTypeOf(&ent.Repo{})).
			Return(&extent.Config{}, nil)

		it := i.NewInteractor(&i.InteractorConfig{
			Store: mock.NewMockStore(ctrl),
			SCM:   scm,
		})

		_, err := it.GetEvaluatedConfig(context.Background(), &ent.User{}, &ent.Repo{}, &extent.EvalValues{})
		if err != nil {
			t.Fatalf("GetEvaluatedConfig returns an error: %v", err)
		}
	})
}
