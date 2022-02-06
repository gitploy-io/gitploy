package repos

import (
	"encoding/json"
	"fmt"
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

func TestDeploymentsAPI_ListChanges(t *testing.T) {
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
			Return([]*extent.Commit{
				{
					SHA: head,
				},
			}, []*extent.CommitFile{}, nil)

		// Ready the router to handle it.
		gin.SetMode(gin.ReleaseMode)

		s := DeploymentsAPI{i: m, log: zap.L()}

		router := gin.New()
		router.GET("/deployments/:number/changes", func(c *gin.Context) {
			// Mocking middlewares to return a user and a repository.
			c.Set(global.KeyUser, &ent.User{})
			c.Set(KeyRepo, &ent.Repo{})
		}, s.ListChanges)

		req, _ := http.NewRequest("GET", fmt.Sprintf("/deployments/%d/changes?page=%d&per_page=%d", input.number, input.page, input.perPage), nil)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("w.Code = %d, wanted %d", w.Code, http.StatusCreated)
			t.FailNow()
		}

		expected := []*extent.Commit{
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
