package interactor

import (
	"context"
	"fmt"
	"time"

	evbus "github.com/asaskevich/EventBus"
	"go.uber.org/zap"
)

type InteractorConfig struct {
	ServerHost       string
	ServerProto      string
	ServerProxyHost  string
	ServerProxyProto string

	OrgEntries    []string
	MemberEntries []string
	AdminUsers    []string

	WebhookSecret string

	LicenseKey string

	Store
	SCM
}

func (c *InteractorConfig) BuildWebhookURL() string {
	if c.ServerProxyProto != "" && c.ServerProxyHost != "" {
		return fmt.Sprintf("%s://%s/hooks", c.ServerProxyProto, c.ServerProxyHost)
	}

	return fmt.Sprintf("%s://%s/hooks", c.ServerProto, c.ServerHost)
}

func (c *InteractorConfig) CheckWebhookSSL() bool {
	if c.ServerProxyProto != "" && c.ServerProxyHost != "" {
		return c.ServerProxyProto == "https"
	}

	return c.ServerProto == "https"
}

type service struct {
	store Store
	scm   SCM
	log   *zap.Logger
}

type Interactor struct {
	Store
	SCM

	// The channel to stop background workers.
	stopCh chan struct{}

	common *service

	// services used for talking to different parts of the entities.
	*ConfigInteractor
	*DeploymentInteractor
	*DeploymentStatisticsInteractor
	*EventInteractor
	*LicenseInteractor
	*LockInteractor
	*RepoInteractor
	*ReviewInteractor
	*UserInteractor
	*PermInteractor
}

func NewInteractor(c *InteractorConfig) *Interactor {
	log := zap.L().Named("interactor")
	defer log.Sync()

	i := &Interactor{
		Store:  c.Store,
		SCM:    c.SCM,
		stopCh: make(chan struct{}),
	}

	i.common = &service{
		store: c.Store,
		scm:   c.SCM,
		log:   log,
	}

	i.ConfigInteractor = (*ConfigInteractor)(i.common)
	i.DeploymentInteractor = (*DeploymentInteractor)(i.common)
	i.DeploymentStatisticsInteractor = (*DeploymentStatisticsInteractor)(i.common)
	i.EventInteractor = &EventInteractor{
		service: i.common,
		events:  evbus.New(),
	}
	i.LicenseInteractor = &LicenseInteractor{
		service:    i.common,
		LicenseKey: c.LicenseKey,
	}
	i.LockInteractor = (*LockInteractor)(i.common)
	i.RepoInteractor = &RepoInteractor{
		service:       i.common,
		WebhookURL:    c.BuildWebhookURL(),
		WebhookSSL:    c.CheckWebhookSSL(),
		WebhookSecret: c.WebhookSecret,
	}
	i.ReviewInteractor = (*ReviewInteractor)(i.common)
	i.UserInteractor = &UserInteractor{
		service:       i.common,
		admins:        c.AdminUsers,
		orgEntries:    c.OrgEntries,
		memberEntries: c.MemberEntries,
	}
	i.PermInteractor = &PermInteractor{
		service:    i.common,
		orgEntries: c.OrgEntries,
	}

	return i
}

func (i *Interactor) Init() {
	log := zap.L().Named("interactor")
	defer log.Sync()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	log.Debug("Resync organization entries.")
	if err := i.ResyncPerms(ctx); err != nil {
		log.Fatal("Failed to resynchronize with the perms.", zap.Error(err))
	}

	go func() {
		log.Debug("Start the working publishing events.")
		i.runPublishingEvents(i.stopCh)
	}()

	go func() {
		log.Debug("Start the worker canceling inactive deployments.")
		i.runClosingInactiveDeployment(i.stopCh)
	}()

	go func() {
		log.Debug("Start the worker for the auto unlock.")
		i.runAutoUnlock(i.stopCh)
	}()
}
