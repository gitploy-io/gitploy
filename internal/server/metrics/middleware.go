package metrics

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	subsystem = "gitploy"
)

var (
	ResponseCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Subsystem: subsystem,
		Name:      "requests_total",
		Help:      "How many HTTP requests processed, partitioned by status code and HTTP method.",
	}, []string{"code", "method", "path"})

	ResponseDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Subsystem: subsystem,
		Name:      "request_duration_seconds",
		Help:      "The HTTP request latencies in seconds.",
	}, []string{"code", "method", "path"})
)

func init() {
	prometheus.MustRegister(ResponseCount)
	prometheus.MustRegister(ResponseDuration)
}

// ReponseMetrics is the middleware to collect metrics about the response.
func ReponseMetrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		status := strconv.Itoa(c.Writer.Status())

		{
			ResponseCount.WithLabelValues(status, c.Request.Method, c.Request.URL.Path).Inc()
		}
		{
			elapsed := float64(time.Since(start)) / float64(time.Second)
			ResponseDuration.WithLabelValues(status, c.Request.Method, c.Request.URL.Path).Observe(elapsed)
		}
	}
}
