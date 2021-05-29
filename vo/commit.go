package vo

type (
	Commit struct {
		Sha           string `json:"sha,omitempty"`
		Message       string `json:"message,omitempty"`
		IsPullRequest bool   `json:"is_pull_request"`
	}

	StatusState string

	Status struct {
		Context   string      `json:"context"`
		AvatarURL string      `json:"avatar_url,omitempty"`
		TargetURL string      `json:"target_url,omitempty"`
		State     StatusState `json:"state"`
	}
)

const (
	StatusStateSuccess StatusState = "success"
	StatusStateFailure StatusState = "failure"
	StatusStatePending StatusState = "pending"
)
