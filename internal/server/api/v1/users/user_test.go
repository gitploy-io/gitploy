package users

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AlekSi/pointer"
	"github.com/gin-gonic/gin"
	"github.com/gitploy-io/gitploy/internal/server/api/v1/users/mock"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/golang/mock/gomock"
)

func TestUser_UpdateUser(t *testing.T) {
	input := struct {
		ID      int64
		Payload *userPatchPayload
	}{
		ID: 1,
		Payload: &userPatchPayload{
			Admin: pointer.ToBool(true),
		},
	}

	ctrl := gomock.NewController(t)
	m := mock.NewMockInteractor(ctrl)

	ctx := gomock.Any()

	t.Log("FindUserByID returns non-admin user.")
	m.
		EXPECT().
		FindUserByID(ctx, input.ID).
		Return(&ent.User{
			ID:    input.ID,
			Admin: false,
		}, nil)

	t.Log("UpdateUser updates the user admin.")
	m.
		EXPECT().
		UpdateUser(ctx, gomock.Eq(&ent.User{
			ID:    input.ID,
			Admin: true,
		})).
		DoAndReturn(func(ctx context.Context, u *ent.User) (*ent.User, error) {
			return u, nil
		})

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.PATCH("/users/:id", NewUserAPI(m).UpdateUser)

	p, _ := json.Marshal(input.Payload)
	req, _ := http.NewRequest("PATCH", fmt.Sprintf("/users/%d", input.ID), bytes.NewBuffer(p))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Code = %v, wanted %v", w.Code, http.StatusOK)
	}
}
