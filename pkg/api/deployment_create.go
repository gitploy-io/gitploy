package api

import (
	"context"
	"fmt"

	"github.com/gitploy-io/gitploy/model/ent"
)

type DeploymentCreateRequest struct {
	Type string `json:"type"`
	Ref  string `json:"ref"`
	Env  string `json:"env"`
}

// Create requests a server to deploy a specific ref(branch, SHA, tag).
func (s *DeploymentsService) Create(ctx context.Context, namespace, name string, body DeploymentCreateRequest) (*ent.Deployment, error) {
	req, err := s.client.NewRequest(
		"POST",
		fmt.Sprintf("api/v1/repos/%s/%s/deployments", namespace, name),
		body,
	)
	if err != nil {
		return nil, err
	}

	var d *ent.Deployment
	err = s.client.Do(ctx, req, &d)
	if err != nil {
		return nil, err
	}

	return d, nil
}
