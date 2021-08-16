package vo

import (
	"strconv"

	"github.com/creasty/defaults"
	"gopkg.in/yaml.v3"
)

type (
	Config struct {
		Envs []*Env `json:"envs" yaml:"envs"`
	}

	Env struct {
		Name                  string    `json:"name" yaml:"name"`
		Task                  string    `json:"task" yaml:"task" default:"deploy"`
		Description           string    `json:"description" yaml:"description"`
		AutoMerge             bool      `default:"true"`
		RequiredContexts      []string  `json:"required_contexts" yaml:"required_contexts"`
		Payload               string    `json:"payload" yaml:"payload"`
		ProductionEnvironment bool      `json:"production_environment" yaml:"production_environment"`
		Approval              *Approval `json:"approval,omitempty" yaml:"approval"`

		// The type of auto_merge must be string to avoid
		// that the value of auto_merge is always set true
		// after processing defaults.Set
		StrAutoMerge string `json:"auto_merge" yaml:"auto_merge"`
	}

	Approval struct {
		Enabled       bool `json:"enabled" yaml:"enabled"`
		RequiredCount int  `json:"required_count" yaml:"required_count"`
	}
)

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

func UnmarshalYAML(content []byte, c *Config) error {
	if err := yaml.Unmarshal([]byte(content), c); err != nil {
		return err
	}

	if err := defaults.Set(c); err != nil {
		return err
	}

	// Set default value manually.
	for _, e := range c.Envs {
		am, err := strconv.ParseBool(e.StrAutoMerge)
		if err != nil {
			continue
		}

		e.AutoMerge = am
	}

	return nil
}
