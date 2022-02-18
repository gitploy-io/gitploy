package interactor_test

import (
	"context"
	"testing"

	gm "github.com/golang/mock/gomock"

	i "github.com/gitploy-io/gitploy/internal/interactor"
	"github.com/gitploy-io/gitploy/internal/interactor/mock"
	"github.com/gitploy-io/gitploy/model/ent"
)

func TestInteractor_RespondReview(t *testing.T) {
	t.Run("Return the review.", func(t *testing.T) {
		t.Log("Start mocking:")
		ctrl := gm.NewController(t)
		store := mock.NewMockStore(ctrl)

		t.Log("\tUpdates the review and dispatches a event.")
		store.EXPECT().
			UpdateReview(gm.Any(), gm.AssignableToTypeOf(&ent.Review{})).
			Return(&ent.Review{}, nil)

		store.EXPECT().
			CreateEvent(gm.Any(), gm.AssignableToTypeOf(&ent.Event{})).
			Return(&ent.Event{}, nil)

		it := i.NewInteractor(&i.InteractorConfig{
			Store: store,
		})
		_, err := it.RespondReview(context.Background(), &ent.Review{})
		if err != nil {
			t.Fatalf("RespondReview returns an error: %v", err)
		}
	})

}
