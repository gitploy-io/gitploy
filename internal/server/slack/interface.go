package slack

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/vo"
)

type (
	Interactor interface {
		FindUserWithChatUserByChatUserID(ctx context.Context, id string) (*ent.User, error)
		SaveChatUser(ctx context.Context, u *ent.User, cu *ent.ChatUser) (*ent.ChatUser, error)
		FindRepoByNamespaceName(ctx context.Context, u *ent.User, namespace, name string) (*ent.Repo, error)
		GetConfig(ctx context.Context, u *ent.User, r *ent.Repo) (*vo.Config, error)
		CreateChatCallback(ctx context.Context, cb *ent.ChatCallback) (*ent.ChatCallback, error)
	}
)
