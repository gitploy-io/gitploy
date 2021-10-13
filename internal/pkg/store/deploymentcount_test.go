package store

import (
	"context"
	"testing"
	"time"

	"github.com/gitploy-io/gitploy/ent/enttest"
	"github.com/gitploy-io/gitploy/ent/migrate"
)

func TestStore_ListDeploymentStatisticssGreaterThanTime(t *testing.T) {
	ctx := context.Background()

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1",
		enttest.WithMigrateOptions(migrate.WithForeignKeys(false)),
	)
	defer client.Close()

	tm := time.Now()

	client.DeploymentStatistics.
		Create().
		SetNamespace("octocat").
		SetName("Hello").
		SetEnv("dev").
		SetUpdatedAt(tm.Add(-time.Hour)).
		SaveX(ctx)

	client.DeploymentStatistics.
		Create().
		SetNamespace("octocat").
		SetName("Hello").
		SetEnv("prod").
		SetUpdatedAt(tm.Add(time.Hour)).
		SaveX(ctx)

	s := NewStore(client)

	dcs, err := s.ListDeploymentStatisticssGreaterThanTime(ctx, tm)
	if err != nil {
		t.Fatalf("ListDeploymentStatisticssGreaterThanTime returns an error: %s", err)
	}

	expected := 1
	if len(dcs) != expected {
		t.Fatalf("ListDeploymentStatisticssGreaterThanTime = %v, wanted %v", len(dcs), expected)
	}
}
