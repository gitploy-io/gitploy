package github

import (
	"context"
	"net/http"

	"github.com/google/go-github/v32/github"
	graphql "github.com/shurcooL/githubv4"

	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/extent"
	"github.com/gitploy-io/gitploy/pkg/e"
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

func (g *Github) ListCommits(ctx context.Context, u *ent.User, r *ent.Repo, branch string, page, perPage int) ([]*extent.Commit, error) {
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

	ret := make([]*extent.Commit, 0)
	for _, cm := range cms {
		ret = append(ret, mapGithubCommitToCommit(cm))
	}

	return ret, nil
}

func (g *Github) CompareCommits(ctx context.Context, u *ent.User, r *ent.Repo, base, head string, page, perPage int) ([]*extent.Commit, []*extent.CommitFile, error) {
	// TODO: Support pagination.
	res, _, err := g.Client(ctx, u.Token).
		Repositories.
		CompareCommits(ctx, r.Namespace, r.Name, base, head)
	if err != nil {
		return nil, nil, err
	}

	cms := make([]*extent.Commit, 0)
	for _, cm := range res.Commits {
		cms = append(cms, mapGithubCommitToCommit(cm))
	}

	cfs := make([]*extent.CommitFile, 0)
	for _, cf := range res.Files {
		cfs = append(cfs, mapGithubCommitFileToCommitFile(cf))
	}

	return cms, cfs, nil
}

func (g *Github) GetCommit(ctx context.Context, u *ent.User, r *ent.Repo, sha string) (*extent.Commit, error) {
	cm, res, err := g.Client(ctx, u.Token).
		Repositories.
		GetCommit(ctx, r.Namespace, r.Name, sha)
	// Github returns Unprocessable entity if the commit is not found.
	if res.StatusCode == http.StatusNotFound || res.StatusCode == http.StatusUnprocessableEntity {
		return nil, e.NewErrorWithMessage(e.ErrorCodeEntityNotFound, "The commit is not found.", err)
	} else if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
	}

	return mapGithubCommitToCommit(cm), nil
}

func (g *Github) ListCommitStatuses(ctx context.Context, u *ent.User, r *ent.Repo, sha string) ([]*extent.Status, error) {
	ss := []*extent.Status{}

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
	result, res, err := client.Checks.ListCheckRunsForRef(ctx, r.Namespace, r.Name, sha, &github.ListCheckRunsOptions{
		ListOptions: github.ListOptions{
			PerPage: 100,
		},
	})
	// check-runs secures the commit is exist.
	if res.StatusCode == http.StatusUnprocessableEntity {
		return nil, e.NewErrorWithMessage(e.ErrorCodeEntityNotFound, "The commit is not found.", err)
	} else if err != nil {
		return nil, e.NewError(
			e.ErrorCodeInternalError,
			err,
		)
	}

	for _, c := range result.CheckRuns {
		ss = append(ss, mapGithubCheckRunToStatus(c))
	}

	return ss, nil
}

func (g *Github) ListBranches(ctx context.Context, u *ent.User, r *ent.Repo, page, perPage int) ([]*extent.Branch, error) {
	bs, _, err := g.Client(ctx, u.Token).
		Repositories.
		ListBranches(ctx, r.Namespace, r.Name, &github.BranchListOptions{
			ListOptions: github.ListOptions{
				Page:    page,
				PerPage: perPage,
			},
		})
	if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
	}

	branches := []*extent.Branch{}
	for _, b := range bs {
		branches = append(branches, mapGithubBranchToBranch(b))
	}

	return branches, nil
}

func (g *Github) GetBranch(ctx context.Context, u *ent.User, r *ent.Repo, branch string) (*extent.Branch, error) {
	b, res, err := g.Client(ctx, u.Token).
		Repositories.
		GetBranch(ctx, r.Namespace, r.Name, branch)
	if res.StatusCode == http.StatusNotFound {
		return nil, e.NewErrorWithMessage(e.ErrorCodeEntityNotFound, "The branch is not found.", err)
	} else if err != nil {
		return nil, e.NewError(e.ErrorCodeInternalError, err)
	}

	return mapGithubBranchToBranch(b), nil
}

// ListTags list up tags as ordered by commit date.
// Github GraphQL explore - https://docs.github.com/en/graphql/overview/explorer
func (g *Github) ListTags(ctx context.Context, u *ent.User, r *ent.Repo, page, perPage int) ([]*extent.Tag, error) {
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

	tags := []*extent.Tag{}
	for _, n := range q.Repository.Refs.Nodes {
		tags = append(tags, &extent.Tag{
			Name:      n.Name,
			CommitSHA: n.Target.Oid,
		})
	}

	return tags, nil
}

func (g *Github) GetTag(ctx context.Context, u *ent.User, r *ent.Repo, tag string) (*extent.Tag, error) {
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
		return nil, e.NewError(
			e.ErrorCodeInternalError,
			err,
		)
	}

	if q.Repository.Refs.TotalCount == 0 {
		return nil, e.NewErrorWithMessage(e.ErrorCodeEntityNotFound, "The tag is not found.", nil)
	}

	n := q.Repository.Refs.Nodes[0]
	t := &extent.Tag{
		Name:      n.Name,
		CommitSHA: n.Target.Oid,
	}

	return t, nil
}

func (g *Github) CreateWebhook(ctx context.Context, u *ent.User, r *ent.Repo, c *extent.WebhookConfig) (int64, error) {
	h, _, err := g.Client(ctx, u.Token).
		Repositories.
		CreateHook(ctx, r.Namespace, r.Name, &github.Hook{
			Config: map[string]interface{}{
				"url":          c.URL,
				"content_type": "json",
				"secret":       c.Secret,
				"insecure_ssl": mapInsecureSSL(c.InsecureSSL),
			},
			Events: []string{"deployment_status", "push"},
			Active: github.Bool(true),
		})
	if err != nil {
		return -1, e.NewErrorWithMessage(e.ErrorCodeInternalError, "It has failed to create a webhook.", err)
	}

	return *h.ID, nil
}

func (g *Github) DeleteWebhook(ctx context.Context, u *ent.User, r *ent.Repo, id int64) error {
	res, err := g.Client(ctx, u.Token).
		Repositories.
		DeleteHook(ctx, r.Namespace, r.Name, id)
	// https://docs.github.com/en/rest/reference/repos#delete-a-repository-webhook
	if res.StatusCode == http.StatusNotFound {
		return e.NewErrorWithMessage(e.ErrorCodeEntityNotFound, "The webhook is not found.", err)
	}

	return err
}
