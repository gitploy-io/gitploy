package global

import "github.com/gin-gonic/gin"

func Response(c *gin.Context, httpCode int, data interface{}) {
	c.JSON(httpCode, data)
}

func ErrorResponse(c *gin.Context, httpCode int, message string) {
	c.JSON(httpCode, map[string]string{
		"message": message,
	})
}
