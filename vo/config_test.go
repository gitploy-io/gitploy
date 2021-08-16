package vo

import (
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestParseConfig(t *testing.T) {
	t.Run("parse required_context field", func(tt *testing.T) {
		s := `
envs:
  - name: dev
    required_contexts:
      - github-action`
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
					Description:           "Gitploy starts to deploy.",
					AutoMerge:             true,
					RequiredContexts:      &[]string{"github-action"},
					Payload:               "",
					ProductionEnvironment: false,
					Approval:              nil,
				},
			},
		}
		if !reflect.DeepEqual(c, e) {
			tt.Errorf("Config = %s, expected %s", spew.Sdump(c), spew.Sdump(e))
		}
	})

	t.Run("parse approval field", func(tt *testing.T) {
		s := `
envs:
  - name: dev
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
					Description:           "Gitploy starts to deploy.",
					AutoMerge:             true,
					RequiredContexts:      nil,
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
    auto_merge: false`
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
					Description:           "Gitploy starts to deploy.",
					StrAutoMerge:          "false",
					AutoMerge:             false,
					RequiredContexts:      nil,
					Payload:               "",
					ProductionEnvironment: false,
					Approval:              nil,
				},
			},
		}
		if !reflect.DeepEqual(c, e) {
			tt.Errorf("Config = %s, expected %s", spew.Sdump(c), spew.Sdump(e))
		}
	})
}
