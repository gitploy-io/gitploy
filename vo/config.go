package vo

type (
	Config struct {
		Envs []*Env `json:"envs" yaml:"envs"`
	}

	Env struct {
		Name             string    `json:"name" yaml:"name"`
		RequiredContexts []string  `json:"required_contexts" yaml:"required_contexts"`
		Approval         *Approval `json:"approval" yaml:"approval"`
	}

	Approval struct {
		Approvers  []string `json:"approvers" yaml:"approvers"`
		WaitMinute int      `json:"wait_minute" yaml:"wait_minute"`
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

func (e *Env) HasApproval() bool {
	if e.Approval == nil {
		return false
	}

	return true
}
