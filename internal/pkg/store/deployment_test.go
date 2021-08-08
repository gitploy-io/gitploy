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

// Set up the repo and the user for deployment test.
func setupDeploymentTestClient(client *ent.Client, r *ent.Repo, u *ent.User) *ent.Client {
	ctx := context.Background()

	client.Repo.
		Create().
		SetID(r.ID).
		SetNamespace(r.Namespace).
		SetName(r.Name).
		SaveX(ctx)

	client.User.
		Create().
		SetID(u.ID).
		SetLogin(u.Login).
		SetToken(u.Token).
		SetRefresh(u.Refresh).
		SetExpiry(time.Time{}).
		SaveX(ctx)

	return client
}

func TestStore_ListDeploymentsOfRepo(t *testing.T) {
	r := &ent.Repo{
		ID: "1",
	}

	u := &ent.User{
		ID:    "1",
		Login: "octocat",
	}

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	client = setupDeploymentTestClient(client, r, u)

	ctx := context.Background()

	client.Deployment.Create().
		SetType(deployment.TypeBranch).
		SetNumber(1).
		SetRef("main").
		SetSha("7e555a2").
		SetEnv("local").
		SetUserID(u.ID).
		SetRepoID(r.ID).
		SetStatus(deployment.StatusCreated).
		SaveX(ctx)

	client.Deployment.Create().
		SetType(deployment.TypeBranch).
		SetNumber(2).
		SetRef("main").
		SetSha("a20052a").
		SetEnv("dev").
		SetUserID(u.ID).
		SetRepoID(r.ID).
		SaveX(ctx)

	client.Deployment.Create().
		SetType(deployment.TypeBranch).
		SetNumber(3).
		SetRef("branch").
		SetSha("7e555a2").
		SetEnv("staging").
		SetUserID(u.ID).
		SetRepoID(r.ID).
		SaveX(ctx)

	s := NewStore(client)

	t.Run("list all deployments", func(tt *testing.T) {
		ds, err := s.ListDeploymentsOfRepo(ctx, r, "", "", 1, 100)
		if err != nil {
			tt.Errorf("failed to list deployments: %s", err)
			return
		}

		e := 3
		if len(ds) != e {
			tt.Errorf("ListDeploymentsOfRepo = len(%v), expected len(%v)", len(ds), e)
		}
	})

	t.Run("list env=local deployments", func(tt *testing.T) {
		ds, err := s.ListDeploymentsOfRepo(ctx, r, "local", "", 1, 100)
		if err != nil {
			tt.Errorf("failed to list deployments: %s", err)
			return
		}

		e := 1
		if len(ds) != e {
			tt.Errorf("ListDeploymentsOfRepo = len(%v), expected len(%v)", len(ds), e)
		}
	})

	t.Run("list status=created deployments", func(tt *testing.T) {
		ds, err := s.ListDeploymentsOfRepo(ctx, r, "", "created", 1, 100)
		if err != nil {
			tt.Errorf("failed to list deployments: %s", err)
			return
		}

		e := 1
		if len(ds) != e {
			tt.Errorf("ListDeploymentsOfRepo = len(%v), expected len(%v)", len(ds), e)
		}
	})

	t.Run("list env=local&status=created deployments", func(tt *testing.T) {
		ds, err := s.ListDeploymentsOfRepo(ctx, r, "local", "created", 1, 100)
		if err != nil {
			tt.Errorf("failed to list deployments: %s", err)
			return
		}

		e := 1
		if len(ds) != e {
			tt.Errorf("ListDeploymentsOfRepo = len(%v), expected len(%v)", len(ds), e)
		}
	})
}

