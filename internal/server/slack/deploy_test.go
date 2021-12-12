package slack

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/slack-go/slack"
	"go.uber.org/zap"

	"github.com/gitploy-io/gitploy/internal/server/slack/mock"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/ent/deployment"
	"github.com/gitploy-io/gitploy/model/extent"
)

func TestSlack_interactDeploy(t *testing.T) {
	t.Run("Create a new deployment with payload.", func(t *testing.T) {
		m := mock.NewMockInteractor(gomock.NewController(t))

		// These values are in "./testdata/deploy-interact.json"
		const (
			callbackID = "nafyVuEqzcchuVmV"
			branch     = "main"
			env        = "prod"
		)

		t.Log("Find the callback which was stored by the Slash command.")
		m.
			EXPECT().
			FindCallbackByHash(gomock.Any(), callbackID).
			Return(&ent.Callback{
				Edges: ent.CallbackEdges{
					Repo: &ent.Repo{ID: 1},
				},
			}, nil)

		t.Log("Get branch to validate the payload.")
		m.
			EXPECT().
			GetBranch(gomock.Any(), gomock.AssignableToTypeOf(&ent.User{}), gomock.AssignableToTypeOf(&ent.Repo{}), branch).
			Return(&extent.Branch{
				Name: branch,
			}, nil)

		t.Log("Get the config file of the repository.")
		m.
			EXPECT().
			GetConfig(gomock.Any(), gomock.AssignableToTypeOf(&ent.User{}), gomock.AssignableToTypeOf(&ent.Repo{})).
			Return(&extent.Config{
				Envs: []*extent.Env{
					{Name: env},
				},
			}, nil)

		t.Log("Deploy with the payload.")
		m.
			EXPECT().
			Deploy(gomock.Any(), gomock.AssignableToTypeOf(&ent.User{}), gomock.AssignableToTypeOf(&ent.Repo{}), &ent.Deployment{
				Type: deployment.TypeBranch,
				Ref:  branch,
				Env:  env,
			}, gomock.AssignableToTypeOf(&extent.Env{})).
			DoAndReturn(func(ctx context.Context, u *ent.User, r *ent.Repo, d *ent.Deployment, e *extent.Env) (*ent.Deployment, error) {
				return d, nil
			})

		t.Log("Create a new event")
		m.
			EXPECT().
			CreateEvent(gomock.Any(), gomock.AssignableToTypeOf(&ent.Event{})).
			Return(&ent.Event{}, nil)

		s := &Slack{
			i:   m,
			log: zap.L(),
		}

		gin.SetMode(gin.ReleaseMode)
		router := gin.New()
		router.POST("/interact", func(c *gin.Context) {
			bytes, _ := ioutil.ReadFile("./testdata/deploy-interact.json")
			intr := slack.InteractionCallback{}
			intr.UnmarshalJSON(bytes)
			c.Set(KeyIntr, intr)
			c.Set(KeyChatUser, &ent.ChatUser{
				Edges: ent.ChatUserEdges{
					User: &ent.User{},
				},
			})
		}, s.interactDeploy)

		req, _ := http.NewRequest("POST", "/interact", nil)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("w.Code = %d, wanted %d. Body = %v", w.Code, http.StatusOK, w.Body)
		}
	})
}
