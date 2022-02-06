package api

import (
	"context"
	"net/http"
	"testing"

	"github.com/gitploy-io/gitploy/model/ent"
	"gopkg.in/h2non/gock.v1"
)

func TestReposService_Get(t *testing.T) {
	t.Run("Return the repository.", func(t *testing.T) {
		repo := &ent.Repo{
			Namespace: "gitploy-io",
			Name:      "gitploy",
		}
		gock.New("https://cloud.gitploy.io").
			Get("/api/v1/repos/gitploy-io/gitploy").
			Reply(200).
			JSON(repo)

		c := NewClient("https://cloud.gitploy.io/", http.DefaultClient)

		ret, err := c.Repos.Get(context.Background(), "gitploy-io", "gitploy")
		if err != nil {
			t.Fatalf("Get returns an error: %s", err)
		}

		if !EqualRepo(repo, ret) {
			t.Fatalf("Get = %v, wanted %v", ret, repo)
		}
	})
}
