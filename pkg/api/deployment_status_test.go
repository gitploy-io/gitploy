package api

import (
	"context"
	"net/http"
	"testing"

	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/extent"
	"gopkg.in/h2non/gock.v1"
)

func TestDeploymentStatus_List(t *testing.T) {
	t.Run("Return the list of statuses.", func(t *testing.T) {
		dss := []*ent.DeploymentStatus{
			{ID: 1},
		}
		gock.New("https://cloud.gitploy.io").
			Get("/api/v1/repos/gitploy-io/gitploy/deployments/1/statuses").
			Reply(200).
			JSON(dss)

		c := NewClient("https://cloud.gitploy.io/", http.DefaultClient)

		_, err := c.DeploymentStatus.List(context.Background(), "gitploy-io", "gitploy", 1, &ListOptions{})
		if err != nil {
			t.Fatalf("Create returns an error: %s", err)
		}
	})
}

func TestDeploymentStatus_Create(t *testing.T) {
	t.Run("Return the deployment statuses.", func(t *testing.T) {
		gock.New("https://cloud.gitploy.io").
			Post("/api/v1/repos/gitploy-io/gitploy/deployments/1/remote-statuses").
			Reply(201).
			JSON(extent.RemoteDeploymentStatus{ID: 1})

		c := NewClient("https://cloud.gitploy.io/", http.DefaultClient)

		_, err := c.DeploymentStatus.CreateRemote(context.Background(), "gitploy-io", "gitploy", 1, &DeploymentStatusCreateRemoteRequest{})
		if err != nil {
			t.Fatalf("Create returns an error: %s", err)
		}
	})
}
