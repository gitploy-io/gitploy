package vo

type (
	Commit struct {
		SHA           string `json:"sha"`
		Message       string `json:"message"`
		IsPullRequest bool   `json:"is_pull_request"`
	}

	StatusState string

	Status struct {
		Context   string      `json:"context"`
		AvatarURL string      `json:"avatar_url"`
		TargetURL string      `json:"target_url"`
		State     StatusState `json:"state"`
	}
)

const (
	StatusStateSuccess StatusState = "success"
	StatusStateFailure StatusState = "failure"
	StatusStatePending StatusState = "pending"
)
