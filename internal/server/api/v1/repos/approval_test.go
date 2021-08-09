package repos

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/internal/server/api/v1/repos/mock"
)

var (
	ctx = gomock.Any()
	any = gomock.Any()
)

func TestRepo_CreateApproval(t *testing.T) {

	t.Run("Create a new approval.", func(t *testing.T) {
		input := struct {
			number  int
			payload *approvalPostPayload
		}{
			number: 7,
			payload: &approvalPostPayload{
				UserID: "4",
			},
		}

		const (
			approvalID = 3
		)

		ctrl := gomock.NewController(t)
		m := mock.NewMockInteractor(ctrl)

		m.
			EXPECT().
			FindDeploymentOfRepoByNumber(ctx, any, input.number).
			Return(&ent.Deployment{
				ID: input.number,
			}, nil)

		// Find the user, and check the permission.
		m.
			EXPECT().
			FindUserByID(ctx, input.payload.UserID).
			Return(&ent.User{
				ID: input.payload.UserID,
			}, nil)

		m.
			EXPECT().
			FindPermOfRepo(ctx, any, &ent.User{
				ID: input.payload.UserID,
			}).
			Return(&ent.Perm{}, nil)

		// Create a new approval and publish.
		m.
			EXPECT().
			CreateApproval(ctx, &ent.Approval{
				DeploymentID: input.number,
				UserID:       input.payload.UserID,
			}).
			DoAndReturn(func(ctx context.Context, a *ent.Approval) (*ent.Approval, error) {
				a.ID = approvalID
				return a, nil
			})

		m.
			EXPECT().
			Publish(ctx, any, any, any, any)

		m.
			EXPECT().
			FindApprovalByID(ctx, approvalID).
			Return(&ent.Approval{
				ID:           approvalID,
				DeploymentID: input.number,
				UserID:       input.payload.UserID,
			}, nil)

		gin.SetMode(gin.ReleaseMode)

		r := NewRepo(RepoConfig{}, m)
		router := gin.New()
		router.POST("/deployments/:number/approvals", func(c *gin.Context) {
			c.Set(KeyRepo, &ent.Repo{})
		}, r.CreateApproval)

		body, _ := json.Marshal(input.payload)
		req, _ := http.NewRequest("POST", fmt.Sprintf("/deployments/%d/approvals", input.number), bytes.NewBuffer(body))

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusCreated {
			t.Fatalf("Code != %d, wanted %d", w.Code, http.StatusOK)
			t.FailNow()
		}
	})
}
