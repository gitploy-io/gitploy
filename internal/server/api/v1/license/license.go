package license

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"go.uber.org/zap"
)

type (
	Licenser struct {
		i   Interactor
		log *zap.Logger
	}
)

func NewLicenser(intr Interactor) *Licenser {
	return &Licenser{
		i:   intr,
		log: zap.L().Named("license"),
	}
}

func (l *Licenser) GetLicense(c *gin.Context) {
	ctx := c.Request.Context()

	lic, err := l.i.GetLicense(ctx)
	if err != nil {
		gb.LogWithError(l.log, "It has failed to get the license.", err)
		gb.ResponseWithError(c, err)
		return
	}

	gb.Response(c, http.StatusOK, lic)
}
