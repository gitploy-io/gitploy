package ent

import "github.com/gitploy-io/gitploy/ent/event"

// CheckEagerLoading checks the event has required edges; Deployment or Approval.
func (e *Event) CheckEagerLoading() error {
	if e.Kind == event.KindDeployment {
		if e.Edges.Deployment == nil {
			return &EagerLoadingError{
				Edge: "deployment",
			}
		}
	}

	if e.Kind == event.KindApproval {
		if e.Edges.Approval == nil {
			return &EagerLoadingError{
				Edge: "approval",
			}
		}
	}

	return nil
}
