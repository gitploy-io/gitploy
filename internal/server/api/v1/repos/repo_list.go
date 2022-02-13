package repos

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/pkg/e"
)

func (s *RepoAPI) List(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		sort      = c.DefaultQuery("sort", "false")
		q         = c.Query("q")
		namespace = c.Query("namespace")
		name      = c.Query("name")
		page      int
		perPage   int
		err       error
	)

	// Validate queries.
	sorted, err := strconv.ParseBool(sort)
	if err != nil {
		gb.ResponseWithError(
			c,
			e.NewErrorWithMessage(e.ErrorCodeParameterInvalid, "The \"sort\" field must be boolean.", err),
		)
		return
	}

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

	v, _ := c.Get(gb.KeyUser)
	u := v.(*ent.User)

	repos, err := s.i.ListReposOfUser(ctx, u, q, namespace, name, sorted, page, perPage)
	if err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to list repositories.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	gb.Response(c, http.StatusOK, repos)
}
