package store

import (
	"context"
	"testing"
	"time"

	"github.com/gitploy-io/gitploy/ent/enttest"
	"github.com/gitploy-io/gitploy/ent/migrate"
)

func TestStore_ListExpiredLocksLessThanTime(t *testing.T) {
	t.Run("Returns expired locks.", func(t *testing.T) {
		ctx := context.Background()

		client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1",
			enttest.WithMigrateOptions(migrate.WithForeignKeys(false)),
		)
		defer client.Close()

		tm := time.Date(2021, 10, 17, 0, 0, 0, 0, time.UTC)

		client.Lock.
			Create().
			SetEnv("dev").
			SetExpiredAt(tm.Add(-time.Hour)).
			SetRepoID(1).
			SetUserID(1).
			SaveX(ctx)

		client.Lock.
			Create().
			SetEnv("staging").
			SetRepoID(1).
			SetUserID(1).
			SaveX(ctx)

		client.Lock.
			Create().
			SetEnv("production").
			SetRepoID(1).
			SetUserID(1).
			SaveX(ctx)

		s := NewStore(client)

		ls, err := s.ListExpiredLocksLessThanTime(ctx, tm)
		if err != nil {
			t.Fatalf("ListExpiredLocksLessThanTime returns an error: %s", err)
		}

		expected := 1
		if len(ls) != expected {
			t.Log("The zero value of time must to be skipped.")
			t.Fatalf("len(ListExpiredLocksLessThanTime) = %v: want %v", len(ls), expected)
		}
	})
}
