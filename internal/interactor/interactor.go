package interactor

import (
	evbus "github.com/asaskevich/EventBus"
	"github.com/hanjunlee/gitploy/vo"
	"go.uber.org/zap"
)

type (
	Interactor struct {
		// Host and protocol of server for Log URL of deployment status.
		ServerHost  string
		ServerProto string

		// License
		licenseKey string
		license    *vo.License

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

		LicenseKey string

		Store
		SCM
	}
)

func NewInteractor(c *InteractorConfig) *Interactor {
	i := &Interactor{
		ServerHost:  c.ServerHost,
		ServerProto: c.ServerProto,
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
