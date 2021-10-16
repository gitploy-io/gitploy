package metrics

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"

	"github.com/gitploy-io/gitploy/ent"
)

const (
	headerAuth = "Authorization"
)

type (
	Metric struct {
		prometheusAuthSecret string
	}

	MetricConfig struct {
		Interactor
		PrometheusAuthSecret string
	}

	collector struct {
		i Interactor

		cache    map[int]*ent.DeploymentStatistics
		lastTime time.Time

		log *zap.Logger
	}
)

func NewMetric(c *MetricConfig) *Metric {
	prometheus.MustRegister(
		newCollector(c.Interactor),
	)

	return &Metric{
		prometheusAuthSecret: c.PrometheusAuthSecret,
	}
}

func (m *Metric) CollectMetrics(c *gin.Context) {
	if m.prometheusAuthSecret != "" {
		if value := strings.TrimPrefix(
			c.GetHeader(headerAuth),
			"Bearer ",
		); m.prometheusAuthSecret != value {
			c.Status(http.StatusUnauthorized)
			return
		}
	}

	h := promhttp.Handler()
	h.ServeHTTP(c.Writer, c.Request)
}

func newCollector(i Interactor) *collector {
	return &collector{
		i:     i,
		cache: map[int]*ent.DeploymentStatistics{},
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
			dcs []*ent.DeploymentStatistics
			err error
		)

		start := time.Now()
		defer func() {
			c.lastTime = start
		}()

		if len(c.cache) == 0 {
			c.log.Debug("List all deployment_count.")
			if dcs, err = c.i.ListAllDeploymentStatistics(ctx); err != nil {
				c.log.Error("It has failed to list all deployment_counts.", zap.Error(err))
				return
			}
		} else {
			c.log.Debug("List deployment_count from the last time.", zap.Time("last", c.lastTime))
			if dcs, err = c.i.ListDeploymentStatisticsGreaterThanTime(ctx, c.lastTime); err != nil {
				c.log.Error("It has failed to list deployment_counts.", zap.Error(err))
				return
			}
		}

		for _, dc := range dcs {
			c.cache[dc.ID] = dc
		}

		for _, dc := range c.cache {
			if dc.Edges.Repo == nil {
				continue
			}

			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(
						namespace,
						"",
						"deployment_count",
					),
					"The count of success deployment of the production environment.",
					[]string{"namespace", "name", "env"},
					nil,
				),
				prometheus.GaugeValue,
				float64(dc.Count),
				dc.Edges.Repo.Namespace, dc.Edges.Repo.Name, dc.Env,
			)

			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(
						namespace,
						"",
						"rollback_count",
					),
					"The count of rollback of the production environment.",
					[]string{"namespace", "name", "env"},
					nil,
				),
				prometheus.GaugeValue,
				float64(dc.RollbackCount),
				dc.Edges.Repo.Namespace, dc.Edges.Repo.Name, dc.Env,
			)

			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(
						namespace,
						"",
						"line_additions",
					),
					"The count of added lines from the latest deployment of the production environment.",
					[]string{"namespace", "name", "env"},
					nil,
				),
				prometheus.GaugeValue,
				float64(dc.Additions),
				dc.Edges.Repo.Namespace, dc.Edges.Repo.Name, dc.Env,
			)

			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(
						namespace,
						"",
						"line_deletions",
					),
					"The count of deleted lines from the latest deployment of the production environment.",
					[]string{"namespace", "name", "env"},
					nil,
				),
				prometheus.GaugeValue,
				float64(dc.Deletions),
				dc.Edges.Repo.Namespace, dc.Edges.Repo.Name, dc.Env,
			)

			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(
						namespace,
						"",
						"line_changes",
					),
					"The count of changed lines from the latest deployment of the production environment.",
					[]string{"namespace", "name", "env"},
					nil,
				),
				prometheus.GaugeValue,
				float64(dc.Changes),
				dc.Edges.Repo.Namespace, dc.Edges.Repo.Name, dc.Env,
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
			prometheus.GaugeValue,
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
			prometheus.GaugeValue,
			float64(lic.MemberLimit),
			string(lic.Kind),
		)
	}
}
