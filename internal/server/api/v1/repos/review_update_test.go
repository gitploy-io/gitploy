package repos

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gitploy-io/gitploy/internal/server/api/v1/repos/mock"
	"github.com/golang/mock/gomock"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func TestReviewService_UpdateMine(t *testing.T) {
	t.Run("Return 400 code when the status is invalid", func(t *testing.T) {
		input := struct {
			payload *reviewPatchPayload
		}{
			payload: &reviewPatchPayload{
				Status: "INVALID",
			},
		}

		ctrl := gomock.NewController(t)
		m := mock.NewMockInteractor(ctrl)

		router := gin.New()

		r := NewRepo(RepoConfig{}, m)
		router.PATCH("/deployments/:number/review", r.UpdateUserReview)

		p, _ := json.Marshal(input.payload)
		req, _ := http.NewRequest("PATCH", "/deployments/1/review", bytes.NewBuffer(p))
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusBadRequest {
			t.Fatalf("Code = %v, wanted %v", w.Code, http.StatusBadRequest)
		}
	})
}
