package store

import (
	"context"
	"testing"

	i "github.com/gitploy-io/gitploy/internal/interactor"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/ent/enttest"
	"github.com/gitploy-io/gitploy/model/ent/migrate"
	"github.com/gitploy-io/gitploy/model/extent"
	"github.com/gitploy-io/gitploy/pkg/e"
)

func TestStore_ListReposOfUser(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1",
		enttest.WithMigrateOptions(migrate.WithForeignKeys(false)),
	)
	defer client.Close()

	ctx := context.Background()

	const (
		u1 = 1
	)

	r1 := client.Repo.
		Create().
		SetID(1).
		SetNamespace("octocat").
		SetName("Hello").
		SetDescription("").
		SaveX(ctx)

	r2 := client.Repo.
		Create().
		SetID(2).
		SetNamespace("octocat").
		SetName("World").
		SetDescription("").
		SaveX(ctx)

	client.Repo.
		Create().
		SetID(3).
		SetNamespace("coco").
		SetName("Bye").
		SetDescription("").
		SaveX(ctx)

	t.Log("Create permissions for r1, r2.")
	client.Perm.
		Create().
		SetUserID(u1).
		SetRepoID(r1.ID).
		SaveX(ctx)

	client.Perm.
		Create().
		SetUserID(u1).
		SetRepoID(r2.ID).
		SaveX(ctx)

	t.Run("List all repositories.", func(t *testing.T) {
		s := NewStore(client)

		rs, err := s.ListReposOfUser(ctx, &ent.User{ID: u1}, &i.ListReposOfUserOptions{
			ListOptions: i.ListOptions{Page: 1, PerPage: 30},
			Query:       "",
			Sorted:      false,
		})
		if err != nil {
			t.Fatalf("ListReposOfUser returns an error: %s", err)
		}

		expected := 2
		if len(rs) != 2 {
			t.Fatalf("ListReposOfUser = %v: %v", len(rs), expected)
		}
	})

	t.Run("Search by the query.", func(t *testing.T) {
		s := NewStore(client)

		rs, err := s.ListReposOfUser(ctx, &ent.User{ID: u1}, &i.ListReposOfUserOptions{
			ListOptions: i.ListOptions{Page: 1, PerPage: 30},
			Query:       "octocat",
			Sorted:      false,
		})
		if err != nil {
			t.Fatalf("ListReposOfUser returns an error: %s", err)
		}

		expected := 2
		if len(rs) != expected {
			t.Fatalf("ListReposOfUser = %v: %v", len(rs), expected)
		}
	})
}

func TestStore_FindRepoOfUserByID(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1",
		enttest.WithMigrateOptions(migrate.WithForeignKeys(false)),
	)
	defer client.Close()

	ctx := context.Background()

	const (
		u1 = 1
	)

	r1 := client.Repo.
		Create().
		SetID(1).
		SetNamespace("octocat").
		SetName("Hello").
		SetDescription("").
		SaveX(ctx)

	r2 := client.Repo.
		Create().
		SetID(2).
		SetNamespace("octocat").
		SetName("World").
		SetDescription("").
		SaveX(ctx)

	client.Repo.
		Create().
		SetID(3).
		SetNamespace("coco").
		SetName("Bye").
		SetDescription("").
		SaveX(ctx)

	t.Log("Create permissions for r1, r2.")
	client.Perm.
		Create().
		SetUserID(u1).
		SetRepoID(r1.ID).
		SaveX(ctx)

	client.Perm.
		Create().
		SetUserID(u1).
		SetRepoID(r2.ID).
		SaveX(ctx)

	t.Run("Find the repo by ID", func(t *testing.T) {
		s := NewStore(client)

		var (
			repo *ent.Repo
			err  error
		)

		if repo, err = s.FindRepoOfUserByID(ctx, &ent.User{ID: u1}, 1); err != nil {
			t.Fatalf("FindRepoOfUserByID returns an error: %s", err)
		}

		if repo.ID != 1 {
			t.Fatalf("FindRepoOfUserByID = %v, wanted %v", repo.ID, 1)
		}
	})

	t.Run("Can not find unauthorized repo by ID", func(t *testing.T) {
		s := NewStore(client)

		var (
			err error
		)

		if _, err = s.FindRepoOfUserByID(ctx, &ent.User{ID: u1}, 3); !ent.IsNotFound(err) {
			t.Fatalf("FindRepoOfUserByID didn't returns an NotFoundError: %s", err)
		}

	})
}

func TestStore_SyncRepo(t *testing.T) {
	t.Run("Create a new repo.", func(t *testing.T) {
		client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1",
			enttest.WithMigrateOptions(migrate.WithForeignKeys(false)),
		)
		defer client.Close()

		s := NewStore(client)
		if _, err := s.SyncRepo(context.Background(), &extent.RemoteRepo{
			ID:          1,
			Namespace:   "octocat",
			Name:        "HelloWorld",
			Description: "nothing",
			Perm:        extent.RemoteRepoPermAdmin,
		}); err != nil {
			t.Fatalf("SyncRepo return an error: %s", err)
		}
	})
}

func TestStore_UpdateRepo(t *testing.T) {
	t.Run("Update the repository name.", func(t *testing.T) {
		client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1",
			enttest.WithMigrateOptions(migrate.WithForeignKeys(false)),
		)
		defer client.Close()

		repo := client.Repo.Create().
			SetNamespace("gitploy-io").
			SetName("gitploy").
			SetDescription("").
			SaveX(context.Background())

		s := NewStore(client)

		// Replace values
		repo.Name = "gitploy-next"
		repo.ConfigPath = "deploy-next.yml"

		var (
			ret *ent.Repo
			err error
		)
		ret, err = s.UpdateRepo(context.Background(), repo)
		if err != nil {
			t.Fatalf("UpdateRepo return an error: %s", err)
		}

		if repo.Name != "gitploy-next" ||
			repo.ConfigPath != "deploy-next.yml" {
			t.Fatalf("UpdateRepo = %s, wanted %s", repo, ret)
		}
	})

	t.Run("Return an error if the same repository name exists.", func(t *testing.T) {
		client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1",
			enttest.WithMigrateOptions(migrate.WithForeignKeys(true)),
		)
		defer client.Close()

		client.Repo.Create().
			SetNamespace("gitploy-io").
			SetName("gitploy-next").
			SetDescription("").
			SaveX(context.Background())

		repo := client.Repo.Create().
			SetNamespace("gitploy-io").
			SetName("gitploy").
			SetDescription("").
			SaveX(context.Background())

		s := NewStore(client)

		repo.Name = "gitploy-next"

		var (
			err error
		)
		repo, err = s.UpdateRepo(context.Background(), repo)
		if !e.HasErrorCode(err, e.ErrorRepoUniqueName) {
			t.Fatalf("UpdateRepo doesn't return the ErrorRepoUniqueName error: %s", err)
		}
	})
}

func TestStore_Activate(t *testing.T) {
	t.Run("Update webhook ID and owner ID.", func(t *testing.T) {
		client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1",
			enttest.WithMigrateOptions(migrate.WithForeignKeys(false)),
		)
		defer client.Close()

		ctx := context.Background()

		r := client.Repo.
			Create().
			SetID(1).
			SetNamespace("octocat").
			SetName("Hello").
			SetDescription("").
			SaveX(ctx)

		r.WebhookID = 1
		r.OwnerID = 1

		s := NewStore(client)
		r, err := s.Activate(ctx, r)
		if err != nil {
			t.Fatalf("Activate returns an error: %s", err)
		}

		if !r.Active {
			t.Fatalf("Activate failed: %v", r)
		}
	})
}
