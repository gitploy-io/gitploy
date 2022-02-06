package api

import (
	"context"
	"fmt"

	"github.com/gitploy-io/gitploy/model/ent"
)

// Get returns the deployment.
func (s *DeploymentsService) Get(ctx context.Context, namespace, name string, number int) (*ent.Deployment, error) {
	req, err := s.client.NewRequest(
		"GET",
		fmt.Sprintf("api/v1/repos/%s/%s/deployments/%d", namespace, name, number),
		nil,
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
