package slack

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/notification"
	"github.com/hanjunlee/gitploy/vo"
)

type (
	Interactor interface {
		FindChatUserByID(ctx context.Context, id string) (*ent.ChatUser, error)
		SaveChatUser(ctx context.Context, u *ent.User, cu *ent.ChatUser) (*ent.ChatUser, error)

		FindPermOfRepo(ctx context.Context, r *ent.Repo, u *ent.User) (*ent.Perm, error)

		FindRepoOfUserByNamespaceName(ctx context.Context, u *ent.User, namespace, name string) (*ent.Repo, error)

		CreateDeployChatCallback(ctx context.Context, cu *ent.ChatUser, repo *ent.Repo, cb *ent.ChatCallback) (*ent.ChatCallback, error)
		FindChatCallbackByState(ctx context.Context, state string) (*ent.ChatCallback, error)
		CloseChatCallback(ctx context.Context, cb *ent.ChatCallback) (*ent.ChatCallback, error)

		ListDeployments(ctx context.Context, r *ent.Repo, env string, status string, page, perPage int) ([]*ent.Deployment, error)
		FindDeploymentByID(ctx context.Context, id int) (*ent.Deployment, error)
		GetNextDeploymentNumberOfRepo(ctx context.Context, r *ent.Repo) (int, error)
		Deploy(ctx context.Context, u *ent.User, re *ent.Repo, d *ent.Deployment, env *vo.Env) (*ent.Deployment, error)
		Rollback(ctx context.Context, u *ent.User, re *ent.Repo, d *ent.Deployment, env *vo.Env) (*ent.Deployment, error)
		GetConfig(ctx context.Context, u *ent.User, r *ent.Repo) (*vo.Config, error)

		Publish(ctx context.Context, typ notification.Type, r *ent.Repo, d *ent.Deployment, a *ent.Approval) error
		Subscribe(func(*ent.User, *ent.Notification)) error

		GetCommit(ctx context.Context, u *ent.User, r *ent.Repo, sha string) (*vo.Commit, error)
		GetBranch(ctx context.Context, u *ent.User, r *ent.Repo, branch string) (*vo.Branch, error)
		GetTag(ctx context.Context, u *ent.User, r *ent.Repo, tag string) (*vo.Tag, error)
	}
)
