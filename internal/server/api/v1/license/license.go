package license

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
)

type (
	Licenser struct {
		i Interactor
	}
)

func NewLicenser(intr Interactor) *Licenser {
	return &Licenser{
		i: intr,
	}
}

func (l *Licenser) GetLicense(c *gin.Context) {
	ctx := c.Request.Context()

	lic, err := l.i.GetLicense(ctx)
	if err != nil {
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to get the license.")
		return
	}

	gb.Response(c, http.StatusOK, lic)
}
