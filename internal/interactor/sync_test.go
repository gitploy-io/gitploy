package interactor_test

import (
	"context"
	"testing"
	"time"

	i "github.com/gitploy-io/gitploy/internal/interactor"
	"github.com/gitploy-io/gitploy/internal/interactor/mock"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/ent/perm"
	"github.com/gitploy-io/gitploy/model/extent"
	"github.com/golang/mock/gomock"
)

func TestInteractor_SyncRemoteRepo(t *testing.T) {
	ctx := gomock.Any()

	t.Run("Synchronization creates a new repository and a new perm when it is first.", func(t *testing.T) {
		input := struct {
			user   *ent.User
			remote *extent.RemoteRepo
			time   time.Time
		}{
			user: &ent.User{
				ID: 2214,
			},
			remote: &extent.RemoteRepo{
				ID:   2214,
				Perm: extent.RemoteRepoPermRead,
			},
			time: time.Now(),
		}

		ctrl := gomock.NewController(t)
		store := mock.NewMockStore(ctrl)

		t.Log("The repo is not found.")
		store.
			EXPECT().
			FindRepoByID(ctx, input.remote.ID).
			Return(nil, &ent.NotFoundError{})

		t.Log("Sync with the remote repo.")
		store.
			EXPECT().
			SyncRepo(ctx, input.remote).
			Return(&ent.Repo{
				ID: input.remote.ID,
			}, nil)

		t.Log("The perm is not found.")
		store.
			EXPECT().
			FindPermOfRepo(ctx, gomock.Eq(&ent.Repo{ID: input.remote.ID}), gomock.Eq(input.user)).
			Return(nil, &ent.NotFoundError{})

		t.Log("Create a new perm for the repo.")
		store.
			EXPECT().
			CreatePerm(ctx, gomock.Eq(&ent.Perm{
				RepoPerm: perm.RepoPerm(input.remote.Perm),
				UserID:   input.user.ID,
				RepoID:   input.remote.ID,
				SyncedAt: input.time,
			})).
			DoAndReturn(func(ctx context.Context, p *ent.Perm) (*ent.Perm, error) {
				return p, nil
			})

		it := i.NewInteractor(&i.InteractorConfig{
			Store: store,
		})
		if err := it.SyncRemoteRepo(context.Background(), input.user, input.remote, input.time); err != nil {
			t.Fatal("SyncRemoteRepo returns error.")
		}
	})

	t.Run("Synchronization updates the perm if it exist.", func(t *testing.T) {
		input := struct {
			user   *ent.User
			remote *extent.RemoteRepo
			time   time.Time
		}{
			user: &ent.User{
				ID: 1,
			},
			remote: &extent.RemoteRepo{
				ID:   1,
				Perm: extent.RemoteRepoPermWrite,
			},
			time: time.Now(),
		}

		ctrl := gomock.NewController(t)
		store := mock.NewMockStore(ctrl)

		t.Log("The repo is found.")
		store.
			EXPECT().
			FindRepoByID(ctx, input.remote.ID).
			Return(&ent.Repo{ID: input.remote.ID}, nil)

		t.Log("The perm is found.")
		store.
			EXPECT().
			FindPermOfRepo(ctx, gomock.Eq(&ent.Repo{ID: input.remote.ID}), gomock.Eq(input.user)).
			Return(&ent.Perm{}, nil)

		t.Log("Update the perm with perm, and synced_at.")
		store.
			EXPECT().
			UpdatePerm(ctx, gomock.Eq(&ent.Perm{
				RepoPerm: perm.RepoPerm(input.remote.Perm),
				SyncedAt: input.time,
			})).
			DoAndReturn(func(ctx context.Context, p *ent.Perm) (*ent.Perm, error) {
				return p, nil
			})

		it := i.NewInteractor(&i.InteractorConfig{
			Store: store,
		})
		if err := it.SyncRemoteRepo(context.Background(), input.user, input.remote, input.time); err != nil {
			t.Fatal("SyncRemoteRepo returns error.")
		}
	})
}
