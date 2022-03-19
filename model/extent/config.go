package extent

import (
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
