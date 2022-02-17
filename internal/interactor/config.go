package interactor

import (
	"context"

	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/extent"
)

type (
	ConfigInteractor service

	ConfigSCM interface {
		GetConfig(ctx context.Context, u *ent.User, r *ent.Repo) (*extent.Config, error)
		GetConfigRedirectURL(ctx context.Context, u *ent.User, r *ent.Repo) (string, error)
		GetNewConfigRedirectURL(ctx context.Context, u *ent.User, r *ent.Repo) (string, error)
	}
)

// GetEvaluatedConfig returns the config after evaluating the variables.
func (i *ConfigInteractor) GetEvaluatedConfig(ctx context.Context, u *ent.User, r *ent.Repo, v *extent.EvalValues) (*extent.Config, error) {
	config, err := i.scm.GetConfig(ctx, u, r)
	if err != nil {
		return nil, err
	}

	if err = config.Eval(v); err != nil {
		return nil, err
	}

	return config, nil
}
