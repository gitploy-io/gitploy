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
		CreateDeployChatCallback(ctx context.Context, cu *ent.ChatUser, repo *ent.Repo, cb *ent.ChatCallback) (*ent.ChatCallback, error)
		FindChatCallbackByID(ctx context.Context, id string) (*ent.ChatCallback, error)
		FindChatCallbackWithEdgesByID(ctx context.Context, id string) (*ent.ChatCallback, error)
		CloseChatCallback(ctx context.Context, cb *ent.ChatCallback) (*ent.ChatCallback, error)
		Deploy(ctx context.Context, u *ent.User, re *ent.Repo, d *ent.Deployment) (*ent.Deployment, error)
		Publish(ctx context.Context, resource interface{}) error
	}
)