func TestStore_GetNextDeploymentNumberOfRepo(t *testing.T) {
	t.Run("Return one when it is the first deployment of the repository.", func(t *testing.T) {
		client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
		defer client.Close()

		ctx := context.Background()

		s := NewStore(client)

		number, err := s.GetNextDeploymentNumberOfRepo(ctx, &ent.Repo{ID: "1"})
		if err != nil {
			t.Fatalf("GetNextDeploymentNumberOfRepo returns an error: %s", err)
			t.FailNow()
		}

		expected := 1
		if number != expected {
			t.Fatalf("GetNextDeploymentNumberOfRepo = %d, want %d", number, expected)
			t.FailNow()
		}
	})

	t.Run("Return two when there is a single deployment.", func(t *testing.T) {
		client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
		defer client.Close()

		ctx := context.Background()

		r1 := client.Repo.Create().
			SetID("1").
			SetNamespace("octocat").
			SetName("HelloWorld").
			SaveX(ctx)

		r2 := client.Repo.Create().
			SetID("2").
			SetNamespace("octocat").
			SetName("GoodBye").
			SaveX(ctx)

		client.User.Create().
			SetID("1").
			SetLogin("octocat").
			SetToken("").
			SetRefresh("").
			SetExpiry(time.Time{}).
			SaveX(ctx)

		client.Deployment.Create().
			SetType(deployment.TypeBranch).
			SetNumber(1).
			SetType("branch").
			SetRef("main").
			SetEnv("local").
			SetUserID("1").
			SetRepoID(r1.ID).
			SetStatus(deployment.StatusCreated).
			SaveX(ctx)

		client.Deployment.Create().
			SetType(deployment.TypeBranch).
			SetNumber(1).
			SetType("branch").
			SetRef("main").
			SetEnv("prod").
			SetUserID("1").
			SetRepoID(r2.ID).
			SetStatus(deployment.StatusCreated).
			SaveX(ctx)

		s := NewStore(client)

		number, err := s.GetNextDeploymentNumberOfRepo(ctx, r1)
		if err != nil {
			t.Fatalf("GetNextDeploymentNumberOfRepo returns an error: %s", err)
			t.FailNow()
		}

		expected := 2
		if number != expected {
			t.Fatalf("GetNextDeploymentNumberOfRepo = %d, want %d", number, expected)
			t.FailNow()
		}
	})
}

func TestStore_FindLatestSuccessfulDeployment(t *testing.T) {
	r := &ent.Repo{
		ID: "1",
	}

	u := &ent.User{
		ID:    "1",
		Login: "octocat",
	}

	t.Run("Return not found error.", func(t *testing.T) {
		client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
		defer client.Close()

		client = setupDeploymentTestClient(client, r, u)

		ctx := context.Background()

		d := client.Deployment.Create().
			SetType(deployment.TypeBranch).
			SetNumber(3).
			SetType("branch").
			SetRef("main").
			SetEnv("prod").
			SetUserID(u.ID).
			SetRepoID(r.ID).
			SetStatus(deployment.StatusCreated).
			SaveX(ctx)

		s := NewStore(client)

		d, err := s.FindLatestSuccessfulDeployment(ctx, d)
		if !ent.IsNotFound(err) {
			t.Fatalf("FindLatestSuccessfulDeployment does not return NotFoundError: %s", err)
			t.FailNow()
		}
	})

	t.Run("Return the latest updated succeed deployment.", func(t *testing.T) {
		client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
		defer client.Close()

		client = setupDeploymentTestClient(client, r, u)

		var (
			ctx = context.Background()
			ca  = time.Now().Add(-2 * time.Hour)
			now = time.Now()
		)

		expected := client.Deployment.Create().
			SetType(deployment.TypeBranch).
			SetNumber(1).
			SetType("branch").
			SetRef("main").
			SetEnv("prod").
			SetUserID(u.ID).
			SetRepoID(r.ID).
			SetStatus(deployment.StatusSuccess).
			SetCreatedAt(ca).
			SetUpdatedAt(now).
			SaveX(ctx)

		client.Deployment.Create().
			SetType(deployment.TypeBranch).
			SetNumber(2).
			SetType("branch").
			SetRef("main").
			SetEnv("prod").
			SetUserID(u.ID).
			SetRepoID(r.ID).
			SetStatus(deployment.StatusSuccess).
			SetCreatedAt(ca).
			SetUpdatedAt(now.Add(-time.Hour)).
			SaveX(ctx)

		d := client.Deployment.Create().
			SetType(deployment.TypeBranch).
			SetNumber(3).
			SetType("branch").
			SetRef("main").
			SetEnv("prod").
			SetUserID(u.ID).
			SetRepoID(r.ID).
			SetStatus(deployment.StatusCreated).
			SaveX(ctx)

		s := NewStore(client)

		d, err := s.FindLatestSuccessfulDeployment(ctx, d)
		if err != nil {
			t.Fatalf("FindLatestSuccessfulDeployment returns an error: %s", err)
			t.FailNow()
		}

		if d.ID != expected.ID {
			t.Fatalf("FindLatestSuccessfulDeployment = %v, wanted %v", d, expected)
			t.FailNow()
		}
	})
}
