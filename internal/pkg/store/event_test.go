package store

import (
	"context"
	"testing"

	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/ent/enttest"
	"github.com/gitploy-io/gitploy/model/ent/event"
	"github.com/gitploy-io/gitploy/model/ent/migrate"
)

func TestStore_CreateEvent(t *testing.T) {

	t.Run("Create a new deployment_status event", func(t *testing.T) {
		ctx := context.Background()

		client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1",
			enttest.WithMigrateOptions(migrate.WithForeignKeys(false)),
		)
		defer client.Close()

		s := NewStore(client)

		_, err := s.CreateEvent(ctx, &ent.Event{
			Kind:               event.KindDeploymentStatus,
			Type:               event.TypeCreated,
			DeploymentStatusID: 1,
		})
		if err != nil {
			t.Fatalf("CreateEvent returns an error: %s", err)
		}
	})
}
