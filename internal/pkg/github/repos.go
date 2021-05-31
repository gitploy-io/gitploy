package github

import (
	"context"

	"github.com/google/go-github/v32/github"
	graphql "github.com/shurcooL/githubv4"

	"github.com/hanjunlee/gitploy/ent"
	reposv1 "github.com/hanjunlee/gitploy/internal/server/v1/repos"
	"github.com/hanjunlee/gitploy/vo"
)

type (
	// Node is tag of Github GraphQL.
	Node struct {
		Name   string
		Target struct {
			Oid       string
			CommitUrl string
		}
	}
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

func (g *Github) ListBranches(ctx context.Context, u *ent.User, r *ent.Repo, page, perPage int) ([]*vo.Branch, error) {
	bs, _, err := g.Client(ctx, u.Token).
		Repositories.
		ListBranches(ctx, r.Namespace, r.Name, &github.BranchListOptions{
			ListOptions: github.ListOptions{
				Page:    page,
				PerPage: perPage,
			},
		})
	if err != nil {
		return nil, err
	}

	branches := []*vo.Branch{}
	for _, b := range bs {
		branches = append(branches, mapGithubBranchToBranch(b))
	}

	return branches, nil
}

func (g *Github) GetBranch(ctx context.Context, u *ent.User, r *ent.Repo, branch string) (*vo.Branch, error) {
	b, _, err := g.Client(ctx, u.Token).
		Repositories.
		GetBranch(ctx, r.Namespace, r.Name, branch)
	if err != nil {
		return nil, err
	}

	return mapGithubBranchToBranch(b), nil
}

// ListTags list up tags as ordered by commit date.
// Github GraphQL explore - https://docs.github.com/en/graphql/overview/explorer
func (g *Github) ListTags(ctx context.Context, u *ent.User, r *ent.Repo, page, perPage int) ([]*vo.Tag, error) {
	var q struct {
		Repository struct {
			Refs struct {
				Nodes    []Node
				PageInfo struct {
					EndCursor   graphql.String
					HasNextPage bool
				}
				TotalCount int
			} `graphql:"refs(refPrefix: \"refs/tags/\", orderBy: {field: TAG_COMMIT_DATE, direction: DESC}, after: $cursor, first: $perPage)"`
		} `graphql:"repository(owner: $namespace, name: $name)"`
	}

	client := g.GraphQLClient(ctx, u.Token)
	v := map[string]interface{}{
		"namespace": graphql.String(r.Namespace),
		"name":      graphql.String(r.Name),
		"perPage":   graphql.Int(perPage),
		"cursor":    (*graphql.String)(nil),
	}

	curPage := 0
	for {
		curPage = curPage + 1
		if err := client.Query(ctx, &q, v); err != nil {
			return nil, err
		}

		if curPage == page || !q.Repository.Refs.PageInfo.HasNextPage {
			break
		}

		v["cursor"] = graphql.NewString(q.Repository.Refs.PageInfo.EndCursor)
	}

	tags := []*vo.Tag{}
	for _, n := range q.Repository.Refs.Nodes {
		tags = append(tags, &vo.Tag{
			Name:      n.Name,
			CommitSha: n.Target.Oid,
		})
	}

	return tags, nil
}

func (g *Github) GetTag(ctx context.Context, u *ent.User, r *ent.Repo, tag string) (*vo.Tag, error) {
	var q struct {
		Repository struct {
			Refs struct {
				Nodes      []Node
				TotalCount int
			} `graphql:"refs(refPrefix: \"refs/tags/\", orderBy: {field: TAG_COMMIT_DATE, direction: DESC}, first: 1, query: $tag)"`
		} `graphql:"repository(owner: $namespace, name: $name)"`
	}

	client := g.GraphQLClient(ctx, u.Token)
	v := map[string]interface{}{
		"namespace": graphql.String(r.Namespace),
		"name":      graphql.String(r.Name),
		"tag":       graphql.String(tag),
	}
	if err := client.Query(ctx, &q, v); err != nil {
		return nil, err
	}

	if q.Repository.Refs.TotalCount == 0 {
		return nil, &reposv1.RefNotFoundError{
			Ref: tag,
		}
	}

	n := q.Repository.Refs.Nodes[0]
	t := &vo.Tag{
		Name:      n.Name,
		CommitSha: n.Target.Oid,
	}

	return t, nil
}
