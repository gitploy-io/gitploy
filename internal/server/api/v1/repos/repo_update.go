package repos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"

	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/pkg/e"
)

type (
	RepoPatchPayload struct {
		Name       *string `json:"name"`
		ConfigPath *string `json:"config_path"`
		Active     *bool   `json:"active"`
	}
)

func (s *RepoAPI) Update(c *gin.Context) {
	ctx := c.Request.Context()

	p := &RepoPatchPayload{}
	var err error
	if err := c.ShouldBindBodyWith(p, binding.JSON); err != nil {
		gb.ResponseWithError(
			c,
			e.NewErrorWithMessage(e.ErrorCodeParameterInvalid, "It has failed to bind the body.", err),
		)
		return
	}

	v, _ := c.Get(gb.KeyUser)
	u := v.(*ent.User)

	rv, _ := c.Get(KeyRepo)
	re := rv.(*ent.Repo)

	// Activate (or Deactivate) the repository:
	// Create a new webhook when it activates the repository,
	// in contrast it remove the webhook when it deactivates.
	if p.Active != nil {
		if *p.Active && !re.Active {
			if re, err = s.i.ActivateRepo(ctx, u, re); err != nil {
				s.log.Check(gb.GetZapLogLevel(err), "Failed to activate the repository.").Write(zap.Error(err))
				gb.ResponseWithError(c, err)
				return
			}
		} else if !*p.Active && re.Active {
			if re, err = s.i.DeactivateRepo(ctx, u, re); err != nil {
				s.log.Check(gb.GetZapLogLevel(err), "Failed to deactivate the repository.").Write(zap.Error(err))
				gb.ResponseWithError(c, err)
				return
			}
		}
	}

	if p.Name != nil {
		s.log.Debug("Set the name field.", zap.String("value", *p.Name))
		re.Name = *p.Name
	}

	if p.ConfigPath != nil {
		s.log.Debug("Set the config_path field.", zap.String("value", *p.ConfigPath))
		re.ConfigPath = *p.ConfigPath
	}

	if re, err = s.i.UpdateRepo(ctx, re); err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to update the repository.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	s.log.Info("Update the repository.", zap.Int64("repo_id", re.ID))
	gb.Response(c, http.StatusOK, re)
}
