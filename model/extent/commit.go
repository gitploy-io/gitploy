package extent

import "time"

type (
	Commit struct {
		SHA           string  `json:"sha"`
		Message       string  `json:"message"`
		IsPullRequest bool    `json:"is_pull_request"`
		HTMLURL       string  `json:"html_url"`
		Author        *Author `json:"author,omitempty"`
	}

	Author struct {
		Login     string    `json:"login"`
		AvatarURL string    `json:"avatar_url"`
		Date      time.Time `json:"date"`
	}

	StatusState string

	Status struct {
		Context   string      `json:"context"`
		AvatarURL string      `json:"avatar_url"`
		TargetURL string      `json:"target_url"`
		State     StatusState `json:"state"`
	}

	CommitFile struct {
		FileName  string `json:"filename"`
		Additions int    `json:"addtitions"`
		Deletions int    `json:"deletions"`
		Changes   int    `json:"changes"`
	}
)

const (
	StatusStateSuccess   StatusState = "success"
	StatusStateFailure   StatusState = "failure"
	StatusStatePending   StatusState = "pending"
	StatusStateCancelled StatusState = "cancelled"
	StatusStateSkipped   StatusState = "skipped"
)
