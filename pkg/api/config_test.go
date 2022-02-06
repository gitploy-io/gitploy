package api

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/gitploy-io/gitploy/model/extent"
	"gopkg.in/h2non/gock.v1"
)

func TestConfig_Get(t *testing.T) {
	t.Run("Return the config.", func(t *testing.T) {
		config := &extent.Config{
			Envs: []*extent.Env{
				{Name: "production"},
			},
		}

		gock.New("https://cloud.gitploy.io").
			Get("/api/v1/repos/gitploy-io/gitploy/config").
			Reply(200).
			JSON(config)

		c := NewClient("https://cloud.gitploy.io/", http.DefaultClient)

		ret, err := c.Config.Get(context.Background(), "gitploy-io", "gitploy")
		if err != nil {
			t.Fatalf("Get returns an error: %s", err)
		}

		if !reflect.DeepEqual(ret, config) {
			t.Fatalf("Get = %v, wanted %v", ret, config)
		}
	})
}
