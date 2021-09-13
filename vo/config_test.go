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
			tt.Errorf("failed to parse: %s", err)
			tt.FailNow()
		}

		e := &Config{
			Envs: []*Env{
				{
					Name:      "dev",
					AutoMerge: pointer.ToBool(true),
				},
			},
		}
		if !reflect.DeepEqual(c, e) {
			tt.Errorf("Config = %s, expected %s", spew.Sdump(c), spew.Sdump(e))
		}
	})
}

func TestEnv_Eval(t *testing.T) {
	t.Run("eval the task.", func(t *testing.T) {
		const (
			deployTask = "deploy"
		)

		cs := []struct {
			env  *Env
			want *Env
		}{
			{
				env: &Env{
					Task: pointer.ToString("${GITPLOY_DEPLOY_TASK}"),
				},
				want: &Env{
					Task: pointer.ToString(deployTask),
				},
			},
			{
				env: &Env{
					Task: pointer.ToString("${GITPLOY_DEPLOY_TASK}:kubernetes"),
				},
				want: &Env{
					Task: pointer.ToString(fmt.Sprintf("%s:kubernetes", deployTask)),
				},
			},
			{
				env: &Env{
					Task: pointer.ToString("${GITPLOY_DEPLOY_TASK}${GITPLOY_ROLLBACK_TASK}"),
				},
				want: &Env{
					Task: pointer.ToString(deployTask),
				},
			},
		}

		for _, c := range cs {
			err := c.env.Eval(&EvalValues{
				DeployTask: deployTask,
			})
			if err != nil {
				t.Fatalf("Eval returns an error: %s", err)
			}
			if !reflect.DeepEqual(c.env, c.want) {
				t.Fatalf("Eval = %v, wanted %v", c.env, c.want)
			}
		}
	})

	t.Run("eval the tag.", func(t *testing.T) {
		const (
			tag = "a/v0.1.0"
		)

		cs := []struct {
			env  *Env
			want *Env
		}{
			{
				env: &Env{
					Task: pointer.ToString("${GITPLOY_TAG#a/}"),
				},
				want: &Env{
					Task: pointer.ToString("v0.1.0"),
				},
			},
		}

		for _, c := range cs {
			err := c.env.Eval(&EvalValues{
				Tag: tag,
			})
			if err != nil {
				t.Fatalf("Eval returns an error: %s", err)
			}
			if !reflect.DeepEqual(c.env, c.want) {
				t.Fatalf("Eval = %v, wanted %v", c.env, c.want)
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
