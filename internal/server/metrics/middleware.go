package metrics

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespace = "gitploy"
)

var (
	RequestCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: "",
		Name:      "requests_total",
		Help:      "How many HTTP requests processed, partitioned by status code and HTTP method.",
	}, []string{"code", "method", "path"})

	RequestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: namespace,
		Subsystem: "",
		Name:      "request_duration_seconds",
		Help:      "The HTTP request latencies in seconds.",
	}, []string{"code", "method", "path"})
)

func init() {
	prometheus.MustRegister(RequestCount)
	prometheus.MustRegister(RequestDuration)
}

// ReponseMetrics is the middleware to collect metrics about the response.
func ReponseMetrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		status := strconv.Itoa(c.Writer.Status())

		{
			RequestCount.WithLabelValues(status, c.Request.Method, c.Request.URL.Path).Inc()
		}
		{
			elapsed := float64(time.Since(start)) / float64(time.Second)
			RequestDuration.WithLabelValues(status, c.Request.Method, c.Request.URL.Path).Observe(elapsed)
		}
	}
}
