package interactor

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/internal/interactor/mock"
)

func TestInteractor_SaveUser(t *testing.T) {
	ctx := gomock.Any()

	t.Run("Creates a new user if user is not found.", func(t *testing.T) {
		input := struct {
			u *ent.User
		}{
			u: &ent.User{
				ID: "1",
			},
		}

		ctrl := gomock.NewController(t)
		store := mock.NewMockStore(ctrl)
		scm := mock.NewMockSCM(ctrl)

		t.Logf("Couldn't find the user.")
		store.
			EXPECT().
			FindUserByID(ctx, input.u.ID).
			Return(nil, &ent.NotFoundError{})

		t.Logf("Create a new user.")
		store.EXPECT().
			CreateUser(ctx, gomock.Eq(&ent.User{ID: input.u.ID})).
			DoAndReturn(func(ctx context.Context, u *ent.User) (*ent.User, error) {
				return u, nil
			})

		i := NewInteractor(&InteractorConfig{
			Store: store,
			SCM:   scm,
		})

		u, err := i.SaveUser(context.Background(), input.u)
		if err != nil {
			t.Fatalf("SaveUser returns an error: %s", err)
		}

		expected := &ent.User{ID: input.u.ID}
		if !reflect.DeepEqual(u, expected) {
			t.Fatalf("SaveUser = %v, wanted %v", u, expected)
		}
	})

	t.Run("Update the user if user is found.", func(t *testing.T) {
		input := struct {
			u *ent.User
		}{
			u: &ent.User{
				ID:    "1",
				Token: "acee",
			},
		}

		ctrl := gomock.NewController(t)
		store := mock.NewMockStore(ctrl)
		scm := mock.NewMockSCM(ctrl)

		t.Logf("Find the user by ID.")
		store.
			EXPECT().
			FindUserByID(ctx, input.u.ID).
			Return(&ent.User{
				ID:    input.u.ID,
				Token: "",
			}, nil)

		t.Logf("Update token of the user.")
		store.EXPECT().
			UpdateUser(ctx, gomock.Eq(&ent.User{
				ID:    input.u.ID,
				Token: input.u.Token,
			})).
			DoAndReturn(func(ctx context.Context, u *ent.User) (*ent.User, error) {
				return u, nil
			})

		i := NewInteractor(&InteractorConfig{
			Store: store,
			SCM:   scm,
		})

		u, err := i.SaveUser(context.Background(), input.u)
		if err != nil {
			t.Fatalf("SaveUser returns an error: %s", err)
		}

		expected := &ent.User{
			ID:    input.u.ID,
			Token: input.u.Token,
		}
		if !reflect.DeepEqual(u, expected) {
			t.Fatalf("SaveUser = %v, wanted %v", u, expected)
		}
	})
}
