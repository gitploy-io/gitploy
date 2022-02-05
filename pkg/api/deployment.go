package api

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/ent/deployment"
)

type (
	DeploymentService service

	DeploymentListOptions struct {
		ListOptions

		Env    string
		Status deployment.Status
	}
)

// List returns the deployment list.
// It returns an error for a bad request.
func (s *DeploymentService) List(ctx context.Context, namespace, name string, options DeploymentListOptions) ([]*ent.Deployment, error) {
	// Build the query.
	vals := url.Values{}

	vals.Add("page", strconv.Itoa(options.ListOptions.Page))
	vals.Add("per_page", strconv.Itoa(options.PerPage))

	if options.Env != "" {
		vals.Add("env", options.Env)
	}

	if options.Status != "" {
		if err := deployment.StatusValidator(options.Status); err != nil {
			return nil, err
		}

		vals.Add("status", string(options.Status))
	}

	// Request a server.
	req, err := s.client.NewRequest(
		"GET",
		fmt.Sprintf("api/v1/repos/%s/%s/deployments?%s", namespace, name, vals.Encode()),
		nil,
	)
	if err != nil {
		return nil, err
	}

	var ds []*ent.Deployment
	err = s.client.Do(ctx, req, &ds)
	if err != nil {
		return nil, err
	}

	return ds, nil
}
