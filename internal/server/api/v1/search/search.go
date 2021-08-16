package search

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/approval"
	"github.com/hanjunlee/gitploy/ent/deployment"
	gb "github.com/hanjunlee/gitploy/internal/server/global"
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
		gb.ErrorResponse(c, http.StatusBadRequest, "Invalid format of \"owned\" parameter.")
		return
	}

	if f, err = time.Parse(time.RFC3339, from); err != nil {
		gb.ErrorResponse(c, http.StatusBadRequest, "Invalid format of \"from\" parameter, RFC3339 format only.")
		return
	}

	if t, err = time.Parse(time.RFC3339, to); err != nil {
		gb.ErrorResponse(c, http.StatusBadRequest, "Invalid format of \"to\" parameter, RFC3339 format only.")
		return
	}

	if p, err = strconv.Atoi(page); err != nil {
		gb.ErrorResponse(c, http.StatusBadRequest, "Invalid format of \"page\" parameter.")
		return
	}

	if pp, err = strconv.Atoi(perPage); err != nil {
		gb.ErrorResponse(c, http.StatusBadRequest, "Invalid format of \"per_page\" parameter.")
		return
	}

	// Search deployments with parameters.
	var (
		ds []*ent.Deployment
	)

	v, _ := c.Get(gb.KeyUser)
	u := v.(*ent.User)

	if ds, err = s.i.SearchDeployments(ctx, u, ss, o, f, t, p, pp); err != nil {
		s.log.Error("It has failed to search deployments.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to search deployments.")
		return
	}

	gb.Response(c, http.StatusOK, ds)
}

func (s *Search) SearchApprovals(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		statuses = c.DefaultQuery("statuses", "")
		from     = c.DefaultQuery("from", time.Now().Add(-activeDuration).Format(time.RFC3339))
		to       = c.DefaultQuery("to", time.Now().Format(time.RFC3339))
		page     = c.DefaultQuery("page", "1")
		perPage  = c.DefaultQuery("per_page", "30")
	)

	var (
		ss  = make([]approval.Status, 0)
		f   time.Time
		t   time.Time
		p   int
		pp  int
		err error
	)

	// Validate query parameters.
	for _, st := range strings.Split(statuses, ",") {
		if st != "" {
			ss = append(ss, approval.Status(st))
		}
	}

	if f, err = time.Parse(time.RFC3339, from); err != nil {
		gb.ErrorResponse(c, http.StatusBadRequest, "Invalid format of \"from\" parameter, RFC3339 format only.")
		return
	}

	if t, err = time.Parse(time.RFC3339, to); err != nil {
		gb.ErrorResponse(c, http.StatusBadRequest, "Invalid format of \"to\" parameter, RFC3339 format only.")
		return
	}

	if p, err = strconv.Atoi(page); err != nil {
		gb.ErrorResponse(c, http.StatusBadRequest, "Invalid format of \"page\" parameter.")
		return
	}

	if pp, err = strconv.Atoi(perPage); err != nil {
		gb.ErrorResponse(c, http.StatusBadRequest, "Invalid format of \"per_page\" parameter.")
		return
	}

	// Search deployments with parameters.
	var (
		ds []*ent.Approval
	)

	v, _ := c.Get(gb.KeyUser)
	u := v.(*ent.User)

	if ds, err = s.i.SearchApprovals(ctx, u, ss, f, t, p, pp); err != nil {
		s.log.Error("It has failed to search deployments.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to search deployments.")
		return
	}

	gb.Response(c, http.StatusOK, ds)
}
