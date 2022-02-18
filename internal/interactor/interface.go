//go:generate sh _mock.sh

package interactor

import (
	"context"
	"time"

	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/extent"
)

type (
	Store interface {
		UserStore
		ChatUserStore
		RepoStore
		PermStore
		DeploymentStore
		DeploymentStatusStore
		DeploymentStatisticsStore
		LockStore
		ReviewStore
		EventStore
	}

	// PermStore defines operations for working with perms.
	PermStore interface {
		ListPermsOfRepo(ctx context.Context, r *ent.Repo, opt *ListPermsOfRepoOptions) ([]*ent.Perm, error)
		FindPermOfRepo(ctx context.Context, r *ent.Repo, u *ent.User) (*ent.Perm, error)
		CreatePerm(ctx context.Context, p *ent.Perm) (*ent.Perm, error)
		UpdatePerm(ctx context.Context, p *ent.Perm) (*ent.Perm, error)
		DeletePermsOfUserLessThanSyncedAt(ctx context.Context, u *ent.User, t time.Time) (int, error)
	}

	ListPermsOfRepoOptions struct {
		ListOptions

		// Query search the 'login' contains the query.
		Query string
	}

	// PermStore defines operations for working with deployment_statuses.
	DeploymentStatusStore interface {
		ListDeploymentStatuses(ctx context.Context, d *ent.Deployment) ([]*ent.DeploymentStatus, error)
		CreateDeploymentStatus(ctx context.Context, s *ent.DeploymentStatus) (*ent.DeploymentStatus, error)
		SyncDeploymentStatus(ctx context.Context, ds *ent.DeploymentStatus) (*ent.DeploymentStatus, error)
	}

	SCM interface {
		UserSCM
		RepoSCM
		DeploymentSCM
		DeploymentStatusSCM
		ConfigSCM
		CommitSCM
		BranchSCM
		TagSCM

		GetRateLimit(ctx context.Context, u *ent.User) (*extent.RateLimit, error)
	}

	// DeploymentStatusSCM defines operations for working with remote deployment status.
	DeploymentStatusSCM interface {
		CreateRemoteDeploymentStatus(ctx context.Context, u *ent.User, r *ent.Repo, d *ent.Deployment, ds *extent.RemoteDeploymentStatus) (*extent.RemoteDeploymentStatus, error)
	}

	// CommitSCM defines operations for working with commit.
	CommitSCM interface {
		ListCommits(ctx context.Context, u *ent.User, r *ent.Repo, branch string, opt *ListOptions) ([]*extent.Commit, error)
		CompareCommits(ctx context.Context, u *ent.User, r *ent.Repo, base, head string, opt *ListOptions) ([]*extent.Commit, []*extent.CommitFile, error)
		GetCommit(ctx context.Context, u *ent.User, r *ent.Repo, sha string) (*extent.Commit, error)
		ListCommitStatuses(ctx context.Context, u *ent.User, r *ent.Repo, sha string) ([]*extent.Status, error)
	}

	BranchSCM interface {
		ListBranches(ctx context.Context, u *ent.User, r *ent.Repo, opt *ListOptions) ([]*extent.Branch, error)
		GetBranch(ctx context.Context, u *ent.User, r *ent.Repo, branch string) (*extent.Branch, error)
	}

	TagSCM interface {
		ListTags(ctx context.Context, u *ent.User, r *ent.Repo, opt *ListOptions) ([]*extent.Tag, error)
		GetTag(ctx context.Context, u *ent.User, r *ent.Repo, tag string) (*extent.Tag, error)
	}
)
