package vo

import (
	"encoding/json"
	"strconv"

	"github.com/drone/envsubst"
	"gopkg.in/yaml.v3"

	eutil "github.com/gitploy-io/gitploy/pkg/e"
)

type (
	Config struct {
		Envs []*Env `json:"envs" yaml:"envs"`
	}

	Env struct {
		Name string `json:"name" yaml:"name"`

		// Github parameters of deployment.
		Task                  *string     `json:"task" yaml:"task"`
		Description           *string     `json:"description" yaml:"description"`
		AutoMerge             *bool       `json:"auto_merge" yaml:"auto_merge"`
		RequiredContexts      *[]string   `json:"required_contexts,omitempty" yaml:"required_contexts"`
		Payload               interface{} `json:"payload" yaml:"payload"`
		ProductionEnvironment *bool       `json:"production_environment" yaml:"production_environment"`

		// Approval is the configuration of Approval,
		// It is disabled when it is empty.
		Approval *Approval `json:"approval,omitempty" yaml:"approval"`

		// Review is the configuration of Review,
		// It is disabled when it is empty.
		Review *Review `json:"review,omitempty" yaml:"review"`
	}

	Approval struct {
		Enabled       bool `json:"enabled" yaml:"enabled"`
		RequiredCount int  `json:"required_count" yaml:"required_count"`
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
	varnameDeployTask   = "GITPLOY_DEPLOY_TASK"
	varnameRollbackTask = "GITPLOY_ROLLBACK_TASK"
	varnameIsRollback   = "GITPLOY_IS_ROLLBACK"
)

const (
	defaultDeployTask   = "deploy"
	defaultRollbackTask = "rollback"
)

func UnmarshalYAML(content []byte, c *Config) error {
	if err := yaml.Unmarshal([]byte(content), c); err != nil {
		return err
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

func (e *Env) IsProductionEnvironment() bool {
	return e.ProductionEnvironment != nil && *e.ProductionEnvironment
}

func (e *Env) IsApprovalEabled() bool {
	if e.Approval == nil {
		return false
	}

	return e.Approval.Enabled
}

func (e *Env) HasReview() bool {
	return e.Review != nil && e.Review.Enabled
}

func (e *Env) Eval(v *EvalValues) error {
	byts, err := json.Marshal(e)
	if err != nil {
		return eutil.NewError(eutil.ErrorCodeConfigParseError, err)
	}

	// Evaluates variables
	mapper := func(vn string) string {
		if vn == varnameDeployTask {
			if !v.IsRollback {
				return defaultDeployTask
			} else {
				return ""
			}
		}

		if vn == varnameRollbackTask {
			if v.IsRollback {
				return defaultRollbackTask
			} else {
				return ""
			}
		}

		if vn == varnameIsRollback {
			return strconv.FormatBool(v.IsRollback)
		}

		return "ERR_NOT_IMPLEMENTED"
	}

	evalued, err := envsubst.Eval(string(byts), mapper)
	if err != nil {
		return eutil.NewError(eutil.ErrorCodeConfigParseError, err)
	}

	if err := json.Unmarshal([]byte(evalued), e); err != nil {
		return eutil.NewError(eutil.ErrorCodeConfigParseError, err)
	}

	return nil
}
