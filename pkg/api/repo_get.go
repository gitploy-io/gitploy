package api

import (
	"context"
	"fmt"

	"github.com/gitploy-io/gitploy/model/ent"
)

// Get returns the repository.
func (s *ReposService) Get(ctx context.Context, namespace, name string) (*ent.Repo, error) {
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
