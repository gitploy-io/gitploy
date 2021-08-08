package vo

type (
	Tag struct {
		Name      string `json:"name,omitempty"`
		CommitSHA string `json:"commit_sha,omitempty"`
	}
)
