package store

import (
	"context"
	"testing"
	"time"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/enttest"
	"github.com/gitploy-io/gitploy/ent/migrate"
	"github.com/gitploy-io/gitploy/ent/perm"
)

func TestStore_DeletePermsOfUserLessThanSyncedAt(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1",
		enttest.WithMigrateOptions(migrate.WithForeignKeys(false)),
	)
	defer client.Close()

	const (
		u1 = 1
		u2 = 2
		r1 = 1
		r2 = 2
	)

	nor := time.Now()

	t.Log("Insert staled perms")
	client.Perm.
		Create().
		SetRepoPerm(perm.RepoPermWrite).
		SetSyncedAt(nor.Add(-1 * time.Hour)).
		SetUserID(u1).
		SetRepoID(r1).
		SaveX(context.Background())

	client.Perm.
		Create().
		SetRepoPerm(perm.RepoPermWrite).
		SetSyncedAt(nor.Add(-1 * time.Hour)).
		SetUserID(u1).
		SetRepoID(r2).
		SaveX(context.Background())

	client.Perm.
		Create().
		SetRepoPerm(perm.RepoPermWrite).
		SetSyncedAt(nor.Add(-1 * time.Hour)).
		SetUserID(u2).
		SetRepoID(r1).
		SaveX(context.Background())

	t.Log("Insert new perms")
	client.Perm.
		Create().
		SetRepoPerm(perm.RepoPermWrite).
		SetSyncedAt(nor.Add(time.Hour)).
		SetUserID(u1).
		SetRepoID(r1).
		SaveX(context.Background())

	client.Perm.
		Create().
		SetRepoPerm(perm.RepoPermWrite).
		SetSyncedAt(nor.Add(time.Hour)).
		SetUserID(u2).
		SetRepoID(r1).
		SaveX(context.Background())

	t.Run("Delete staled perms.", func(t *testing.T) {
		s := NewStore(client)

		cnt, err := s.DeletePermsOfUserLessThanSyncedAt(context.Background(), &ent.User{ID: u1}, nor)
		if err != nil {
			t.Fatalf("DeletePermsOfUserLessThanSyncedAt returns an error: %s", err)
		}

		expected := 2
		if cnt != expected {
			t.Fatalf("DeletePermsOfUserLessThanSyncedAt = %v: %v", cnt, expected)
		}
	})
}
