package slack

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/deployment"
	"github.com/hanjunlee/gitploy/ent/notification"
	"github.com/hanjunlee/gitploy/internal/server/slack/mock"
	"github.com/hanjunlee/gitploy/vo"
	"go.uber.org/zap"
)

func TestSlack_interactDeploy(t *testing.T) {
	t.Run("deploy the branch and request to approvers successfully.", func(t *testing.T) {
		r := &ent.Repo{
			ID: "1",
		}

		u := &ent.User{
			ID: "1",
		}

		cu := &ent.ChatUser{
			ID:     "U025KUBB2",
			UserID: u.ID,
			Edges: ent.ChatUserEdges{
				User: u,
			},
		}

		cb := &ent.ChatCallback{
			ID:         1,
			RepoID:     r.ID,
			ChatUserID: cu.ID,
			Edges: ent.ChatCallbackEdges{
				ChatUser: cu,
				Repo:     r,
			},
		}

		m := mock.NewMockInteractor(gomock.NewController(t))

		// These values are equal to the mocking payload.
		const (
			callbackID = "nafyVuEqzcchuVmV"
			chatUserID = "U025KUBB2"
			branch     = "main"
			env        = "prod"
			number     = 4
		)

		m.
			EXPECT().
			FindChatCallbackByHash(gomock.Any(), callbackID).
			Return(cb, nil)

		m.
			EXPECT().
			FindChatUserByID(gomock.Any(), chatUserID).
			Return(cu, nil)

		m.
			EXPECT().
			GetBranch(gomock.Any(), u, r, branch).
			Return(&vo.Branch{
				Name:      branch,
				CommitSHA: "commit_sha",
			}, nil)

		m.
			EXPECT().
			GetConfig(gomock.Any(), u, r).
			Return(&vo.Config{
				Envs: []*vo.Env{
					{
						Name: env,
						Approval: &vo.Approval{
							Enabled: false,
						},
					},
				},
			}, nil)

		m.
			EXPECT().
			GetNextDeploymentNumberOfRepo(gomock.Any(), r).
			Return(number, nil)

		m.
			EXPECT().
			Deploy(gomock.Any(), u, r, &ent.Deployment{
				Number: number,
				Type:   deployment.TypeBranch,
				Ref:    branch,
				Env:    env,
			}, gomock.Any()).
			Return(&ent.Deployment{
				ID:     10,
				Number: number,
				Type:   deployment.TypeBranch,
				Ref:    branch,
				Env:    env,
			}, nil)

		m.
			EXPECT().
			Publish(gomock.Any(), notification.TypeDeploymentCreated, r, gomock.Any(), nil).
			Return(nil)

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
