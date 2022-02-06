package api

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"gopkg.in/h2non/gock.v1"

	"github.com/AlekSi/pointer"
	"github.com/gitploy-io/gitploy/model/ent"
)

func EqualRepo(a, b *ent.Repo) bool {
	return a.Namespace == b.Namespace && a.Name == b.Name
}

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

func TestReposService_Update(t *testing.T) {
	t.Run("Update the 'config_path' field.", func(t *testing.T) {
		gock.New("https://cloud.gitploy.io").
			Patch("/api/v1/repos/gitploy-io/gitploy").
			AddMatcher(func(req *http.Request, ereq *gock.Request) (bool, error) {
				defer req.Body.Close()
				output, _ := ioutil.ReadAll(req.Body)

				b := RepoUpdateRequest{}
				if err := json.Unmarshal(output, &b); err != nil {
					return false, err
				}

				// Verify the field "config_path" in the body.
				return *b.ConfigPath == "new_deploy.yml", nil
			}).
			Reply(200).
			JSON(&ent.Repo{
				Namespace:  "gitploy-io",
				Name:       "gitploy",
				ConfigPath: "new_deploy.yml",
			})

		c := NewClient("https://cloud.gitploy.io/", http.DefaultClient)

		_, err := c.Repos.Update(context.Background(), "gitploy-io", "gitploy", RepoUpdateRequest{
			ConfigPath: pointer.ToString("new_deploy.yml"),
		})
		if err != nil {
			t.Fatalf("Update returns an error: %s", err)
		}
	})
}
