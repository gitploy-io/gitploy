//go:generate mockgen -source ./interface.go -destination ./mock/pkg.go -package mock

package interactor

import (
	"context"
	"time"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/approval"
	"github.com/gitploy-io/gitploy/ent/deployment"
	"github.com/gitploy-io/gitploy/vo"
)

type (
	Store interface {
		CountUsers(context.Context) (int, error)
		ListUsers(ctx context.Context, login string, page, perPage int) ([]*ent.User, error)
		FindUserByID(ctx context.Context, id int64) (*ent.User, error)
		FindUserByHash(ctx context.Context, hash string) (*ent.User, error)
		FindUserByLogin(ctx context.Context, login string) (*ent.User, error)
		CreateUser(ctx context.Context, u *ent.User) (*ent.User, error)
		UpdateUser(ctx context.Context, u *ent.User) (*ent.User, error)
		DeleteUser(ctx context.Context, u *ent.User) error

		FindChatUserByID(ctx context.Context, id string) (*ent.ChatUser, error)
		CreateChatUser(ctx context.Context, cu *ent.ChatUser) (*ent.ChatUser, error)
		UpdateChatUser(ctx context.Context, cu *ent.ChatUser) (*ent.ChatUser, error)
		DeleteChatUser(ctx context.Context, cu *ent.ChatUser) error

		CountActiveRepos(ctx context.Context) (int, error)
		CountRepos(ctx context.Context) (int, error)
		ListReposOfUser(ctx context.Context, u *ent.User, q, namespace, name string, sorted bool, page, perPage int) ([]*ent.Repo, error)
		FindRepoOfUserByID(ctx context.Context, u *ent.User, id int64) (*ent.Repo, error)
		FindRepoOfUserByNamespaceName(ctx context.Context, u *ent.User, namespace, name string) (*ent.Repo, error)
		FindRepoByID(ctx context.Context, id int64) (*ent.Repo, error)
		SyncRepo(ctx context.Context, r *vo.RemoteRepo) (*ent.Repo, error)
		UpdateRepo(ctx context.Context, r *ent.Repo) (*ent.Repo, error)
		Activate(ctx context.Context, r *ent.Repo) (*ent.Repo, error)
		Deactivate(ctx context.Context, r *ent.Repo) (*ent.Repo, error)

		ListPermsOfRepo(ctx context.Context, r *ent.Repo, q string, page, perPage int) ([]*ent.Perm, error)
		FindPermOfRepo(ctx context.Context, r *ent.Repo, u *ent.User) (*ent.Perm, error)
		CreatePerm(ctx context.Context, p *ent.Perm) (*ent.Perm, error)
		UpdatePerm(ctx context.Context, p *ent.Perm) (*ent.Perm, error)
		DeletePermsOfUserLessThanSyncedAt(ctx context.Context, u *ent.User, t time.Time) (int, error)

		CountDeployments(ctx context.Context) (int, error)
		SearchDeployments(ctx context.Context, u *ent.User, s []deployment.Status, owned bool, from time.Time, to time.Time, page, perPage int) ([]*ent.Deployment, error)
		ListInactiveDeploymentsLessThanTime(ctx context.Context, t time.Time, page, perPage int) ([]*ent.Deployment, error)
		ListDeploymentsOfRepo(ctx context.Context, r *ent.Repo, env string, status string, page, perPage int) ([]*ent.Deployment, error)
		FindDeploymentByID(ctx context.Context, id int) (*ent.Deployment, error)
		FindDeploymentByUID(ctx context.Context, uid int64) (*ent.Deployment, error)
		FindDeploymentOfRepoByNumber(ctx context.Context, r *ent.Repo, number int) (*ent.Deployment, error)
		FindPrevSuccessDeployment(ctx context.Context, d *ent.Deployment) (*ent.Deployment, error)
		GetNextDeploymentNumberOfRepo(ctx context.Context, r *ent.Repo) (int, error)
		CreateDeployment(ctx context.Context, d *ent.Deployment) (*ent.Deployment, error)
		UpdateDeployment(ctx context.Context, d *ent.Deployment) (*ent.Deployment, error)

		CreateDeploymentStatus(ctx context.Context, s *ent.DeploymentStatus) (*ent.DeploymentStatus, error)
		SyncDeploymentStatus(ctx context.Context, ds *ent.DeploymentStatus) (*ent.DeploymentStatus, error)

		FindDeploymentStatisticsOfRepoByEnv(ctx context.Context, r *ent.Repo, env string) (*ent.DeploymentStatistics, error)
		CreateDeploymentStatistics(ctx context.Context, s *ent.DeploymentStatistics) (*ent.DeploymentStatistics, error)
		UpdateDeploymentStatistics(ctx context.Context, s *ent.DeploymentStatistics) (*ent.DeploymentStatistics, error)

		ListAllDeploymentStatistics(ctx context.Context) ([]*ent.DeploymentStatistics, error)
		ListDeploymentStatisticsGreaterThanTime(ctx context.Context, updated time.Time) ([]*ent.DeploymentStatistics, error)

		CreateCallback(ctx context.Context, cb *ent.Callback) (*ent.Callback, error)
		FindCallbackByHash(ctx context.Context, hash string) (*ent.Callback, error)

		SearchApprovals(ctx context.Context, u *ent.User, s []approval.Status, from time.Time, to time.Time, page, perPage int) ([]*ent.Approval, error)
		ListApprovals(ctx context.Context, d *ent.Deployment) ([]*ent.Approval, error)
		FindApprovalByID(ctx context.Context, id int) (*ent.Approval, error)
		FindApprovalOfUser(ctx context.Context, d *ent.Deployment, u *ent.User) (*ent.Approval, error)
		CreateApproval(ctx context.Context, a *ent.Approval) (*ent.Approval, error)
		UpdateApproval(ctx context.Context, a *ent.Approval) (*ent.Approval, error)
		DeleteApproval(ctx context.Context, a *ent.Approval) error

		SearchReviews(ctx context.Context, u *ent.User) ([]*ent.Review, error)
		ListReviews(ctx context.Context, d *ent.Deployment) ([]*ent.Review, error)
		FindReviewOfUser(ctx context.Context, u *ent.User, d *ent.Deployment) (*ent.Review, error)
		CreateReview(ctx context.Context, rv *ent.Review) (*ent.Review, error)
		UpdateReview(ctx context.Context, rv *ent.Review) (*ent.Review, error)

		ListExpiredLocksLessThanTime(ctx context.Context, t time.Time) ([]*ent.Lock, error)
		ListLocksOfRepo(ctx context.Context, r *ent.Repo) ([]*ent.Lock, error)
		FindLockOfRepoByEnv(ctx context.Context, r *ent.Repo, env string) (*ent.Lock, error)
		HasLockOfRepoForEnv(ctx context.Context, r *ent.Repo, env string) (bool, error)
		FindLockByID(ctx context.Context, id int) (*ent.Lock, error)
		CreateLock(ctx context.Context, l *ent.Lock) (*ent.Lock, error)
		UpdateLock(ctx context.Context, l *ent.Lock) (*ent.Lock, error)
		DeleteLock(ctx context.Context, l *ent.Lock) error

		ListEventsGreaterThanTime(ctx context.Context, t time.Time) ([]*ent.Event, error)
		CreateEvent(ctx context.Context, e *ent.Event) (*ent.Event, error)
		CheckNotificationRecordOfEvent(ctx context.Context, e *ent.Event) bool
	}

	SCM interface {
		GetRemoteUserByToken(ctx context.Context, token string) (*vo.RemoteUser, error)
		ListRemoteOrgsByToken(ctx context.Context, token string) ([]string, error)

		ListRemoteRepos(ctx context.Context, u *ent.User) ([]*vo.RemoteRepo, error)

		// SCM returns the deployment with UID and SHA.
		CreateRemoteDeployment(ctx context.Context, u *ent.User, r *ent.Repo, d *ent.Deployment, e *vo.Env) (*vo.RemoteDeployment, error)
		CancelDeployment(ctx context.Context, u *ent.User, r *ent.Repo, d *ent.Deployment, s *ent.DeploymentStatus) error
		GetConfig(ctx context.Context, u *ent.User, r *ent.Repo) (*vo.Config, error)

		CreateWebhook(ctx context.Context, u *ent.User, r *ent.Repo, c *vo.WebhookConfig) (int64, error)
		DeleteWebhook(ctx context.Context, u *ent.User, r *ent.Repo, id int64) error

		ListCommits(ctx context.Context, u *ent.User, r *ent.Repo, branch string, page, perPage int) ([]*vo.Commit, error)
		CompareCommits(ctx context.Context, u *ent.User, r *ent.Repo, base, head string, page, perPage int) ([]*vo.Commit, []*vo.CommitFile, error)
		GetCommit(ctx context.Context, u *ent.User, r *ent.Repo, sha string) (*vo.Commit, error)
		ListCommitStatuses(ctx context.Context, u *ent.User, r *ent.Repo, sha string) ([]*vo.Status, error)

		ListBranches(ctx context.Context, u *ent.User, r *ent.Repo, page, perPage int) ([]*vo.Branch, error)
		GetBranch(ctx context.Context, u *ent.User, r *ent.Repo, branch string) (*vo.Branch, error)

		ListTags(ctx context.Context, u *ent.User, r *ent.Repo, page, perPage int) ([]*vo.Tag, error)
		GetTag(ctx context.Context, u *ent.User, r *ent.Repo, tag string) (*vo.Tag, error)

		GetRateLimit(ctx context.Context, u *ent.User) (*vo.RateLimit, error)
	}
)
