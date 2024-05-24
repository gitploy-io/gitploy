package shared

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gitploy-io/gitploy/internal/server/api/shared/mock"
	"github.com/gitploy-io/gitploy/model/extent"
	"github.com/golang/mock/gomock"
)

func TestMiddleware_IsLicenseExpired(t *testing.T) {
	month := 30 * 24 * time.Hour

	t.Run("Return 200 when the license is OSS.", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := mock.NewMockInteractor(ctrl)

		m.
			EXPECT().
			GetLicense(gomock.Any()).
			Return(extent.NewOSSLicense(), nil)

		gin.SetMode(gin.ReleaseMode)
		router := gin.New()

		lm := NewMiddleware(m)
		router.GET("/repos", lm.IsLicenseExpired(), func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		req, _ := http.NewRequest("GET", "/repos", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("IsLicenseExpired = %v, wanted %v", w.Code, http.StatusOK)
		}
	})

	t.Run("Return 200 when the count of member is under the limit.", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := mock.NewMockInteractor(ctrl)

		m.
			EXPECT().
			GetLicense(gomock.Any()).
			Return(extent.NewTrialLicense(extent.TrialMemberLimit, extent.TrialDeploymentLimit), nil)

		gin.SetMode(gin.ReleaseMode)
		router := gin.New()

		lm := NewMiddleware(m)
		router.GET("/repos", lm.IsLicenseExpired(), func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		req, _ := http.NewRequest("GET", "/repos", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("IsLicenseExpired = %v, wanted %v", w.Code, http.StatusOK)
		}
	})

	t.Run("Return 403 when the license is expired.", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := mock.NewMockInteractor(ctrl)

		m.
			EXPECT().
			GetLicense(gomock.Any()).
			Return(extent.NewStandardLicense(10, &extent.SigningData{
				MemberLimit: 20,
				ExpiredAt:   time.Now().Add(-2 * month),
			}), nil)

		gin.SetMode(gin.ReleaseMode)
		router := gin.New()

		lm := NewMiddleware(m)
		router.GET("/repos", lm.IsLicenseExpired(), func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		req, _ := http.NewRequest("GET", "/repos", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusPaymentRequired {
			t.Fatalf("IsLicenseExpired = %v, wanted %v", w.Code, http.StatusPaymentRequired)
		}
	})

	t.Run("Return 200 when the license is not expired.", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := mock.NewMockInteractor(ctrl)

		t.Logf("Set expired at yesterday")
		m.
			EXPECT().
			GetLicense(gomock.Any()).
			Return(extent.NewStandardLicense(10, &extent.SigningData{
				MemberLimit: 20,
				ExpiredAt:   time.Now().Add(-24 * time.Hour),
			}), nil)

		gin.SetMode(gin.ReleaseMode)
		router := gin.New()

		lm := NewMiddleware(m)
		router.GET("/repos", lm.IsLicenseExpired(), func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		req, _ := http.NewRequest("GET", "/repos", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("IsLicenseExpired = %v, wanted %v", w.Code, http.StatusOK)
		}
	})
}
