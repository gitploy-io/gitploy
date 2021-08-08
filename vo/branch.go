package vo

type (
	Branch struct {
		Name      string `json:"name,omitempty"`
		CommitSHA string `json:"commit_sha,omitempty"`
	}
)
