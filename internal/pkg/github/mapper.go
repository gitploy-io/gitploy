package github

import (
	"github.com/google/go-github/v32/github"
	"github.com/hanjunlee/gitploy/ent"
)

func mapGithubUserToUser(u *github.User) *ent.User {
	return &ent.User{
		Login:  *u.Login,
		Avatar: *u.AvatarURL,
	}
}
