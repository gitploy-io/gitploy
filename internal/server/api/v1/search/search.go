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
		statuses = c.DefaultQuery("statuses", "")
		owned    = c.DefaultQuery("owned", "true")
		from     = c.DefaultQuery("from", time.Now().Add(-activeDuration).Format(time.RFC3339))
		to       = c.DefaultQuery("to", time.Now().Format(time.RFC3339))
		page     = c.DefaultQuery("page", "1")
		perPage  = c.DefaultQuery("per_page", "30")
	)

	var (
		ss  = make([]deployment.Status, 0)
		o   bool
		f   time.Time
		t   time.Time
		p   int
		pp  int
		err error
	)

	// Validate query parameters.
	for _, st := range strings.Split(statuses, ",") {
		if st != "" {
			ss = append(ss, deployment.Status(st))
		}
	}

	if o, err = strconv.ParseBool(owned); err != nil {
		gb.ResponseWithError(
			c,
			e.NewErrorWithMessage(e.ErrorCodeParameterInvalid, "The owned must be boolean.", err),
		)
		return
	}

	if f, err = time.Parse(time.RFC3339, from); err != nil {
		gb.ResponseWithError(
			c,
			e.NewErrorWithMessage(e.ErrorCodeParameterInvalid, "Invalid format of \"from\" parameter, RFC3339 format only.", err),
		)
		return
	}

	if t, err = time.Parse(time.RFC3339, to); err != nil {
		gb.ResponseWithError(
			c,
			e.NewErrorWithMessage(e.ErrorCodeParameterInvalid, "Invalid format of \"to\" parameter, RFC3339 format only.", err),
		)
		return
	}

	if p, err = strconv.Atoi(page); err != nil {
		gb.ResponseWithError(
			c,
			e.NewErrorWithMessage(e.ErrorCodeParameterInvalid, "Invalid format of \"page\" parameter.", err),
		)
		return
	}

	if pp, err = strconv.Atoi(perPage); err != nil {
		gb.ResponseWithError(
			c,
			e.NewErrorWithMessage(e.ErrorCodeParameterInvalid, "Invalid format of \"per_page\" parameter.", err),
		)
		return
	}

	// Search deployments with parameters.
	var (
		ds []*ent.Deployment
	)

	v, _ := c.Get(gb.KeyUser)
	u := v.(*ent.User)

	if ds, err = s.i.SearchDeploymentsOfUser(ctx, u, &i.SearchDeploymentsOfUserOptions{
		ListOptions: i.ListOptions{Page: p, PerPage: pp},
		Statuses:    ss,
		Owned:       o,
		From:        f,
		To:          t,
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
