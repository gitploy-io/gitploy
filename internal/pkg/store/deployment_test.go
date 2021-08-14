package store

import (
	"context"
	"testing"
	"time"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/deployment"
	"github.com/hanjunlee/gitploy/ent/enttest"
	"github.com/hanjunlee/gitploy/ent/migrate"

	_ "github.com/mattn/go-sqlite3"
)

func TestStore_SearchDeployments(t *testing.T) {

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1",
		enttest.WithMigrateOptions(migrate.WithForeignKeys(false)),
	)
	defer client.Close()

	ctx := context.Background()

	const (
		u1 = "1"
		u2 = "2"
		r1 = "1"
		r2 = "2"
	)

	client.Perm.
		Create().
		SetUserID(u1).
		SetRepoID(r1).
		SaveX(ctx)

	client.Deployment.Create().
		SetType(deployment.TypeBranch).
		SetNumber(1).
		SetRef("main").
		SetSha("7e555a2").
		SetEnv("local").
		SetUserID(u1).
		SetRepoID(r1).
		SaveX(ctx)

	client.Deployment.Create().
		SetType(deployment.TypeBranch).
		SetNumber(2).
		SetRef("main").
		SetSha("a20052a").
		SetEnv("dev").
		SetUserID(u2).
		SetRepoID(r1).
		SaveX(ctx)

	t.Run("u1 searchs all deployments of r1.", func(t *testing.T) {
		const (
			owned   = false
			page    = 1
			perPage = 2
		)

		store := NewStore(client)

		res, err := store.SearchDeployments(ctx,
			&ent.User{ID: u1},
			[]deployment.Status{
				deployment.StatusWaiting,
			},
			owned,
			time.Now().Add(-time.Minute),
			time.Now(),
			page,
			perPage)
		if err != nil {
			t.Fatalf("SearchDeployments return an error: %s", err)
			t.FailNow()
		}

		expected := 2
		if len(res) != expected {
			t.Fatalf("SearchDeployments = %v, wanted %v", res, expected)
			t.FailNow()
		}
	})

	t.Run("u1 searchs owned deployments of r1.", func(t *testing.T) {
		const (
			owned   = true
			page    = 1
			perPage = 2
		)

		store := NewStore(client)

		res, err := store.SearchDeployments(ctx,
			&ent.User{ID: u1},
			[]deployment.Status{
				deployment.StatusWaiting,
			},
			owned,
			time.Now().Add(-time.Minute),
			time.Now(),
			page,
			perPage)
		if err != nil {
			t.Fatalf("SearchDeployments return an error: %s", err)
			t.FailNow()
		}

		expected := 1
		if len(res) != expected {
			t.Fatalf("SearchDeployments = %v, wanted %v", res, expected)
			t.FailNow()
		}
	})
}

