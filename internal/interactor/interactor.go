package interactor

import (
	"context"

	evbus "github.com/asaskevich/EventBus"
	"go.uber.org/zap"

	"github.com/hanjunlee/gitploy/ent"
)

type (
	Interactor struct {
		Store
		SCM
		Chat

		// Notification events
		stopCh chan struct{}
		events evbus.Bus

		log *zap.Logger
	}

	// Chat is optional function.
	// by provide FakeChat you can disable chat.
	FakeChat struct{}
)

func NewInteractor(store Store, scm SCM, chat Chat) *Interactor {
	i := &Interactor{
		Store:  store,
		SCM:    scm,
		Chat:   chat,
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

func NewFakeChat() *FakeChat {
	return &FakeChat{}
}

func (c *FakeChat) Notify(ctx context.Context, cu *ent.ChatUser, n *ent.Notification) error {
	return nil
}
