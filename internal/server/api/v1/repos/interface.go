package repos

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/notification"
	"github.com/hanjunlee/gitploy/vo"
)

type (
	Interactor interface {
		FindUserByID(ctx context.Context, id string) (*ent.User, error)

		ListPermsOfRepo(ctx context.Context, r *ent.Repo, q string, page, perPage int) ([]*ent.Perm, error)
		FindPermOfRepo(ctx context.Context, r *ent.Repo, u *ent.User) (*ent.Perm, error)

		ListReposOfUser(ctx context.Context, u *ent.User, sorted bool, q string, page, perPage int) ([]*ent.Repo, error)
		FindRepoOfUserByID(ctx context.Context, u *ent.User, id string) (*ent.Repo, error)
		FindRepoOfUserByNamespaceName(ctx context.Context, u *ent.User, namespace, name string) (*ent.Repo, error)
		UpdateRepo(ctx context.Context, r *ent.Repo) (*ent.Repo, error)
		ActivateRepo(ctx context.Context, u *ent.User, r *ent.Repo, c *vo.WebhookConfig) (*ent.Repo, error)
		DeactivateRepo(ctx context.Context, u *ent.User, r *ent.Repo) (*ent.Repo, error)

		ListDeployments(ctx context.Context, r *ent.Repo, env string, status string, page, perPage int) ([]*ent.Deployment, error)
		FindDeploymentByID(ctx context.Context, id int) (*ent.Deployment, error)
		FindDeploymentOfRepoByNumber(ctx context.Context, r *ent.Repo, number int) (*ent.Deployment, error)
		GetNextDeploymentNumberOfRepo(ctx context.Context, r *ent.Repo) (int, error)
		Deploy(ctx context.Context, u *ent.User, re *ent.Repo, d *ent.Deployment, env *vo.Env) (*ent.Deployment, error)
		IsApproved(ctx context.Context, d *ent.Deployment) bool
		CreateDeploymentToSCM(ctx context.Context, u *ent.User, re *ent.Repo, d *ent.Deployment, env *vo.Env) (*ent.Deployment, error)
		Rollback(ctx context.Context, u *ent.User, re *ent.Repo, d *ent.Deployment, env *vo.Env) (*ent.Deployment, error)
		GetConfig(ctx context.Context, u *ent.User, r *ent.Repo) (*vo.Config, error)

		ListApprovals(ctx context.Context, d *ent.Deployment) ([]*ent.Approval, error)
		FindApprovalByID(ctx context.Context, id int) (*ent.Approval, error)
		FindApprovalOfUser(ctx context.Context, d *ent.Deployment, u *ent.User) (*ent.Approval, error)
		CreateApproval(ctx context.Context, a *ent.Approval) (*ent.Approval, error)
		UpdateApproval(ctx context.Context, a *ent.Approval) (*ent.Approval, error)
		DeleteApproval(ctx context.Context, a *ent.Approval) error

		Publish(ctx context.Context, typ notification.Type, r *ent.Repo, d *ent.Deployment, a *ent.Approval) error

		ListCommits(ctx context.Context, u *ent.User, r *ent.Repo, branch string, page, perPage int) ([]*vo.Commit, error)
		GetCommit(ctx context.Context, u *ent.User, r *ent.Repo, sha string) (*vo.Commit, error)
		ListCommitStatuses(ctx context.Context, u *ent.User, r *ent.Repo, sha string) ([]*vo.Status, error)

		ListBranches(ctx context.Context, u *ent.User, r *ent.Repo, page, perPage int) ([]*vo.Branch, error)
		GetBranch(ctx context.Context, u *ent.User, r *ent.Repo, branch string) (*vo.Branch, error)

		ListTags(ctx context.Context, u *ent.User, r *ent.Repo, page, perPage int) ([]*vo.Tag, error)
		GetTag(ctx context.Context, u *ent.User, r *ent.Repo, tag string) (*vo.Tag, error)
	}
)
