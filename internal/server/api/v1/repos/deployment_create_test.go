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
	"github.com/gitploy-io/gitploy/model/extent"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func TestDeploymentAPI_Create(t *testing.T) {
	t.Run("Create a new deployment.", func(t *testing.T) {
		t.Log("Start mocking:")
		ctrl := gomock.NewController(t)
		m := mock.NewMockInteractor(ctrl)

		t.Log("\tGet the evaludated config.")
		m.
			EXPECT().
			GetEvaluatedConfig(gomock.Any(), gomock.AssignableToTypeOf(&ent.User{}), gomock.AssignableToTypeOf(&ent.Repo{}), gomock.AssignableToTypeOf(&extent.EvalValues{})).
			Return(&extent.Config{
				Envs: []*extent.Env{
					{Name: "prod"},
				}}, nil)

		t.Log("\tCreate a new deployment, and dispatch a event.")
		m.
			EXPECT().
			Deploy(gomock.Any(),
				gomock.AssignableToTypeOf(&ent.User{}),
				gomock.AssignableToTypeOf(&ent.Repo{}),
				gomock.AssignableToTypeOf(&ent.Deployment{}),
				gomock.AssignableToTypeOf(&extent.Env{})).
			Return(&ent.Deployment{}, nil)

		m.
			EXPECT().
			FindDeploymentByID(gomock.Any(), gomock.Any()).
			Return(&ent.Deployment{}, nil)

		s := DeploymentAPI{i: m, log: zap.L()}

		router := gin.New()
		router.POST("/repos/:id/deployments", func(c *gin.Context) {
			c.Set(global.KeyUser, &ent.User{})
			c.Set(KeyRepo, &ent.Repo{})
		}, s.Create)

		body, _ := json.Marshal(&DeploymentPostPayload{
			Type: "branch",
			Ref:  "main",
			Env:  "prod",
		})
		req, _ := http.NewRequest("POST", "/repos/1/deployments", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		if w.Code != http.StatusCreated {
			t.Fatalf("Code = %v, wanted %v. Body=%v", w.Code, http.StatusCreated, w.Body)
		}
	})
}
