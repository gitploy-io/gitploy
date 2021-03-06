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

func (s *TagAPI) List(c *gin.Context) {
	ctx := c.Request.Context()

	var (
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

	uv, _ := c.Get(gb.KeyUser)
	u := uv.(*ent.User)

	rv, _ := c.Get(KeyRepo)
	repo := rv.(*ent.Repo)

	tags, err := s.i.ListTags(ctx, u, repo, &i.ListOptions{
		Page:    page,
		PerPage: perPage,
	})
	if err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to list tags.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	gb.Response(c, http.StatusOK, tags)
}
