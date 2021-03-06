package interactor_test

import (
	"context"
	"testing"

	i "github.com/gitploy-io/gitploy/internal/interactor"
	"github.com/gitploy-io/gitploy/internal/interactor/mock"
	"github.com/gitploy-io/gitploy/model/extent"
	"github.com/golang/mock/gomock"
)

func TestInteractor_GetLicense(t *testing.T) {
	t.Run("Return the trial license when the signing data is nil.", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		store := mock.NewMockStore(ctrl)

		t.Log("MOCK - return the count of users.")
		store.
			EXPECT().
			CountUsers(gomock.AssignableToTypeOf(context.Background())).
			Return(extent.TrialMemberLimit, nil)

		store.
			EXPECT().
			CountDeployments(gomock.AssignableToTypeOf(context.Background())).
			Return(extent.TrialDeploymentLimit, nil)

		it := i.NewInteractor(&i.InteractorConfig{
			Store: store,
		})

		lic, err := it.GetLicense(context.Background())
		if err != nil {
			t.Fatalf("GetLicense returns an error: %s", err)
		}

		if !lic.IsTrial() {
			t.Fatalf("GetLicense = %v, wanted %v", lic.Kind, extent.LicenseKindTrial)
		}
	})
}
