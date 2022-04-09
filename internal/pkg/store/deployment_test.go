package store

import (
	"context"
	"testing"
	"time"

	i "github.com/gitploy-io/gitploy/internal/interactor"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/ent/deployment"
	"github.com/gitploy-io/gitploy/model/ent/enttest"
	"github.com/gitploy-io/gitploy/model/ent/migrate"
	"github.com/gitploy-io/gitploy/pkg/e"

	_ "github.com/mattn/go-sqlite3"
)

func TestStore_SearchDeployments(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1",
		enttest.WithMigrateOptions(migrate.WithForeignKeys(false)),
	)
	defer client.Close()

	ctx := context.Background()

	const (
		u1 = 1
		u2 = 2
		r1 = 1
		r2 = 2
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
		SetEnv("local").
		SetUserID(u1).
		SetRepoID(r1).
		SaveX(ctx)

	client.Deployment.Create().
		SetType(deployment.TypeBranch).
		SetNumber(2).
		SetRef("main").
		SetEnv("dev").
		SetUserID(u2).
		SetRepoID(r1).
		SaveX(ctx)

	client.Deployment.Create().
		SetType(deployment.TypeBranch).
		SetNumber(3).
		SetRef("main").
		SetEnv("dev").
		SetUserID(u2).
		SetRepoID(r1).
		SetStatus(deployment.StatusCreated).
		SaveX(ctx)

	t.Run("u1 searchs all deployments of r1.", func(t *testing.T) {
		const (
			owned   = false
			page    = 1
			perPage = 30
		)

		store := NewStore(client)

		res, err := store.SearchDeploymentsOfUser(ctx,
			&ent.User{ID: u1},
			&i.SearchDeploymentsOfUserOptions{
				ListOptions: i.ListOptions{Page: page, PerPage: perPage},
				Statuses:    []deployment.Status{},
				Owned:       owned,
				From:        time.Now().UTC().Add(-time.Minute),
				To:          time.Now().UTC(),
			})
		if err != nil {
			t.Fatalf("SearchDeployments return an error: %s", err)
		}

		expected := 3
		if len(res) != expected {
			t.Fatalf("SearchDeployments = %v, wanted %v", res, expected)
		}
	})

	t.Run("u1 searchs waiting deployments of r1.", func(t *testing.T) {
		const (
			owned   = false
			page    = 1
			perPage = 30
		)

		store := NewStore(client)

		res, err := store.SearchDeploymentsOfUser(ctx,
			&ent.User{ID: u1},
			&i.SearchDeploymentsOfUserOptions{
				ListOptions: i.ListOptions{Page: page, PerPage: perPage},
				Statuses:    []deployment.Status{deployment.StatusWaiting},
				Owned:       owned,
				From:        time.Now().UTC().Add(-time.Minute),
				To:          time.Now().UTC(),
			})
		if err != nil {
			t.Fatalf("SearchDeployments return an error: %s", err)
		}

		expected := 2
		if len(res) != expected {
			t.Fatalf("SearchDeployments = %v, wanted %v", len(res), expected)
		}
	})

	t.Run("u1 searchs owned deployments of r1.", func(t *testing.T) {
		const (
			owned   = true
			page    = 1
			perPage = 30
		)

		store := NewStore(client)

		res, err := store.SearchDeploymentsOfUser(ctx,
			&ent.User{ID: u1},
			&i.SearchDeploymentsOfUserOptions{
				ListOptions: i.ListOptions{Page: page, PerPage: perPage},
				Statuses:    []deployment.Status{},
				Owned:       owned,
				From:        time.Now().UTC().Add(-time.Minute),
				To:          time.Now().UTC(),
			})
		if err != nil {
			t.Fatalf("SearchDeployments return an error: %s", err)
		}

		expected := 1
		if len(res) != expected {
			t.Fatalf("SearchDeployments = %v, wanted %v", res, expected)
		}
	})
}

