package metrics

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type (
	Metric struct{}
)

func NewMetric() *Metric {
	return &Metric{}
}

func (m *Metric) CollectMetrics(c *gin.Context) {
	h := promhttp.Handler()
	h.ServeHTTP(c.Writer, c.Request)
}
