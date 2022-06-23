package repos

import (
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

func TestDeploymentAPI_Rollback(t *testing.T) {
	t.Run("Create a new roll-back.", func(t *testing.T) {
		t.Log("Start mocking:")
		ctrl := gomock.NewController(t)
		m := mock.NewMockInteractor(ctrl)

		t.Log("\tFind the deployment by the number.")
		m.
			EXPECT().
			FindDeploymentOfRepoByNumber(gomock.Any(), gomock.AssignableToTypeOf(&ent.Repo{}), gomock.AssignableToTypeOf(1)).
			Return(&ent.Deployment{Env: "prod"}, nil)

		t.Log("\tGet the evaludated config.")
		m.
			EXPECT().
			GetEvaluatedConfig(gomock.Any(), gomock.AssignableToTypeOf(&ent.User{}), gomock.AssignableToTypeOf(&ent.Repo{}), gomock.AssignableToTypeOf(&extent.EvalValues{})).
			Return(&extent.Config{Envs: []*extent.Env{{Name: "prod"}}}, nil)

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
		router.POST("/deployments/:number/rollback", func(c *gin.Context) {
			c.Set(global.KeyUser, &ent.User{})
			c.Set(KeyRepo, &ent.Repo{})
		}, s.Rollback)

		req, _ := http.NewRequest("POST", "/deployments/1/rollback", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		if w.Code != http.StatusCreated {
			t.Fatalf("Code = %v, wanted %v. Body=%v", w.Code, http.StatusCreated, w.Body)
		}
	})
}