func TestStore_ListDeploymentsOfRepo(t *testing.T) {
	const (
		r1 = 1
		u1 = 1
		u2 = 2
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
		ds, err := s.ListDeploymentsOfRepo(ctx, &ent.Repo{ID: r1}, &i.ListDeploymentsOfRepoOptions{
			ListOptions: i.ListOptions{Page: 1, PerPage: 100},
			Env:         "",
			Status:      "",
		})
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
		ds, err := s.ListDeploymentsOfRepo(ctx, &ent.Repo{ID: r1}, &i.ListDeploymentsOfRepoOptions{
			ListOptions: i.ListOptions{Page: 1, PerPage: 100},
			Env:         "local",
			Status:      "",
		})
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
		ds, err := s.ListDeploymentsOfRepo(ctx, &ent.Repo{ID: r1}, &i.ListDeploymentsOfRepoOptions{
			ListOptions: i.ListOptions{Page: 1, PerPage: 100},
			Env:         "",
			Status:      "created",
		})
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
		ds, err := s.ListDeploymentsOfRepo(ctx, &ent.Repo{ID: r1}, &i.ListDeploymentsOfRepoOptions{
			ListOptions: i.ListOptions{Page: 1, PerPage: 100},
			Env:         "local",
			Status:      "created",
		})
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

		number, err := s.GetNextDeploymentNumberOfRepo(ctx, &ent.Repo{ID: 1})
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
			SetUserID(1).
			SetRepoID(1).
			SetStatus(deployment.StatusCreated).
			SaveX(ctx)

		client.Deployment.Create().
			SetType(deployment.TypeBranch).
			SetNumber(1).
			SetType("branch").
			SetRef("main").
			SetEnv("prod").
			SetUserID(1).
			SetRepoID(2).
			SetStatus(deployment.StatusCreated).
			SaveX(ctx)

		s := NewStore(client)

		number, err := s.GetNextDeploymentNumberOfRepo(ctx, &ent.Repo{ID: 1})
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

func TestStore_FindPrevRunningDeployment(t *testing.T) {
	t.Run("Returns an not_found error if there's no deployment.", func(t *testing.T) {
		client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1",
			enttest.WithMigrateOptions(migrate.WithForeignKeys(false)),
		)
		defer client.Close()

		client.Deployment.Create().
			SetType(deployment.TypeBranch).
			SetNumber(1).
			SetType("branch").
			SetRef("main").
			SetEnv("prod").
			SetStatus(deployment.StatusSuccess).
			SetCreatedAt(time.Now().Add(-1 * time.Hour)).
			SetRepoID(1).
			SetUserID(1).
			SaveX(context.Background())

		s := NewStore(client)

		_, err := s.FindPrevRunningDeployment(context.Background(), &ent.Deployment{
			Env:    "prod",
			RepoID: 1,
		})
		if !e.HasErrorCode(err, e.ErrorCodeEntityNotFound) {
			t.Fatalf("FindPrevRunningDeployment error != NotFoundError: %s", err)
		}
	})

	t.Run("Return the latest running deployment.", func(t *testing.T) {
		client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1",
			enttest.WithMigrateOptions(migrate.WithForeignKeys(false)),
		)
		defer client.Close()

		client.Deployment.Create().
			SetType(deployment.TypeBranch).
			SetNumber(1).
			SetType("branch").
			SetRef("main").
			SetEnv("prod").
			SetStatus(deployment.StatusRunning).
			SetCreatedAt(time.Now().Add(-1 * time.Hour)).
			SetRepoID(1).
			SetUserID(1).
			SaveX(context.Background())

		s := NewStore(client)

		_, err := s.FindPrevRunningDeployment(context.Background(), &ent.Deployment{
			Env:    "prod",
			RepoID: 1,
		})
		if err != nil {
			t.Fatalf("FindPrevRunningDeployment returns an error: %s", err)
		}
	})
}

func TestStore_FindPrevSuccessDeployment(t *testing.T) {
	ca := time.Now()

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1",
		enttest.WithMigrateOptions(migrate.WithForeignKeys(false)),
	)
	defer client.Close()

	ctx := context.Background()

	first := client.Deployment.Create().
		SetType(deployment.TypeBranch).
		SetNumber(1).
		SetType("branch").
		SetRef("main").
		SetEnv("prod").
		SetStatus(deployment.StatusSuccess).
		SetCreatedAt(ca.Add(-2 * time.Hour)).
		SetUserID(1).
		SetRepoID(1).
		SaveX(ctx)

	client.Deployment.Create().
		SetType(deployment.TypeBranch).
		SetNumber(2).
		SetType("branch").
		SetRef("main").
		SetEnv("prod").
		SetStatus(deployment.StatusSuccess).
		SetCreatedAt(ca.Add(-time.Hour)).
		SetUserID(1).
		SetRepoID(1).
		SaveX(ctx)

	latest := client.Deployment.Create().
		SetType(deployment.TypeBranch).
		SetNumber(3).
		SetType("branch").
		SetRef("main").
		SetEnv("prod").
		SetStatus(deployment.StatusSuccess).
		SetCreatedAt(ca).
		SetUserID(1).
		SetRepoID(1).
		SaveX(ctx)

	t.Run("First deployment returns not found error.", func(t *testing.T) {
		s := NewStore(client)

		_, err := s.FindPrevSuccessDeployment(ctx, first)
		if !ent.IsNotFound(err) {
			t.Fatalf("FindPrevSuccessDeployment does not return NotFoundError: %s", err)
		}
	})

	t.Run("Return the latest succeed deployment.", func(t *testing.T) {
		s := NewStore(client)

		d, err := s.FindPrevSuccessDeployment(ctx, latest)
		if err != nil {
			t.Fatalf("FindPrevSuccessDeployment returns an error: %s", err)
		}

		expected := 2
		if d.ID != expected {
			t.Fatalf("FindPrevSuccessDeployment = %v, wanted %v", d.ID, expected)
		}
	})
}

func TestStore_UpdateDeployment(t *testing.T) {
	t.Run("Update the status of deployment.", func(t *testing.T) {
		ctx := context.Background()

		client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1",
			enttest.WithMigrateOptions(migrate.WithForeignKeys(false)),
		)
		defer client.Close()

		d := client.Deployment.Create().
			SetType(deployment.TypeBranch).
			SetNumber(1).
			SetType("branch").
			SetRef("main").
			SetEnv("prod").
			SetStatus(deployment.StatusCreated).
			SetUserID(1).
			SetRepoID(1).
			SaveX(ctx)

		s := NewStore(client)

		d.Status = deployment.StatusRunning

		d, err := s.UpdateDeployment(ctx, d)
		if err != nil {
			t.Fatalf("UpdateDeployment returns an error: %s", err)
		}

		expected := deployment.StatusRunning
		if d.Status != expected {
			t.Fatalf("UpdateDeployment = %v, wanted %v", d.Status, expected)
		}
	})
}
