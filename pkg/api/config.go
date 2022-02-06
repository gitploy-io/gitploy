package api

import (
	"context"
	"fmt"

	"github.com/gitploy-io/gitploy/model/extent"
)

type (
	ConfigService service
)

func (s *ConfigService) Get(ctx context.Context, namespace, name string) (*extent.Config, error) {
	req, err := s.client.NewRequest(
		"GET",
		fmt.Sprintf("api/v1/repos/%s/%s/config", namespace, name),
		nil,
	)
	if err != nil {
		return nil, err
	}

	var config *extent.Config
	if err := s.client.Do(ctx, req, &config); err != nil {
		return nil, err
	}

	return config, nil
}
