package repos

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/internal/server/global"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/vo"
)

type (
	lockPostPayload struct {
		Env string `json:"env"`
	}
)

func (r *Repo) ListLocks(c *gin.Context) {
	ctx := c.Request.Context()

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	locks, err := r.i.ListLocksOfRepo(ctx, re)
	if err != nil {
		r.log.Error("It has failed to list locks.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to list locks.")
		return
	}

	gb.Response(c, http.StatusOK, locks)
}

func (r *Repo) CreateLock(c *gin.Context) {
	ctx := c.Request.Context()

	p := &lockPostPayload{}
	if err := c.ShouldBindBodyWith(p, binding.JSON); err != nil {
		r.log.Error("It has failed to bind the payload.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusBadRequest, "It has failed to bind the payload.")
		return
	}

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	vu, _ := c.Get(global.KeyUser)
	u := vu.(*ent.User)

	// Validate the payload, it check whether the env exist or not in deploy.yml.
	cfg, err := r.i.GetConfig(ctx, u, re)
	if vo.IsConfigNotFoundError(err) || vo.IsConfigParseError(err) {
		r.log.Warn("The config is invalid.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusUnprocessableEntity, "The config is invalid.")
		return
	} else if err != nil {
		r.log.Error("It has failed to get the config file.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to get the config file.")
		return
	}

	if !cfg.HasEnv(p.Env) {
		r.log.Warn("The env is not found.", zap.String("env", p.Env))
		gb.ErrorResponse(c, http.StatusUnprocessableEntity, fmt.Sprintf("The '%s' env is not found.", p.Env))
		return
	}

	if ok, err := r.i.HasLockOfRepoForEnv(ctx, re, p.Env); ok {
		r.log.Warn("The lock already exist.", zap.String("env", p.Env))
		gb.ErrorResponse(c, http.StatusUnprocessableEntity, "The lock already exist.")
		return
	} else if err != nil {
		r.log.Error("It has failed to check the lock.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to check the lock.")
		return
	}

	// Lock the environment.
	l, err := r.i.CreateLock(ctx, &ent.Lock{
		Env:    p.Env,
		UserID: u.ID,
		RepoID: re.ID,
	})
	if err != nil {
		r.log.Error("It has failed to lock the env.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to lock the env.")
		return
	}

	r.log.Debug("Lock the env.", zap.String("env", p.Env))
	gb.Response(c, http.StatusCreated, l)
}

func (r *Repo) DeleteLock(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		sid = c.Param("lockID")
	)

	id, err := strconv.Atoi(sid)
	if err != nil {
		r.log.Error("The lock ID must to be number.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusBadRequest, "The lock ID must to be number.")
		return
	}

	l, err := r.i.FindLockByID(ctx, id)
	if ent.IsNotFound(err) {
		r.log.Warn("The lock is not found.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusUnprocessableEntity, "The lock is not found.")
		return
	} else if err != nil {
		r.log.Error("It has failed to find the lock.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to find the lock.")
		return
	}

	if err := r.i.DeleteLock(ctx, l); err != nil {
		r.log.Error("It has failed to delete the lock.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to delete the lock.")
		return
	}

	gb.Response(c, http.StatusOK, nil)
}
