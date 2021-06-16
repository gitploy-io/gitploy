package repos

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/vo"
)

type (
	Interactor interface {
		FindPermByRepoID(ctx context.Context, u *ent.User, repoID string) (*ent.Perm, error)

		FindRepoByID(ctx context.Context, u *ent.User, id string) (*ent.Repo, error)
		ListRepos(ctx context.Context, u *ent.User, sorted bool, q string, page, perPage int) ([]*ent.Repo, error)
		FindRepoByNamespaceName(ctx context.Context, u *ent.User, namespace, name string) (*ent.Repo, error)
		PatchRepo(ctx context.Context, r *ent.Repo, p *RepoPayload) (*ent.Repo, error)
		ActivateRepo(ctx context.Context, u *ent.User, r *ent.Repo, c *vo.WebhookConfig) (*ent.Repo, error)
		DeactivateRepo(ctx context.Context, u *ent.User, r *ent.Repo) (*ent.Repo, error)

		ListDeployments(ctx context.Context, r *ent.Repo, env string, status string, page, perPage int) ([]*ent.Deployment, error)
		FindLatestDeployment(ctx context.Context, r *ent.Repo, env string) (*ent.Deployment, error)
		Deploy(ctx context.Context, u *ent.User, re *ent.Repo, d *ent.Deployment) (*ent.Deployment, error)
		GetConfig(ctx context.Context, u *ent.User, r *ent.Repo) (*vo.Config, error)

		ListCommits(ctx context.Context, u *ent.User, r *ent.Repo, branch string, page, perPage int) ([]*vo.Commit, error)
		GetCommit(ctx context.Context, u *ent.User, r *ent.Repo, sha string) (*vo.Commit, error)
		ListCommitStatuses(ctx context.Context, u *ent.User, r *ent.Repo, sha string) ([]*vo.Status, error)

		ListBranches(ctx context.Context, u *ent.User, r *ent.Repo, page, perPage int) ([]*vo.Branch, error)
		GetBranch(ctx context.Context, u *ent.User, r *ent.Repo, branch string) (*vo.Branch, error)

		ListTags(ctx context.Context, u *ent.User, r *ent.Repo, page, perPage int) ([]*vo.Tag, error)
		GetTag(ctx context.Context, u *ent.User, r *ent.Repo, tag string) (*vo.Tag, error)
	}
)
