package store

import (
	"context"
	"testing"

	"github.com/gitploy-io/gitploy/model/ent/enttest"
	"github.com/gitploy-io/gitploy/model/ent/migrate"
	"github.com/gitploy-io/gitploy/model/ent/review"
	"github.com/gitploy-io/gitploy/pkg/e"
)

func TestStore_UpdateReview(t *testing.T) {
	t.Run("Return an unprocessible entity error when the vaildation is failed.", func(t *testing.T) {
		ctx := context.Background()

		client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1",
			enttest.WithMigrateOptions(migrate.WithForeignKeys(false)),
		)
		defer client.Close()

		r := client.Review.
			Create().
			SetDeploymentID(1).
			SetUserID(1).
			SaveX(ctx)

		s := NewStore(client)

		r.Status = review.Status("UNPROCESSIBLE")
		_, err := s.UpdateReview(ctx, r)
		if !e.HasErrorCode(err, e.ErrorCodeEntityUnprocessable) {
			t.Fatalf("UpdateReview error code = %v, wanted unprocessable_entity", err)
		}
	})

	t.Run("Update the review.", func(t *testing.T) {
		ctx := context.Background()

		client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1",
			enttest.WithMigrateOptions(migrate.WithForeignKeys(false)),
		)
		defer client.Close()

		r := client.Review.
			Create().
			SetDeploymentID(1).
			SetUserID(1).
			SaveX(ctx)

		s := NewStore(client)

		r.Status = review.StatusApproved
		_, err := s.UpdateReview(ctx, r)
		if err != nil {
			t.Fatalf("UpdateReview returns an error: %v", err)
		}
	})
}
