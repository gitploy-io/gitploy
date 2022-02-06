package api

import (
	"context"
	"fmt"

	"github.com/gitploy-io/gitploy/model/ent"
)

// Update requests to trigger the 'waiting' deployment.
func (s *DeploymentsService) Update(ctx context.Context, namespace, name string, number int) (*ent.Deployment, error) {
	req, err := s.client.NewRequest(
		"PUT",
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
