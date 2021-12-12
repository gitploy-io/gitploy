package github

import (
	"context"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/extent"
	"github.com/google/go-github/v32/github"
)

func (g *Github) GetRateLimit(ctx context.Context, u *ent.User) (*extent.RateLimit, error) {
	rl, _, err := g.Client(ctx, u.Token).
		RateLimits(ctx)
	if err != nil {
		return nil, err
	}

	return mapGithubRateLimitToRateLimit(rl), nil
}

func mapGithubRateLimitToRateLimit(gr *github.RateLimits) *extent.RateLimit {
	return &extent.RateLimit{
		Limit:     gr.Core.Limit,
		Remaining: gr.Core.Remaining,
		Reset:     gr.Core.Reset.Time,
	}
}
