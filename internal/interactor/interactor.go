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

		Store
		SCM

		// The channel to stop background workers.
		stopCh chan struct{}
		events evbus.Bus

		log *zap.Logger

		common *service

		// services used for talking to different parts of the entities.
		*LicenseInteractor
		*UsersInteractor
		*DeploymentsInteractor
		*DeploymentStatisticsInteractor
	}

	InteractorConfig struct {
		ServerHost  string
		ServerProto string

		OrgEntries    []string
		MemberEntries []string
		AdminUsers    []string

		LicenseKey string

		Store
		SCM
	}

	service struct {
		store Store
		scm   SCM
		log   *zap.Logger
	}
)

func NewInteractor(c *InteractorConfig) *Interactor {
	i := &Interactor{
		ServerHost:  c.ServerHost,
		ServerProto: c.ServerProto,
		Store:       c.Store,
		SCM:         c.SCM,
		stopCh:      make(chan struct{}),
		events:      evbus.New(),
		log:         zap.L().Named("interactor"),
	}

	i.common = &service{
		store: c.Store,
		scm:   c.SCM,
		log:   zap.L(),
	}

	i.LicenseInteractor = &LicenseInteractor{
		service:    i.common,
		LicenseKey: c.LicenseKey,
	}
	i.UsersInteractor = &UsersInteractor{
		service:       i.common,
		admins:        c.AdminUsers,
		orgEntries:    c.OrgEntries,
		memberEntries: c.MemberEntries,
	}
	i.DeploymentsInteractor = (*DeploymentsInteractor)(i.common)
	i.DeploymentStatisticsInteractor = (*DeploymentStatisticsInteractor)(i.common)

	go func() {
		i.log.Info("Start the working publishing events.")
		i.runPublishingEvents(i.stopCh)
	}()

	go func() {
		i.log.Info("Start the worker canceling inactive deployments.")
		i.runClosingInactiveDeployment(i.stopCh)
	}()

	go func() {
		i.log.Info("Start the worker for the auto unlock.")
		i.runAutoUnlock(i.stopCh)
	}()

	return i
}
