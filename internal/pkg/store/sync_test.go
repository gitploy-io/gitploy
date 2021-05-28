package store

import (
	"context"
	"testing"
	"time"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/enttest"
	"github.com/hanjunlee/gitploy/ent/perm"

	_ "github.com/mattn/go-sqlite3"
)

func TestStore_SyncPerm(t *testing.T) {
	t.Run("sync by creating a new repo and a new perm", func(tt *testing.T) {
		client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
		defer client.Close()

		ctx := context.Background()

		u, err := client.User.Create().
			SetID("1").
			SetLogin("octocat").
			SetToken("access_token").
			SetRefresh("refresh_token").
			SetExpiry(time.Time{}).
			Save(ctx)
		if err != nil {
			tt.Error("failed to create a new user")
			return
		}

		p := &ent.Perm{
			RepoPerm: perm.RepoPermWrite,
			Edges: ent.PermEdges{
				Repo: &ent.Repo{
					ID:        "1",
					Namespace: "octocat",
					Name:      "HelloWorld",
				},
			},
		}

		s := NewStore(client)

		err = s.SyncPerm(ctx, u, p, time.Now())
		if err != nil {
			tt.Errorf("failed to sync perm: %s", err)
		}

		if cnt, _ := client.Repo.Query().Count(ctx); cnt != 1 {
			tt.Error("repo was not created.")
		}

		if cnt, _ := client.Perm.Query().Count(ctx); cnt != 1 {
			tt.Error("perm was not created.")
		}
	})

	t.Run("sync by updating the repo and the perm", func(tt *testing.T) {
		client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
		defer client.Close()

		ctx := context.Background()

		u, err := client.User.Create().
			SetID("1").
			SetLogin("octocat").
			SetToken("access_token").
			SetRefresh("refresh_token").
			SetExpiry(time.Time{}).
			Save(ctx)
		if err != nil {
			tt.Error("failed to create a new user")
			return
		}

		_, err = client.Repo.Create().
			SetID("1").
			SetNamespace("octocat").
			SetName("HelloWorld").
			Save(ctx)
		if err != nil {
			tt.Error("failed to create a new repo")
			return
		}

		_, err = client.Perm.Create().
			SetUserID("1").
			SetRepoID("1").
			Save(ctx)
		if err != nil {
			tt.Error("failed to create a new perm")
			return
		}

		p := &ent.Perm{
			RepoPerm: perm.RepoPermWrite,
			Edges: ent.PermEdges{
				Repo: &ent.Repo{
					Namespace: "octocat",
					Name:      "HelloWorld",
				},
			},
		}

		s := NewStore(client)

		err = s.SyncPerm(ctx, u, p, time.Now())
		if err != nil {
			tt.Errorf("failed to sync perm: %s", err)
		}
	})
}
