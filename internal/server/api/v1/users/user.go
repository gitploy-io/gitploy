package users

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/hanjunlee/gitploy/ent"
	gb "github.com/hanjunlee/gitploy/internal/server/global"
)

type (
	User struct {
		i   Interactor
		log *zap.Logger
	}

	userData struct {
		ID        string        `json:"id,omitempty"`
		Login     string        `json:"login,omitempty"`
		Avatar    string        `json:"avatar,omitempty"`
		Admin     bool          `json:"admin"`
		CreatedAt time.Time     `json:"created_at,omitempty"`
		UpdatedAt time.Time     `json:"updated_at,omitempty"`
		Edges     userEdgesData `json:"edges"`
	}

	userEdgesData struct {
		ChatUserData *chatUserData `json:"chat_user,omitempty"`
	}

	chatUserData struct {
		ID        string    `json:"id,omitempty"`
		CreatedAt time.Time `json:"created_at,omitempty"`
		UpdatedAt time.Time `json:"updated_at,omitempty"`
	}
)

func NewUser(i Interactor) *User {
	return &User{
		i:   i,
		log: zap.L().Named("users"),
	}
}

func (u *User) Me(c *gin.Context) {
	v, _ := c.Get(gb.KeyUser)
	uv, _ := v.(*ent.User)

	ctx := c.Request.Context()

	uv, err := u.i.FindUserByID(ctx, uv.ID)
	if err != nil {
		u.log.Error("failed to find the user.", zap.Error(err))
		gb.ErrorResponse(c, http.StatusInternalServerError, "It has failed to find the user.")
		return
	}

	gb.Response(c, http.StatusOK, mapUserToUserData(uv))
}

func mapUserToUserData(u *ent.User) *userData {
	var cud *chatUserData
	if cu := u.Edges.ChatUser; cu != nil {
		cud = &chatUserData{
			ID:        cu.ID,
			CreatedAt: cu.CreatedAt,
			UpdatedAt: cu.UpdatedAt,
		}
	}

	return &userData{
		ID:        u.ID,
		Login:     u.Login,
		Avatar:    u.Avatar,
		Admin:     u.Admin,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Edges: userEdgesData{
			ChatUserData: cud,
		},
	}
}
