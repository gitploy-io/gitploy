package github

import (
	"context"
	"reflect"
	"testing"

	"github.com/gitploy-io/gitploy/model/ent"
	"gopkg.in/h2non/gock.v1"
)

func TestGithub_GetConfigRedirectURL(t *testing.T) {
	t.Run("Return the link of the configuration file.", func(t *testing.T) {
		t.Log("Mocking the GET repo API.")
		gock.New("https://api.github.com").
			Get("/repos/octocat/hello-world").
			Reply(200).
			File("./testdata/repo.get.json")

		g := NewGithub(&GithubConfig{})

		link, err := g.GetConfigRedirectURL(
			context.Background(),
			&ent.User{},
			&ent.Repo{
				Namespace:  "octocat",
				Name:       "hello-world",
				ConfigPath: "config.yml",
			},
		)
		if err != nil {
			t.Fatalf("GetConfigRedirectURL returns an error: %s", err)
		}

		want := "https://github.com/octocat/Hello-World/blob/master/config.yml"
		if !reflect.DeepEqual(link, want) {
			t.Fatalf("GetConfigRedirectURL = %v, wanted %v", link, want)
		}
	})
}

func TestGithub_GetNewConfigRedirectURL(t *testing.T) {
	t.Run("Return the link of the configuration file.", func(t *testing.T) {
		t.Log("Mocking the GET repo API.")
		gock.New("https://api.github.com").
			Get("/repos/octocat/hello-world").
			Reply(200).
			File("./testdata/repo.get.json")

		g := NewGithub(&GithubConfig{})

		link, err := g.GetNewConfigRedirectURL(
			context.Background(),
			&ent.User{},
			&ent.Repo{
				Namespace:  "octocat",
				Name:       "hello-world",
				ConfigPath: "config.yml",
			},
		)
		if err != nil {
			t.Fatalf("GetNewConfigRedirectURL returns an error: %s", err)
		}

		want := "https://github.com/octocat/Hello-World/new/master/config.yml"
		if !reflect.DeepEqual(link, want) {
			t.Fatalf("GetNewConfigRedirectURL = %v, wanted %v", link, want)
		}
	})
}
