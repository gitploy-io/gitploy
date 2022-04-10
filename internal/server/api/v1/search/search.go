package search

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	i "github.com/gitploy-io/gitploy/internal/interactor"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/ent/deployment"
	"github.com/gitploy-io/gitploy/pkg/e"
)

const (
	activeDuration = 30 * time.Minute
)

type (
	Search struct {
		i   Interactor
		log *zap.Logger
	}
)

func NewSearch(i Interactor) *Search {
	return &Search{
		i:   i,
		log: zap.L().Named("search"),
	}
}

func (s *Search) SearchDeployments(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		statuses       = []deployment.Status{}
		owned          bool
		productionOnly bool
		from, to       time.Time
		page, perPage  int
		err            error
	)

	// Validate query parameters.
	for _, s := range strings.Split(c.DefaultQuery("statuses", ""), ",") {
		if s != "" {
			statuses = append(statuses, deployment.Status(s))
		}
	}

	if owned, err = strconv.ParseBool(c.DefaultQuery("owned", "true")); err != nil {
		gb.ResponseWithError(c, e.NewErrorWithMessage(e.ErrorCodeParameterInvalid, "The owned must be boolean.", err))
		return
	}

	if productionOnly, err = strconv.ParseBool(c.DefaultQuery("production_only", "false")); err != nil {
		gb.ResponseWithError(c, e.NewErrorWithMessage(e.ErrorCodeParameterInvalid, "The production must be boolean.", err))
		return
	}

	if from, err = time.Parse(time.RFC3339, c.DefaultQuery("from", time.Now().Add(-activeDuration).Format(time.RFC3339))); err != nil {
		gb.ResponseWithError(
			c,
			e.NewErrorWithMessage(e.ErrorCodeParameterInvalid, "Invalid format of \"from\" parameter, RFC3339 format only.", err),
		)
		return
	}

	if to, err = time.Parse(time.RFC3339, c.DefaultQuery("to", time.Now().Format(time.RFC3339))); err != nil {
		gb.ResponseWithError(
			c,
			e.NewErrorWithMessage(e.ErrorCodeParameterInvalid, "Invalid format of \"to\" parameter, RFC3339 format only.", err),
		)
		return
	}

	if page, err = strconv.Atoi(c.DefaultQuery("page", "1")); err != nil {
		gb.ResponseWithError(c, e.NewErrorWithMessage(e.ErrorCodeParameterInvalid, "The page must be number.", err))
		return
	}

	if perPage, err = strconv.Atoi(c.DefaultQuery("per_page", "1")); err != nil {
		gb.ResponseWithError(c, e.NewErrorWithMessage(e.ErrorCodeParameterInvalid, "The per_page must be number.", err))
		return
	}

	// Search deployments with parameters.
	var (
		ds []*ent.Deployment
	)

	v, _ := c.Get(gb.KeyUser)
	u := v.(*ent.User)

	if ds, err = s.i.SearchDeploymentsOfUser(ctx, u, &i.SearchDeploymentsOfUserOptions{
		ListOptions:    i.ListOptions{Page: page, PerPage: perPage},
		Statuses:       statuses,
		Owned:          owned,
		ProductionOnly: productionOnly,
		From:           from,
		To:             to,
	}); err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to search deployments.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	gb.Response(c, http.StatusOK, ds)
}

func (s *Search) SearchAssignedReviews(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		rvs []*ent.Review
		err error
	)

	v, _ := c.Get(gb.KeyUser)
	u := v.(*ent.User)

	if rvs, err = s.i.SearchReviews(ctx, u); err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to search reviews.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	gb.Response(c, http.StatusOK, rvs)
}
