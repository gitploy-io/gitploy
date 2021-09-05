package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
)

const (
	extraDuration time.Duration = 30 * 24 * time.Hour
)

type (
	LicenseMiddleware struct {
		i Interactor
	}
)

func NewLicenseMiddleware(intr Interactor) *LicenseMiddleware {
	return &LicenseMiddleware{
		i: intr,
	}
}

func (lm *LicenseMiddleware) IsExpired() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		lic, err := lm.i.GetLicense(ctx)
		if err != nil {
			gb.AbortWithErrorResponse(c, http.StatusInternalServerError, "It has failed to get the license.")
			return
		}

		if lic.IsOverLimit() {
			gb.AbortWithErrorResponse(c, http.StatusPaymentRequired, "The member count is over the limit.")
			return
		}

		if !lic.IsTrial() && lic.IsExpired() {
			now := time.Now()
			if lic.ExpiredAt.Add(extraDuration).Before(now) {
				gb.AbortWithErrorResponse(c, http.StatusPaymentRequired, "The license is expired.")
				return
			}
		}
	}
}
