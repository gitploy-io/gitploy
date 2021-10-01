package interactor

import (
	"context"
	"testing"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/perm"
	"github.com/gitploy-io/gitploy/internal/interactor/mock"
	"github.com/gitploy-io/gitploy/vo"
	"github.com/golang/mock/gomock"
)

func TestInteractor_SyncRemoteRepo(t *testing.T) {
	ctx := gomock.Any()

	t.Run("Synchronization creates a new repository and a new perm when it is first.", func(t *testing.T) {
		input := struct {
			user   *ent.User
			remote *vo.RemoteRepo
		}{
			user: &ent.User{
				ID: 2214,
			},
			remote: &vo.RemoteRepo{
				ID:   2214,
				Perm: vo.RemoteRepoPermRead,
			},
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
			})).
			DoAndReturn(func(ctx context.Context, p *ent.Perm) (*ent.Perm, error) {
				return p, nil
			})

		i := &Interactor{Store: store}
		if err := i.SyncRemoteRepo(context.Background(), input.user, input.remote); err != nil {
			t.Fatal("SyncRemoteRepo returns error.")
		}
	})
}
