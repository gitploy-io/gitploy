package extent

type (
	RemoteDeploymentStatus struct {
		ID          int64  `json:"id"`
		Status      string `json:"status"`
		Description string `json:"description"`
		LogURL      string `json:"log_url"`
	}
)
