package repos

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	i "github.com/gitploy-io/gitploy/internal/interactor"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/pkg/e"
)

func (s *DeploymentAPI) ListChanges(c *gin.Context) {
	var (
		number  int
		page    int
		perPage int
		err     error
	)

	if number, err = strconv.Atoi(c.Param("number")); err != nil {
		s.log.Warn("Invalid parameter: number must be integer.", zap.Error(err))
		gb.ResponseWithError(c, e.NewError(e.ErrorCodeParameterInvalid, err))
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

	vu, _ := c.Get(gb.KeyUser)
	u := vu.(*ent.User)

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	ctx := context.WithValue(c.Request.Context(), i.KeyUser, u)

	d, err := s.i.FindDeploymentOfRepoByNumber(ctx, re, number)
	if err != nil {
		s.log.Warn("Failed to find the deployment.", zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	commits, err := s.i.CompareCommitsFromLastestDeployment(ctx, re, d, &i.ListOptions{
		Page:    page,
		PerPage: perPage,
	})
	if err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to compare two commits.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	s.log.Info("Get the changed commits for the deployment.", zap.Int("number", number))
	gb.Response(c, http.StatusOK, commits)
}
