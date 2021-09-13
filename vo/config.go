package vo

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/drone/envsubst"
	"gopkg.in/yaml.v3"
)

type (
	Config struct {
		Envs []*Env `json:"envs" yaml:"envs"`
	}

	Env struct {
		Name string `json:"name" yaml:"name"`

		// Github parameters of deployment.
		Task                  *string   `json:"task" yaml:"task"`
		Description           *string   `json:"description" yaml:"description"`
		AutoMerge             *bool     `json:"auto_merge" yaml:"auto_merge"`
		RequiredContexts      *[]string `json:"required_contexts,omitempty" yaml:"required_contexts"`
		Payload               *string   `json:"payload" yaml:"payload"`
		ProductionEnvironment *bool     `json:"production_environment" yaml:"production_environment"`

		// Approval is the configuration of Approval,
		// It is disabled when it is empty.
		Approval *Approval `json:"approval,omitempty" yaml:"approval"`
	}

	Approval struct {
		Enabled       bool `json:"enabled" yaml:"enabled"`
		RequiredCount int  `json:"required_count" yaml:"required_count"`
	}

	EvalValues struct {
		DeployTask   string
		RollbackTask string
		Tag          string
		IsRollback   bool
	}
)

const (
	varnameDeployTask   = "GITPLOY_DEPLOY_TASK"
	varnameRollbackTask = "GITPLOY_ROLLBACK_TASK"
	varnameTag          = "GITPLOY_TAG"
	varnameIsRollback   = "GITPLOY_IS_ROLLBACK"
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

func (e *Env) IsApprovalEabled() bool {
	if e.Approval == nil {
		return false
	}

	return e.Approval.Enabled
}

func (e *Env) Eval(v *EvalValues) error {
	byts, err := json.Marshal(e)
	if err != nil {
		return fmt.Errorf("failed to marshal the env: %w", err)
	}

	// Evaluates variables
	mapper := func(vn string) string {
		switch vn {
		case varnameDeployTask:
			return v.DeployTask
		case varnameRollbackTask:
			return v.RollbackTask
		case varnameIsRollback:
			return strconv.FormatBool(v.IsRollback)
		case varnameTag:
			return v.Tag
		default:
			return "ERR_NOT_IMPLEMENTED"
		}
	}

	evalued, err := envsubst.Eval(string(byts), mapper)
	if err != nil {
		return fmt.Errorf("failed to eval variables: %w", err)
	}

	ne := &Env{}
	if err := json.Unmarshal([]byte(evalued), ne); err != nil {
		return fmt.Errorf("failed to unmarshal to the env: %w", err)
	}

	return nil
}
