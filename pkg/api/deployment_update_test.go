package api

import (
	"context"
	"net/http"
	"testing"

	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/ent/deployment"
	"gopkg.in/h2non/gock.v1"
)

func TestDeploymentsService_Update(t *testing.T) {
	t.Run("Verify the request.", func(t *testing.T) {
		d := &ent.Deployment{
			ID: 2, Number: 1, Type: deployment.TypeBranch, Env: "production", Ref: "main",
		}
		gock.New("https://cloud.gitploy.io").
			Put("/api/v1/repos/gitploy-io/gitploy/deployments/1").
			Reply(201).
			JSON(d)

		c := NewClient("https://cloud.gitploy.io/", http.DefaultClient)

		_, err := c.Deployments.Update(context.Background(), "gitploy-io", "gitploy", 1)
		if err != nil {
			t.Fatalf("Create returns an error: %s", err)
		}
	})
}
