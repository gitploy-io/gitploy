package store

import (
	"context"
	"testing"
	"time"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/enttest"
	"github.com/hanjunlee/gitploy/ent/migrate"
	"github.com/hanjunlee/gitploy/ent/perm"
)

func TestStore_DeletePermsOfUserLessThanUpdatedAt(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1",
		enttest.WithMigrateOptions(migrate.WithForeignKeys(false)),
	)
	defer client.Close()

	const (
		u1 = "1"
		u2 = "2"
		r  = "1"
	)

	nor := time.Now()

	t.Log("Insert staled perms")
	client.Perm.
		Create().
		SetRepoPerm(perm.RepoPermWrite).
		SetUserID(u1).
		SetRepoID(r).
		SetUpdatedAt(nor.Add(-1 * time.Hour)).
		SaveX(context.Background())

	client.Perm.
		Create().
		SetRepoPerm(perm.RepoPermWrite).
		SetUserID(u2).
		SetRepoID(r).
		SetUpdatedAt(nor.Add(-1 * time.Hour)).
		SaveX(context.Background())

	t.Log("Insert new perms")
	client.Perm.
		Create().
		SetRepoPerm(perm.RepoPermWrite).
		SetUserID(u1).
		SetRepoID(r).
		SetUpdatedAt(nor.Add(time.Hour)).
		SaveX(context.Background())

	t.Run("Delete staled perms.", func(t *testing.T) {
		s := NewStore(client)

		cnt, err := s.DeletePermsOfUserLessThanUpdatedAt(context.Background(), &ent.User{ID: u1}, nor)
		if err != nil {
			t.Fatalf("DeletePermsOfUserLessThanUpdatedAt returns an error: %s", err)
		}

		expected := 1
		if cnt != expected {
			t.Fatalf("DeletePermsOfUserLessThanUpdatedAt = %v: %v", cnt, expected)
		}
	})
}
