package extent

import (
	"regexp"
	"strings"
	"time"

	"github.com/gitploy-io/cronexpr"
	eutil "github.com/gitploy-io/gitploy/pkg/e"
)

type (
	Env struct {
		Name string `json:"name" yaml:"name"`

		// GitHub parameters of deployment.
		Task                  *string     `json:"task" yaml:"task"`
		Description           *string     `json:"description" yaml:"description"`
		AutoMerge             *bool       `json:"auto_merge" yaml:"auto_merge"`
		RequiredContexts      *[]string   `json:"required_contexts,omitempty" yaml:"required_contexts"`
		Payload               interface{} `json:"payload" yaml:"payload"`
		ProductionEnvironment *bool       `json:"production_environment" yaml:"production_environment"`

		// DeployableRef validates the ref is deployable or not.
		DeployableRef *string `json:"deployable_ref" yaml:"deployable_ref"`

		// AutoDeployOn deploys automatically when the pattern is matched.
		AutoDeployOn *string `json:"auto_deploy_on" yaml:"auto_deploy_on"`

		// Serialization verify if there is a running deployment.
		Serialization *bool `json:"serialization" yaml:"serialization"`

		// Review is the configuration of Review,
		// It is disabled when it is empty.
		Review *Review `json:"review,omitempty" yaml:"review"`

		// FrozenWindows is the list of windows to freeze deployments.
		FrozenWindows []FrozenWindow `json:"frozen_windows" yaml:"frozen_windows"`
	}

	Review struct {
		Enabled   bool     `json:"enabled" yaml:"enabled"`
		Reviewers []string `json:"reviewers" yaml:"reviewers"`
	}

	FrozenWindow struct {
		Start    string `json:"start" yaml:"start"`
		Duration string `json:"duration" yaml:"duration"`
		Location string `json:"location" yaml:"location"`
	}
)

// IsProductionEnvironment verifies whether the environment is production or not.
func (e *Env) IsProductionEnvironment() bool {
	return e.ProductionEnvironment != nil && *e.ProductionEnvironment
}

// IsDeployableRef verifies the ref is deployable.
func (e *Env) IsDeployableRef(ref string) (bool, error) {
	if e.DeployableRef == nil {
		return true, nil
	}

	matched, err := regexp.MatchString(*e.DeployableRef, ref)
	if err != nil {
		return false, eutil.NewError(eutil.ErrorCodeConfigInvalid, err)
	}

	return matched, nil
}

// IsAutoDeployOn verifies the ref is matched with 'auto_deploy_on'.
func (e *Env) IsAutoDeployOn(ref string) (bool, error) {
	if e.AutoDeployOn == nil {
		return false, nil
	}

	matched, err := regexp.MatchString(*e.AutoDeployOn, ref)
	if err != nil {
		return false, eutil.NewError(eutil.ErrorCodeConfigInvalid, err)
	}

	return matched, nil
}

// HasReview check whether the review is enabled or not.
func (e *Env) HasReview() bool {
	return e.Review != nil && e.Review.Enabled
}

// IsFreezed verifies whether the current time is in a freeze window.
// It returns an error when parsing an expression is failed.
func (e *Env) IsFreezed(t time.Time) (bool, error) {
	if len(e.FrozenWindows) == 0 {
		return false, nil
	}

	for _, w := range e.FrozenWindows {
		s, err := cronexpr.ParseInLocation(strings.TrimSpace(w.Start), w.Location)
		if err != nil {
			return false, eutil.NewErrorWithMessage(
				eutil.ErrorCodeConfigInvalid,
				"The crontab expression of the freeze window is invalid.",
				err,
			)
		}

		d, err := time.ParseDuration(w.Duration)
		if err != nil {
			return false, eutil.NewErrorWithMessage(
				eutil.ErrorCodeConfigInvalid,
				"The duration of the freeze window is invalid.",
				err,
			)
		}

		// Add one minute to include the starting time.
		start := s.Prev(t.Add(time.Minute))
		end := start.Add(d)
		if t.After(start) && t.Before(end) {
			return true, nil
		}
	}

	return false, nil
}
