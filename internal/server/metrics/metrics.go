package metrics

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"

	"github.com/gitploy-io/gitploy/ent"
)

type (
	Metric struct{}

	MetricConfig struct {
		Interactor
	}

	collector struct {
		i Interactor

		cache    map[int]*ent.DeploymentCount
		lastTime time.Time

		log *zap.Logger
	}
)

func NewMetric(c *MetricConfig) *Metric {
	prometheus.MustRegister(
		newCollector(c.Interactor),
	)

	return &Metric{}
}

func (m *Metric) CollectMetrics(c *gin.Context) {
	h := promhttp.Handler()
	h.ServeHTTP(c.Writer, c.Request)
}

func newCollector(i Interactor) *collector {
	return &collector{
		i:     i,
		cache: map[int]*ent.DeploymentCount{},
		log:   zap.L().Named("collector"),
	}
}

func (c *collector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *collector) Collect(ch chan<- prometheus.Metric) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	{
		var (
			dcs []*ent.DeploymentCount
			err error
		)

		start := time.Now()
		defer func() {
			c.lastTime = start
		}()

		if len(c.cache) == 0 {
			c.log.Debug("List all deployment_count.")
			if dcs, err = c.i.ListAllDeploymentCounts(ctx); err != nil {
				c.log.Error("It has failed to list all deployment_counts.", zap.Error(err))
				return
			}
		} else {
			c.log.Debug("List deployment_count from the last time.", zap.Time("last", c.lastTime))
			if dcs, err = c.i.ListDeploymentCountsGreaterThanTime(ctx, c.lastTime); err != nil {
				c.log.Error("It has failed to list deployment_counts.", zap.Error(err))
				return
			}
		}

		for _, dc := range dcs {
			c.cache[dc.ID] = dc
		}

		for _, dc := range c.cache {
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(
						namespace,
						"",
						"deployment_count",
					),
					"The count of success deployment for each environment, respectively.",
					[]string{"namespace", "name", "env"},
					nil,
				),
				prometheus.CounterValue,
				float64(dc.Count),
				dc.Namespace, dc.Name, dc.Env,
			)
		}

		c.log.Debug("Collect deployment_count metrics successfully.")
	}

	{
		lic, err := c.i.GetLicense(ctx)
		if err != nil {
			c.log.Error("It has failed to get the license.", zap.Error(err))
			return
		}

		ch <- prometheus.MustNewConstMetric(
			prometheus.NewDesc(
				prometheus.BuildFQName(
					namespace,
					"",
					"member_count",
				),
				"The total count of members.",
				nil,
				nil,
			),
			prometheus.CounterValue,
			float64(lic.MemberCount),
		)

		ch <- prometheus.MustNewConstMetric(
			prometheus.NewDesc(
				prometheus.BuildFQName(
					namespace,
					"",
					"member_limit",
				),
				"The limit count of members.",
				[]string{"kind"},
				nil,
			),
			prometheus.CounterValue,
			float64(lic.MemberLimit),
			string(lic.Kind),
		)
	}
}
