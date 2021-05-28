package github

import (
	"strconv"
	"strings"

	"github.com/google/go-github/v32/github"
	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/perm"
)

func mapGithubUserToUser(u *github.User) *ent.User {
	return &ent.User{
		ID:     strconv.FormatInt(*u.ID, 10),
		Login:  *u.Login,
		Avatar: *u.AvatarURL,
	}
}

func mapGithubRepoToRepo(r *github.Repository) *ent.Repo {
	namespace, name := splitNamespaceName(*r.FullName)

	if r.Description == nil {
		r.Description = github.String("")
	}

	return &ent.Repo{
		ID:          strconv.FormatInt(*r.ID, 10),
		Namespace:   namespace,
		Name:        name,
		Description: *r.Description,
	}
}

func splitNamespaceName(fullName string) (string, string) {
	ss := strings.Split(fullName, "/")
	return ss[0], ss[1]
}

func mapGithubPermToPerm(perms map[string]bool) *ent.Perm {
	var p perm.RepoPerm

	if admin, ok := perms["admin"]; ok && admin {
		p = perm.RepoPermAdmin
	} else if push, ok := perms["push"]; ok && push {
		p = perm.RepoPermWrite
	} else if pull, ok := perms["pull"]; ok && pull {
		p = perm.RepoPermRead
	} else {
		p = perm.RepoPermRead
	}

	return &ent.Perm{
		RepoPerm: p,
	}
}
