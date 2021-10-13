package interactor

import (
	"context"
	"testing"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/internal/interactor/mock"
	"github.com/golang/mock/gomock"
)

func TestInteractor_ProduceDeploymentStatisticsOfRepo(t *testing.T) {
	t.Run("Create the statistics when the record is not found.", func(t *testing.T) {
		input := struct {
			repo *ent.Repo
			d    *ent.Deployment
		}{
			repo: &ent.Repo{
				Namespace: "octocat",
				Name:      "HelloWorld",
			},
			d: &ent.Deployment{
				Env: "production",
			},
		}

		ctrl := gomock.NewController(t)
		store := mock.NewMockStore(ctrl)
		scm := mock.NewMockSCM(ctrl)

		t.Log("MOCK - Find the deployment_statistics by the environment.")
		store.
			EXPECT().
			FindDeploymentStatisticsOfRepoByEnv(gomock.Any(), gomock.Eq(input.repo), gomock.Eq(input.d.Env)).
			Return(nil, &ent.NotFoundError{})

		t.Log("MOCK - Create a new statistics.")
		store.
			EXPECT().
			CreateDeploymentStatistics(gomock.Any(), gomock.AssignableToTypeOf(&ent.DeploymentStatistics{})).
			Return(&ent.DeploymentStatistics{ID: 1}, nil)

		i := newMockInteractor(store, scm)

		_, err := i.ProduceDeploymentStatisticsOfRepo(context.Background(), input.repo, input.d)
		if err != nil {
			t.Fatalf("ProduceDeploymentStatisticsOfRepo returns an error: %s", err)
		}
	})

	t.Run("Increate the rollback_count when the deployment is rollback.", func(t *testing.T) {
		input := struct {
			repo *ent.Repo
			d    *ent.Deployment
		}{
			repo: &ent.Repo{
				Namespace: "octocat",
				Name:      "HelloWorld",
			},
			d: &ent.Deployment{
				Env:        "production",
				IsRollback: true,
			},
		}

		ctrl := gomock.NewController(t)
		store := mock.NewMockStore(ctrl)
		scm := mock.NewMockSCM(ctrl)

		t.Log("MOCK - Find the deployment_statistics by the environment.")
		store.
			EXPECT().
			FindDeploymentStatisticsOfRepoByEnv(gomock.Any(), gomock.Eq(input.repo), gomock.Eq(input.d.Env)).
			Return(&ent.DeploymentStatistics{ID: 1, RollbackCount: 1}, nil)

		t.Log("MOCK - Increase the rollback_count.")
		store.
			EXPECT().
			UpdateDeploymentStatistics(gomock.Any(), gomock.Eq(&ent.DeploymentStatistics{ID: 1, RollbackCount: 2})).
			Return(&ent.DeploymentStatistics{ID: 1, RollbackCount: 2}, nil)

		i := newMockInteractor(store, scm)

		_, err := i.ProduceDeploymentStatisticsOfRepo(context.Background(), input.repo, input.d)
		if err != nil {
			t.Fatalf("ProduceDeploymentStatisticsOfRepo returns an error: %s", err)
		}
	})
}
