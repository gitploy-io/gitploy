package vo

type (
	Notification struct {
		Kind       NotificationKind `json:"kind"`
		Type       NotificationType `json:"type"`
		Repo       *RepoData        `json:"repo,omitempty"`
		Deployment *DeploymentData  `json:"deployment,omitempty"`
		Approval   *ApprovalData    `json:"approval,omitempty"`
	}

	NotificationKind string

	NotificationType string

	RepoData struct {
		Namespace string `json:"namespace"`
		Name      string `json:"name"`
	}

	DeploymentData struct {
		Number int    `json:"number"`
		Type   string `json:"type"`
		Ref    string `json:"ref"`
		Env    string `json:"env"`
		Status string `json:"status"`
		Login  string `json:"login"`
	}

	ApprovalData struct {
		Status string `json:"status"`
		Login  string `json:"login"`
	}
)

const (
	KindDeployment NotificationKind = "deployment"
	KindApproval   NotificationKind = "approval"
)

const (
	TypeCreated NotificationType = "created"
	TypeUpdated NotificationType = "updated"
)
