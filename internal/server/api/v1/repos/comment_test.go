package repos

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/comment"
	"github.com/gitploy-io/gitploy/internal/server/api/v1/repos/mock"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func TestRepo_CreateComment(t *testing.T) {
	t.Run("Return 400 when the number parameter is invalid", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := mock.NewMockInteractor(ctrl)

		r := NewRepo(RepoConfig{}, m)
		router := gin.New()
		router.POST("/deployments/:number/comments", r.CreateComment)

		req, _ := http.NewRequest("POST", "/deployments/foo/comments", nil)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusBadRequest {
			t.Fatalf("Code != %d, wanted %d", w.Code, http.StatusBadRequest)
		}
	})

	t.Run("Return 201 when the payload is valid", func(t *testing.T) {
		input := struct {
			payload *commentPostPayload
		}{
			payload: &commentPostPayload{
				Status:  "approved",
				Comment: "LGTM",
			},
		}

		ctrl := gomock.NewController(t)
		m := mock.NewMockInteractor(ctrl)

		t.Log("MOCK - Find the deployment.")
		m.
			EXPECT().
			FindDeploymentOfRepoByNumber(gomock.Any(), gomock.AssignableToTypeOf(&ent.Repo{}), gomock.Any()).
			Return(&ent.Deployment{}, nil)

		t.Log("MOCK - Create a new deployment with the payload.")
		m.
			EXPECT().
			CreateComment(gomock.Any(), gomock.Eq(&ent.Comment{
				Status:  comment.StatusApproved,
				Comment: "LGTM",
			})).
			DoAndReturn(func(ctx context.Context, cmt *ent.Comment) (*ent.Comment, error) {
				return cmt, nil
			})

		m.
			EXPECT().
			FindCommentByID(gomock.Any(), gomock.Any()).
			Return(&ent.Comment{}, nil)

		r := NewRepo(RepoConfig{}, m)
		router := gin.New()
		router.POST("/deployments/:number/comments", func(c *gin.Context) {
			c.Set(gb.KeyUser, &ent.User{})
			c.Set(KeyRepo, &ent.Repo{})
		}, r.CreateComment)

		body, _ := json.Marshal(input.payload)
		req, _ := http.NewRequest("POST", "/deployments/1/comments", bytes.NewBuffer(body))

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusCreated {
			t.Fatalf("Code = %v, wanted %v", w.Code, http.StatusCreated)
		}
	})
}
