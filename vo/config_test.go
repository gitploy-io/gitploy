package vo

import (
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestParseConfig(t *testing.T) {
	t.Run("parse the config file", func(tt *testing.T) {
		s := `
envs:
  - name: dev
    required_contexts:
      - github-action
    approval:
      enabled: true`
		c := &Config{}

		err := UnmarshalYAML([]byte(s), c)
		if err != nil {
			tt.Errorf("failed to parse: %s", err)
			tt.FailNow()
		}

		e := &Config{
			Envs: []*Env{
				{
					Name:                  "dev",
					Task:                  "deploy",
					Description:           "",
					AutoMerge:             true,
					RequiredContexts:      []string{"github-action"},
					Payload:               "",
					ProductionEnvironment: false,
					Approval: &Approval{
						Enabled: true,
					},
				},
			},
		}
		if !reflect.DeepEqual(c, e) {
			tt.Errorf("Config = %s, expected %s", spew.Sdump(c), spew.Sdump(e))
		}
	})

	t.Run("parse auto_merge field.", func(tt *testing.T) {
		s := `
envs:
  - name: dev
    auto_merge: false
    required_contexts:
      - github-action
    approval:
	  enabled: true
	  required_count: 1`
		c := &Config{}

		err := UnmarshalYAML([]byte(s), c)
		if err != nil {
			tt.Errorf("failed to parse: %s", err)
			tt.FailNow()
		}

		e := &Config{
			Envs: []*Env{
				{
					Name:                  "dev",
					Task:                  "deploy",
					Description:           "",
					StrAutoMerge:          "false",
					AutoMerge:             false,
					RequiredContexts:      []string{"github-action"},
					Payload:               "",
					ProductionEnvironment: false,
					Approval: &Approval{
						Enabled:       true,
						RequiredCount: 1,
					},
				},
			},
		}
		if !reflect.DeepEqual(c, e) {
			tt.Errorf("Config = %s, expected %s", spew.Sdump(c), spew.Sdump(e))
		}
	})
}
