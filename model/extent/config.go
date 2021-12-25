package extent

import (
	"regexp"
	"strconv"

	"github.com/drone/envsubst"
	"gopkg.in/yaml.v3"

	eutil "github.com/gitploy-io/gitploy/pkg/e"
)

type (
	Config struct {
		Envs []*Env `json:"envs" yaml:"envs"`

		// save the content of the configuration file
		// when it is unmarshalled.
		source []byte
	}

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

		// Review is the configuration of Review,
		// It is disabled when it is empty.
		Review *Review `json:"review,omitempty" yaml:"review"`
	}

	Review struct {
		Enabled   bool     `json:"enabled" yaml:"enabled"`
		Reviewers []string `json:"reviewers" yaml:"reviewers"`
	}

	EvalValues struct {
		IsRollback bool
	}
)

const (
	VarnameDeployTask   = "GITPLOY_DEPLOY_TASK"
	VarnameRollbackTask = "GITPLOY_ROLLBACK_TASK"
	VarnameIsRollback   = "GITPLOY_IS_ROLLBACK"
)

const (
	// DefaultDeployTask is the value of the 'GITPLOY_DEPLOY_TASK' variable.
	DefaultDeployTask = "deploy"
	// DefaultRollbackTask is the value of the 'GITPLOY_ROLLBACK_TASK' variable.
	DefaultRollbackTask = "rollback"
)

func UnmarshalYAML(content []byte, c *Config) error {
	if err := yaml.Unmarshal([]byte(content), c); err != nil {
		return eutil.NewError(eutil.ErrorCodeConfigInvalid, err)
	}

	c.source = content

	return nil
}

func (c *Config) Eval(v *EvalValues) error {
	// Evaluates variables
	mapper := func(vn string) string {
		if vn == VarnameDeployTask {
			if !v.IsRollback {
				return DefaultDeployTask
			} else {
				return ""
			}
		}

		if vn == VarnameRollbackTask {
			if v.IsRollback {
				return DefaultRollbackTask
			} else {
				return ""
			}
		}

		if vn == VarnameIsRollback {
			return strconv.FormatBool(v.IsRollback)
		}

		return "ERR_NOT_IMPLEMENTED"
	}

	evalued, err := envsubst.Eval(string(c.source), mapper)
	if err != nil {
		return eutil.NewError(eutil.ErrorCodeConfigInvalid, err)
	}

	if err := yaml.Unmarshal([]byte(evalued), c); err != nil {
		return eutil.NewError(eutil.ErrorCodeConfigInvalid, err)
	}

	return nil
}

func (c *Config) HasEnv(name string) bool {
	for _, e := range c.Envs {
		if e.Name == name {
			return true
		}
	}

	return false
}

func (c *Config) GetEnv(name string) *Env {
	for _, e := range c.Envs {
		if e.Name == name {
			return e
		}
	}

	return nil
}

// IsProductionEnvironment check whether the environment is production or not.
func (e *Env) IsProductionEnvironment() bool {
	return e.ProductionEnvironment != nil && *e.ProductionEnvironment
}

// IsDeployableRef validate the ref is deployable.
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

// IsAutoDeployOn validate the ref is matched with 'auto_deploy_on'.
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
