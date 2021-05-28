package vo

type (
	Commit struct {
		Sha           string `json:"sha,omitempty"`
		Message       string `json:"message,omitempty"`
		IsPullRequest bool   `json:"is_pull_request"`
	}
)
