package interactor_test

import (
	"context"
	"testing"
	"time"

	i "github.com/gitploy-io/gitploy/internal/interactor"
	"github.com/gitploy-io/gitploy/internal/interactor/mock"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/extent"
	"github.com/golang/mock/gomock"
)

func TestInteractor_IsAdminUser(t *testing.T) {
	t.Run("Return false when admins property is nil", func(t *testing.T) {
		i := i.NewInteractor(&i.InteractorConfig{
			AdminUsers: nil,
		})

		expected := false
		if ret := i.IsAdminUser(context.Background(), "octocat"); ret != expected {
			t.Fatalf("IsAdminUser = %v, wanted %v", ret, expected)
		}
	})
}

func TestInteractor_IsEntryMember(t *testing.T) {
	t.Run("Return false when the user's login is not included.", func(t *testing.T) {
		it := i.NewInteractor(&i.InteractorConfig{
			MemberEntries: []string{"octocat"},
		})

		want := false
		if ret := it.IsEntryMember(context.Background(), "coco"); ret != want {
			t.Fatalf("IsEntryMember = %v, wanted %v", ret, want)
		}
	})

	t.Run("Return true when the user's login is included.", func(t *testing.T) {
		it := i.NewInteractor(&i.InteractorConfig{
			MemberEntries: []string{"octocat"},
		})

		want := true
		if ret := it.IsEntryMember(context.Background(), "octocat"); ret != want {
			t.Fatalf("IsEntryMember = %v, wanted %v", ret, want)
		}
	})
}

func TestInteractor_IsOrgMember(t *testing.T) {
	t.Run("Return false when the org is not included.", func(t *testing.T) {
		it := i.NewInteractor(&i.InteractorConfig{
			MemberEntries: []string{"gitploy-io"},
		})

		want := false
		if ret := it.IsOrgMember(context.Background(), []string{"github"}); ret != want {
			t.Fatalf("IsEntryMember = %v, wanted %v", ret, want)
		}
	})

	t.Run("Return true when the org is included.", func(t *testing.T) {
		it := i.NewInteractor(&i.InteractorConfig{
			MemberEntries: []string{"gitploy-io"},
		})

		want := true
		if ret := it.IsOrgMember(context.Background(), []string{"gitploy-io"}); ret != want {
			t.Fatalf("IsEntryMember = %v, wanted %v", ret, want)
		}
	})
}

func TestInteractor_SyncRemoteRepo(t *testing.T) {
	ctx := gomock.Any()

	t.Run("Synchronization creates a new repository and a new perm when it is first.", func(t *testing.T) {
		t.Log("Start mocking:")
		ctrl := gomock.NewController(t)
		store := mock.NewMockStore(ctrl)

		t.Log("\tThe repository is not found, create a new repository.")
		store.
			EXPECT().
			FindRepoByID(ctx, gomock.Any()).
			Return(nil, &ent.NotFoundError{})

		store.
			EXPECT().
			SyncRepo(ctx, gomock.AssignableToTypeOf(&extent.RemoteRepo{})).
			Return(&ent.Repo{}, nil)

		t.Log("\tThe perm is not found, create a new perm.")
		store.
			EXPECT().
			FindPermOfRepo(ctx, gomock.AssignableToTypeOf(&ent.Repo{}), gomock.AssignableToTypeOf(&ent.User{})).
			Return(nil, &ent.NotFoundError{})

		store.
			EXPECT().
			CreatePerm(ctx, gomock.AssignableToTypeOf(&ent.Perm{})).
			DoAndReturn(func(ctx context.Context, p *ent.Perm) (*ent.Perm, error) {
				return p, nil
			})

		it := i.NewInteractor(&i.InteractorConfig{
			Store: store,
		})
		if err := it.SyncRemoteRepo(context.Background(), &ent.User{}, &extent.RemoteRepo{}, time.Now()); err != nil {
			t.Fatal("SyncRemoteRepo returns error.")
		}
	})

	t.Run("Synchronization updates the perm if it exist.", func(t *testing.T) {
		t.Log("Start mocking:")
		ctrl := gomock.NewController(t)
		store := mock.NewMockStore(ctrl)

		store.
			EXPECT().
			FindRepoByID(ctx, gomock.Any()).
			Return(&ent.Repo{}, nil)

		t.Log("\tThe perm exists, the interactor updates the perm.")
		store.
			EXPECT().
			FindPermOfRepo(ctx, gomock.AssignableToTypeOf(&ent.Repo{}), gomock.AssignableToTypeOf(&ent.User{})).
			Return(&ent.Perm{}, nil)

		store.
			EXPECT().
			UpdatePerm(ctx, gomock.AssignableToTypeOf(&ent.Perm{})).
			DoAndReturn(func(ctx context.Context, p *ent.Perm) (*ent.Perm, error) {
				return p, nil
			})

		it := i.NewInteractor(&i.InteractorConfig{
			Store: store,
		})
		if err := it.SyncRemoteRepo(context.Background(), &ent.User{}, &extent.RemoteRepo{}, time.Now()); err != nil {
			t.Fatal("SyncRemoteRepo returns error.")
		}
	})
}
