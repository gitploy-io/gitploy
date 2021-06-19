package slack

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/slack-go/slack"
	"go.uber.org/zap"
)

type (
	SlackMiddleware struct {
		secret string
		log    *zap.Logger
	}
)

func NewSlackMiddleware(secret string) *SlackMiddleware {
	return &SlackMiddleware{
		secret: secret,
		log:    zap.L().Named("slack-middleware"),
	}
}

func (m *SlackMiddleware) Verify() gin.HandlerFunc {
	return func(c *gin.Context) {
		v, err := slack.NewSecretsVerifier(c.Request.Header, m.secret)
		if err != nil {
			m.log.Error("failed to generate the verifier.")
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		d := copyRawData(c.Request)
		v.Write(d)

		if err := v.Ensure(); err != nil {
			m.log.Error("invalid request.", zap.Error(err))
			c.AbortWithStatus(http.StatusBadRequest)
		}
	}
}

func copyRawData(req *http.Request) []byte {
	bodyBytes, _ := ioutil.ReadAll(req.Body)
	req.Body.Close()

	req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	return bodyBytes
}
