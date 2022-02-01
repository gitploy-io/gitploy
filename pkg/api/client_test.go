package api

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gitploy-io/gitploy/pkg/e"
)

func TestClient_Do(t *testing.T) {
	t.Run("Return an internal error", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, `{"code": "entity_not_found", "message": "It is not found."}`)
		}))
		defer ts.Close()

		// Append '/' to avoid an trailing slash error.
		baseURL, _ := url.Parse(ts.URL + "/")
		c := &client{httpClient: http.DefaultClient, BaseURL: baseURL}

		req, err := c.NewRequest("GET", baseURL.Path, nil)
		if err != nil {
			t.Fatalf("Failed to build a request: %s", err)
		}

		err = c.Do(context.Background(), req, nil)
		if !e.HasErrorCode(err, e.ErrorCodeEntityNotFound) {
			t.Fatalf("Do = %v, want ErrorCodeEntityNotFound", err)
		}
	})

	t.Run("Return OK response", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))
		defer ts.Close()

		// Append '/' to avoid an trailing slash error.
		baseURL, _ := url.Parse(ts.URL + "/")
		c := &client{httpClient: http.DefaultClient, BaseURL: baseURL}

		req, err := c.NewRequest("GET", baseURL.Path, nil)
		if err != nil {
			t.Fatalf("Failed to build a request: %s", err)
		}

		err = c.Do(context.Background(), req, nil)
		if err != nil {
			t.Fatalf("Do returns an error: %s", err)
		}
	})
}
