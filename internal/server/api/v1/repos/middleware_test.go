package repos

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"

	i "github.com/gitploy-io/gitploy/internal/interactor"
	"github.com/gitploy-io/gitploy/internal/server/api/v1/repos/mock"
	"github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/ent/perm"
)

func TestRepoMiddleware_RepoWritePerm(t *testing.T) {
	ctx := gomock.Any()

	t.Run("Return 403 error when the permission is read.", func(t *testing.T) {
		t.Log("Start mocking:")
		ctrl := gomock.NewController(t)
		m := mock.NewMockInteractor(ctrl)

		t.Logf("\tFind the repository, and get the read permission.")
		m.
			EXPECT().
			FindRepoOfUserByNamespaceName(ctx, gomock.AssignableToTypeOf(&ent.User{}), gomock.AssignableToTypeOf(&i.FindRepoOfUserByNamespaceNameOptions{})).
			Return(&ent.Repo{
				Namespace: "octocat",
				Name:      "hello-world",
			}, nil)

		m.
			EXPECT().
			FindPermOfRepo(ctx, gomock.AssignableToTypeOf(&ent.Repo{}), gomock.AssignableToTypeOf(&ent.User{})).
			Return(&ent.Perm{
				RepoPerm: perm.RepoPermRead,
			}, nil)

		gin.SetMode(gin.ReleaseMode)
		router := gin.New()

		rm := NewRepoMiddleware(m)
		router.PATCH("/repos/:namespace/:name", func(c *gin.Context) {
			// Mocking middlewares to return a user and a repository.
			c.Set(global.KeyUser, &ent.User{})
		}, rm.RepoWritePerm(), func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		req, _ := http.NewRequest("PATCH", fmt.Sprintf("/repos/%s/%s", "octocat", "hello-world"), nil)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusForbidden {
			t.Fatalf("RepoWritePerm = %v, wanted %v", w.Code, http.StatusForbidden)
		}
	})

	t.Run("Return 200 when the permission is write.", func(t *testing.T) {
		t.Log("Start mocking:")
		ctrl := gomock.NewController(t)
		m := mock.NewMockInteractor(ctrl)

		t.Logf("\tFind the repository, and get the write permission.")
		m.
			EXPECT().
			FindRepoOfUserByNamespaceName(ctx, gomock.AssignableToTypeOf(&ent.User{}), gomock.AssignableToTypeOf(&i.FindRepoOfUserByNamespaceNameOptions{})).
			Return(&ent.Repo{
				Namespace: "octocat",
				Name:      "hello-world",
			}, nil)

		m.
			EXPECT().
			FindPermOfRepo(ctx, gomock.AssignableToTypeOf(&ent.Repo{}), gomock.AssignableToTypeOf(&ent.User{})).
			Return(&ent.Perm{
				RepoPerm: perm.RepoPermWrite,
			}, nil)

		gin.SetMode(gin.ReleaseMode)
		router := gin.New()

		rm := NewRepoMiddleware(m)
		router.PATCH("/repos/:namespace/:name", func(c *gin.Context) {
			// Mocking middlewares to return a user and a repository.
			c.Set(global.KeyUser, &ent.User{})
		}, rm.RepoWritePerm(), func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		req, _ := http.NewRequest("PATCH", fmt.Sprintf("/repos/%s/%s", "octocat", "hello-world"), nil)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("RepoWritePerm = %v, wanted %v", w.Code, http.StatusOK)
		}
	})
}

func TestRepoMiddleware_RepoAdminPerm(t *testing.T) {
	ctx := gomock.Any()

	t.Run("Return 200 when the permission is admin.", func(t *testing.T) {
		t.Log("Start mocking:")
		ctrl := gomock.NewController(t)
		m := mock.NewMockInteractor(ctrl)

		t.Logf("\tFind the repository, and get the admin permission.")
		m.
			EXPECT().
			FindRepoOfUserByNamespaceName(ctx, gomock.AssignableToTypeOf(&ent.User{}), gomock.AssignableToTypeOf(&i.FindRepoOfUserByNamespaceNameOptions{})).
			Return(&ent.Repo{
				Namespace: "octocat",
				Name:      "hello-world",
			}, nil)

		t.Logf("It returns the read permission.")
		m.
			EXPECT().
			FindPermOfRepo(ctx, gomock.AssignableToTypeOf(&ent.Repo{}), gomock.AssignableToTypeOf(&ent.User{})).
			Return(&ent.Perm{
				RepoPerm: perm.RepoPermAdmin,
			}, nil)

		gin.SetMode(gin.ReleaseMode)
		router := gin.New()

		rm := NewRepoMiddleware(m)
		router.PATCH("/repos/:namespace/:name", func(c *gin.Context) {
			// Mocking middlewares to return a user and a repository.
			c.Set(global.KeyUser, &ent.User{})
		}, rm.RepoWritePerm(), func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		req, _ := http.NewRequest("PATCH", fmt.Sprintf("/repos/%s/%s", "octocat", "hello-world"), nil)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("RepoWritePerm = %v, wanted %v", w.Code, http.StatusOK)
		}
	})
}
