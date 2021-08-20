package hooks

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/internal/server/hooks/mock"
)

func init() {
	gin.SetMode("release")
}

func TestHook_HandleHook(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockInteractor(ctrl)

	m.
		EXPECT().
		FindDeploymentByUID(gomock.Any(), gomock.Eq(int64(145988746))).
		Return(&ent.Deployment{
			ID:  1,
			UID: 145988746,
		}, nil)

	t.Run("", func(t *testing.T) {
		m.
			EXPECT().
			CreateDeploymentStatus(gomock.Any(), gomock.Eq(&ent.DeploymentStatus{
				Status:       "success",
				Description:  "Deployed successfully.",
				DeploymentID: 1,
			})).
			Return(&ent.DeploymentStatus{
				ID:           1,
				Status:       "success",
				Description:  "Deployed successfully.",
				DeploymentID: 1,
			}, nil)

		m.
			EXPECT().
			UpdateDeployment(gomock.Any(), gomock.Eq(&ent.Deployment{
				ID:     1,
				UID:    145988746,
				Status: "success",
			})).
			Return(&ent.Deployment{
				ID:     1,
				UID:    145988746,
				Status: "success",
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
