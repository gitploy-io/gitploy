package repos

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"go.uber.org/zap"

	"github.com/gitploy-io/gitploy/internal/server/api/v1/repos/mock"
	"github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/extent"
)

func TestLockAPI_Create(t *testing.T) {
	t.Run("Return 422 when the environment is undefined.", func(t *testing.T) {
		input := struct {
			payload *LockPostPayload
		}{
			payload: &LockPostPayload{
				Env: "production",
			},
		}

		ctrl := gomock.NewController(t)
		m := mock.NewMockInteractor(ctrl)

		m.
			EXPECT().
			GetConfig(gomock.Any(), gomock.AssignableToTypeOf(&ent.User{}), gomock.AssignableToTypeOf(&ent.Repo{})).
			Return(&extent.Config{
				Envs: []*extent.Env{
					{
						Name: "dev",
					},
				}}, nil)

		gin.SetMode(gin.ReleaseMode)

		s := LockAPI{i: m, log: zap.L()}
		router := gin.New()
		router.POST("repos/:id/locks", func(c *gin.Context) {
			c.Set(global.KeyUser, &ent.User{})
			c.Set(KeyRepo, &ent.Repo{})
		}, s.Create)

		body, _ := json.Marshal(input.payload)
		req, _ := http.NewRequest("POST", "/repos/1/locks", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		if w.Code != http.StatusUnprocessableEntity {
			t.Fatalf("Code = %v, wanted %v. Body=%v", w.Code, http.StatusUnprocessableEntity, w.Body)
		}
	})

	t.Run("Lock the env", func(t *testing.T) {
		input := struct {
			payload *LockPostPayload
		}{
			payload: &LockPostPayload{
				Env: "production",
			},
		}

		ctrl := gomock.NewController(t)
		m := mock.NewMockInteractor(ctrl)

		m.
			EXPECT().
			GetConfig(gomock.Any(), gomock.AssignableToTypeOf(&ent.User{}), gomock.AssignableToTypeOf(&ent.Repo{})).
			Return(&extent.Config{
				Envs: []*extent.Env{
					{
						Name: "production",
					},
				}}, nil)

		t.Log("Lock the env.")
		m.
			EXPECT().
			CreateLock(gomock.Any(), gomock.AssignableToTypeOf(&ent.Lock{})).
			Return(&ent.Lock{ID: 1}, nil)

		t.Log("Get the lock with edges")
		m.
			EXPECT().
			FindLockByID(gomock.Any(), 1).
			Return(&ent.Lock{ID: 1}, nil)

		s := LockAPI{i: m, log: zap.L()}
		gin.SetMode(gin.ReleaseMode)
		router := gin.New()
		router.POST("repos/:id/locks", func(c *gin.Context) {
			c.Set(global.KeyUser, &ent.User{})
			c.Set(KeyRepo, &ent.Repo{})
		}, s.Create)

		body, _ := json.Marshal(input.payload)
		req, _ := http.NewRequest("POST", "/repos/1/locks", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		if w.Code != http.StatusCreated {
			t.Fatalf("Code = %v, wanted %v. Body=%v", w.Code, http.StatusCreated, w.Body)
		}
	})
}
