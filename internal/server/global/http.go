package global

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gitploy-io/gitploy/pkg/e"
)

func Response(c *gin.Context, httpCode int, data interface{}) {
	c.JSON(httpCode, data)
}

func ResponseWithError(c *gin.Context, err error) {
	if ge, ok := err.(*e.Error); ok {
		c.JSON(e.GetHttpCode(ge.Code), map[string]string{
			"code":    string(ge.Code),
			"message": ge.Message,
		})
		return
	}

	c.JSON(http.StatusInternalServerError, map[string]string{
		"code":    string(e.ErrorCodeInternalError),
		"message": err.Error(),
	})
}

// ResponseWithStatusAndError overrides the HTTP status.
func ResponseWithStatusAndError(c *gin.Context, status int, err error) {
	if ge, ok := err.(*e.Error); ok {
		c.JSON(status, map[string]string{
			"code":    string(ge.Code),
			"message": ge.Message,
		})
		return
	}

	c.JSON(http.StatusInternalServerError, map[string]string{
		"code":    string(e.ErrorCodeInternalError),
		"message": err.Error(),
	})
}

func AbortWithError(c *gin.Context, err error) {
	if ge, ok := err.(*e.Error); ok {
		c.AbortWithStatusJSON(e.GetHttpCode(ge.Code), map[string]string{
			"code":    string(ge.Code),
			"message": ge.Message,
		})
		return
	}

	c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
		"code":    string(e.ErrorCodeInternalError),
		"message": err.Error(),
	})
}

// AbortWithStatusAndError overrides the HTTP status.
func AbortWithStatusAndError(c *gin.Context, status int, err error) {
	if ge, ok := err.(*e.Error); ok {
		c.AbortWithStatusJSON(status, map[string]string{
			"code":    string(ge.Code),
			"message": ge.Message,
		})
		return
	}

	c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
		"code":    string(e.ErrorCodeInternalError),
		"message": err.Error(),
	})
}
