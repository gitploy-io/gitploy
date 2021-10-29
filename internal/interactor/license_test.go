package interactor

import (
	"context"
	"testing"

	"github.com/gitploy-io/gitploy/internal/interactor/mock"
	"github.com/gitploy-io/gitploy/vo"
	"github.com/golang/mock/gomock"
)

func TestStore_GetLicense(t *testing.T) {
	t.Run("Return the trial license when the signing data is nil.", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		store := mock.NewMockStore(ctrl)

		t.Log("MOCK - return the count of users.")
		store.
			EXPECT().
			CountUsers(gomock.AssignableToTypeOf(context.Background())).
			Return(vo.TrialMemberLimit, nil)

		store.
			EXPECT().
			CountDeployments(gomock.AssignableToTypeOf(context.Background())).
			Return(vo.TrialDeploymentLimit, nil)

		i := &Interactor{Store: store}

		lic, err := i.GetLicense(context.Background())
		if err != nil {
			t.Fatalf("GetLicense returns an error: %s", err)
		}

		if !lic.IsTrial() {
			t.Fatalf("GetLicense = %v, wanted %v", lic.Kind, vo.LicenseKindTrial)
		}
	})
}
