package api

import (
	"context"
	"fmt"

	"github.com/gitploy-io/gitploy/model/ent"
)

type (
	RepoService service

	RepoListOptions struct {
		ListOptions
	}

	RepoUpdateRequest struct {
		ConfigPath *string `json:"config_path,omitempty"`
		Active     *bool   `json:"active,omitempty"`
	}
)

// ListAll returns all repositories.
func (s *RepoService) ListAll(ctx context.Context) ([]*ent.Repo, error) {
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
func (s *RepoService) List(ctx context.Context, options RepoListOptions) ([]*ent.Repo, error) {
	req, err := s.client.NewRequest(
		"GET",
		fmt.Sprintf("api/v1/repos?page=%d&per_page=%d", options.Page, options.PerPage),
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

// Get returns the repository.
func (s *RepoService) Get(ctx context.Context, namespace, name string) (*ent.Repo, error) {
	req, err := s.client.NewRequest(
		"GET",
		fmt.Sprintf("api/v1/repos/%s/%s", namespace, name),
		nil)
	if err != nil {
		return nil, err
	}

	var repo *ent.Repo
	if err := s.client.Do(ctx, req, &repo); err != nil {
		return nil, err
	}

	return repo, nil
}

func (s *RepoService) Update(ctx context.Context, namespace, name string, options RepoUpdateRequest) (*ent.Repo, error) {
	req, err := s.client.NewRequest(
		"PATCH",
		fmt.Sprintf("api/v1/repos/%s/%s", namespace, name),
		options,
	)
	if err != nil {
		return nil, err
	}

	var repo *ent.Repo
	if err := s.client.Do(ctx, req, &repo); err != nil {
		return nil, err
	}

	return repo, nil
}
