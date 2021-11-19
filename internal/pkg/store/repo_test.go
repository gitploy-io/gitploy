package store

import (
	"context"
	"testing"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/enttest"
	"github.com/gitploy-io/gitploy/ent/migrate"
	"github.com/gitploy-io/gitploy/vo"
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

		rs, err := s.ListReposOfUser(ctx, &ent.User{ID: u1}, "", "", "", false, 1, 30)
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

		rs, err := s.ListReposOfUser(ctx, &ent.User{ID: u1}, "octocat", "", "", false, 1, 30)
		if err != nil {
			t.Fatalf("ListReposOfUser returns an error: %s", err)
		}

		expected := 2
		if len(rs) != expected {
			t.Fatalf("ListReposOfUser = %v: %v", len(rs), expected)
		}
	})

	t.Run("Search by the query and the namespace.", func(t *testing.T) {
		s := NewStore(client)

		rs, err := s.ListReposOfUser(ctx, &ent.User{ID: u1}, "octocat", "coco", "", false, 1, 30)
		if err != nil {
			t.Fatalf("ListReposOfUser returns an error: %s", err)
		}

		expected := 0
		if len(rs) != expected {
			t.Fatalf("ListReposOfUser = %v: %v", len(rs), expected)
		}
	})

	t.Run("Search by the query and the name.", func(t *testing.T) {
		s := NewStore(client)

		rs, err := s.ListReposOfUser(ctx, &ent.User{ID: u1}, "octocat", "", "Hello", false, 1, 30)
		if err != nil {
			t.Fatalf("ListReposOfUser returns an error: %s", err)
		}

		expected := 1
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
		if _, err := s.SyncRepo(context.Background(), &vo.RemoteRepo{
			ID:          1,
			Namespace:   "octocat",
			Name:        "HelloWorld",
			Description: "nothing",
			Perm:        vo.RemoteRepoPermAdmin,
		}); err != nil {
			t.Fatalf("SyncRepo return an error: %s", err)
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
