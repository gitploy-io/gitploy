package repos

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	i "github.com/gitploy-io/gitploy/internal/interactor"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/pkg/e"
)

func (s *PermAPI) List(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		q       = c.DefaultQuery("q", "")
		page    int
		perPage int
		err     error
	)

	// Validate quries
	if page, err = strconv.Atoi(c.DefaultQuery("page", defaultQueryPage)); err != nil {
		s.log.Warn("Invalid parameter: page is not integer.", zap.Error(err))
		gb.ResponseWithError(c, e.NewError(e.ErrorCodeParameterInvalid, err))
		return
	}

	if perPage, err = strconv.Atoi(c.DefaultQuery("per_page", defaultQueryPerPage)); err != nil {
		s.log.Warn("Invalid parameter: per_page is not integer.", zap.Error(err))
		gb.ResponseWithError(c, e.NewError(e.ErrorCodeParameterInvalid, err))
		return
	}

	v, _ := c.Get(KeyRepo)
	re := v.(*ent.Repo)

	if perPage > 100 {
		perPage = 100
	}

	perms, err := s.i.ListPermsOfRepo(ctx, re, &i.ListPermsOfRepoOptions{
		ListOptions: i.ListOptions{Page: page, PerPage: perPage},
		Query:       q,
	})
	if err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to list permissions.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	gb.Response(c, http.StatusOK, perms)
}
