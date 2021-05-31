package vo

import (
	"reflect"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestParseConfig(t *testing.T) {
	t.Run("parse the config file", func(tt *testing.T) {
		s := `
envs:
  - name: dev
    required_contexts:
      - github-action
    approval:
      approvers:
        - hanjunlee
      wait_minute: 60`
		c := &Config{}

		err := yaml.Unmarshal([]byte(s), c)
		if err != nil {
			tt.Errorf("failed to parse: %s", err)
			tt.FailNow()
		}

		e := &Config{
			Envs: []*Env{
				{
					Name:             "dev",
					RequiredContexts: []string{"github-action"},
					Approval: &Approval{
						Approvers:  []string{"hanjunlee"},
						WaitMinute: 60,
					},
				},
			},
		}
		if !reflect.DeepEqual(c, e) {
			tt.Errorf("Config = %v, expected %v", c, e)
		}
	})

	t.Run("parse the config file without required_contexts", func(tt *testing.T) {
		s := `
envs:
  - name: dev
    approval:
      approvers:
        - hanjunlee
      wait_minute: 60`
		c := &Config{}

		err := yaml.Unmarshal([]byte(s), c)
		if err != nil {
			tt.Errorf("failed to parse: %s", err)
			tt.FailNow()
		}

		e := &Config{
			Envs: []*Env{
				{
					Name:             "dev",
					RequiredContexts: nil,
					Approval: &Approval{
						Approvers:  []string{"hanjunlee"},
						WaitMinute: 60,
					},
				},
			},
		}
		if !reflect.DeepEqual(c, e) {
			tt.Errorf("Config = %v, expected %v", *c, *e)
		}
	})
}
