package ent

// CheckEagerLoading checks the approval has required edges; Deployment and User.
func (a *Approval) CheckEagerLoading() error {
	if a.Edges.Deployment == nil {
		return &EagerLoadingError{
			Edge: "deployment",
		}
	}

	if a.Edges.User == nil {
		return &EagerLoadingError{
			Edge: "user",
		}
	}

	return nil
}
