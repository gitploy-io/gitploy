package github

import (
	"context"
	"testing"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/pkg/e"
	"gopkg.in/h2non/gock.v1"
)

func TestGithub_DeleteWebhook(t *testing.T) {
	t.Run("Return the ErrorCodeEntityNotFound error when the webhook is not found.", func(t *testing.T) {
		t.Log("Mocking the delete webhook API")
		gock.New("https://api.github.com").
			Delete("/repos/gitploy-io/gitploy/hooks/1").
			Reply(404)

		g := NewGithub(&GithubConfig{})

		const hookID = 1

		err := g.DeleteWebhook(
			context.Background(),
			&ent.User{},
			&ent.Repo{
				Namespace: "gitploy-io",
				Name:      "gitploy",
			},
			hookID,
		)

		if !e.HasErrorCode(err, e.ErrorCodeEntityNotFound) {
			t.Fatalf("DeleteWebhook doesn't returns an ErrorCodeEntityNotFound error: %v", err)
		}
	})

	t.Run("Delete the webhook.", func(t *testing.T) {
		t.Log("Mocking the delete webhook API")
		gock.New("https://api.github.com").
			Delete("/repos/gitploy-io/gitploy/hooks/1").
			Reply(200)

		g := NewGithub(&GithubConfig{})

		const hookID = 1

		err := g.DeleteWebhook(
			context.Background(),
			&ent.User{},
			&ent.Repo{
				Namespace: "gitploy-io",
				Name:      "gitploy",
			},
			hookID,
		)

		if err != nil {
			t.Fatalf("DeleteWebhook returns an error: %v", err)
		}
	})
}
