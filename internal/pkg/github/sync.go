package github

import (
	"context"

	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/extent"
	"github.com/gitploy-io/gitploy/pkg/e"
	"github.com/google/go-github/v32/github"
)

func (g *Github) ListRemoteRepos(ctx context.Context, u *ent.User) ([]*extent.RemoteRepo, error) {
	grs, err := g.listRemoteRepos(ctx, u)
	if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
	}

	remotes := make([]*extent.RemoteRepo, 0)
	for _, r := range grs {
		remotes = append(remotes, mapGithubRepoToRemotePerm(r))
	}

	return remotes, nil
}

func (g *Github) listRemoteRepos(ctx context.Context, u *ent.User) ([]*github.Repository, error) {
	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}

	all := make([]*github.Repository, 0)
	for {
		remotes, res, err := g.Client(ctx, u.Token).
			Repositories.
			List(ctx, "", opt)
		if err != nil {
			return nil, err
		}

		all = append(all, remotes...)
		if res.NextPage == 0 {
			break
		}

		opt.Page = res.NextPage
	}

	return all, nil
}
