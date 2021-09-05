package interactor

import (
	"context"

	"github.com/gitploy-io/gitploy/vo"
)

func (i *Interactor) GetRemoteUserByToken(ctx context.Context, token string) (*vo.RemoteUser, error) {
	return i.SCM.GetUser(ctx, token)
}

func (i *Interactor) IsAdminUser(ctx context.Context, login string) bool {
	for _, admin := range i.admins {
		if login == admin {
			return true
		}
	}

	return false
}
