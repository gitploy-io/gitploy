package repos

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AlekSi/pointer"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/internal/server/api/v1/repos/mock"
	"github.com/hanjunlee/gitploy/internal/server/global"
	"github.com/hanjunlee/gitploy/vo"
)

func TestRepo_UpdateRepo(t *testing.T) {
	t.Run("Patch config_path field.", func(t *testing.T) {
		input := struct {
			payload *repoPatchPayload
		}{
			payload: &repoPatchPayload{
				ConfigPath: pointer.ToString("deploy.yml"),
			},
		}

		const (
			r1 = "1"
		)

		ctrl := gomock.NewController(t)
		m := mock.NewMockInteractor(ctrl)

		t.Log("Update the config_path field with the payload.")
		m.
			EXPECT().
			UpdateRepo(gomock.Any(), gomock.Eq(&ent.Repo{
				ID:         r1,
				ConfigPath: *input.payload.ConfigPath,
			})).
			DoAndReturn(func(ctx context.Context, r *ent.Repo) (*ent.Repo, error) {
				return r, nil
			})

		gin.SetMode(gin.ReleaseMode)
		router := gin.New()

		r := NewRepo(RepoConfig{}, m)
		router.PATCH("/repos/:id", func(c *gin.Context) {
			t.Log("Set up fake middleware")
			c.Set(global.KeyUser, &ent.User{})
			c.Set(KeyRepo, &ent.Repo{ID: r1})
		}, r.UpdateRepo)

		p, _ := json.Marshal(input.payload)
		req, _ := http.NewRequest("PATCH", fmt.Sprintf("/repos/%s", r1), bytes.NewBuffer(p))
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("Code = %v, wanted %v", w.Code, http.StatusOK)
		}
	})

	t.Run("Patch active field.", func(t *testing.T) {
		input := struct {
			payload *repoPatchPayload
		}{
			payload: &repoPatchPayload{
				Active: pointer.ToBool(true),
			},
		}

		const (
			r1 = "1"
		)

		ctrl := gomock.NewController(t)
		m := mock.NewMockInteractor(ctrl)

		t.Log("Check to call ActivateRepo.")
		m.
			EXPECT().
			ActivateRepo(gomock.Any(), gomock.AssignableToTypeOf(&ent.User{}), gomock.Eq(&ent.Repo{
				ID: r1,
			}), gomock.AssignableToTypeOf(&vo.WebhookConfig{})).
			DoAndReturn(func(ctx context.Context, u *ent.User, r *ent.Repo, c *vo.WebhookConfig) (*ent.Repo, error) {
				return r, nil
			})

		gin.SetMode(gin.ReleaseMode)
		router := gin.New()

		r := NewRepo(RepoConfig{}, m)
		router.PATCH("/repos/:id", func(c *gin.Context) {
			t.Log("Set up fake middleware")
			c.Set(global.KeyUser, &ent.User{})
			c.Set(KeyRepo, &ent.Repo{ID: r1})
		}, r.UpdateRepo)

		p, _ := json.Marshal(input.payload)
		req, _ := http.NewRequest("PATCH", fmt.Sprintf("/repos/%s", r1), bytes.NewBuffer(p))
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("Code = %v, wanted %v", w.Code, http.StatusOK)
		}
	})
}
