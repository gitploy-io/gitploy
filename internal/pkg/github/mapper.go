package github

import (
	"strconv"
	"strings"

	"github.com/google/go-github/v32/github"
	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/perm"
	"github.com/hanjunlee/gitploy/vo"
)

func mapGithubUserToUser(u *github.User) *vo.RemoteUser {
	return &vo.RemoteUser{
		ID:        strconv.FormatInt(*u.ID, 10),
		Login:     *u.Login,
		AvatarURL: *u.AvatarURL,
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

func mapGithubCommitToCommit(cm *github.RepositoryCommit) *vo.Commit {
	isPullRequest := false
	if cm.Commit.Author != nil && cm.Commit.Committer != nil {
		if *cm.Commit.Author.Name != *cm.Commit.Committer.Name {
			isPullRequest = true
		}
	}

	return &vo.Commit{
		SHA:           *cm.SHA,
		Message:       *cm.Commit.Message,
		IsPullRequest: isPullRequest,
		HTMLURL:       *cm.HTMLURL,
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

	switch *c.Conclusion {
	case "failure":
		state = vo.StatusStateFailure
	case "cancelled":
		state = vo.StatusStateFailure
	case "timed_out":
		state = vo.StatusStateFailure
	case "success":
		state = vo.StatusStateSuccess
	default:
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
