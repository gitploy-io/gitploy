package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	i "github.com/gitploy-io/gitploy/internal/interactor"
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

	r, err := w.i.FindRepoOfUserByNamespaceName(ctx, u, &i.FindRepoOfUserByNamespaceNameOptions{
		Namespace: namespace,
		Name:      name,
	})
	if err != nil {
		w.log.Check(gb.GetZapLogLevel(err), "Failed to get the repository.").Write(zap.Error(err))
		c.String(http.StatusForbidden, "It has failed to get the repository.")
		return
	}

	url, err := w.i.GetConfigRedirectURL(ctx, u, r)
	if err != nil {
		w.log.Check(gb.GetZapLogLevel(err), "Failed to get the redirect URL for the configuration file.").Write(zap.Error(err))
		c.String(http.StatusInternalServerError, "It has failed to get the redirect URL for the configuration file.")
		return
	}

	w.log.Debug("Redirect to the URL.", zap.String("URL", url))
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

	r, err := w.i.FindRepoOfUserByNamespaceName(ctx, u, &i.FindRepoOfUserByNamespaceNameOptions{
		Namespace: namespace,
		Name:      name,
	})
	if err != nil {
		w.log.Check(gb.GetZapLogLevel(err), "Failed to get the repository.").Write(zap.Error(err))
		c.String(http.StatusForbidden, "It has failed to get the repository.")
		return
	}

	url, err := w.i.GetNewConfigRedirectURL(ctx, u, r)
	if err != nil {
		w.log.Check(gb.GetZapLogLevel(err), "Failed to get the redirect URL to create a new file.").Write(zap.Error(err))
		c.String(http.StatusInternalServerError, "It has failed to get the redirect URL for the configuration file.")
		return
	}

	w.log.Debug("Redirect to the URL.", zap.String("URL", url))
	c.Redirect(http.StatusFound, url)
}
