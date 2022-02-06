package api

import (
	"context"
	"fmt"

	"github.com/gitploy-io/gitploy/model/ent"
)

type RepoUpdateRequest struct {
	ConfigPath *string `json:"config_path,omitempty"`
	Active     *bool   `json:"active,omitempty"`
}

func (s *ReposService) Update(ctx context.Context, namespace, name string, options RepoUpdateRequest) (*ent.Repo, error) {
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
