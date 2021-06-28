package store

import (
	"context"
	"testing"
	"time"

	"github.com/hanjunlee/gitploy/ent"
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
		t.Errorf("failed to create a new repo: %s", err)
		return
	}

	_, err = client.User.Create().
		SetID("1").
		SetLogin("octocat").
		SetToken("").
		SetRefresh("").
		SetExpiry(time.Time{}).
		Save(ctx)
	if err != nil {
		t.Errorf("failed to create a new user: %s", err)
		return
	}

	_, err = client.Deployment.Create().
		SetType(deployment.TypeBranch).
		SetNumber(1).
		SetRef("main").
		SetEnv("local").
		SetUserID("1").
		SetRepoID("1").
		SetStatus(deployment.StatusCreated).
		Save(ctx)
	if err != nil {
		t.Errorf("failed to create a new repo: %s", err)
		return
	}

	_, err = client.Deployment.Create().
		SetType(deployment.TypeBranch).
		SetNumber(2).
		SetRef("main").
		SetEnv("dev").
		SetUserID("1").
		SetRepoID("1").
		Save(ctx)
	if err != nil {
		t.Errorf("failed to create a new repo: %s", err)
		return
	}

	_, err = client.Deployment.Create().
		SetType(deployment.TypeBranch).
		SetNumber(3).
		SetRef("branch").
		SetEnv("staging").
		SetUserID("1").
		SetRepoID("1").
		Save(ctx)
	if err != nil {
		t.Errorf("failed to create a new repo: %s", err)
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

func TestStore_CreateDeployment(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	ctx := context.Background()
	r, err := client.Repo.Create().
		SetID("1").
		SetNamespace("octocat").
		SetName("HelloWorld").
		Save(ctx)
	if err != nil {
		t.Errorf("failed to create a new repo: %s", err)
		return
	}

	u, err := client.User.Create().
		SetID("1").
		SetLogin("octocat").
		SetToken("").
		SetRefresh("").
		SetExpiry(time.Time{}).
		Save(ctx)
	if err != nil {
		t.Errorf("failed to create a new user: %s", err)
		return
	}

	_, err = client.Deployment.Create().
		SetType(deployment.TypeBranch).
		SetNumber(1).
		SetType("branch").
		SetRef("main").
		SetEnv("local").
		SetUserID("1").
		SetRepoID("1").
		SetStatus(deployment.StatusCreated).
		Save(ctx)
	if err != nil {
		t.Errorf("failed to create a new repo: %s", err)
		return
	}

	s := NewStore(client)

	t.Run("The next number is 2", func(tt *testing.T) {
		d, err := s.CreateDeployment(ctx, u, r, &ent.Deployment{
			Type: "branch",
			Ref:  "main",
			Env:  "prod",
		})
		if err != nil {
			t.Errorf("CreateDeployment return error: %s", err)
		}

		if d.Number != 2 {
			t.Errorf("CreateDeployment = %d, 2", d.Number)
		}
	})
}
