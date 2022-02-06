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
	"go.uber.org/zap"

	"github.com/gitploy-io/gitploy/internal/server/api/v1/repos/mock"
	"github.com/gitploy-io/gitploy/model/ent"
)

func TestLocksAPI_Update(t *testing.T) {
	t.Run("Update the expired time.", func(t *testing.T) {
		expiredAt := time.Date(2021, 10, 17, 0, 0, 0, 0, time.UTC)

		input := struct {
			id      int
			payload *LockPatchPayload
		}{
			id: 1,
			payload: &LockPatchPayload{
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

		s := LocksAPI{i: m, log: zap.L()}
		gin.SetMode(gin.ReleaseMode)
		router := gin.New()
		router.PATCH("repos/:id/locks/:lockID", s.Update)

		body, _ := json.Marshal(input.payload)
		req, _ := http.NewRequest("PATCH", fmt.Sprintf("/repos/1/locks/%d", input.id), bytes.NewBuffer(body))

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("Code = %v, wanted %v. Body=%v", w.Code, http.StatusCreated, w.Body)
		}
	})
}
