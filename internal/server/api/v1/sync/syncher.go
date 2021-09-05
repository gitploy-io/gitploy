package sync

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gitploy-io/gitploy/ent"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"go.uber.org/zap"
)

type (
	// Syncher syncronizes local repositories with remote repositories.
	Syncher struct {
		i   Interactor
		log *zap.Logger
	}
)

// NewSyncher create a new syncher.
func NewSyncher(i Interactor) *Syncher {
	return &Syncher{
		i:   i,
		log: zap.L().Named("sync"),
	}
}

// Sync synchronize local repositories with remote repositories,
// It create a new local repository if doesn't exist, but
// if exist it updates it.
func (s *Syncher) Sync(c *gin.Context) {
	ctx := c.Request.Context()

	v, _ := c.Get(gb.KeyUser)
	u := v.(*ent.User)

	remotes, err := s.i.ListRemoteRepos(ctx, u)
	if err != nil {
		s.log.Error("It has failed to list remote repositories.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to list remote repositories.")
		return
	}

	syncTime := time.Now()
	syncCnt := 0
	for _, re := range remotes {
		// Skip un-selected repositories.
		if !s.i.IsEntryOrg(ctx, re.Namespace) {
			continue
		}

		if err := s.i.SyncRemoteRepo(ctx, u, re); err != nil {
			s.log.Error("It has failed to sync with the remote repository.", zap.Error(err), zap.String("repo_id", re.ID))
			continue
		}
		syncCnt = syncCnt + 1
	}
	s.log.Debug(fmt.Sprintf("Processed to schronize with %d repositories.", syncCnt), zap.String("user_id", u.ID))

	// Delete staled perms.
	var cnt int
	if cnt, err = s.i.DeletePermsOfUserLessThanUpdatedAt(ctx, u, syncTime); err != nil {
		s.log.Error("It has failed to delete staled repositories.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to delete staled repositories.")
		return
	}
	s.log.Debug(fmt.Sprintf("Delete %d staled perms.", cnt))

	s.log.Debug("Success to synchronize.", zap.String("user", u.Login))
	c.Status(http.StatusOK)
}
