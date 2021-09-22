package store

import (
	"context"
	"testing"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/enttest"
	"github.com/gitploy-io/gitploy/ent/event"
	"github.com/gitploy-io/gitploy/ent/migrate"
)

func TestStore_CreateEvent(t *testing.T) {

	t.Run("Create a new deleted event", func(t *testing.T) {
		ctx := context.Background()

		client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1",
			enttest.WithMigrateOptions(migrate.WithForeignKeys(false)),
		)
		defer client.Close()

		s := NewStore(client)

		e, err := s.CreateEvent(ctx, &ent.Event{
			Kind:      event.KindApproval,
			Type:      event.TypeDeleted,
			DeletedID: 1,
		})
		if err != nil {
			t.Fatalf("CreateEvent returns an error: %s", err)
		}

		if e.DeletedID != 1 {
			t.Fatalf("event.DeletedID = %v, wanted %v", e.DeletedID, 1)
		}
	})
}
