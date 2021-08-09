package store

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/deployment"
	"github.com/hanjunlee/gitploy/ent/enttest"
	"github.com/hanjunlee/gitploy/ent/notification"

	_ "github.com/mattn/go-sqlite3"
)

func setupNotificationTestClient(client *ent.Client, u *ent.User) *ent.Client {
	ctx := context.Background()

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

func TestStore_ListPublishingNotificaitonsGreaterThanTime(t *testing.T) {
	t.Run("Returns notifications that notified is false.", func(t *testing.T) {
		u := &ent.User{
			ID: "1",
		}

		client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
		client = setupNotificationTestClient(client, u)
		defer client.Close()

		var (
			ctx = context.Background()
			ct  = time.Now().Add(-time.Second)
		)

		n := client.Notification.
			Create().
			SetType(notification.TypeDeploymentCreated).
			SetRepoNamespace("octocat").
			SetRepoName("HelloWorld").
			SetDeploymentNumber(3).
			SetDeploymentType("commit").
			SetDeploymentRef("ee23abb").
			SetDeploymentEnv("local").
			SetDeploymentStatus(string(deployment.StatusCreated)).
			SetDeploymentLogin("octocat").
			SetCreatedAt(ct).
			SetUpdatedAt(ct).
			SetUserID(u.ID).
			SaveX(ctx)

		s := NewStore(client)

		ns, err := s.ListPublishingNotificaitonsGreaterThanTime(ctx, ct.Add(-3*time.Second))
		if err != nil {
			t.Fatalf("ListPublishingNotificaitonsGreaterThanTime returns an error: %s", err)
			t.FailNow()
		}

		expected := []*ent.Notification{n}
		if reflect.DeepEqual(ns, expected) {
			t.Fatalf("ListPublishingNotificaitonsGreaterThanTime = %v, wanted %v", ns, expected)
			t.FailNow()
		}

		// Expect zero when the other list notificaitons.
		ns, _ = s.ListPublishingNotificaitonsGreaterThanTime(ctx, ct.Add(-3*time.Second))
		if len(ns) != 0 {
			t.Fatalf("ListPublishingNotificaitonsGreaterThanTime count = %d, wanted %d", len(ns), 0)
			t.FailNow()
		}
	})
}
