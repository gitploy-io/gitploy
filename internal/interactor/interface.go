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

		FindDeploymentStatisticsOfRepoByEnv(ctx context.Context, r *ent.Repo, env string) (*ent.DeploymentStatistics, error)
		CreateDeploymentStatistics(ctx context.Context, s *ent.DeploymentStatistics) (*ent.DeploymentStatistics, error)
		UpdateDeploymentStatistics(ctx context.Context, s *ent.DeploymentStatistics) (*ent.DeploymentStatistics, error)

		ListAllDeploymentStatistics(ctx context.Context) ([]*ent.DeploymentStatistics, error)
		ListDeploymentStatisticsGreaterThanTime(ctx context.Context, updated time.Time) ([]*ent.DeploymentStatistics, error)

		SearchReviews(ctx context.Context, u *ent.User) ([]*ent.Review, error)
		ListReviews(ctx context.Context, d *ent.Deployment) ([]*ent.Review, error)
		FindReviewOfUser(ctx context.Context, u *ent.User, d *ent.Deployment) (*ent.Review, error)
		FindReviewByID(ctx context.Context, id int) (*ent.Review, error)
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

		CreateRemoteDeploymentStatus(ctx context.Context, u *ent.User, r *ent.Repo, d *ent.Deployment, ds *extent.RemoteDeploymentStatus) (*extent.RemoteDeploymentStatus, error)

		ListCommits(ctx context.Context, u *ent.User, r *ent.Repo, branch string, page, perPage int) ([]*extent.Commit, error)
		CompareCommits(ctx context.Context, u *ent.User, r *ent.Repo, base, head string, page, perPage int) ([]*extent.Commit, []*extent.CommitFile, error)
		GetCommit(ctx context.Context, u *ent.User, r *ent.Repo, sha string) (*extent.Commit, error)
		ListCommitStatuses(ctx context.Context, u *ent.User, r *ent.Repo, sha string) ([]*extent.Status, error)

		ListBranches(ctx context.Context, u *ent.User, r *ent.Repo, page, perPage int) ([]*extent.Branch, error)
		GetBranch(ctx context.Context, u *ent.User, r *ent.Repo, branch string) (*extent.Branch, error)

		ListTags(ctx context.Context, u *ent.User, r *ent.Repo, page, perPage int) ([]*extent.Tag, error)
		GetTag(ctx context.Context, u *ent.User, r *ent.Repo, tag string) (*extent.Tag, error)

		GetRateLimit(ctx context.Context, u *ent.User) (*extent.RateLimit, error)
	}
)
