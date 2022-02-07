package interactor

import (
	evbus "github.com/asaskevich/EventBus"
	"go.uber.org/zap"
)

type (
	Interactor struct {
		Store
		SCM

		// The channel to stop background workers.
		stopCh chan struct{}

		common *service

		// services used for talking to different parts of the entities.
		*DeploymentsInteractor
		*DeploymentStatisticsInteractor
		*EventsInteractor
		*LicenseInteractor
		*LocksInteractor
		*ReposInteractor
		*SyncInteractor
		*UsersInteractor
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
		Store:  c.Store,
		SCM:    c.SCM,
		stopCh: make(chan struct{}),
	}

	log := zap.L().Named("interactor")

	i.common = &service{
		store: c.Store,
		scm:   c.SCM,
		log:   log,
	}

	i.DeploymentsInteractor = (*DeploymentsInteractor)(i.common)
	i.DeploymentStatisticsInteractor = (*DeploymentStatisticsInteractor)(i.common)
	i.EventsInteractor = &EventsInteractor{
		service: i.common,
		events:  evbus.New(),
	}
	i.LicenseInteractor = &LicenseInteractor{
		service:    i.common,
		LicenseKey: c.LicenseKey,
	}
	i.LocksInteractor = (*LocksInteractor)(i.common)
	i.ReposInteractor = (*ReposInteractor)(i.common)
	i.SyncInteractor = &SyncInteractor{
		service:    i.common,
		orgEntries: c.OrgEntries,
	}
	i.UsersInteractor = &UsersInteractor{
		service:       i.common,
		admins:        c.AdminUsers,
		orgEntries:    c.OrgEntries,
		memberEntries: c.MemberEntries,
	}

	go func() {
		log.Info("Start the working publishing events.")
		i.runPublishingEvents(i.stopCh)
	}()

	go func() {
		log.Info("Start the worker canceling inactive deployments.")
		i.runClosingInactiveDeployment(i.stopCh)
	}()

	go func() {
		log.Info("Start the worker for the auto unlock.")
		i.runAutoUnlock(i.stopCh)
	}()

	return i
}
