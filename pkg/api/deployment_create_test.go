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

func TestDeploymentsService_Create(t *testing.T) {
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

		_, err := c.Deployments.Create(context.Background(), "gitploy-io", "gitploy", DeploymentCreateRequest{
			Type: "branch",
			Env:  "production",
			Ref:  "main",
		})
		if err != nil {
			t.Fatalf("Create returns an error: %s", err)
		}
	})
}
