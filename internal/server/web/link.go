package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
)

// RedirectToConfig redirects to the URL to read the configuration file.
func (w *Web) RedirectToConfig(c *gin.Context) {
	var (
		namespace = c.Param("namespace")
		name      = c.Param("name")
	)

	ctx := c.Request.Context()

	uv, _ := c.Get(gb.KeyUser)
	u := uv.(*ent.User)

	r, err := w.i.FindRepoOfUserByNamespaceName(ctx, u, namespace, name)
	if err != nil {
		w.log.Check(gb.GetZapLogLevel(err), "Failed to get the repository.").Write(zap.Error(err))
		gb.AbortWithError(c, err)
		return
	}

	url, err := w.i.GetConfigRedirectURL(ctx, u, r)
	if err != nil {
		w.log.Check(gb.GetZapLogLevel(err), "Failed to get the redirect URL for the configuration file.").Write(zap.Error(err))
		gb.AbortWithError(c, err)
		return
	}

	c.Redirect(http.StatusMovedPermanently, url)
}

// RedirectToNewConfig redirect to the URL to create a new file.
func (w *Web) RedirectToNewConfig(c *gin.Context) {
	var (
		namespace = c.Param("namespace")
		name      = c.Param("name")
	)

	ctx := c.Request.Context()

	uv, _ := c.Get(gb.KeyUser)
	u := uv.(*ent.User)

	r, err := w.i.FindRepoOfUserByNamespaceName(ctx, u, namespace, name)
	if err != nil {
		w.log.Check(gb.GetZapLogLevel(err), "Failed to get the repository.").Write(zap.Error(err))
		gb.AbortWithError(c, err)
		return
	}

	url, err := w.i.GetNewFileRedirectURL(ctx, u, r)
	if err != nil {
		w.log.Check(gb.GetZapLogLevel(err), "Failed to get the redirect URL to create a new file.").Write(zap.Error(err))
		gb.AbortWithError(c, err)
		return
	}

	c.Redirect(http.StatusMovedPermanently, url)
}
