package global

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gitploy-io/gitploy/pkg/e"
)

func Response(c *gin.Context, httpCode int, data interface{}) {
	c.JSON(httpCode, data)
}

func ErrorResponse(c *gin.Context, httpCode int, message string) {
	c.JSON(httpCode, map[string]string{
		"message": message,
	})
}

func ResponseWithError(c *gin.Context, err error) {
	if ge, ok := err.(*e.Error); ok {
		c.JSON(e.GetHttpCode(ge.Code), map[string]string{
			"code":    string(ge.Code),
			"message": e.GetMessage(ge.Code),
		})
		return
	}

	c.JSON(http.StatusInternalServerError, map[string]string{
		"code":    string(e.ErrorCodeInternalError),
		"message": err.Error(),
	})
}

func AbortWithErrorResponse(c *gin.Context, httpCode int, message string) {
	c.AbortWithStatusJSON(httpCode, map[string]string{
		"message": message,
	})
}
