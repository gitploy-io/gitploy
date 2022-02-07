package interactor

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"go.uber.org/zap"

	"github.com/gitploy-io/gitploy/internal/interactor/mock"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/pkg/e"
)

func TestInteractor_DeactivateRepo(t *testing.T) {
	t.Run("Deactivate successfully even if the webhook is not found.", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		store := mock.NewMockStore(ctrl)
		scm := mock.NewMockSCM(ctrl)

		t.Log("Mocking DeleteWebhook to return an EntityNotFound error.")
		scm.
			EXPECT().
			DeleteWebhook(gomock.Any(), gomock.AssignableToTypeOf(&ent.User{}), gomock.AssignableToTypeOf(&ent.Repo{}), gomock.Any()).
			Return(e.NewError(e.ErrorCodeEntityNotFound, nil))

		store.
			EXPECT().
			Deactivate(gomock.Any(), gomock.AssignableToTypeOf(&ent.Repo{})).
			Return(&ent.Repo{}, nil)

		i := &ReposInteractor{
			store: store,
			scm:   scm,
			log:   zap.L(),
		}

		_, err := i.DeactivateRepo(context.Background(), &ent.User{}, &ent.Repo{})
		if err != nil {
			t.Fatalf("DeactivateRepo returns an error: %v", err)
		}
	})
}
