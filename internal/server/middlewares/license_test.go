package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gitploy-io/gitploy/internal/server/middlewares/mock"
	"github.com/gitploy-io/gitploy/vo"
	"github.com/golang/mock/gomock"
)

func TestLicenseMiddleware_IsExpired(t *testing.T) {
	month := 30 * 24 * time.Hour

	t.Run("Return 402 error when the count of member is over the limit.", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := mock.NewMockInteractor(ctrl)

		m.
			EXPECT().
			GetLicense(gomock.Any()).
			Return(vo.NewTrialLicense(7), nil)

		gin.SetMode(gin.ReleaseMode)
		router := gin.New()

		lm := NewLicenseMiddleware(m)
		router.GET("/repos", lm.IsExpired(), func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		req, _ := http.NewRequest("GET", "/repos", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusPaymentRequired {
			t.Fatalf("IsExpired = %v, wanted %v", w.Code, http.StatusPaymentRequired)
		}
	})

	t.Run("Return 200 when the count of member is under the limit.", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := mock.NewMockInteractor(ctrl)

		m.
			EXPECT().
			GetLicense(gomock.Any()).
			Return(vo.NewTrialLicense(vo.TrialMemberLimit), nil)

		gin.SetMode(gin.ReleaseMode)
		router := gin.New()

		lm := NewLicenseMiddleware(m)
		router.GET("/repos", lm.IsExpired(), func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		req, _ := http.NewRequest("GET", "/repos", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("IsExpired = %v, wanted %v", w.Code, http.StatusOK)
		}
	})

	t.Run("Return 403 when the license is expired.", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := mock.NewMockInteractor(ctrl)

		m.
			EXPECT().
			GetLicense(gomock.Any()).
			Return(vo.NewStandardLicense(10, &vo.SigningData{
				MemberLimit: 20,
				ExpiredAt:   time.Now().Add(-2 * month),
			}), nil)

		gin.SetMode(gin.ReleaseMode)
		router := gin.New()

		lm := NewLicenseMiddleware(m)
		router.GET("/repos", lm.IsExpired(), func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		req, _ := http.NewRequest("GET", "/repos", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusPaymentRequired {
			t.Fatalf("IsExpired = %v, wanted %v", w.Code, http.StatusPaymentRequired)
		}
	})

	t.Run("Return 200 when the license is not expired.", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := mock.NewMockInteractor(ctrl)

		t.Logf("Set expired at yesterday")
		m.
			EXPECT().
			GetLicense(gomock.Any()).
			Return(vo.NewStandardLicense(10, &vo.SigningData{
				MemberLimit: 20,
				ExpiredAt:   time.Now().Add(-24 * time.Hour),
			}), nil)

		gin.SetMode(gin.ReleaseMode)
		router := gin.New()

		lm := NewLicenseMiddleware(m)
		router.GET("/repos", lm.IsExpired(), func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		req, _ := http.NewRequest("GET", "/repos", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("IsExpired = %v, wanted %v", w.Code, http.StatusOK)
		}
	})
}
