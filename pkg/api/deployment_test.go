package api

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/ent/deployment"
	"gopkg.in/h2non/gock.v1"
)

func TestDeploymentService_List(t *testing.T) {
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

		ret, err := c.Deployment.List(context.Background(), "gitploy-io", "gitploy", DeploymentListOptions{
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

func TestDeploymentService_Create(t *testing.T) {
	t.Run("Verify the body of a request.", func(t *testing.T) {
		d := &ent.Deployment{
			ID: 2, Number: 1, Type: deployment.TypeBranch, Env: "production", Ref: "main",
		}
		gock.New("https://cloud.gitploy.io").
			Post("/api/v1/repos/gitploy-io/gitploy/deployments").
			AddMatcher(func(req *http.Request, ereq *gock.Request) (bool, error) {
				defer req.Body.Close()
				output, _ := ioutil.ReadAll(req.Body)

				b := DeploymentCreateRequest{}
				if err := json.Unmarshal(output, &b); err != nil {
					return false, err
				}

				// Verify the fields of the body.
				return b.Type == "branch" && b.Env == "production" && b.Ref == "main", nil
			}).
			Reply(201).
			JSON(d)

		c := NewClient("https://cloud.gitploy.io/", http.DefaultClient)

		_, err := c.Deployment.Create(context.Background(), "gitploy-io", "gitploy", DeploymentCreateRequest{
			Type: "branch",
			Env:  "production",
			Ref:  "main",
		})
		if err != nil {
			t.Fatalf("Create returns an error: %s", err)
		}
	})
}

func TestDeploymentService_Update(t *testing.T) {
	t.Run("Verify the request.", func(t *testing.T) {
		d := &ent.Deployment{
			ID: 2, Number: 1, Type: deployment.TypeBranch, Env: "production", Ref: "main",
		}
		gock.New("https://cloud.gitploy.io").
			Put("/api/v1/repos/gitploy-io/gitploy/deployments/1").
			Reply(201).
			JSON(d)

		c := NewClient("https://cloud.gitploy.io/", http.DefaultClient)

		_, err := c.Deployment.Update(context.Background(), "gitploy-io", "gitploy", 1)
		if err != nil {
			t.Fatalf("Create returns an error: %s", err)
		}
	})
}
