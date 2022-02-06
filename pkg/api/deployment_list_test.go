package api

import (
	"context"
	"net/http"
	"testing"

	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/ent/deployment"
	"gopkg.in/h2non/gock.v1"
)

func TestDeploymentsService_List(t *testing.T) {
	t.Run("Verify the query of a request.", func(t *testing.T) {
		ds := []*ent.Deployment{
			{ID: 1, Env: "production", Status: deployment.StatusWaiting},
			{ID: 2, Env: "production", Status: deployment.StatusWaiting},
		}
		gock.New("https://cloud.gitploy.io").
			Get("/api/v1/repos/gitploy-io/gitploy/deployments").
			MatchParam("env", "production").
			MatchParam("status", "waiting").
			Reply(200).
			JSON(ds)

		c := NewClient("https://cloud.gitploy.io/", http.DefaultClient)

		ret, err := c.Deployments.List(context.Background(), "gitploy-io", "gitploy", DeploymentListOptions{
			ListOptions: ListOptions{Page: 1, PerPage: 30},
			Env:         "production",
			Status:      "waiting",
		})
		if err != nil {
			t.Fatalf("List returns an error: %s", err)
		}

		for idx := range ret {
			if ret[idx].ID != ds[idx].ID {
				t.Fatalf("List = %v, wanted %v", ret, ds)
			}
		}
	})
}
