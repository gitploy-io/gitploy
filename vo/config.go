package vo

type (
	Config struct {
		Envs []*Env
	}

	Env struct {
		Name             string
		RequiredContexts []string
		Approval         *Approval
	}

	Approval struct {
		Approvers []string
		WaitTime  int
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
