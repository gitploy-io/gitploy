package api

import (
	"context"
	"fmt"

	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/extent"
)

type (
	DeploymentStatusService service

	DeploymentStatusCreateRemoteRequest struct {
		Status      string `json:"status"`
		Description string `json:"description"`
		LogURL      string `json:"log_url"`
	}
)

// List returns the list of deployment statuses.
// It returns an error for a bad request.
func (s *DeploymentStatusService) List(ctx context.Context, namespace, name string, number int, opt *ListOptions) ([]*ent.DeploymentStatus, error) {
	req, err := s.client.NewRequest(
		"GET",
		fmt.Sprintf("api/v1/repos/%s/%s/deployments/%d/statuses", namespace, name, number),
		nil,
	)
	if err != nil {
		return nil, err
	}

	var dss []*ent.DeploymentStatus
	if err := s.client.Do(ctx, req, &dss); err != nil {
		return nil, err
	}

	return dss, nil
}

// CreateRemote returns the remote status.
// It returns an error for a bad request.
func (s *DeploymentStatusService) CreateRemote(ctx context.Context, namespace, name string, number int, body *DeploymentStatusCreateRemoteRequest) (*extent.RemoteDeploymentStatus, error) {
	req, err := s.client.NewRequest(
		"POST",
		fmt.Sprintf("api/v1/repos/%s/%s/deployments/%d/remote-statuses", namespace, name, number),
		body,
	)
	if err != nil {
		return nil, err
	}

	var ds *extent.RemoteDeploymentStatus
	if err := s.client.Do(ctx, req, &ds); err != nil {
		return nil, err
	}

	return ds, nil
}
