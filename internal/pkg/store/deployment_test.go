package store

import (
	"context"
	"testing"

	"github.com/hanjunlee/gitploy/ent/deployment"
	"github.com/hanjunlee/gitploy/ent/enttest"

	_ "github.com/mattn/go-sqlite3"
)

func TestStore_ListDeployments(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	ctx := context.Background()
	r, err := client.Repo.Create().
		SetID("1").
		SetNamespace("octocat").
		SetName("HelloWorld").
		Save(ctx)
	if err != nil {
		t.Error("failed to create a new repo")
		return
	}

	_, err = client.Deployment.Create().
		SetType(deployment.TypeBranch).
		SetRef("main").
		SetEnv("local").
		SetRepoID("1").
		SetStatus(deployment.StatusCreated).
		Save(ctx)
	if err != nil {
		t.Error("failed to create a new repo")
		return
	}

	_, err = client.Deployment.Create().
		SetType(deployment.TypeBranch).
		SetRef("main").
		SetEnv("dev").
		SetRepoID("1").
		Save(ctx)
	if err != nil {
		t.Error("failed to create a new repo")
		return
	}

	_, err = client.Deployment.Create().
		SetType(deployment.TypeBranch).
		SetRef("branch").
		SetEnv("staging").
		SetRepoID("1").
		Save(ctx)
	if err != nil {
		t.Error("failed to create a new repo")
		return
	}

	s := NewStore(client)

	t.Run("list all deployments", func(tt *testing.T) {
		ds, err := s.ListDeployments(ctx, r, "", "", 1, 100)
		if err != nil {
			tt.Errorf("failed to list deployments: %s", err)
			return
		}

		e := 3
		if len(ds) != e {
			tt.Errorf("ListDeployments = len(%v), expected len(%v)", len(ds), e)
		}
	})

	t.Run("list env=local deployments", func(tt *testing.T) {
		ds, err := s.ListDeployments(ctx, r, "local", "", 1, 100)
		if err != nil {
			tt.Errorf("failed to list deployments: %s", err)
			return
		}

		e := 1
		if len(ds) != e {
			tt.Errorf("ListDeployments = len(%v), expected len(%v)", len(ds), e)
		}
	})

	t.Run("list status=created deployments", func(tt *testing.T) {
		ds, err := s.ListDeployments(ctx, r, "", "created", 1, 100)
		if err != nil {
			tt.Errorf("failed to list deployments: %s", err)
			return
		}

		e := 1
		if len(ds) != e {
			tt.Errorf("ListDeployments = len(%v), expected len(%v)", len(ds), e)
		}
	})

	t.Run("list env=local&status=created deployments", func(tt *testing.T) {
		ds, err := s.ListDeployments(ctx, r, "local", "created", 1, 100)
		if err != nil {
			tt.Errorf("failed to list deployments: %s", err)
			return
		}

		e := 1
		if len(ds) != e {
			tt.Errorf("ListDeployments = len(%v), expected len(%v)", len(ds), e)
		}
	})
}
