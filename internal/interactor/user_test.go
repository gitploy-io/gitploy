package interactor

import (
	"context"
	"testing"
)

func TestInteractor_IsAdminUser(t *testing.T) {
	t.Run("Return false when admins property is nil", func(t *testing.T) {
		i := &Interactor{
			admins: nil,
		}

		expected := false
		if ret := i.IsAdminUser(context.Background(), "octocat"); ret != expected {
			t.Fatalf("IsAdminUser = %v, wanted %v", ret, expected)
		}
	})
}
