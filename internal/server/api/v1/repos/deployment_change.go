package repos

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/ent/deployment"
	"github.com/gitploy-io/gitploy/model/extent"
	"github.com/gitploy-io/gitploy/pkg/e"
)

func (s *DeploymentsAPI) ListChanges(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		number  = c.Param("number")
		page    = c.DefaultQuery("page", "1")
		perPage = c.DefaultQuery("per_page", "30")
	)

	vu, _ := c.Get(gb.KeyUser)
	u := vu.(*ent.User)

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	d, err := s.i.FindDeploymentOfRepoByNumber(ctx, re, atoi(number))
	if err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to find the deployments.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	ld, err := s.i.FindPrevSuccessDeployment(ctx, d)
	if e.HasErrorCode(err, e.ErrorCodeEntityNotFound) {
		s.log.Debug("The previous deployment is not found.")
		gb.Response(c, http.StatusOK, []*extent.Commit{})
		return
	} else if err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to find the deployments.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	// Get SHA when the status of deployment is waiting.
	sha := d.Sha
	if sha == "" {
		sha, err = s.getCommitSha(ctx, u, re, d.Type, d.Ref)
		if err != nil {
			s.log.Check(gb.GetZapLogLevel(err), "It has failed to get the commit SHA.").Write(zap.Error(err))
			gb.ResponseWithError(c, err)
			return
		}
	}

	commits, _, err := s.i.CompareCommits(ctx, u, re, ld.Sha, sha, atoi(page), atoi(perPage))
	if err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to compare two commits.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	gb.Response(c, http.StatusOK, commits)
}

func (s *DeploymentsAPI) getCommitSha(ctx context.Context, u *ent.User, re *ent.Repo, typ deployment.Type, ref string) (string, error) {
	switch typ {
	case deployment.TypeCommit:
		c, err := s.i.GetCommit(ctx, u, re, ref)
		if err != nil {
			return "", err
		}

		return c.SHA, nil
	case deployment.TypeBranch:
		b, err := s.i.GetBranch(ctx, u, re, ref)
		if err != nil {
			return "", err
		}

		return b.CommitSHA, nil
	case deployment.TypeTag:
		t, err := s.i.GetTag(ctx, u, re, ref)
		if err != nil {
			return "", err
		}

		return t.CommitSHA, nil
	default:
		return "", fmt.Errorf("Type must be one of commit, branch, tag.")
	}
}
