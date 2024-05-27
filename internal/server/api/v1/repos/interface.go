//go:generate mockgen -source ./interface.go -destination ./mock/interactor.go -package mock

package repos

import (
	"context"

	i "github.com/gitploy-io/gitploy/internal/interactor"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/extent"
)

type (
	Interactor interface {
		FindUserByID(ctx context.Context, id int64) (*ent.User, error)

		ListPermsOfRepo(ctx context.Context, r *ent.Repo, opt *i.ListPermsOfRepoOptions) ([]*ent.Perm, error)
		FindPermOfRepo(ctx context.Context, r *ent.Repo, u *ent.User) (*ent.Perm, error)

		ListReposOfUser(ctx context.Context, u *ent.User, opt *i.ListReposOfUserOptions) ([]*ent.Repo, error)
		FindRepoOfUserByID(ctx context.Context, u *ent.User, id int64) (*ent.Repo, error)
		FindRepoOfUserByNamespaceName(ctx context.Context, u *ent.User, opt *i.FindRepoOfUserByNamespaceNameOptions) (*ent.Repo, error)
		UpdateRepo(ctx context.Context, r *ent.Repo) (*ent.Repo, error)
		ActivateRepo(ctx context.Context, u *ent.User, r *ent.Repo) (*ent.Repo, error)
		DeactivateRepo(ctx context.Context, u *ent.User, r *ent.Repo) (*ent.Repo, error)

		ListDeploymentsOfRepo(ctx context.Context, r *ent.Repo, opt *i.ListDeploymentsOfRepoOptions) ([]*ent.Deployment, error)
		FindDeploymentByID(ctx context.Context, id int) (*ent.Deployment, error)
		FindDeploymentOfRepoByNumber(ctx context.Context, r *ent.Repo, number int) (*ent.Deployment, error)
		FindPrevSuccessDeployment(ctx context.Context, d *ent.Deployment) (*ent.Deployment, error)
		Deploy(ctx context.Context, u *ent.User, re *ent.Repo, d *ent.Deployment, env *extent.Env) (*ent.Deployment, error)
		DeployToRemote(ctx context.Context, u *ent.User, r *ent.Repo, d *ent.Deployment, env *extent.Env) (*ent.Deployment, error)
		GetConfig(ctx context.Context, u *ent.User, r *ent.Repo) (*extent.Config, error)
		GetEvaluatedConfig(ctx context.Context, u *ent.User, r *ent.Repo, v *extent.EvalValues) (*extent.Config, error)

		ListDeploymentStatuses(ctx context.Context, d *ent.Deployment) ([]*ent.DeploymentStatus, error)
		CreateRemoteDeploymentStatus(ctx context.Context, u *ent.User, r *ent.Repo, d *ent.Deployment, ds *extent.RemoteDeploymentStatus) (*extent.RemoteDeploymentStatus, error)

		ListReviews(ctx context.Context, d *ent.Deployment) ([]*ent.Review, error)
		FindReviewOfUser(ctx context.Context, u *ent.User, d *ent.Deployment) (*ent.Review, error)
		UpdateReview(ctx context.Context, rv *ent.Review) (*ent.Review, error)
		RespondReview(ctx context.Context, rv *ent.Review) (*ent.Review, error)

		ListLocksOfRepo(ctx context.Context, r *ent.Repo) ([]*ent.Lock, error)
		HasLockOfRepoForEnv(ctx context.Context, r *ent.Repo, env string) (bool, error)
		FindLockByID(ctx context.Context, id int) (*ent.Lock, error)
		CreateLock(ctx context.Context, l *ent.Lock) (*ent.Lock, error)
		UpdateLock(ctx context.Context, l *ent.Lock) (*ent.Lock, error)
		DeleteLock(ctx context.Context, l *ent.Lock) error

		ListCommits(ctx context.Context, u *ent.User, r *ent.Repo, branch string, opt *i.ListOptions) ([]*extent.Commit, error)
		CompareCommits(ctx context.Context, u *ent.User, r *ent.Repo, base, head string, opt *i.ListOptions) ([]*extent.Commit, []*extent.CommitFile, error)
		GetCommit(ctx context.Context, u *ent.User, r *ent.Repo, sha string) (*extent.Commit, error)
		ListCommitStatuses(ctx context.Context, u *ent.User, r *ent.Repo, sha string) ([]*extent.Status, error)

		ListBranches(ctx context.Context, u *ent.User, r *ent.Repo, opt *i.ListOptions) ([]*extent.Branch, error)
		GetBranch(ctx context.Context, u *ent.User, r *ent.Repo, branch string) (*extent.Branch, error)
		GetDefaultBranch(ctx context.Context, u *ent.User, r *ent.Repo) (*extent.Branch, error)

		ListTags(ctx context.Context, u *ent.User, r *ent.Repo, opt *i.ListOptions) ([]*extent.Tag, error)
		GetTag(ctx context.Context, u *ent.User, r *ent.Repo, tag string) (*extent.Tag, error)
	}
)
