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
	"github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
	gm "github.com/golang/mock/gomock"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func TestReviewAPI_UpdateMine(t *testing.T) {
	t.Run("Return 400 code when the status is invalid", func(t *testing.T) {
		ctrl := gm.NewController(t)
		s := &ReviewAPI{i: mock.NewMockInteractor(ctrl), log: zap.L()}

		router := gin.New()
		router.PATCH("/deployments/:number/review", s.UpdateMine)

		p, _ := json.Marshal(&ReviewPatchPayload{
			Status: "INVALID",
		})
		req, _ := http.NewRequest("PATCH", "/deployments/1/review", bytes.NewBuffer(p))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusBadRequest {
			t.Fatalf("Code = %v, wanted %v", w.Code, http.StatusBadRequest)
		}
	})

	t.Run("Return 200 when the update has succeeded.", func(t *testing.T) {
		t.Log("Start mocking:")
		ctrl := gm.NewController(t)
		i := mock.NewMockInteractor(ctrl)

		t.Log("\tFind the user's review.")
		i.EXPECT().
			FindDeploymentOfRepoByNumber(gm.Any(), gm.AssignableToTypeOf(&ent.Repo{}), gm.AssignableToTypeOf(1)).
			Return(&ent.Deployment{}, nil)

		i.EXPECT().
			FindReviewOfUser(gm.Any(), gm.Any(), gm.AssignableToTypeOf(&ent.Deployment{})).
			Return(&ent.Review{}, nil)

		t.Log("\tRespond the review.")
		i.EXPECT().
			RespondReview(gm.Any(), gm.AssignableToTypeOf(&ent.Review{})).
			Return(&ent.Review{}, nil)

		s := &ReviewAPI{i: i, log: zap.L()}
		router := gin.New()
		router.PATCH("/deployments/:number/review", func(c *gin.Context) {
			c.Set(global.KeyUser, &ent.User{})
			c.Set(KeyRepo, &ent.Repo{})
		}, s.UpdateMine)

		p, _ := json.Marshal(&ReviewPatchPayload{
			Status: "approved",
		})
		req, _ := http.NewRequest("PATCH", "/deployments/1/review", bytes.NewBuffer(p))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("Code = %v, wanted %v", w.Code, http.StatusOK)
		}
	})
}
