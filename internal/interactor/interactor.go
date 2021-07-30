package interactor

import (
	evbus "github.com/asaskevich/EventBus"
	"go.uber.org/zap"
)

type (
	Interactor struct {
		Store
		SCM

		// Notification events
		stopCh chan struct{}
		events evbus.Bus

		log *zap.Logger
	}

	// Chat is optional function.
	// by provide FakeChat you can disable chat.
	FakeChat struct{}
)

func NewInteractor(store Store, scm SCM) *Interactor {
	i := &Interactor{
		Store:  store,
		SCM:    scm,
		stopCh: make(chan struct{}),
		events: evbus.New(),
		log:    zap.L().Named("interactor"),
	}

	go func() {
		i.log.Info("start to polling notification.")
		i.polling(i.stopCh)
	}()
	return i
}
