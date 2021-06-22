package stream

import "github.com/hanjunlee/gitploy/ent"

type (
	Interactor interface {
		Subscribe(func(*ent.User, *ent.Notification)) error
		Unsubscribe(func(*ent.User, *ent.Notification)) error
	}
)
