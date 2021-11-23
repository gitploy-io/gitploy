package repos

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/deployment"
	"github.com/gitploy-io/gitploy/internal/server/api/v1/repos/mock"
	"github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/vo"
)

func TestRepo_ListDeploymentChanges(t *testing.T) {
	ctx := gomock.Any()
	any := gomock.Any()

	t.Run("Return commits successfully", func(t *testing.T) {
		input := struct {
			number  int
			page    int
			perPage int
		}{
			number:  5,
			page:    1,
			perPage: 30,
		}

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := mock.NewMockInteractor(ctrl)

		const (
			base = "ee42de2"
			head = "231eed1"
		)

		m.
			EXPECT().
			FindDeploymentOfRepoByNumber(ctx, any, gomock.Eq(input.number)).
			Return(&ent.Deployment{
				ID:     7,
				Number: input.number,
				Sha:    head,
				Status: deployment.StatusCreated,
			}, nil)

		m.
			EXPECT().
			FindPrevSuccessDeployment(ctx, any).
			Return(&ent.Deployment{
				ID:     5,
				Sha:    base,
				Status: deployment.StatusSuccess,
			}, nil)

		m.
			EXPECT().
			CompareCommits(ctx, any, any, base, head, gomock.Eq(input.page), gomock.Eq(input.perPage)).
			Return([]*vo.Commit{
				{
					SHA: head,
				},
			}, []*vo.CommitFile{}, nil)

		// Ready the router to handle it.
		gin.SetMode(gin.ReleaseMode)

		repos := NewRepo(RepoConfig{}, m)
		router := gin.New()
		router.GET("/deployments/:number/changes", func(c *gin.Context) {
			// Mocking middlewares to return a user and a repository.
			c.Set(global.KeyUser, &ent.User{})
			c.Set(KeyRepo, &ent.Repo{})
		}, repos.ListDeploymentChanges)

		req, _ := http.NewRequest("GET", fmt.Sprintf("/deployments/%d/changes?page=%d&per_page=%d", input.number, input.page, input.perPage), nil)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("w.Code = %d, wanted %d", w.Code, http.StatusCreated)
			t.FailNow()
		}

		expected := []*vo.Commit{
			{
				SHA: head,
			},
		}
		eb, _ := json.Marshal(expected)
		if bytes := w.Body.Bytes(); string(bytes) != string(eb) {
			t.Errorf("w.Body = %s, wanted %s", string(bytes), string(eb))
			t.FailNow()
		}
	})
}

func TestRepo_CreateDeployment(t *testing.T) {
	t.Run("a new deployment entity.", func(t *testing.T) {
		input := struct {
			payload *deploymentPostPayload
		}{
			payload: &deploymentPostPayload{
				Type: "branch",
				Ref:  "main",
				Env:  "prod",
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
						Name: "prod",
					},
				}}, nil)

		t.Log("Deploy with the payload successfully.")
		m.
			EXPECT().
			Deploy(gomock.Any(), gomock.AssignableToTypeOf(&ent.User{}), gomock.AssignableToTypeOf(&ent.Repo{}), gomock.Eq(&ent.Deployment{
				Type: deployment.Type(input.payload.Type),
				Env:  input.payload.Env,
				Ref:  input.payload.Ref,
			}), gomock.AssignableToTypeOf(&vo.Env{})).
			Return(&ent.Deployment{}, nil)

		t.Log("Dispatch the event.")
		m.
			EXPECT().
			CreateEvent(gomock.Any(), gomock.AssignableToTypeOf(&ent.Event{})).
			Return(&ent.Event{}, nil)

		t.Log("Read the deployment with edges.")
		m.
			EXPECT().
			FindDeploymentByID(gomock.Any(), gomock.Any()).
			Return(&ent.Deployment{}, nil)

		r := NewRepo(RepoConfig{}, m)

		gin.SetMode(gin.ReleaseMode)
		router := gin.New()
		router.POST("/repos/:id/deployments", func(c *gin.Context) {
			c.Set(global.KeyUser, &ent.User{})
			c.Set(KeyRepo, &ent.Repo{})
		}, r.CreateDeployment)

		body, _ := json.Marshal(input.payload)
		req, _ := http.NewRequest("POST", "/repos/1/deployments", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		if w.Code != http.StatusCreated {
			t.Fatalf("Code = %v, wanted %v. Body=%v", w.Code, http.StatusCreated, w.Body)
		}
	})
}
