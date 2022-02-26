package hooks

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/AlekSi/pointer"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v42/github"

	"github.com/gitploy-io/gitploy/internal/server/hooks/mock"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/ent/deployment"
	"github.com/gitploy-io/gitploy/model/extent"
)

func init() {
	gin.SetMode("release")
}

func TestHook_HandleHook(t *testing.T) {
	t.Run("Listen the deployment event.", func(t *testing.T) {
		e := &github.DeploymentStatusEvent{}
		bytes, _ := ioutil.ReadFile("./testdata/github.deployment_status.json")
		if err := json.Unmarshal(bytes, &e); err != nil {
			t.Fatalf("It has failed to unmarshal: %s", err)
		}

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := mock.NewMockInteractor(ctrl)

		m.
			EXPECT().
			FindDeploymentByUID(gomock.Any(), gomock.Eq(int64(*e.Deployment.ID))).
			Return(&ent.Deployment{
				ID:  1,
				UID: *e.Deployment.ID,
			}, nil)

		m.
			EXPECT().
			SyncDeploymentStatus(gomock.Any(), gomock.Eq(&ent.DeploymentStatus{
				Status:       *e.DeploymentStatus.State,
				Description:  *e.DeploymentStatus.Description,
				CreatedAt:    e.DeploymentStatus.CreatedAt.Time.UTC(),
				UpdatedAt:    e.DeploymentStatus.UpdatedAt.Time.UTC(),
				DeploymentID: 1,
			})).
			Return(&ent.DeploymentStatus{
				ID:           1,
				Status:       *e.DeploymentStatus.State,
				Description:  *e.DeploymentStatus.Description,
				CreatedAt:    e.DeploymentStatus.CreatedAt.Time.UTC(),
				UpdatedAt:    e.DeploymentStatus.UpdatedAt.Time.UTC(),
				DeploymentID: 1,
			}, nil)

		m.
			EXPECT().
			UpdateDeployment(gomock.Any(), gomock.Eq(&ent.Deployment{
				ID:     1,
				UID:    *e.Deployment.ID,
				Status: deployment.StatusSuccess,
			})).
			Return(&ent.Deployment{
				ID:     1,
				UID:    *e.Deployment.ID,
				Status: deployment.StatusSuccess,
			}, nil)

		m.
			EXPECT().
			CreateEvent(gomock.Any(), gomock.Any()).
			Return(&ent.Event{}, nil)

		h := NewHooks(&ConfigHooks{}, m)
		r := gin.New()
		r.POST("/hooks", h.HandleHook)

		// Build the Github webhook.
		json, err := os.Open("./testdata/github.deployment_status.json")
		if err != nil {
			t.Errorf("It has failed to open the JSON file: %s", err)
			t.FailNow()
		}
		req, _ := http.NewRequest("POST", "/hooks", json)
		req.Header.Set(headerGithubDelivery, "72d3162e")
		req.Header.Set(headerGtihubEvent, "deployment_status")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Validate the result.
		if w.Code != http.StatusCreated {
			t.Errorf("w.Code = %d, wanted %d", w.Code, http.StatusCreated)
			t.Logf("w.Body = %v", w.Body)
			t.FailNow()
		}
	})

	t.Run("Listen the push event.", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := mock.NewMockInteractor(ctrl)

		m.
			EXPECT().
			FindRepoByID(gomock.Any(), gomock.Any()).
			Return(&ent.Repo{
				Edges: ent.RepoEdges{
					Owner: &ent.User{},
				},
			}, nil)

		t.Log("Return the auto-deployment environment.")
		m.
			EXPECT().
			GetConfig(gomock.Any(), gomock.AssignableToTypeOf(&ent.User{}), gomock.AssignableToTypeOf(&ent.Repo{})).
			Return(&extent.Config{
				Envs: []*extent.Env{
					{
						Name: "dev",
					},
					{
						Name:         "production",
						AutoDeployOn: pointer.ToString("refs/tags/.*"),
					},
				},
			}, nil)

		m.
			EXPECT().
			Deploy(
				gomock.Any(),
				gomock.AssignableToTypeOf(&ent.User{}),
				gomock.AssignableToTypeOf(&ent.Repo{}),
				gomock.Eq(&ent.Deployment{
					Type: deployment.TypeTag,
					Ref:  "simple-tag",
					Env:  "production",
				}),
				gomock.AssignableToTypeOf(&extent.Env{}),
			).
			Return(&ent.Deployment{}, nil)

		m.
			EXPECT().
			CreateEvent(gomock.Any(), gomock.AssignableToTypeOf(&ent.Event{})).
			Return(&ent.Event{}, nil)

		h := NewHooks(&ConfigHooks{}, m)
		r := gin.New()
		r.POST("/hooks", h.HandleHook)

		// Build the Github webhook.
		json, err := os.Open("./testdata/github.push.json")
		if err != nil {
			t.Errorf("It has failed to open the JSON file: %s", err)
			t.FailNow()
		}
		req, _ := http.NewRequest("POST", "/hooks", json)
		req.Header.Set(headerGithubDelivery, "72d3162e")
		req.Header.Set(headerGtihubEvent, "push")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Validate the result.
		if w.Code != http.StatusOK {
			t.Errorf("w.Code = %d, wanted %d", w.Code, http.StatusOK)
			t.Logf("w.Body = %v", w.Body)
			t.FailNow()
		}
	})
}
