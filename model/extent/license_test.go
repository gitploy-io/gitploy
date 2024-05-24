package extent

import (
	"testing"
	"time"
)

func TestLicense_IsOverLimit(t *testing.T) {
	t.Run("Return false when the license is OSS.", func(t *testing.T) {
		l := NewOSSLicense()

		expected := false
		if finished := l.IsOverLimit(); finished != expected {
			t.Fatalf("IsOverLimit = %v, wanted %v", finished, expected)
		}
	})

	t.Run("Return false when the trial license is less than or equal to the limit.", func(t *testing.T) {
		l := NewTrialLicense(TrialMemberLimit, TrialDeploymentLimit)

		expected := false
		if finished := l.IsOverLimit(); finished != expected {
			t.Fatalf("IsOverLimit = %v, wanted %v", finished, expected)
		}
	})

	t.Run("Return true when the standard license is less than the limit.", func(t *testing.T) {
		sl := NewStandardLicense(20, &SigningData{
			MemberLimit: 20,
			ExpiredAt:   time.Now(),
		})

		if finished := sl.IsOverLimit(); finished != false {
			t.Fatalf("IsOverLimit = %v, wanted %v", finished, false)
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
