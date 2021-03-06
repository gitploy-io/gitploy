package interactor_test

import (
	"testing"

	i "github.com/gitploy-io/gitploy/internal/interactor"
)

func TestInteractorConfig_BuildWebhookURL(t *testing.T) {
	t.Run("Return the webhook URL built with the proxy host.", func(t *testing.T) {
		c := &i.InteractorConfig{
			ServerProxyHost:  "hook.cloud.gitploy.io",
			ServerProxyProto: "https",
		}
		wanted := "https://hook.cloud.gitploy.io/hooks"
		if ret := c.BuildWebhookURL(); ret != wanted {
			t.Fatalf("BuildWebhookURL = %v, wanted %v", ret, wanted)
		}
	})

	t.Run("Return the webhook URL built with the server host.", func(t *testing.T) {
		c := &i.InteractorConfig{
			ServerHost:  "cloud.gitploy.io",
			ServerProto: "https",
		}
		wanted := "https://cloud.gitploy.io/hooks"
		if ret := c.BuildWebhookURL(); ret != wanted {
			t.Fatalf("BuildWebhookURL = %v, wanted %v", ret, wanted)
		}
	})
}

func TestInteractorConfig_CheckWebhookURL(t *testing.T) {
	t.Run("Return true when the proxy has SSL verification.", func(t *testing.T) {
		c := &i.InteractorConfig{
			ServerProxyHost:  "hook.cloud.gitploy.io",
			ServerProxyProto: "https",
			ServerProto:      "http",
		}
		wanted := true
		if ret := c.CheckWebhookSSL(); ret != wanted {
			t.Fatalf("BuildWebhookURL = %v, wanted %v", ret, wanted)
		}
	})

	t.Run("Return true when the server has SSL verification.", func(t *testing.T) {
		c := &i.InteractorConfig{
			ServerHost:  "cloud.gitploy.io",
			ServerProto: "https",
		}
		wanted := true
		if ret := c.CheckWebhookSSL(); ret != wanted {
			t.Fatalf("BuildWebhookURL = %v, wanted %v", ret, wanted)
		}
	})
}
