package repos

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"go.uber.org/zap"

	"github.com/gitploy-io/gitploy/internal/server/api/v1/repos/mock"
	"github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/ent/deployment"
	"github.com/gitploy-io/gitploy/model/extent"
)

func TestDeploymentService_Create(t *testing.T) {
	t.Run("a new deployment entity.", func(t *testing.T) {
		input := struct {
			payload *DeploymentPostPayload
		}{
			payload: &DeploymentPostPayload{
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
			Return(&extent.Config{
				Envs: []*extent.Env{
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
			}), gomock.AssignableToTypeOf(&extent.Env{})).
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

		s := DeploymentService{i: m, log: zap.L()}

		gin.SetMode(gin.ReleaseMode)
		router := gin.New()
		router.POST("/repos/:id/deployments", func(c *gin.Context) {
			c.Set(global.KeyUser, &ent.User{})
			c.Set(KeyRepo, &ent.Repo{})
		}, s.Create)

		body, _ := json.Marshal(input.payload)
		req, _ := http.NewRequest("POST", "/repos/1/deployments", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		if w.Code != http.StatusCreated {
			t.Fatalf("Code = %v, wanted %v. Body=%v", w.Code, http.StatusCreated, w.Body)
		}
	})
}
