package repos

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/internal/server/api/v1/repos/mock"
	"github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/pkg/e"
	"github.com/gitploy-io/gitploy/vo"
)

func TestRepo_CreateLock(t *testing.T) {
	t.Run("Return 422 when the environment is undefined.", func(t *testing.T) {
		input := struct {
			payload *lockPostPayload
		}{
			payload: &lockPostPayload{
				Env: "production",
			},
		}

		ctrl := gomock.NewController(t)
		m := mock.NewMockInteractor(ctrl)

		m.
			EXPECT().
			GetConfig(gomock.Any(), gomock.AssignableToTypeOf(&ent.User{}), gomock.AssignableToTypeOf(&ent.Repo{})).
			Return(nil, e.NewError(e.ErrorCodeConfigUndefinedEnv, nil))

		r := NewRepo(RepoConfig{}, m)

		gin.SetMode(gin.ReleaseMode)

		router := gin.New()
		router.POST("repos/:id/locks", func(c *gin.Context) {
			c.Set(global.KeyUser, &ent.User{})
			c.Set(KeyRepo, &ent.Repo{})
		}, r.CreateLock)

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
			payload *lockPostPayload
		}{
			payload: &lockPostPayload{
				Env: "production",
			},
		}

		ctrl := gomock.NewController(t)
		m := mock.NewMockInteractor(ctrl)

		m.
			EXPECT().
			GetConfig(gomock.Any(), gomock.AssignableToTypeOf(&ent.User{}), gomock.AssignableToTypeOf(&ent.Repo{})).
			Return(&vo.Config{
				Envs: []*vo.Env{
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

		r := NewRepo(RepoConfig{}, m)

		gin.SetMode(gin.ReleaseMode)
		router := gin.New()
		router.POST("repos/:id/locks", func(c *gin.Context) {
			c.Set(global.KeyUser, &ent.User{})
			c.Set(KeyRepo, &ent.Repo{})
		}, r.CreateLock)

		body, _ := json.Marshal(input.payload)
		req, _ := http.NewRequest("POST", "/repos/1/locks", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		if w.Code != http.StatusCreated {
			t.Fatalf("Code = %v, wanted %v. Body=%v", w.Code, http.StatusCreated, w.Body)
		}
	})
}

func TestRepo_UpdateLock(t *testing.T) {
	t.Run("Update the expired time.", func(t *testing.T) {
		expiredAt := time.Date(2021, 10, 17, 0, 0, 0, 0, time.UTC)

		input := struct {
			id      int
			payload *lockPatchPayload
		}{
			id: 1,
			payload: &lockPatchPayload{
				ExpiredAt: pointer.ToString(expiredAt.Format(time.RFC3339)),
			},
		}

		ctrl := gomock.NewController(t)
		m := mock.NewMockInteractor(ctrl)

		t.Log("MOCK - Find the lock by ID.")
		m.
			EXPECT().
			FindLockByID(gomock.Any(), input.id).
			Return(&ent.Lock{ID: input.id, Env: "production"}, nil).
			MaxTimes(2)

		t.Log("MOCK - Update the expired_at field.")
		m.
			EXPECT().
			UpdateLock(gomock.Any(), gomock.Eq(&ent.Lock{
				ID:        input.id,
				Env:       "production",
				ExpiredAt: &expiredAt,
			})).
			DoAndReturn(func(_ context.Context, l *ent.Lock) (*ent.Lock, error) {
				return l, nil
			})

		r := NewRepo(RepoConfig{}, m)

		gin.SetMode(gin.ReleaseMode)
		router := gin.New()
		router.PATCH("repos/:id/locks/:lockID", r.UpdateLock)

		body, _ := json.Marshal(input.payload)
		req, _ := http.NewRequest("PATCH", fmt.Sprintf("/repos/1/locks/%d", input.id), bytes.NewBuffer(body))

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("Code = %v, wanted %v. Body=%v", w.Code, http.StatusCreated, w.Body)
		}
	})
}

func TestRepo_DeleteLock(t *testing.T) {
	t.Run("Unlock the env", func(t *testing.T) {
		input := struct {
			id int
		}{
			id: 1,
		}

		ctrl := gomock.NewController(t)
		m := mock.NewMockInteractor(ctrl)

		t.Log("Find the lock")
		m.
			EXPECT().
			FindLockByID(gomock.Any(), input.id).
			Return(&ent.Lock{ID: input.id}, nil)

		t.Log("Delete the lock")
		m.
			EXPECT().
			DeleteLock(gomock.Any(), gomock.Eq(&ent.Lock{ID: input.id})).
			Return(nil)

		r := NewRepo(RepoConfig{}, m)

		gin.SetMode(gin.ReleaseMode)
		router := gin.New()
		router.DELETE("repos/:id/locks/:lockID", r.DeleteLock)

		req, _ := http.NewRequest("DELETE", fmt.Sprintf("/repos/1/locks/%d", input.id), nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		if w.Code != http.StatusOK {
			t.Fatalf("Code = %v, wanted %v. Body=%v", w.Code, http.StatusCreated, w.Body)
		}
	})
}
