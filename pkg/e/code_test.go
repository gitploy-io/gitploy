package e

import (
	"fmt"
	"net/http"
	"testing"
)

func TestError_GetHTTPError(t *testing.T) {
	t.Run("Return the matche HTTP code.", func(t *testing.T) {
		err := NewError(ErrorCodeInternalError, nil)
		if err.GetHTTPCode() != http.StatusInternalServerError {
			t.Fatalf("GetHTTPCode = %v, wanted %v", err.GetHTTPCode(), http.StatusInternalServerError)
		}
	})
}

func Test_IsError(t *testing.T) {
	t.Run("Return true when the type of error is Error.", func(t *testing.T) {
		err := NewError(ErrorCodeInternalError, nil)
		if ok := IsError(err); !ok {
			t.Fatalf("IsError = %v, wanted %v", ok, true)
		}
	})

	t.Run("Return false when the type of error is not Error.", func(t *testing.T) {
		err := fmt.Errorf("fmt.Error")
		if ok := IsError(err); ok {
			t.Fatalf("IsError = %v, wanted %v", ok, false)
		}
	})
}

func Test_HasErrorCode(t *testing.T) {
	t.Run("Return true when the type of error is Error and the code is internal_error.", func(t *testing.T) {
		err := NewError(ErrorCodeInternalError, nil)
		if ok := HasErrorCode(err, ErrorCodeInternalError); !ok {
			t.Fatalf("IsError = %v, wanted %v", ok, true)
		}
	})
}
