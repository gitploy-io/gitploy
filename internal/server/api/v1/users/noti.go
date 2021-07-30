package users

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/hanjunlee/gitploy/ent"
	gb "github.com/hanjunlee/gitploy/internal/server/global"
)

type (
	notificationData struct {
		ID         int            `json:"id"`
		Type       string         `json:"type"`
		Repo       repoData       `json:"repo"`
		Deployment deploymentData `json:"deployment"`
		Approval   approvalData   `json:"approval"`
		Notified   bool           `json:"notified"`
		Checked    bool           `json:"checked"`
		CreatedAt  time.Time      `json:"created_at"`
		UpdatedAt  time.Time      `json:"updated_at"`
	}

	repoData struct {
		Namespace string `json:"namespace"`
		Name      string `json:"name"`
	}

	deploymentData struct {
		Number int    `json:"number"`
		Type   string `json:"type"`
		Ref    string `json:"ref"`
		Env    string `json:"env"`
		Status string `json:"status"`
		Login  string `json:"login"`
	}

	approvalData struct {
		Status string `json:"status"`
		Login  string `json:"login"`
	}

	notiPayload struct {
		Checked bool `json:"checked"`
	}
)

func (u *User) ListNotifications(c *gin.Context) {
	var (
		page    = c.DefaultQuery("page", "1")
		perPage = c.DefaultQuery("perPage", "30")
	)
	v, _ := c.Get(gb.KeyUser)
	vu, _ := v.(*ent.User)

	ctx := c.Request.Context()
	ns, err := u.i.ListNotifications(ctx, vu, atoi(page), atoi(perPage))
	if err != nil {
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to list notifications.")
		return
	}

	ds := []notificationData{}
	for _, n := range ns {
		ds = append(ds, mapToNotificationData(n))
	}

	gb.Response(c, http.StatusOK, ds)
}

func (u *User) UpdateNotification(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		id = c.Param("id")
	)

	v, _ := c.Get(gb.KeyUser)
	vu, _ := v.(*ent.User)

	p := &notiPayload{}
	var err error
	if err := c.ShouldBindBodyWith(p, binding.JSON); err != nil {
		gb.ErrorResponse(c, http.StatusBadRequest, "It has failed to bind the body")
		return
	}

	nf, err := u.i.FindNotificationByID(ctx, atoi(id))
	if ent.IsNotFound(err) {
		gb.ErrorResponse(c, http.StatusNotFound, "There is no notification.")
		return
	} else if err != nil {
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to find the notification.")
		return
	}

	if nf.UserID != vu.ID {
		gb.ErrorResponse(c, http.StatusForbidden, "There is no permission for the notification.")
		return
	}

	if nf.Checked != p.Checked {
		nf.Checked = p.Checked
		if nf, err = u.i.UpdateNotification(ctx, nf); err != nil {
			gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to update the notification.")
			return
		}
	}

	gb.Response(c, http.StatusOK, mapToNotificationData(nf))
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func mapToNotificationData(n *ent.Notification) notificationData {
	return notificationData{
		ID:   n.ID,
		Type: string(n.Type),
		Repo: repoData{
			Namespace: n.RepoNamespace,
			Name:      n.RepoName,
		},
		Deployment: deploymentData{
			Number: n.DeploymentNumber,
			Type:   n.DeploymentType,
			Ref:    n.DeploymentRef,
			Env:    n.DeploymentEnv,
			Status: n.DeploymentStatus,
			Login:  n.DeploymentLogin,
		},
		Approval: approvalData{
			Status: n.ApprovalStatus,
			Login:  n.ApprovalLogin,
		},
		Notified:  n.Notified,
		Checked:   n.Checked,
		CreatedAt: n.CreatedAt,
		UpdatedAt: n.UpdatedAt,
	}
}
