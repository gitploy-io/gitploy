package github

import (
	"testing"
)

func Test_NewGithub(t *testing.T) {
	t.Run("Create a new Github with the base URL.", func(t *testing.T) {
		url := "https://github.gitploy.io/"

		g := NewGithub(&GithubConfig{
			BaseURL: url,
		})

		if g.baseURL != url {
			t.Fatalf("NewGithub.baseURL = %v, wanted %v", g.baseURL, url)
		}
	})
}
