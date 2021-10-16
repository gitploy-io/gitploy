package github

import (
	"strings"

	"github.com/gitploy-io/gitploy/vo"
	"github.com/google/go-github/v32/github"
)

func mapGithubUserToUser(u *github.User) *vo.RemoteUser {
	return &vo.RemoteUser{
		ID:        *u.ID,
		Login:     *u.Login,
		AvatarURL: *u.AvatarURL,
	}
}

func splitNamespaceName(fullName string) (string, string) {
	ss := strings.Split(fullName, "/")
	return ss[0], ss[1]
}

func mapGithubRepoToRemotePerm(r *github.Repository) *vo.RemoteRepo {
	namespace, name := splitNamespaceName(*r.FullName)

	if r.Description == nil {
		r.Description = github.String("")
	}

	return &vo.RemoteRepo{
		ID:          *r.ID,
		Namespace:   namespace,
		Name:        name,
		Description: *r.Description,
		Perm:        mapGithubPermToRepoPerm(*r.Permissions),
	}
}

func mapGithubPermToRepoPerm(perms map[string]bool) vo.RemoteRepoPerm {
	var p vo.RemoteRepoPerm

	// Github represent the permission of the repository with these enums:
	// "admin", "push", and "pull".
	// https://docs.github.com/en/rest/reference/repos#list-repositories-for-the-authenticated-user
	if admin, ok := perms["admin"]; ok && admin {
		p = vo.RemoteRepoPermAdmin
	} else if push, ok := perms["push"]; ok && push {
		p = vo.RemoteRepoPermWrite
	} else if pull, ok := perms["pull"]; ok && pull {
		p = vo.RemoteRepoPermRead
	} else {
		p = vo.RemoteRepoPermRead
	}

	return p
}

func mapGithubCommitToCommit(cm *github.RepositoryCommit) *vo.Commit {
	var author *vo.Author
	if cm.Author != nil && cm.Commit.Author != nil {
		author = &vo.Author{
			Login:     *cm.Author.Login,
			AvatarURL: *cm.Author.AvatarURL,
			Date:      *cm.Commit.Author.Date,
		}
	}

	return &vo.Commit{
		SHA:     *cm.SHA,
		Message: *cm.Commit.Message,
		HTMLURL: *cm.HTMLURL,
		Author:  author,
	}
}

func mapGithubCommitFileToCommitFile(cf *github.CommitFile) *vo.CommitFile {
	return &vo.CommitFile{
		FileName:  *cf.Filename,
		Additions: *cf.Additions,
		Deletions: *cf.Deletions,
		Changes:   *cf.Changes,
	}
}

func mapGithubStatusToStatus(s *github.RepoStatus) *vo.Status {
	var (
		state vo.StatusState
	)
	switch *s.State {
	case "pending":
		state = vo.StatusStatePending
	case "failure":
		state = vo.StatusStateFailure
	case "error":
		state = vo.StatusStateFailure
	case "success":
		state = vo.StatusStateSuccess
	default:
		state = vo.StatusStatePending
	}

	return &vo.Status{
		Context: *s.Context,
		// TODO: fix
		AvatarURL: "",
		TargetURL: *s.TargetURL,
		State:     state,
	}
}

func mapGithubCheckRunToStatus(c *github.CheckRun) *vo.Status {
	var (
		state vo.StatusState
	)

	// Conclusion exist when the status is 'completed', only.
	if c.Conclusion == nil {
		return &vo.Status{
			Context:   *c.Name,
			AvatarURL: *c.App.Owner.AvatarURL,
			TargetURL: *c.HTMLURL,
			State:     vo.StatusStatePending,
		}
	}

	if *c.Conclusion == "success" {
		state = vo.StatusStateSuccess
	} else if *c.Conclusion == "failure" {
		state = vo.StatusStateFailure
	} else {
		state = vo.StatusStatePending
	}

	return &vo.Status{
		Context:   *c.Name,
		AvatarURL: *c.App.Owner.AvatarURL,
		TargetURL: *c.HTMLURL,
		State:     state,
	}
}

func mapGithubBranchToBranch(b *github.Branch) *vo.Branch {
	return &vo.Branch{
		Name:      *b.Name,
		CommitSHA: *b.Commit.SHA,
	}
}

func mapInsecureSSL(ssl bool) int {
	if ssl {
		return 0
	}
	return 1
}
