// +build oss

package stream

import (
	"time"

	"github.com/gin-gonic/gin"
)

func (s *Stream) GetEvents(c *gin.Context) {
	w := c.Writer

L:
	for {
		select {
		case <-w.CloseNotify():
			break L
		case <-time.After(time.Minute):
			break L
		}
	}
}
