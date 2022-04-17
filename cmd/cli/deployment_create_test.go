package main

import (
	"reflect"
	"testing"

	"github.com/gitploy-io/gitploy/model/extent"
)

func Test_buildDyanmicPayload(t *testing.T) {
	t.Run("Return an error when syntax is invalid.", func(t *testing.T) {
		_, err := buildDyanmicPayload([]string{
			"foo",
		}, &extent.Env{
			DynamicPayload: &extent.DynamicPayload{
				Enabled: true,
			},
		})

		if err == nil {
			t.Fatalf("buildDyanmicPayload dosen't return an error")
		}
	})

	t.Run("Return a payload with default values.", func(t *testing.T) {
		var qux interface{} = "qux"

		payload, err := buildDyanmicPayload([]string{}, &extent.Env{
			DynamicPayload: &extent.DynamicPayload{
				Enabled: true,
				Inputs: map[string]extent.Input{
					"foo": {
						Type: extent.InputTypeString,
					},
					"baz": {
						Type:    extent.InputTypeString,
						Default: &qux,
					},
				},
			},
		})

		if err != nil {
			t.Fatalf("buildDyanmicPayload returns an error")
		}

		if expected := map[string]interface{}{
			"baz": "qux",
		}; !reflect.DeepEqual(payload, expected) {
			t.Fatalf("buildDyanmicPayload = %v, wanted %v", payload, expected)
		}
	})
}
