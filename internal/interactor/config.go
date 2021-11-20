package interactor

import (
	"context"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/pkg/e"
	"github.com/gitploy-io/gitploy/vo"
)

func (i *Interactor) GetEnv(ctx context.Context, u *ent.User, r *ent.Repo, env string) (*vo.Env, error) {
	config, err := i.SCM.GetConfig(ctx, u, r)
	if err != nil {
		return nil, err
	}

	if !config.HasEnv(env) {
		return nil, e.NewError(e.ErrorCodeConfigUndefinedEnv, nil)
	}

	return config.GetEnv(env), nil
}
