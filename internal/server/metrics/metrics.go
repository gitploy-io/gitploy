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

		start := time.Now().UTC()
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
						"total_deployment_count",
					),
					"The total deployment count of the production deployments.",
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
						"total_rollback_count",
					),
					"The total rollback count of the production deployments.",
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
						"total_line_additions",
					),
					"The total added lines of the production deployments.",
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
						"total_line_deletions",
					),
					"The total deleted lines of the production deployments.",
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
						"total_line_changes",
					),
					"The total changed lines of the production deployments.",
					[]string{"namespace", "name", "env"},
					nil,
				),
				prometheus.GaugeValue,
				float64(dc.Changes),
				dc.Edges.Repo.Namespace, dc.Edges.Repo.Name, dc.Env,
			)

			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(
						namespace,
						"",
						"total_lead_time_seconds",
					),
					"The total amount of time it takes a commit to get into the production environments.",
					[]string{"namespace", "name", "env"},
					nil,
				),
				prometheus.GaugeValue,
				float64(dc.LeadTimeSeconds),
				dc.Edges.Repo.Namespace, dc.Edges.Repo.Name, dc.Env,
			)

			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(
						namespace,
						"",
						"total_commit_count",
					),
					"The total commit count of production deployments.",
					[]string{"namespace", "name", "env"},
					nil,
				),
				prometheus.GaugeValue,
				float64(dc.CommitCount),
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

	{
		acnt, err := c.i.CountActiveRepos(ctx)
		if err != nil {
			c.log.Error("It has failed to count active repositories.", zap.Error(err))
			return
		}

		cnt, err := c.i.CountRepos(ctx)
		if err != nil {
			c.log.Error("It has failed to count total repositories.", zap.Error(err))
			return
		}

		ch <- prometheus.MustNewConstMetric(
			prometheus.NewDesc(
				prometheus.BuildFQName(
					namespace,
					"",
					"total_active_repo_count",
				),
				"The count of active repositories.",
				nil,
				nil,
			),
			prometheus.GaugeValue,
			float64(acnt),
		)

		ch <- prometheus.MustNewConstMetric(
			prometheus.NewDesc(
				prometheus.BuildFQName(
					namespace,
					"",
					"total_repo_count",
				),
				"The count of repositories.",
				nil,
				nil,
			),
			prometheus.GaugeValue,
			float64(cnt),
		)
	}
}
