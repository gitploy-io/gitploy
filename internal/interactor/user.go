package interactor

import "github.com/hanjunlee/gitploy/ent"

func (i *Interactor) FindUser() (*ent.User, error) {
	return i.store.FindUser()
}
