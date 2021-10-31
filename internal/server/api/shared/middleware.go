package shared

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/pkg/e"
)

const (
	extraDuration time.Duration = 30 * 24 * time.Hour
)

type (
	Middleware struct {
		i Interactor
	}
)

func NewMiddleware(intr Interactor) *Middleware {
	return &Middleware{
		i: intr,
	}
}

func (m *Middleware) IsLicenseExpired() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		lic, err := m.i.GetLicense(ctx)
		if err != nil {
			gb.AbortWithStatusAndError(c, http.StatusInternalServerError, err)
			return
		}

		if lic.IsOSS() {
			return
		}

		if lic.IsOverLimit() {
			gb.AbortWithStatusAndError(
				c,
				http.StatusPaymentRequired,
				e.NewErrorWithMessage(e.ErrorCodeLicenseRequired, "The license is over the limit.", nil),
			)
			return
		}

		if lic.IsStandard() && lic.IsExpired() {
			now := time.Now()
			if lic.ExpiredAt.Add(extraDuration).Before(now) {
				gb.AbortWithStatusAndError(
					c,
					http.StatusPaymentRequired,
					e.NewErrorWithMessage(e.ErrorCodeLicenseRequired, "The license is expired.", nil),
				)
				return
			}
		}
	}
}

func (m *Middleware) OnlyAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, ok := c.Get(gb.KeyUser)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
				"message": "Unauthorized user",
			})
		}
	}
}
