package vo

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/AlekSi/pointer"
	"github.com/davecgh/go-spew/spew"
)

func TestUnmarshalYAML(t *testing.T) {
	t.Run("unmarhsal the required_context field", func(tt *testing.T) {
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

	t.Run("unmarshal auto_merge: false ", func(tt *testing.T) {
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

	t.Run("unmarshal auto_merge: true", func(tt *testing.T) {
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
	t.Run("Evaluate the configuration.", func(t *testing.T) {
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

	t.Run("Evaluate the configuration with the regexp.", func(t *testing.T) {
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
}

func TestEnv_IsProductionEnvironment(t *testing.T) {
	t.Run("Reutrn false when the production environment is nil", func(t *testing.T) {
		e := &Env{}

		expected := false
		if e.IsProductionEnvironment() != expected {
			t.Errorf("IsProductionEnvironment = %v, wanted %v", e.IsProductionEnvironment(), expected)
		}
	})

	t.Run("Reutrn true when the production environment is true", func(t *testing.T) {
		e := &Env{
			ProductionEnvironment: pointer.ToBool(true),
		}

		expected := true
		if e.IsProductionEnvironment() != expected {
			t.Errorf("IsProductionEnvironment = %v, wanted %v", e.IsProductionEnvironment(), expected)
		}
	})
}

func TestEnv_IsDeployableRef(t *testing.T) {
	t.Run("Return true when 'deployable_ref' is not defined.", func(t *testing.T) {
		e := &Env{}

		ret, err := e.IsDeployableRef("")
		if err != nil {
			t.Fatalf("IsDeployableRef returns an error: %s", err)
		}

		expected := true
		if ret != expected {
			t.Fatalf("IsDeployableRef = %v, wanted %v", ret, expected)
		}
	})

	t.Run("Return true when 'deployable_ref' is matched.", func(t *testing.T) {
		e := &Env{
			DeployableRef: pointer.ToString("main"),
		}

		ret, err := e.IsDeployableRef("main")
		if err != nil {
			t.Fatalf("IsDeployableRef returns an error: %s", err)
		}

		expected := true
		if ret != expected {
			t.Fatalf("IsDeployableRef = %v, wanted %v", ret, expected)
		}
	})

	t.Run("Return false when 'deployable_ref' is not matched.", func(t *testing.T) {
		e := &Env{
			DeployableRef: pointer.ToString("main"),
		}

		ret, err := e.IsDeployableRef("branch")
		if err != nil {
			t.Fatalf("IsDeployableRef returns an error: %s", err)
		}

		expected := false
		if ret != expected {
			t.Fatalf("IsDeployableRef = %v, wanted %v", ret, expected)
		}
	})
}

func TestEnv_Eval(t *testing.T) {
	t.Run("eval the task.", func(t *testing.T) {
		cs := []struct {
			env  *Env
			want *Env
		}{
			{
				env: &Env{
					Task: pointer.ToString("${GITPLOY_DEPLOY_TASK}"),
				},
				want: &Env{
					Task: pointer.ToString(DefaultDeployTask),
				},
			},
			{
				env: &Env{
					Task: pointer.ToString("${GITPLOY_DEPLOY_TASK}:kubernetes"),
				},
				want: &Env{
					Task: pointer.ToString(fmt.Sprintf("%s:kubernetes", DefaultDeployTask)),
				},
			},
			{
				env: &Env{
					Task: pointer.ToString("${GITPLOY_DEPLOY_TASK}${GITPLOY_ROLLBACK_TASK}"),
				},
				want: &Env{
					Task: pointer.ToString(DefaultDeployTask),
				},
			},
		}

		for _, c := range cs {
			err := c.env.Eval(&EvalValues{})
			if err != nil {
				t.Fatalf("Eval returns an error: %s", err)
			}
			if !reflect.DeepEqual(c.env, c.want) {
				t.Fatalf("Eval = %v, wanted %v", *c.env.Task, *c.want.Task)
			}
		}
	})

	t.Run("eval the is_rollback.", func(t *testing.T) {
		const (
			isRollback = true
		)

		cs := []struct {
			env  *Env
			want *Env
		}{
			{
				env: &Env{
					Payload: pointer.ToString("{\"is_rollback\": ${GITPLOY_IS_ROLLBACK}}"),
				},
				want: &Env{
					Payload: pointer.ToString("{\"is_rollback\": true}"),
				},
			},
		}

		for _, c := range cs {
			err := c.env.Eval(&EvalValues{
				IsRollback: isRollback,
			})
			if err != nil {
				t.Fatalf("Eval returns an error: %s", err)
			}
			if !reflect.DeepEqual(c.env, c.want) {
				t.Fatalf("Eval = %v, wanted %v", c.env, c.want)
			}
		}
	})
}
