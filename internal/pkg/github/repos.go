package github

import (
	"context"

	"github.com/google/go-github/v32/github"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/vo"
)

func (g *Github) ListCommits(ctx context.Context, u *ent.User, r *ent.Repo, branch string, page, perPage int) ([]*vo.Commit, error) {
	cms, _, err := g.Client(ctx, u.Token).
		Repositories.
		ListCommits(ctx, r.Namespace, r.Name, &github.CommitsListOptions{
			SHA: branch,
			ListOptions: github.ListOptions{
				Page:    page,
				PerPage: perPage,
			},
		})
	if err != nil {
		return nil, err
	}

	ret := make([]*vo.Commit, 0)
	for _, cm := range cms {
		ret = append(ret, mapGithubCommitToCommit(cm))
	}

	return ret, nil
}

func (g *Github) GetCommit(ctx context.Context, u *ent.User, r *ent.Repo, sha string) (*vo.Commit, error) {
	cm, _, err := g.Client(ctx, u.Token).
		Repositories.
		GetCommit(ctx, r.Namespace, r.Name, sha)
	if err != nil {
		return nil, err
	}

	return mapGithubCommitToCommit(cm), nil
}

func (g *Github) ListCommitStatuses(ctx context.Context, u *ent.User, r *ent.Repo, sha string) ([]*vo.Status, error) {
	ss := []*vo.Status{}

	client := g.Client(ctx, u.Token)

	// Repo status
	cs, _, err := client.Repositories.GetCombinedStatus(ctx, r.Namespace, r.Name, sha, &github.ListOptions{
		PerPage: 100,
	})
	if err != nil {
		return nil, err
	}

	for _, rs := range cs.Statuses {
		ss = append(ss, mapGithubStatusToStatus(rs))
	}

	// Check-run
	result, _, err := client.Checks.ListCheckRunsForRef(ctx, r.Namespace, r.Name, sha, &github.ListCheckRunsOptions{
		ListOptions: github.ListOptions{
			PerPage: 100,
		},
	})
	if err != nil {
		return nil, err
	}

	for _, c := range result.CheckRuns {
		if c.Conclusion == nil {
			continue
		}

		ss = append(ss, mapGithubCheckRunToStatus(c))
	}

	return ss, nil
}
