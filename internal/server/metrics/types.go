package metrics

type (
	Metric struct {
		prometheusAuthSecret string
	}

	MetricConfig struct {
		Interactor
		PrometheusAuthSecret string
	}
)
