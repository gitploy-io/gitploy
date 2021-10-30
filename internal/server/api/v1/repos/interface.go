//go:generate mockgen -source ./interface.go -destination ./mock/interactor.go -package mock

package repos

import (
	"context"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/vo"
)

type (
	Interactor interface {
		FindUserByID(ctx context.Context, id int64) (*ent.User, error)

		ListPermsOfRepo(ctx context.Context, r *ent.Repo, q string, page, perPage int) ([]*ent.Perm, error)
		FindPermOfRepo(ctx context.Context, r *ent.Repo, u *ent.User) (*ent.Perm, error)

		ListReposOfUser(ctx context.Context, u *ent.User, q, namespace, name string, sorted bool, page, perPage int) ([]*ent.Repo, error)
		FindRepoOfUserByID(ctx context.Context, u *ent.User, id int64) (*ent.Repo, error)
		FindRepoOfUserByNamespaceName(ctx context.Context, u *ent.User, namespace, name string) (*ent.Repo, error)
		UpdateRepo(ctx context.Context, r *ent.Repo) (*ent.Repo, error)
		ActivateRepo(ctx context.Context, u *ent.User, r *ent.Repo, c *vo.WebhookConfig) (*ent.Repo, error)
		DeactivateRepo(ctx context.Context, u *ent.User, r *ent.Repo) (*ent.Repo, error)

		ListDeploymentsOfRepo(ctx context.Context, r *ent.Repo, env string, status string, page, perPage int) ([]*ent.Deployment, error)
		FindDeploymentByID(ctx context.Context, id int) (*ent.Deployment, error)
		FindDeploymentOfRepoByNumber(ctx context.Context, r *ent.Repo, number int) (*ent.Deployment, error)
		FindPrevSuccessDeployment(ctx context.Context, d *ent.Deployment) (*ent.Deployment, error)
		IsApproved(ctx context.Context, d *ent.Deployment) bool
		Deploy(ctx context.Context, u *ent.User, re *ent.Repo, d *ent.Deployment, env *vo.Env) (*ent.Deployment, error)
		DeployToRemote(ctx context.Context, u *ent.User, r *ent.Repo, d *ent.Deployment, env *vo.Env) (*ent.Deployment, error)
		GetConfig(ctx context.Context, u *ent.User, r *ent.Repo) (*vo.Config, error)

		ListApprovals(ctx context.Context, d *ent.Deployment) ([]*ent.Approval, error)
		FindApprovalByID(ctx context.Context, id int) (*ent.Approval, error)
		FindApprovalOfUser(ctx context.Context, d *ent.Deployment, u *ent.User) (*ent.Approval, error)
		CreateApproval(ctx context.Context, a *ent.Approval) (*ent.Approval, error)
		UpdateApproval(ctx context.Context, a *ent.Approval) (*ent.Approval, error)
		DeleteApproval(ctx context.Context, a *ent.Approval) error

		ListLocksOfRepo(ctx context.Context, r *ent.Repo) ([]*ent.Lock, error)
		HasLockOfRepoForEnv(ctx context.Context, r *ent.Repo, env string) (bool, error)
		FindLockByID(ctx context.Context, id int) (*ent.Lock, error)
		CreateLock(ctx context.Context, l *ent.Lock) (*ent.Lock, error)
		UpdateLock(ctx context.Context, l *ent.Lock) (*ent.Lock, error)
		DeleteLock(ctx context.Context, l *ent.Lock) error

		CreateEvent(ctx context.Context, e *ent.Event) (*ent.Event, error)

		ListCommits(ctx context.Context, u *ent.User, r *ent.Repo, branch string, page, perPage int) ([]*vo.Commit, error)
		CompareCommits(ctx context.Context, u *ent.User, r *ent.Repo, base, head string, page, perPage int) ([]*vo.Commit, []*vo.CommitFile, error)
		GetCommit(ctx context.Context, u *ent.User, r *ent.Repo, sha string) (*vo.Commit, error)
		ListCommitStatuses(ctx context.Context, u *ent.User, r *ent.Repo, sha string) ([]*vo.Status, error)

		ListBranches(ctx context.Context, u *ent.User, r *ent.Repo, page, perPage int) ([]*vo.Branch, error)
		GetBranch(ctx context.Context, u *ent.User, r *ent.Repo, branch string) (*vo.Branch, error)

		ListTags(ctx context.Context, u *ent.User, r *ent.Repo, page, perPage int) ([]*vo.Tag, error)
		GetTag(ctx context.Context, u *ent.User, r *ent.Repo, tag string) (*vo.Tag, error)
	}
)
