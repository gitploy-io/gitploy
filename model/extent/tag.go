package extent

type (
	Tag struct {
		Name      string `json:"name,omitempty"`
		CommitSHA string `json:"commit_sha,omitempty"`
	}
)
