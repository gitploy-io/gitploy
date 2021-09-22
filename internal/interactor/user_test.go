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

func TestInteractor_IsEntryMember(t *testing.T) {
	t.Run("Return false when the user's login is not includes.", func(t *testing.T) {
		i := &Interactor{
			memberEntries: []string{"octocat"},
		}

		want := false
		if ret := i.IsEntryMember(context.Background(), "coco"); ret != want {
			t.Fatalf("IsEntryMember = %v, wanted %v", ret, want)
		}
	})
}
