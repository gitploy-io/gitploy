package extent

import (
	"testing"
	"time"

	"github.com/AlekSi/pointer"
)

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
					FrozenWindows: []FrozenWindow{
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
					FrozenWindows: []FrozenWindow{
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
			FrozenWindows: []FrozenWindow{
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
			FrozenWindows: []FrozenWindow{
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

func TestDynamicPayload_Validate(t *testing.T) {
	t.Run("Return an error when the required field is not exist.", func(t *testing.T) {
		c := DynamicPayload{
			Inputs: map[string]Input{
				"foo": {
					Type:     InputTypeString,
					Required: pointer.ToBool(true),
				},
			},
		}

		err := c.Validate(map[string]interface{}{})
		if err == nil {
			t.Fatalf("Validate doesn't return an error.")
		}
	})

	t.Run("Skip the optional field if there is no value.", func(t *testing.T) {
		c := DynamicPayload{
			Inputs: map[string]Input{
				"foo": {
					Type: InputTypeString,
				},
			},
		}

		err := c.Validate(map[string]interface{}{})
		if err != nil {
			t.Fatalf("Validate return an error: %s", err)
		}
	})

	t.Run("Return an error when the selected value is not in the options.", func(t *testing.T) {
		c := DynamicPayload{
			Inputs: map[string]Input{
				"foo": {
					Type:     InputTypeSelect,
					Required: pointer.ToBool(true),
					Options:  &[]string{"option1", "option2"},
				},
			},
		}

		input := map[string]interface{}{"foo": "value"}

		err := c.Validate(input)
		if err == nil {
			t.Fatalf("Validate doesn't return an error.")
		}
	})

	t.Run("Return nil if validation has succeed.", func(t *testing.T) {
		c := DynamicPayload{
			Inputs: map[string]Input{
				"foo": {
					Type:     InputTypeSelect,
					Required: pointer.ToBool(true),
					Options:  &[]string{"option1", "option2"},
				},
				"bar": {
					Type:     InputTypeNumber,
					Required: pointer.ToBool(true),
				},
				"baz": {
					Type: InputTypeNumber,
				},
				"qux": {
					Type: InputTypeBoolean,
				},
			},
		}

		input := map[string]interface{}{
			"foo": "option1",
			"bar": 1,
			"baz": 4.2,
			"qux": false,
		}

		err := c.Validate(input)
		if err != nil {
			t.Fatalf("Evaluate return an error: %s", err)
		}
	})
}
