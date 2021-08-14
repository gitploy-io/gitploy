package store

import (
	"context"
	"testing"
	"time"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/approval"
	"github.com/hanjunlee/gitploy/ent/enttest"
	"github.com/hanjunlee/gitploy/ent/migrate"
)

func TestStore_SearchApprovals(t *testing.T) {
	client := enttest.Open(t,
		"sqlite3",
		"file:ent?mode=memory&cache=shared&_fk=1",
		enttest.WithOptions(ent.Debug()),
		enttest.WithMigrateOptions(migrate.WithForeignKeys(false)),
	)
	defer client.Close()

	ctx := context.Background()

	const (
		u1 = "1"
		d1 = 1
		d2 = 2
		d3 = 3
	)

	client.Approval.
		Create().
		SetStatus(approval.StatusApproved).
		SetUserID(u1).
		SetDeploymentID(d1).
		SaveX(ctx)

	client.Approval.
		Create().
		SetStatus(approval.StatusPending).
		SetUserID(u1).
		SetDeploymentID(d2).
		SaveX(ctx)

	client.Approval.
		Create().
		SetStatus(approval.StatusPending).
		SetUserID(u1).
		SetDeploymentID(d3).
		SaveX(ctx)

	t.Run("u1 searchs requested approvals of the deployment.", func(t *testing.T) {
		const (
			owned   = false
			page    = 1
			perPage = 2
		)

		store := NewStore(client)

		res, err := store.SearchApprovals(ctx,
			&ent.User{ID: u1},
			[]approval.Status{
				approval.StatusPending,
			},
			time.Now().Add(-time.Minute),
			time.Now(),
			page,
			perPage,
		)
		if err != nil {
			t.Fatalf("SearchApprovals return an error: %s", err)
			t.FailNow()
		}

		expected := 2
		if len(res) != expected {
			t.Fatalf("SearchApprovals = %v, wanted %v", res, expected)
			t.FailNow()
		}
	})

}