package github

import (
	"context"

	"github.com/google/go-github/v32/github"
	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/vo"
)

func (g *Github) GetAllPermsWithRepo(ctx context.Context, token string) ([]*ent.Perm, error) {
	remotes, err := g.getAllRepositories(ctx, token)
	if err != nil {
		return nil, err
	}

	perms := make([]*ent.Perm, 0)
	for _, remote := range remotes {
		perm := mapGithubPermToPerm(*remote.Permissions)
		local := mapGithubRepoToRepo(remote)
		perm.Edges.Repo = local
		perms = append(perms, perm)
	}

	return perms, nil
}

func (g *Github) getAllRepositories(ctx context.Context, token string) ([]*github.Repository, error) {
	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}

	all := make([]*github.Repository, 0)
	for {
		remotes, res, err := g.Client(ctx, token).Repositories.List(ctx, "", opt)
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

func (g *Github) ListRemoteRepos(ctx context.Context, u *ent.User) ([]*vo.RemoteRepo, error) {
	grs, err := g.listRemoteRepos(ctx, u)
	if err != nil {
		return nil, err
	}

	remotes := make([]*vo.RemoteRepo, 0)
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
