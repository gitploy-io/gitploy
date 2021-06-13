package interactor

import "go.uber.org/zap"

type (
	Interactor struct {
		store Store
		scm   SCM
		log   *zap.Logger
	}
)

func NewInteractor(store Store, scm SCM) *Interactor {
	return &Interactor{
		store: store,
		scm:   scm,
	}
}
