package api

import (
	"context"
	"net/http"
	"testing"

	"github.com/gitploy-io/gitploy/model/ent"
	"gopkg.in/h2non/gock.v1"
)

func TestReposService_List(t *testing.T) {
	t.Run("Return repositories.", func(t *testing.T) {
		repos := []*ent.Repo{
			{
				Namespace: "gitploy-io",
				Name:      "gitploy",
			},
			{
				Namespace: "gitploy-io",
				Name:      "website",
			},
		}
		gock.New("https://cloud.gitploy.io").
			Get("/api/v1/repos").
			Reply(200).
			JSON(repos)

		c := NewClient("https://cloud.gitploy.io/", http.DefaultClient)

		ret, err := c.Repos.List(context.Background(), RepoListOptions{
			ListOptions: ListOptions{Page: 1, PerPage: 30},
		})
		if err != nil {
			t.Fatalf("List returns an error: %s", err)
		}

		for idx := range ret {
			if !EqualRepo(repos[idx], ret[idx]) {
				t.Fatalf("List = %v, wanted %v", ret, repos)
			}
		}
	})
}
