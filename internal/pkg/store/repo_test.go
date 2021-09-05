package store

import (
	"context"
	"testing"
	"time"

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
		u1 = "1"
	)

	r1 := client.Repo.
		Create().
		SetID("1").
		SetNamespace("octocat").
		SetName("Hello").
		SetDescription("").
		SaveX(ctx)

	r2 := client.Repo.
		Create().
		SetID("2").
		SetNamespace("octocat").
		SetName("World").
		SetDescription("").
		SaveX(ctx)

	client.Repo.
		Create().
		SetID("3").
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

	t.Run("Returns two repos.", func(t *testing.T) {
		s := NewStore(client)

		var (
			rs  []*ent.Repo
			err error
		)

		if rs, err = s.ListReposOfUser(ctx, &ent.User{ID: u1}, "", 1, 30); err != nil {
			t.Fatalf("ListReposOfUser returns an error: %s", err)
		}

		expected := 2
		if len(rs) != 2 {
			t.Fatalf("ListReposOfUser = %v: %v", len(rs), expected)
		}
	})

	t.Run("Returns a single repo when it queries.", func(t *testing.T) {
		s := NewStore(client)

		var (
			rs  []*ent.Repo
			err error
		)

		const (
			query = "Hello"
		)

		if rs, err = s.ListReposOfUser(ctx, &ent.User{ID: u1}, query, 1, 30); err != nil {
			t.Fatalf("ListReposOfUser returns an error: %s", err)
		}

		expected := 1
		if len(rs) != 1 {
			t.Fatalf("ListReposOfUser = %v: %v", len(rs), expected)
		}
	})

}

func TestStore_ListSortedReposOfUser(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1",
		enttest.WithMigrateOptions(migrate.WithForeignKeys(false)),
	)
	defer client.Close()

	ctx := context.Background()

	const (
		u1 = "1"
	)

	r1 := client.Repo.
		Create().
		SetID("1").
		SetNamespace("octocat").
		SetName("Hello").
		SetDescription("").
		SaveX(ctx)

	r2 := client.Repo.
		Create().
		SetID("2").
		SetNamespace("octocat").
		SetName("World").
		SetDescription("").
		SaveX(ctx)

	client.Repo.
		Create().
		SetID("3").
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

	t.Run("Returns repos in inserted-order manager.", func(t *testing.T) {
		s := NewStore(client)

		var (
			rs  []*ent.Repo
			err error
		)

		if rs, err = s.ListSortedReposOfUser(ctx, &ent.User{ID: u1}, "", 1, 30); err != nil {
			t.Fatalf("ListSortedReposOfUser returns an error: %s", err)
		}

		expected := struct {
			Count       int
			FirstRepoID string
		}{
			Count:       2,
			FirstRepoID: r1.ID,
		}
		if len(rs) != 2 && rs[0].ID == expected.FirstRepoID {
			t.Fatalf("ListSortedReposOfUser = %v: %v", len(rs), expected)
		}
	})

	t.Run("Returns repos in deployed-order manager.", func(t *testing.T) {
		t.Log("Create a new deployment for r2.")
		client.Deployment.
			Create().
			SetNumber(1).
			SetRef("xde333e").
			SetEnv("local").
			SetUserID(u1).
			SetRepoID(r2.ID).
			SetCreatedAt(time.Now()).
			SaveX(ctx)

		s := NewStore(client)

		var (
			rs  []*ent.Repo
			err error
		)

		if rs, err = s.ListSortedReposOfUser(ctx, &ent.User{ID: u1}, "", 1, 30); err != nil {
			t.Fatalf("ListSortedReposOfUser returns an error: %s", err)
		}

		expected := struct {
			Count       int
			FirstRepoID string
		}{
			Count:       2,
			FirstRepoID: r2.ID,
		}
		if len(rs) != expected.Count && rs[0].ID == expected.FirstRepoID {
			t.Fatalf("ListSortedReposOfUser = %v: %v", len(rs), expected)
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
		u1 = "1"
	)

	r1 := client.Repo.
		Create().
		SetID("1").
		SetNamespace("octocat").
		SetName("Hello").
		SetDescription("").
		SaveX(ctx)

	r2 := client.Repo.
		Create().
		SetID("2").
		SetNamespace("octocat").
		SetName("World").
		SetDescription("").
		SaveX(ctx)

	client.Repo.
		Create().
		SetID("3").
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

		if repo, err = s.FindRepoOfUserByID(ctx, &ent.User{ID: u1}, "1"); err != nil {
			t.Fatalf("FindRepoOfUserByID returns an error: %s", err)
		}

		expected := "1"
		if repo.ID != expected {
			t.Fatalf("FindRepoOfUserByID = %v, wanted %v", repo.ID, expected)
		}
	})

	t.Run("Can not find unauthorized repo by ID", func(t *testing.T) {
		s := NewStore(client)

		var (
			err error
		)

		if _, err = s.FindRepoOfUserByID(ctx, &ent.User{ID: u1}, "3"); !ent.IsNotFound(err) {
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
			ID:          "1",
			Namespace:   "octocat",
			Name:        "HelloWorld",
			Description: "nothing",
			Perm:        vo.RemoteRepoPermAdmin,
		}); err != nil {
			t.Fatalf("SyncRepo return an error: %s", err)
		}
	})
}
