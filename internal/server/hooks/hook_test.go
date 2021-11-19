package hooks

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v32/github"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/deployment"
	"github.com/gitploy-io/gitploy/internal/server/hooks/mock"
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
		json, err := os.Open("./testdata/github.hook.json")
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
}
