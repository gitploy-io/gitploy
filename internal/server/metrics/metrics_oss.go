// +build oss

package metrics

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewMetric(c *MetricConfig) *Metric {
	return &Metric{}
}

func (m *Metric) CollectMetrics(c *gin.Context) {
	c.Status(http.StatusOK)
}
