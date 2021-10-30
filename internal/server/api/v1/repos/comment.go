package repos

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/comment"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/pkg/e"
	"go.uber.org/zap"
)

type (
	commentPostPayload struct {
		Status  string `json:"status"`
		Comment string `json:"comment"`
	}
)

func (r *Repo) ListComments(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		number int
		err    error
	)

	if number, err = strconv.Atoi(c.Param("number")); err != nil {
		gb.ResponseWithError(
			c,
			e.NewErrorWithMessage(e.ErrorCodeInvalidRequest, "The number must be integer.", nil),
		)
		return
	}

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	d, err := r.i.FindDeploymentOfRepoByNumber(ctx, re, number)
	if err != nil {
		gb.LogWithError(r.log, "Failed to find the deployment.", err)
		gb.ResponseWithError(c, err)
		return
	}

	cmts, err := r.i.ListCommentsOfDeployment(ctx, d)
	if err != nil {
		gb.LogWithError(r.log, "Failed to list comments.", err)
		gb.ResponseWithError(c, err)
		return
	}

	gb.Response(c, http.StatusOK, cmts)
}

func (r *Repo) GetComment(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		id  int
		err error
	)
	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		gb.ResponseWithError(
			c,
			e.NewErrorWithMessage(e.ErrorCodeInvalidRequest, "The id must be integer.", nil),
		)
		return
	}

	cmt, err := r.i.FindCommentByID(ctx, id)
	if err != nil {
		gb.LogWithError(r.log, "Failed to find the comment.", err)
		gb.ResponseWithError(c, err)
		return
	}

	gb.Response(c, http.StatusOK, cmt)
}

func (r *Repo) CreateComment(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		number int
		err    error
	)

	if number, err = strconv.Atoi(c.Param("number")); err != nil {
		r.log.Warn("Failed to parse 'number'.", zap.Error(err))
		gb.ResponseWithError(
			c,
			e.NewErrorWithMessage(e.ErrorCodeInvalidRequest, "The number must be integer.", err),
		)
		return
	}

	p := &commentPostPayload{}
	if err := c.ShouldBindBodyWith(p, binding.JSON); err != nil {
		r.log.Warn("Failed to parse the payload.", zap.Error(err))
		gb.ResponseWithError(
			c,
			e.NewError(e.ErrorCodeInvalidRequest, err),
		)
		return
	}

	vu, _ := c.Get(gb.KeyUser)
	u := vu.(*ent.User)

	vr, _ := c.Get(KeyRepo)
	re := vr.(*ent.Repo)

	d, err := r.i.FindDeploymentOfRepoByNumber(ctx, re, number)
	if err != nil {
		gb.LogWithError(r.log, "Failed to find the deployment.", err)
		gb.ResponseWithError(c, err)
		return
	}

	cmt, err := r.i.CreateComment(ctx, &ent.Comment{
		Status:       comment.Status(p.Status),
		Comment:      p.Comment,
		UserID:       u.ID,
		DeploymentID: d.ID,
	})
	if err != nil {
		gb.LogWithError(r.log, "Failed to create a new comment.", err)
		gb.ResponseWithError(c, err)
		return
	}

	cmt, _ = r.i.FindCommentByID(ctx, cmt.ID)
	gb.Response(c, http.StatusCreated, cmt)
}
