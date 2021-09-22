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
	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/internal/server/api/v1/repos/mock"
	"github.com/golang/mock/gomock"
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
			CreateEvent(ctx, any).
			Return(&ent.Event{}, nil)

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
		}
	})
}

func TestRepo_DeleteApproval(t *testing.T) {
	t.Run("Delete the approval.", func(t *testing.T) {
		input := struct {
			aid int
		}{
			aid: 7,
		}

		ctrl := gomock.NewController(t)
		m := mock.NewMockInteractor(ctrl)

		t.Log("Find the approval by ID.")
		m.
			EXPECT().
			FindApprovalByID(gomock.Any(), input.aid).
			Return(&ent.Approval{ID: input.aid}, nil)

		t.Log("Delete the approval.")
		m.
			EXPECT().
			DeleteApproval(gomock.Any(), gomock.AssignableToTypeOf(&ent.Approval{})).
			Return(nil)

		t.Log("Create a deleted event.")
		m.
			EXPECT().
			CreateEvent(gomock.Any(), gomock.AssignableToTypeOf(&ent.Event{})).
			Return(&ent.Event{ID: 1}, nil)

		gin.SetMode(gin.ReleaseMode)

		r := NewRepo(RepoConfig{}, m)
		router := gin.New()
		router.DELETE("/deployments/:number/approvals/:aid", r.DeleteApproval)

		req, _ := http.NewRequest("DELETE", fmt.Sprintf("/deployments/1/approvals/%d", input.aid), nil)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("Code != %d, wanted %d", w.Code, http.StatusOK)
		}
	})
}
