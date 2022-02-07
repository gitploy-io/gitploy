// Copyright 2021 Gitploy.IO Inc. All rights reserved.
// Use of this source code is governed by the Gitploy Non-Commercial License
// that can be found in the LICENSE file.

//go:build !oss

package repos

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/gitploy-io/gitploy/internal/server/api/v1/repos/mock"
	"github.com/golang/mock/gomock"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func TestReviewsAPI_UpdateMine(t *testing.T) {
	t.Run("Return 400 code when the status is invalid", func(t *testing.T) {
		input := struct {
			payload *ReviewPatchPayload
		}{
			payload: &ReviewPatchPayload{
				Status: "INVALID",
			},
		}

		ctrl := gomock.NewController(t)
		m := mock.NewMockInteractor(ctrl)

		router := gin.New()

		s := &ReviewsAPI{i: m, log: zap.L()}
		router.PATCH("/deployments/:number/review", s.UpdateMine)

		p, _ := json.Marshal(input.payload)
		req, _ := http.NewRequest("PATCH", "/deployments/1/review", bytes.NewBuffer(p))
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusBadRequest {
			t.Fatalf("Code = %v, wanted %v", w.Code, http.StatusBadRequest)
		}
	})
}
