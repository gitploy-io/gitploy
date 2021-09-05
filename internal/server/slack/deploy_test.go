package slack

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/deployment"
	"github.com/gitploy-io/gitploy/internal/server/slack/mock"
	"github.com/gitploy-io/gitploy/vo"
	"github.com/golang/mock/gomock"
	"go.uber.org/zap"
)

func TestSlack_interactDeploy(t *testing.T) {
	t.Run("Create a new deployment with payload.", func(t *testing.T) {
		m := mock.NewMockInteractor(gomock.NewController(t))

		// These values are in "./testdata/deploy-interact.json"
		const (
			callbackID = "nafyVuEqzcchuVmV"
			chatUserID = "U025KUBB2"
			branch     = "main"
			env        = "prod"
		)

		t.Log("Find the callback which was stored by the Slash command.")
		m.
			EXPECT().
			FindCallbackByHash(gomock.Any(), callbackID).
			Return(&ent.Callback{}, nil)

		t.Log("Find the chat-user who sent the payload.")
		m.
			EXPECT().
			FindChatUserByID(gomock.Any(), chatUserID).
			Return(&ent.ChatUser{}, nil)

		t.Log("Get branch to validate the payload.")
		m.
			EXPECT().
			GetBranch(gomock.Any(), gomock.AssignableToTypeOf(&ent.User{}), gomock.AssignableToTypeOf(&ent.Repo{}), branch).
			Return(&vo.Branch{
				Name: branch,
			}, nil)

		t.Log("Get the config file of the repository.")
		m.
			EXPECT().
			GetConfig(gomock.Any(), gomock.AssignableToTypeOf(&ent.User{}), gomock.AssignableToTypeOf(&ent.Repo{})).
			Return(&vo.Config{
				Envs: []*vo.Env{
					{Name: env},
				},
			}, nil)

		t.Log("Get the next number of deployment.")
		m.
			EXPECT().
			GetNextDeploymentNumberOfRepo(gomock.Any(), gomock.AssignableToTypeOf(&ent.Repo{})).
			Return(4, nil)

		t.Log("Deploy with the payload.")
		m.
			EXPECT().
			Deploy(gomock.Any(), gomock.AssignableToTypeOf(&ent.User{}), gomock.AssignableToTypeOf(&ent.Repo{}), &ent.Deployment{
				Number: 4,
				Type:   deployment.TypeBranch,
				Ref:    branch,
				Env:    env,
			}, gomock.AssignableToTypeOf(&vo.Env{})).
			DoAndReturn(func(ctx context.Context, u *ent.User, r *ent.Repo, d *ent.Deployment, e *vo.Env) (*ent.Deployment, error) {
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
		router.POST("/interact", s.interactDeploy)

		// Build the Slack payload.
		bytes, err := ioutil.ReadFile("./testdata/deploy-interact.json")
		if err != nil {
			t.Errorf("It has failed to open the JSON file: %s", err)
			t.FailNow()
		}

		form := url.Values{}
		form.Add("payload", string(bytes))
		req, _ := http.NewRequest("POST", "/interact", strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("w.Code = %d, wanted %d", w.Code, http.StatusOK)
			t.Logf("w.Body = %v", w.Body)
			t.FailNow()
		}
	})
}
