package github

import (
	"context"

	"github.com/google/go-github/v32/github"
	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/vo"
)

func (g *Github) GetRateLimit(ctx context.Context, u *ent.User) (*vo.RateLimit, error) {
	rl, _, err := g.Client(ctx, u.Token).
		RateLimits(ctx)
	if err != nil {
		return nil, err
	}

	return mapGithubRateLimitToRateLimit(rl), nil
}

func mapGithubRateLimitToRateLimit(gr *github.RateLimits) *vo.RateLimit {
	return &vo.RateLimit{
		Limit:     gr.Core.Limit,
		Remaining: gr.Core.Remaining,
		Reset:     gr.Core.Reset.Time,
	}
}
