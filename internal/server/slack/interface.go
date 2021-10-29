//go:generate mockgen -source ./interface.go -destination ./mock/interactor.go -package mock

package slack

import (
	"context"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/vo"
)

type (
	Interactor interface {
		FindUserByID(ctx context.Context, id int64) (*ent.User, error)

		FindChatUserByID(ctx context.Context, id string) (*ent.ChatUser, error)
		CreateChatUser(ctx context.Context, cu *ent.ChatUser) (*ent.ChatUser, error)
		UpdateChatUser(ctx context.Context, cu *ent.ChatUser) (*ent.ChatUser, error)
		DeleteChatUser(ctx context.Context, cu *ent.ChatUser) error

		ListPermsOfRepo(ctx context.Context, r *ent.Repo, q string, page, perPage int) ([]*ent.Perm, error)
		FindPermOfRepo(ctx context.Context, r *ent.Repo, u *ent.User) (*ent.Perm, error)

		FindRepoOfUserByNamespaceName(ctx context.Context, u *ent.User, namespace, name string) (*ent.Repo, error)
		UpdateRepo(ctx context.Context, r *ent.Repo) (*ent.Repo, error)

		CreateCallback(ctx context.Context, cb *ent.Callback) (*ent.Callback, error)
		FindCallbackByHash(ctx context.Context, hash string) (*ent.Callback, error)

		ListDeploymentsOfRepo(ctx context.Context, r *ent.Repo, env string, status string, page, perPage int) ([]*ent.Deployment, error)
		FindDeploymentByID(ctx context.Context, id int) (*ent.Deployment, error)
		Deploy(ctx context.Context, u *ent.User, re *ent.Repo, d *ent.Deployment, env *vo.Env) (*ent.Deployment, error)
		GetConfig(ctx context.Context, u *ent.User, r *ent.Repo) (*vo.Config, error)

		CreateApproval(ctx context.Context, a *ent.Approval) (*ent.Approval, error)

		ListLocksOfRepo(ctx context.Context, r *ent.Repo) ([]*ent.Lock, error)
		FindLockOfRepoByEnv(ctx context.Context, r *ent.Repo, env string) (*ent.Lock, error)
		HasLockOfRepoForEnv(ctx context.Context, r *ent.Repo, env string) (bool, error)
		FindLockByID(ctx context.Context, id int) (*ent.Lock, error)
		CreateLock(ctx context.Context, l *ent.Lock) (*ent.Lock, error)
		DeleteLock(ctx context.Context, l *ent.Lock) error

		SubscribeEvent(fn func(e *ent.Event)) error
		UnsubscribeEvent(fn func(e *ent.Event)) error

		CheckNotificationRecordOfEvent(ctx context.Context, e *ent.Event) bool
		CreateEvent(ctx context.Context, e *ent.Event) (*ent.Event, error)

		GetCommit(ctx context.Context, u *ent.User, r *ent.Repo, sha string) (*vo.Commit, error)
		GetBranch(ctx context.Context, u *ent.User, r *ent.Repo, branch string) (*vo.Branch, error)
		GetTag(ctx context.Context, u *ent.User, r *ent.Repo, tag string) (*vo.Tag, error)
	}
)
