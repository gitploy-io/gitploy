package vo

import (
	"testing"
	"time"
)

func TestLicense_IsOverLimit(t *testing.T) {
	t.Run("Return false when the count of member is over the limit.", func(t *testing.T) {
		l := NewTrialLicense(6)

		expected := true
		if finished := l.IsOverLimit(); finished != expected {
			t.Fatalf("IsOverLimit = %v, wanted %v", finished, expected)
		}
	})

	t.Run("Return true when the count of member is under the limit.", func(t *testing.T) {
		l := NewTrialLicense(5)

		expected := false
		if finished := l.IsOverLimit(); finished != expected {
			t.Fatalf("IsOverLimit = %v, wanted %v", finished, expected)
		}
	})
}

func TestLicense_IsExpired(t *testing.T) {
	t.Run("Return true when the license is expired.", func(t *testing.T) {
		tm := time.Now()

		t.Log("Build the license with the expired time.")
		l := NewStandardLicense(5, &SigningData{
			MemberLimit: 20,
			ExpiredAt:   tm.Add(-24 * time.Hour),
		})

		expected := true
		if expired := l.IsExpired(); expired != expected {
			t.Fatalf("IsExpired = %v, wanted %v", expired, expected)
		}
	})

	t.Run("Return false when the license is not expired.", func(t *testing.T) {
		tm := time.Now()

		l := NewStandardLicense(5, &SigningData{
			MemberLimit: 20,
			ExpiredAt:   tm.Add(time.Hour),
		})

		expected := false
		if expired := l.IsExpired(); expired != expected {
			t.Fatalf("IsExpired = %v, wanted %v", expired, expected)
		}
	})
}
