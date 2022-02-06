package api

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/gitploy-io/gitploy/model/ent"
)

type RepoListOptions struct {
	ListOptions
}

// ListAll returns all repositories.
func (s *ReposService) ListAll(ctx context.Context) ([]*ent.Repo, error) {
	// Max value for 'perPage'.
	const perPage = 100

	repos := make([]*ent.Repo, 0)
	page := 1

	for {
		rs, err := s.List(ctx, RepoListOptions{
			ListOptions{
				Page:    page,
				PerPage: perPage,
			},
		})
		if err != nil {
			return nil, err
		}

		repos = append(repos, rs...)

		// Met the end of pages.
		if len(rs) < perPage {
			break
		}

		page = page + 1
	}

	return repos, nil
}

// List returns repositories which are on the page.
func (s *ReposService) List(ctx context.Context, options RepoListOptions) ([]*ent.Repo, error) {
	// Build the query.
	vals := url.Values{}
	vals.Add("page", strconv.Itoa(options.Page))
	vals.Add("per_page", strconv.Itoa(options.PerPage))

	req, err := s.client.NewRequest(
		"GET",
		fmt.Sprintf("api/v1/repos?%s", vals.Encode()),
		nil)
	if err != nil {
		return nil, err
	}

	var repos []*ent.Repo
	if err := s.client.Do(ctx, req, &repos); err != nil {
		return nil, err
	}

	return repos, nil
}