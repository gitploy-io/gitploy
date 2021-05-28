package sync

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	gb "github.com/hanjunlee/gitploy/internal/server/global"
	"go.uber.org/zap"
)

type (
	// Syncher syncronizes local repositories with remote repositories.
	Syncher struct {
		scm   SCMHandler
		store StoreHandler
		log   *zap.Logger
	}
)

// NewSyncher create a new syncher.
func NewSyncher(scm SCMHandler, store StoreHandler) *Syncher {
	return &Syncher{
		scm:   scm,
		store: store,
		log:   zap.L().Named("sync"),
	}
}

// Sync synchronize local repositories with remote repositories,
// It create a new local repository if doesn't exist, but
// if exist it updates it.
func (s *Syncher) Sync(c *gin.Context) {
	ctx := c.Request.Context()

	u, _ := s.store.FindUserByHash(ctx, c.GetString(gb.KeySession))

	perms, err := s.scm.GetAllPermsWithRepo(c, u.Token)
	if err != nil {
		s.log.Error("failed to get all repositories.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to get all repositories.")
		return
	}

	sync := time.Now()

	for _, perm := range perms {
		re := perm.Edges.Repo
		if err := s.store.SyncPerm(c, u, perm, sync); err != nil {
			s.log.Error("failed to sync with the perm.", zap.String("repo", re.Name), zap.Error(err))
			gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to failed to sync with perms.")
			return
		}
	}

	gb.Response(c, http.StatusOK, nil)
}
