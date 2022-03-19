package extent

import (
	"reflect"
	"testing"

	"github.com/AlekSi/pointer"
	"github.com/davecgh/go-spew/spew"
)

func TestUnmarshalYAML(t *testing.T) {
	t.Run("Unmarhsal the required_context field", func(tt *testing.T) {
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
					Name:             "dev",
					RequiredContexts: &[]string{"github-action"},
				},
			},
			source: []byte(s),
		}
		if !reflect.DeepEqual(c, e) {
			tt.Errorf("Config = %s, expected %s", spew.Sdump(c), spew.Sdump(e))
		}
	})

	t.Run("Unmarshal 'auto_merge: false'", func(tt *testing.T) {
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
					Name:      "dev",
					AutoMerge: pointer.ToBool(false),
				},
			},
			source: []byte(s),
		}
		if !reflect.DeepEqual(c, e) {
			tt.Errorf("Config = %s, expected %s", spew.Sdump(c), spew.Sdump(e))
		}
	})

	t.Run("Unmarshal 'auto_merge: true'", func(tt *testing.T) {
		s := `
envs:
  - name: dev
    auto_merge: true`
		c := &Config{}

		err := UnmarshalYAML([]byte(s), c)
		if err != nil {
			tt.Fatalf("failed to parse: %s", err)
		}

		e := &Config{
			Envs: []*Env{
				{
					Name:      "dev",
					AutoMerge: pointer.ToBool(true),
				},
			},
			source: []byte(s),
		}
		if !reflect.DeepEqual(c, e) {
			tt.Errorf("Config = %s, expected %s", spew.Sdump(c), spew.Sdump(e))
		}
	})
}

func TestConfig_Eval(t *testing.T) {
	t.Run("Umarshal the task with the variable template.", func(t *testing.T) {
		s := `
envs:
  - name: dev
    task: ${GITPLOY_DEPLOY_TASK}:kubernetes`

		c := &Config{}
		if err := UnmarshalYAML([]byte(s), c); err != nil {
			t.Fatalf("Failed to parse the configuration file: %v", err)
		}

		err := c.Eval(&EvalValues{})
		if err != nil {
			t.Fatalf("Eval returns an error: %v", err)
		}

		e := &Config{
			Envs: []*Env{
				{
					Name: "dev",
					Task: pointer.ToString("deploy:kubernetes"),
				},
			},
			source: []byte(s),
		}
		if !reflect.DeepEqual(c, e) {
			t.Errorf("Config = %v expected %v", spew.Sdump(c), spew.Sdump(e))
		}
	})

	t.Run("Unmarshal the deployable_ref field with a regexp.", func(t *testing.T) {
		s := `
envs:
  - name: dev
    task: ${GITPLOY_DEPLOY_TASK}:kubernetes
    deployable_ref: 'v.*\..*\..*'`

		c := &Config{}
		if err := UnmarshalYAML([]byte(s), c); err != nil {
			t.Fatalf("Failed to parse the configuration file: %v", err)
		}

		err := c.Eval(&EvalValues{})
		if err != nil {
			t.Fatalf("Eval returns an error: %v", err)
		}

		e := &Config{
			Envs: []*Env{
				{
					Name:          "dev",
					Task:          pointer.ToString("deploy:kubernetes"),
					DeployableRef: pointer.ToString(`v.*\..*\..*`),
				},
			},
			source: []byte(s),
		}
		if !reflect.DeepEqual(c, e) {
			t.Errorf("Config = %v expected %v", spew.Sdump(c), spew.Sdump(e))
		}
	})

	t.Run("Unmarshal the frozen_windows field", func(t *testing.T) {
		s := `
envs:
  - name: dev
    frozen_windows:
      - start: "55 23 * * *"
        duration: "10m"`

		c := &Config{}
		if err := UnmarshalYAML([]byte(s), c); err != nil {
			t.Fatalf("Failed to parse the configuration file: %v", err)
		}

		e := &Config{
			Envs: []*Env{
				{
					Name: "dev",
					FrozenWindows: []FrozenWindow{
						{
							Start:    "55 23 * * *",
							Duration: "10m",
						},
					},
				},
			},
			source: []byte(s),
		}
		if !reflect.DeepEqual(c, e) {
			t.Errorf("Config = %v expected %v", spew.Sdump(c), spew.Sdump(e))
		}
	})
}
