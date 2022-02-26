package github

import (
	"strings"

	"github.com/gitploy-io/gitploy/model/extent"
	"github.com/google/go-github/v42/github"
)

func mapGithubUserToUser(u *github.User) *extent.RemoteUser {
	return &extent.RemoteUser{
		ID:        *u.ID,
		Login:     *u.Login,
		AvatarURL: *u.AvatarURL,
	}
}

func splitNamespaceName(fullName string) (string, string) {
	ss := strings.Split(fullName, "/")
	return ss[0], ss[1]
}

func mapGithubRepoToRemotePerm(r *github.Repository) *extent.RemoteRepo {
	namespace, name := splitNamespaceName(*r.FullName)

	if r.Description == nil {
		r.Description = github.String("")
	}

	return &extent.RemoteRepo{
		ID:          *r.ID,
		Namespace:   namespace,
		Name:        name,
		Description: *r.Description,
		Perm:        mapGithubPermToRepoPerm(r.Permissions),
	}
}

func mapGithubPermToRepoPerm(perms map[string]bool) extent.RemoteRepoPerm {
	var p extent.RemoteRepoPerm

	// Github represent the permission of the repository with these enums:
	// "admin", "push", and "pull".
	// https://docs.github.com/en/rest/reference/repos#list-repositories-for-the-authenticated-user
	if admin, ok := perms["admin"]; ok && admin {
		p = extent.RemoteRepoPermAdmin
	} else if push, ok := perms["push"]; ok && push {
		p = extent.RemoteRepoPermWrite
	} else if pull, ok := perms["pull"]; ok && pull {
		p = extent.RemoteRepoPermRead
	} else {
		p = extent.RemoteRepoPermRead
	}

	return p
}

func mapGithubCommitToCommit(cm *github.RepositoryCommit) *extent.Commit {
	var author *extent.Author
	if cm.Author != nil && cm.Commit.Author != nil {
		author = &extent.Author{
			Login:     *cm.Author.Login,
			AvatarURL: *cm.Author.AvatarURL,
			Date:      *cm.Commit.Author.Date,
		}
	}

	return &extent.Commit{
		SHA:     *cm.SHA,
		Message: *cm.Commit.Message,
		HTMLURL: *cm.HTMLURL,
		Author:  author,
	}
}

func mapGithubCommitFileToCommitFile(cf *github.CommitFile) *extent.CommitFile {
	return &extent.CommitFile{
		FileName:  *cf.Filename,
		Additions: *cf.Additions,
		Deletions: *cf.Deletions,
		Changes:   *cf.Changes,
	}
}

func mapGithubStatusToStatus(s *github.RepoStatus) *extent.Status {
	var (
		state extent.StatusState
	)
	switch *s.State {
	case "pending":
		state = extent.StatusStatePending
	case "failure":
		state = extent.StatusStateFailure
	case "error":
		state = extent.StatusStateFailure
	case "success":
		state = extent.StatusStateSuccess
	default:
		state = extent.StatusStatePending
	}

	status := &extent.Status{
		Context:   *s.Context,
		AvatarURL: "",
		State:     state,
	}
	if s.TargetURL != nil {
		status.TargetURL = *s.TargetURL
	}

	return status
}

func mapGithubCheckRunToStatus(c *github.CheckRun) *extent.Status {
	var (
		state extent.StatusState
	)

	// Conclusion exist when the status is 'completed', only.
	if c.Conclusion == nil {
		state = extent.StatusStatePending
	} else if *c.Conclusion == "success" {
		state = extent.StatusStateSuccess
	} else if *c.Conclusion == "failure" {
		state = extent.StatusStateFailure
	} else if *c.Conclusion == "cancelled" {
		state = extent.StatusStateCancelled
	} else if *c.Conclusion == "skipped" {
		state = extent.StatusStateSkipped
	} else {
		state = extent.StatusStatePending
	}

	return &extent.Status{
		Context:   *c.Name,
		AvatarURL: *c.App.Owner.AvatarURL,
		TargetURL: *c.HTMLURL,
		State:     state,
	}
}

func mapGithubBranchToBranch(b *github.Branch) *extent.Branch {
	return &extent.Branch{
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
