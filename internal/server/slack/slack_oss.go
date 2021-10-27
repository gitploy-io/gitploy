// +build oss

package slack

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewSlack(c *SlackConfig) *Slack {
	return &Slack{}
}

func (s *Slack) Cmd(c *gin.Context) {
	c.Status(http.StatusInternalServerError)
}

func (s *Slack) Interact(c *gin.Context) {
	c.Status(http.StatusInternalServerError)
}
