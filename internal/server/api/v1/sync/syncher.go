package sync

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
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
		s.log.Check(gb.GetZapLogLevel(err), "Failed to list remote repositories.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	syncTime := time.Now().UTC()
	syncCnt := 0
	for _, re := range remotes {
		// Skip un-selected repositories.
		if !s.i.IsEntryOrg(ctx, re.Namespace) {
			continue
		}

		if err := s.i.SyncRemoteRepo(ctx, u, re, syncTime); err != nil {
			s.log.Error("It has failed to sync with the remote repository.", zap.Error(err), zap.Int64("repo_id", re.ID))
			continue
		}
		syncCnt = syncCnt + 1
	}
	s.log.Debug(fmt.Sprintf("Processed to schronize with %d repositories.", syncCnt), zap.Int64("user_id", u.ID))

	// Delete staled perms.
	var cnt int
	if cnt, err = s.i.DeletePermsOfUserLessThanSyncedAt(ctx, u, syncTime); err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to delete staled repositories.").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}
	s.log.Debug(fmt.Sprintf("Delete %d staled perms.", cnt))

	s.log.Debug("Success to synchronize.", zap.String("user", u.Login))
	c.Status(http.StatusOK)
}
