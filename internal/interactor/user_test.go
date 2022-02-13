package interactor_test

import (
	"context"
	"testing"

	i "github.com/gitploy-io/gitploy/internal/interactor"
)

func TestInteractor_IsAdminUser(t *testing.T) {
	t.Run("Return false when admins property is nil", func(t *testing.T) {
		i := i.NewInteractor(&i.InteractorConfig{
			AdminUsers: nil,
		})

		expected := false
		if ret := i.IsAdminUser(context.Background(), "octocat"); ret != expected {
			t.Fatalf("IsAdminUser = %v, wanted %v", ret, expected)
		}
	})
}

func TestInteractor_IsEntryMember(t *testing.T) {
	t.Run("Return false when the user's login is not included.", func(t *testing.T) {
		it := i.NewInteractor(&i.InteractorConfig{
			MemberEntries: []string{"octocat"},
		})

		want := false
		if ret := it.IsEntryMember(context.Background(), "coco"); ret != want {
			t.Fatalf("IsEntryMember = %v, wanted %v", ret, want)
		}
	})

	t.Run("Return true when the user's login is included.", func(t *testing.T) {
		it := i.NewInteractor(&i.InteractorConfig{
			MemberEntries: []string{"octocat"},
		})

		want := true
		if ret := it.IsEntryMember(context.Background(), "octocat"); ret != want {
			t.Fatalf("IsEntryMember = %v, wanted %v", ret, want)
		}
	})
}

func TestInteractor_IsOrgMember(t *testing.T) {
	t.Run("Return false when the org is not included.", func(t *testing.T) {
		it := i.NewInteractor(&i.InteractorConfig{
			MemberEntries: []string{"gitploy-io"},
		})

		want := false
		if ret := it.IsOrgMember(context.Background(), []string{"github"}); ret != want {
			t.Fatalf("IsEntryMember = %v, wanted %v", ret, want)
		}
	})

	t.Run("Return true when the org is included.", func(t *testing.T) {
		it := i.NewInteractor(&i.InteractorConfig{
			MemberEntries: []string{"gitploy-io"},
		})

		want := true
		if ret := it.IsOrgMember(context.Background(), []string{"gitploy-io"}); ret != want {
			t.Fatalf("IsEntryMember = %v, wanted %v", ret, want)
		}
	})
}
