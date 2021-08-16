package vo

type (
	Branch struct {
		Name      string `json:"name"`
		CommitSHA string `json:"commit_sha"`
	}
)
