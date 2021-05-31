package vo

type (
	Tag struct {
		Name      string `json:"name,omitempty"`
		CommitSha string `json:"commit_sha,omitempty"`
	}
)
