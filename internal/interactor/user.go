package interactor

import (
	"context"

	"github.com/hanjunlee/gitploy/vo"
)

func (i *Interactor) GetRemoteUserByToken(ctx context.Context, token string) (*vo.RemoteUser, error) {
	return i.SCM.GetUser(ctx, token)
}
