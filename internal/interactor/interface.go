package interactor

import (
	"context"
	"time"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/vo"
)

type (
	Store interface {
		FindUserByID(ctx context.Context, id string) (*ent.User, error)
		FindUserByHash(ctx context.Context, hash string) (*ent.User, error)
		FindUserByLogin(ctx context.Context, login string) (*ent.User, error)
		CreateUser(ctx context.Context, u *ent.User) (*ent.User, error)
		UpdateUser(ctx context.Context, u *ent.User) (*ent.User, error)

		FindChatUserByID(ctx context.Context, id string) (*ent.ChatUser, error)
		CreateChatUser(ctx context.Context, u *ent.User, cu *ent.ChatUser) (*ent.ChatUser, error)
		UpdateChatUser(ctx context.Context, u *ent.User, cu *ent.ChatUser) (*ent.ChatUser, error)

		ListReposOfUser(ctx context.Context, u *ent.User, q string, page, perPage int) ([]*ent.Repo, error)
		ListSortedReposOfUser(ctx context.Context, u *ent.User, q string, page, perPage int) ([]*ent.Repo, error)
		FindRepoByID(ctx context.Context, id string) (*ent.Repo, error)
		FindRepoByNamespaceName(ctx context.Context, namespace, name string) (*ent.Repo, error)
		UpdateRepo(ctx context.Context, r *ent.Repo) (*ent.Repo, error)
		Activate(ctx context.Context, r *ent.Repo) (*ent.Repo, error)
		Deactivate(ctx context.Context, r *ent.Repo) (*ent.Repo, error)

		FindPermOfRepo(ctx context.Context, r *ent.Repo, u *ent.User) (*ent.Perm, error)
		SyncPerm(ctx context.Context, user *ent.User, perm *ent.Perm, sync time.Time) error

		ListDeployments(ctx context.Context, r *ent.Repo, env string, status string, page, perPage int) ([]*ent.Deployment, error)
		FindDeploymentByID(ctx context.Context, id int) (*ent.Deployment, error)
		FindDeploymentByUID(ctx context.Context, uid int64) (*ent.Deployment, error)
		FindDeploymentOfRepoByNumber(ctx context.Context, r *ent.Repo, number int) (*ent.Deployment, error)
		GetNextDeploymentNumberOfRepo(ctx context.Context, r *ent.Repo) (int, error)
		CreateDeployment(ctx context.Context, d *ent.Deployment) (*ent.Deployment, error)
		UpdateDeployment(ctx context.Context, d *ent.Deployment) (*ent.Deployment, error)

		CreateDeploymentStatus(ctx context.Context, s *ent.DeploymentStatus) (*ent.DeploymentStatus, error)

		CreateDeployChatCallback(ctx context.Context, cu *ent.ChatUser, repo *ent.Repo, cb *ent.ChatCallback) (*ent.ChatCallback, error)
		FindChatCallbackByState(ctx context.Context, state string) (*ent.ChatCallback, error)
		CloseChatCallback(ctx context.Context, cb *ent.ChatCallback) (*ent.ChatCallback, error)

		ListApprovals(ctx context.Context, d *ent.Deployment) ([]*ent.Approval, error)
		FindApprovalOfUser(ctx context.Context, d *ent.Deployment, u *ent.User) (*ent.Approval, error)
		CreateApproval(ctx context.Context, a *ent.Approval) (*ent.Approval, error)
		UpdateApproval(ctx context.Context, a *ent.Approval) (*ent.Approval, error)

		ListNotifications(ctx context.Context, u *ent.User, page, perPage int) ([]*ent.Notification, error)
		ListNotificationsFromTime(ctx context.Context, t time.Time) ([]*ent.Notification, error)
		FindNotificationByID(ctx context.Context, id int) (*ent.Notification, error)
		CreateNotification(ctx context.Context, n *ent.Notification) (*ent.Notification, error)
		UpdateNotification(ctx context.Context, n *ent.Notification) (*ent.Notification, error)
	}

	SCM interface {
		GetUser(ctx context.Context, token string) (*ent.User, error)
		// TODO: fix type of return value to prevent using the ent package.
		GetAllPermsWithRepo(ctx context.Context, token string) ([]*ent.Perm, error)

		// SCM returns the deployment with UID and SHA.
		CreateDeployment(ctx context.Context, u *ent.User, r *ent.Repo, d *ent.Deployment, e *vo.Env) (int64, error)
		GetConfig(ctx context.Context, u *ent.User, r *ent.Repo) (*vo.Config, error)

		CreateWebhook(ctx context.Context, u *ent.User, r *ent.Repo, c *vo.WebhookConfig) (int64, error)
		DeleteWebhook(ctx context.Context, u *ent.User, r *ent.Repo, id int64) error

		ListCommits(ctx context.Context, u *ent.User, r *ent.Repo, branch string, page, perPage int) ([]*vo.Commit, error)
		GetCommit(ctx context.Context, u *ent.User, r *ent.Repo, sha string) (*vo.Commit, error)
		ListCommitStatuses(ctx context.Context, u *ent.User, r *ent.Repo, sha string) ([]*vo.Status, error)

		ListBranches(ctx context.Context, u *ent.User, r *ent.Repo, page, perPage int) ([]*vo.Branch, error)
		GetBranch(ctx context.Context, u *ent.User, r *ent.Repo, branch string) (*vo.Branch, error)

		ListTags(ctx context.Context, u *ent.User, r *ent.Repo, page, perPage int) ([]*vo.Tag, error)
		GetTag(ctx context.Context, u *ent.User, r *ent.Repo, tag string) (*vo.Tag, error)
	}

	Chat interface {
		NotifyDeployment(ctx context.Context, cu *ent.ChatUser, d *ent.Deployment) error
	}
)
