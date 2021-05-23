package github

import (
	"strconv"

	"github.com/google/go-github/v32/github"
	"github.com/hanjunlee/gitploy/ent"
)

func mapGithubUserToUser(u *github.User) *ent.User {
	return &ent.User{
		ID:     strconv.FormatInt(*u.ID, 10),
		Login:  *u.Login,
		Avatar: *u.AvatarURL,
	}
}