func TestStore_ListDeploymentsOfRepo(t *testing.T) {
	const (
		r1 = "1"
		u1 = "1"
		u2 = "2"
	)

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1",
		enttest.WithMigrateOptions(migrate.WithForeignKeys(false)),
	)
	defer client.Close()

	ctx := context.Background()

	client.Deployment.Create().
		SetType(deployment.TypeBranch).
		SetNumber(1).
		SetRef("main").
		SetSha("7e555a2").
		SetEnv("local").
		SetUserID(u1).
		SetRepoID(r1).
		SetStatus(deployment.StatusCreated).
		SaveX(ctx)

	client.Deployment.Create().
		SetType(deployment.TypeBranch).
		SetNumber(2).
		SetRef("main").
		SetSha("a20052a").
		SetEnv("dev").
		SetUserID(u2).
		SetRepoID(r1).
		SaveX(ctx)

	client.Deployment.Create().
		SetType(deployment.TypeBranch).
		SetNumber(3).
		SetRef("branch").
		SetSha("7e555a2").
		SetEnv("staging").
		SetUserID(u1).
		SetRepoID(r1).
		SaveX(ctx)

	s := NewStore(client)

	t.Run("list all deployments", func(tt *testing.T) {
		ds, err := s.ListDeploymentsOfRepo(ctx, &ent.Repo{ID: r1}, "", "", 1, 100)
		if err != nil {
			tt.Errorf("failed to list deployments: %s", err)
			return
		}

		e := 3
		if len(ds) != e {
			tt.Errorf("ListDeploymentsOfRepo = %v, expected %v", len(ds), e)
		}
	})

	t.Run("list env=local deployments", func(tt *testing.T) {
		ds, err := s.ListDeploymentsOfRepo(ctx, &ent.Repo{ID: r1}, "local", "", 1, 100)
		if err != nil {
			tt.Errorf("failed to list deployments: %s", err)
			return
		}

		e := 1
		if len(ds) != e {
			tt.Errorf("ListDeploymentsOfRepo = %v, expected %v", len(ds), e)
		}
	})

	t.Run("list status=created deployments", func(tt *testing.T) {
		ds, err := s.ListDeploymentsOfRepo(ctx, &ent.Repo{ID: r1}, "", "created", 1, 100)
		if err != nil {
			tt.Errorf("failed to list deployments: %s", err)
			return
		}

		e := 1
		if len(ds) != e {
			tt.Errorf("ListDeploymentsOfRepo = %v, expected %v", len(ds), e)
		}
	})

	t.Run("list env=local&status=created deployments", func(tt *testing.T) {
		ds, err := s.ListDeploymentsOfRepo(ctx, &ent.Repo{ID: r1}, "local", "created", 1, 100)
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
		client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1",
			enttest.WithMigrateOptions(migrate.WithForeignKeys(false)),
		)
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
		const (
			u1 = "1"
			r1 = "1"
			r2 = "2"
		)

		client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1",
			enttest.WithMigrateOptions(migrate.WithForeignKeys(false)),
		)
		defer client.Close()

		ctx := context.Background()

		client.Deployment.Create().
			SetType(deployment.TypeBranch).
			SetNumber(1).
			SetType("branch").
			SetRef("main").
			SetEnv("local").
			SetUserID("1").
			SetRepoID(r1).
			SetStatus(deployment.StatusCreated).
			SaveX(ctx)

		client.Deployment.Create().
			SetType(deployment.TypeBranch).
			SetNumber(1).
			SetType("branch").
			SetRef("main").
			SetEnv("prod").
			SetUserID("1").
			SetRepoID(r2).
			SetStatus(deployment.StatusCreated).
			SaveX(ctx)

		s := NewStore(client)

		number, err := s.GetNextDeploymentNumberOfRepo(ctx, &ent.Repo{ID: r1})
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
	const (
		u1 = "1"
		r1 = "1"
	)

	t.Run("Return not found error.", func(t *testing.T) {
		client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1",
			enttest.WithMigrateOptions(migrate.WithForeignKeys(false)),
		)
		defer client.Close()

		ctx := context.Background()

		d := client.Deployment.Create().
			SetType(deployment.TypeBranch).
			SetNumber(3).
			SetType("branch").
			SetRef("main").
			SetEnv("prod").
			SetUserID(u1).
			SetRepoID(r1).
			SetStatus(deployment.StatusCreated).
			SaveX(ctx)

		s := NewStore(client)

		_, err := s.FindLatestSuccessfulDeployment(ctx, d)
		if !ent.IsNotFound(err) {
			t.Fatalf("FindLatestSuccessfulDeployment does not return NotFoundError: %s", err)
			t.FailNow()
		}
	})

	t.Run("Return the latest updated succeed deployment.", func(t *testing.T) {
		client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1",
			enttest.WithMigrateOptions(migrate.WithForeignKeys(false)),
		)
		defer client.Close()

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
			SetUserID(u1).
			SetRepoID(r1).
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
			SetUserID(u1).
			SetRepoID(r1).
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
			SetUserID(u1).
			SetRepoID(r1).
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
