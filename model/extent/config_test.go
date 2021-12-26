package extent

import (
	"reflect"
	"testing"
	"time"

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

	t.Run("Unmarshal the freeze_windows field", func(t *testing.T) {
		s := `
envs:
  - name: dev
    freeze_windows:
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
					FreezeWindows: []FreezeWindow{
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

func TestEnv_IsFreezed(t *testing.T) {
	t.Run("Return true when the time is in the window", func(t *testing.T) {
		runs := []struct {
			t    time.Time
			e    *Env
			want bool
		}{
			{
				t: time.Date(2012, 12, 1, 23, 55, 10, 0, time.UTC),
				e: &Env{
					FreezeWindows: []FreezeWindow{
						{
							Start:    "55 23 * Dec *",
							Duration: "10m",
						},
					},
				},
				want: true,
			},
			{
				t: time.Date(2012, 1, 1, 0, 3, 0, 0, time.UTC),
				e: &Env{
					FreezeWindows: []FreezeWindow{
						{
							Start:    "55 23 * Dec *",
							Duration: "10m",
						},
					},
				},
				want: true,
			},
		}
		e := &Env{
			FreezeWindows: []FreezeWindow{
				{
					Start:    "55 23 * Dec *",
					Duration: "10m",
				},
			},
		}

		for _, r := range runs {
			freezed, err := e.IsFreezed(r.t)
			if err != nil {
				t.Fatalf("IsFreezed returns an error: %s", err)
			}

			if freezed != r.want {
				t.Fatalf("IsFreezed = %v, wanted %v", freezed, r.want)
			}
		}
	})

	t.Run("Return false when the time is out of the window", func(t *testing.T) {
		e := &Env{
			FreezeWindows: []FreezeWindow{
				{
					Start:    "55 23 * Dec *",
					Duration: "10m",
				},
			},
		}

		freezed, err := e.IsFreezed(time.Date(2012, 1, 1, 0, 10, 0, 0, time.UTC))
		if err != nil {
			t.Fatalf("IsFreezed returns an error: %s", err)
		}

		if freezed != false {
			t.Fatalf("IsFreezed = %v, wanted %v", freezed, false)
		}
	})
}
