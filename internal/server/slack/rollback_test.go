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

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/deployment"
	"github.com/gitploy-io/gitploy/internal/server/slack/mock"
	"github.com/gitploy-io/gitploy/vo"
)

func TestSlack_interactRollback(t *testing.T) {
	t.Run("Rollback with the returned deployment.", func(t *testing.T) {
		m := mock.NewMockInteractor(gomock.NewController(t))

		// These values are in "./testdata/rollback-interact.json"
		const (
			callbackID   = "hZUZvJgWhxYvdekUGESXKjSusKWWIRKr"
			chatUserID   = "U025KUBB2"
			deploymentID = 33
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

		t.Log("Find the deployment by ID.")
		m.
			EXPECT().
			FindDeploymentByID(gomock.Any(), deploymentID).
			Return(&ent.Deployment{
				ID:   deploymentID,
				Type: deployment.TypeCommit,
				Ref:  "main",
				Sha:  "ee411aa",
				Env:  "prod",
			}, nil)

		t.Log("Get the config file of the repository.")
		m.
			EXPECT().
			GetConfig(gomock.Any(), gomock.AssignableToTypeOf(&ent.User{}), gomock.AssignableToTypeOf(&ent.Repo{})).
			Return(&vo.Config{
				Envs: []*vo.Env{
					{Name: "prod"},
				},
			}, nil)

		t.Log("Check the lock.")
		m.
			EXPECT().
			HasLockOfRepoForEnv(gomock.Any(), gomock.AssignableToTypeOf(&ent.Repo{}), "prod").
			Return(false, nil)

		t.Log("Get the next number of deployment.")
		m.
			EXPECT().
			GetNextDeploymentNumberOfRepo(gomock.Any(), gomock.AssignableToTypeOf(&ent.Repo{})).
			Return(4, nil)

		t.Log("Roll back with the returned deployment.")
		m.
			EXPECT().
			Deploy(gomock.Any(), gomock.AssignableToTypeOf(&ent.User{}), gomock.AssignableToTypeOf(&ent.Repo{}), &ent.Deployment{
				Number:     4,
				Type:       deployment.TypeCommit,
				Ref:        "main",
				Sha:        "ee411aa",
				Env:        "prod",
				IsRollback: true,
			}, gomock.AssignableToTypeOf(&vo.Env{})).
			DoAndReturn(func(ctx context.Context, u *ent.User, r *ent.Repo, d *ent.Deployment, e *vo.Env) (*ent.Deployment, error) {
				d.ID = deploymentID + 1
				return d, nil
			})

		t.Log("Create a new event")
		m.
			EXPECT().
			CreateEvent(gomock.Any(), gomock.AssignableToTypeOf(&ent.Event{})).
			Return(&ent.Event{}, nil)

		s := &Slack{i: m, log: zap.L()}

		gin.SetMode(gin.ReleaseMode)
		router := gin.New()
		router.POST("/interact", func(c *gin.Context) {
			bytes, _ := ioutil.ReadFile("./testdata/rollback-interact.json")
			intr := slack.InteractionCallback{}
			intr.UnmarshalJSON(bytes)
			c.Set(KeyIntr, intr)
			c.Set(KeyChatUser, &ent.ChatUser{
				Edges: ent.ChatUserEdges{
					User: &ent.User{},
				},
			})
		}, s.interactRollback)

		req, _ := http.NewRequest("POST", "/interact", nil)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("w.Code = %d, wanted %d. Body = %v", w.Code, http.StatusOK, w.Body)
		}
	})
}
