package repos

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gitploy-io/gitploy/internal/server/api/v1/repos/mock"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/golang/mock/gomock"
	"go.uber.org/zap"
)

func TestLockService_Delete(t *testing.T) {
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

		s := LockService{i: m, log: zap.L()}
		gin.SetMode(gin.ReleaseMode)
		router := gin.New()
		router.DELETE("repos/:id/locks/:lockID", s.Delete)

		req, _ := http.NewRequest("DELETE", fmt.Sprintf("/repos/1/locks/%d", input.id), nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		if w.Code != http.StatusOK {
			t.Fatalf("Code = %v, wanted %v. Body=%v", w.Code, http.StatusCreated, w.Body)
		}
	})
}
