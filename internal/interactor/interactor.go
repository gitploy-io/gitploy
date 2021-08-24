package interactor

import (
	evbus "github.com/asaskevich/EventBus"
	"go.uber.org/zap"
)

type (
	Interactor struct {
		// Host and protocol of server for Log URL of deployment status.
		ServerHost  string
		ServerProto string

		orgEntries []string
		// Admin Users
		admins []string

		// License
		licenseKey string

		Store
		SCM

		// The channel to stop background workers.
		stopCh chan struct{}
		events evbus.Bus

		log *zap.Logger
	}

	InteractorConfig struct {
		ServerHost  string
		ServerProto string

		OrgEntries []string
		AdminUsers []string

		LicenseKey string

		Store
		SCM
	}
)

func NewInteractor(c *InteractorConfig) *Interactor {
	i := &Interactor{
		ServerHost:  c.ServerHost,
		ServerProto: c.ServerProto,
		orgEntries:  c.OrgEntries,
		admins:      c.AdminUsers,
		licenseKey:  c.LicenseKey,
		Store:       c.Store,
		SCM:         c.SCM,
		stopCh:      make(chan struct{}),
		events:      evbus.New(),
		log:         zap.L().Named("interactor"),
	}

	go func() {
		i.log.Info("Start the working publishing events.")
		i.runPublishingEvents(i.stopCh)
	}()

	go func() {
		i.log.Info("Start the worker canceling inactive deployments.")
		i.runClosingInactiveDeployment(i.stopCh)
	}()
	return i
}
