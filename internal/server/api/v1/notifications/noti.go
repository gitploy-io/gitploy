package notifications

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"

	"github.com/hanjunlee/gitploy/ent"
	gb "github.com/hanjunlee/gitploy/internal/server/global"
)

type (
	Noti struct {
		i   Interactor
		log *zap.Logger
	}

	notiPayload struct {
		Checked bool `json:"checked"`
	}
)

func NewNoti(i Interactor) *Noti {
	return &Noti{
		i:   i,
		log: zap.L().Named("notifications"),
	}
}

func (n *Noti) ListNotifications(c *gin.Context) {
	var (
		page    = c.DefaultQuery("page", "1")
		perPage = c.DefaultQuery("perPage", "30")
	)
	v, _ := c.Get(gb.KeyUser)
	u, _ := v.(*ent.User)

	ctx := c.Request.Context()
	ns, err := n.i.ListNotifications(ctx, u, atoi(page), atoi(perPage))
	if err != nil {
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to list notifications.")
		return
	}

	gb.Response(c, http.StatusOK, ns)
}

func (n *Noti) UpdateNotification(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		id = c.Param("id")
	)

	v, _ := c.Get(gb.KeyUser)
	u, _ := v.(*ent.User)

	p := &notiPayload{}
	var err error
	if err := c.ShouldBindBodyWith(p, binding.JSON); err != nil {
		gb.ErrorResponse(c, http.StatusBadRequest, "It has failed to bind the body")
		return
	}

	nf, err := n.i.FindNotificationByID(ctx, atoi(id))
	if ent.IsNotFound(err) {
		gb.ErrorResponse(c, http.StatusNotFound, "There is no notification.")
		return
	} else if err != nil {
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to find the notification.")
		return
	}

	if nf.UserID != u.ID {
		gb.ErrorResponse(c, http.StatusForbidden, "There is no permission for the notification.")
		return
	}

	if nf.Checked != p.Checked {
		nf.Checked = p.Checked
		if nf, err = n.i.UpdateNotification(ctx, nf); err != nil {
			gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to update the notification.")
			return
		}
	}

	gb.Response(c, http.StatusOK, nf)
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
