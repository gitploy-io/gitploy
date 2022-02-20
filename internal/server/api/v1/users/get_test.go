package users

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gitploy-io/gitploy/internal/server/api/v1/users/mock"
	"github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/model/ent"
	gm "github.com/golang/mock/gomock"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func TestUsers_GetMe(t *testing.T) {
	t.Run("Return user's information with the 'hash' field.", func(t *testing.T) {
		const hash = "HASH_VALUE"

		t.Log("Start mocking:")
		ctrl := gm.NewController(t)
		i := mock.NewMockInteractor(ctrl)

		t.Log("\tFind the user.")
		i.EXPECT().
			FindUserByID(gm.Any(), gm.AssignableToTypeOf(int64(1))).
			Return(&ent.User{Hash: hash}, nil)

		api := NewUserAPI(i)
		r := gin.New()
		r.GET("/user",
			func(c *gin.Context) {
				c.Set(global.KeyUser, &ent.User{})
			},
			api.GetMe)

		req, _ := http.NewRequest("GET", "/user", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		t.Log("Evaluate the return value.")
		if w.Code != http.StatusOK {
			t.Fatalf("Code = %v, wanted %v", w.Code, http.StatusOK)
		}

		d := extendedUserData{}
		if err := json.Unmarshal(w.Body.Bytes(), &d); err != nil {
			t.Fatalf("Failed to unmarshal: %v", err)
		}

		if d.Hash != hash {
			t.Fatalf("Hash = %v, wanted %v", d.Hash, hash)
		}
	})
}
