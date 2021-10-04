package sync

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/internal/server/api/v1/sync/mock"
	"github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/vo"
	"github.com/golang/mock/gomock"
)

func TestSyncher_Sync(t *testing.T) {
	ctx := gomock.Any()

	t.Run("Synchronize with remote repositories", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := mock.NewMockInteractor(ctrl)

		t.Log("List remote repositories")
		m.
			EXPECT().
			ListRemoteRepos(ctx, gomock.AssignableToTypeOf(&ent.User{})).
			Return([]*vo.RemoteRepo{
				{
					ID:        1,
					Namespace: "octocat",
					Name:      "HelloWorld",
				},
				{
					ID:        1,
					Namespace: "cat",
					Name:      "GoodBye",
				},
			}, nil)

		t.Log("Only octocat is trusted namespace.")
		m.
			EXPECT().
			IsEntryOrg(ctx, gomock.Any()).
			DoAndReturn(func(ctx context.Context, namespace string) bool {
				return namespace == "octocat"
			}).
			AnyTimes()

		t.Log("Sync with HelloWorld only.")
		m.
			EXPECT().
			SyncRemoteRepo(ctx, gomock.Any(), gomock.Eq(&vo.RemoteRepo{
				ID:        1,
				Namespace: "octocat",
				Name:      "HelloWorld",
			}), gomock.AssignableToTypeOf(time.Time{})).
			Return(nil)

		t.Log("Delete staled perms.")
		m.
			EXPECT().
			DeletePermsOfUserLessThanSyncedAt(ctx, gomock.Any(), gomock.Any()).
			Return(0, nil)

		gin.SetMode(gin.ReleaseMode)
		router := gin.New()

		s := NewSyncher(m)
		router.POST("/sync", func(c *gin.Context) {
			c.Set(global.KeyUser, &ent.User{})
		}, s.Sync)

		req, _ := http.NewRequest("POST", "/sync", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		if w.Code != http.StatusOK {
			t.Fatalf("Sync = %v, wanted %v", w.Code, http.StatusOK)
		}
	})
}
