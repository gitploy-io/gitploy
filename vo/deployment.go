package vo

type (
	RemoteDeployment struct {
		UID     int64  `json:"uid"`
		SHA     string `json:"sha"`
		HTLMURL string `json:"html_url"`
	}
)
