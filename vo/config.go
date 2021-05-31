package vo

type (
	Config struct {
		Envs []*Env `yaml:"envs"`
	}

	Env struct {
		Name             string    `yaml:"name"`
		RequiredContexts []string  `yaml:"required_contexts"`
		Approval         *Approval `yaml:"approval"`
	}

	Approval struct {
		Approvers  []string `yaml:"approvers"`
		WaitMinute int      `yaml:"wait_minute"`
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
