package ent

import (
	"github.com/gitploy-io/gitploy/model/ent/deployment"
)

func (d *Deployment) GetShortRef() string {
	const maxlen = 7

	if d.Type == deployment.TypeCommit &&
		len(d.Ref) > maxlen {
		return d.Ref[:7]
	}

	return d.Ref
}
