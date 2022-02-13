package store

import (
	"context"
	"testing"
	"time"

	i "github.com/gitploy-io/gitploy/internal/interactor"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/ent/enttest"
	"github.com/gitploy-io/gitploy/model/ent/migrate"
)

func TestStore_Searchusers(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1",
		enttest.WithMigrateOptions(migrate.WithForeignKeys(false)),
	)
	defer client.Close()

	ctx := context.Background()

	u1 := client.User.
		Create().
		SetID(1).
		SetLogin("banana").
		SetAvatar("").
		SetToken("").
		SetRefresh("").
		SetExpiry(time.Time{}).
		SaveX(ctx)

	u2 := client.User.
		Create().
		SetID(2).
		SetLogin("apple").
		SetAvatar("").
		SetToken("").
		SetRefresh("").
		SetExpiry(time.Time{}).
		SaveX(ctx)

	t.Run("Returns users in login-sorted manager.", func(t *testing.T) {
		s := NewStore(client)

		var (
			us  []*ent.User
			err error
		)

		if us, err = s.SearchUsers(ctx, &i.SearchUsersOptions{
			Query:       "",
			ListOptions: i.ListOptions{Page: 1, PerPage: 30},
		}); err != nil {
			t.Fatalf("ListUsers returns an error: %s", err)
		}

		expected := []*ent.User{u2, u1}
		if !(len(us) == 2 && us[0].ID == expected[0].ID) {
			t.Fatalf("ListUsers = %v, wanted %v", us, expected)
		}
	})

	t.Run("Returns user filtered by login.", func(t *testing.T) {
		s := NewStore(client)

		var (
			us  []*ent.User
			err error
		)

		if us, err = s.SearchUsers(ctx, &i.SearchUsersOptions{
			Query: "pple",
			ListOptions: i.ListOptions{
				Page:    1,
				PerPage: 30,
			}}); err != nil {
			t.Fatalf("ListUsers returns an error: %s", err)
		}

		expected := u2
		if !(len(us) == 1 && us[0].ID == expected.ID) {
			t.Fatalf("ListUsers = %v, wanted %v", us, expected)
		}
	})
}
