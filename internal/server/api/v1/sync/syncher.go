package sync

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gb "github.com/hanjunlee/gitploy/internal/server/global"
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

	u, _ := s.i.FindUserByHash(ctx, c.GetString(gb.KeySession))

	if err := s.i.Sync(ctx, u); err != nil {
		s.log.Error("failed to synchronize.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to synchronize.")
	}

	s.log.Debug("success to synchronize.", zap.String("user", u.Login))
	c.Status(http.StatusOK)
}
