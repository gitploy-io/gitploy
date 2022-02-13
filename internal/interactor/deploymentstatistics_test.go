package interactor_test

import (
	"context"
	"testing"

	i "github.com/gitploy-io/gitploy/internal/interactor"
	"github.com/gitploy-io/gitploy/internal/interactor/mock"
	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/model/extent"
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

		t.Log("MOCK - The prev deployment is not found.")
		store.
			EXPECT().
			FindPrevSuccessDeployment(gomock.Any(), gomock.AssignableToTypeOf(&ent.Deployment{})).
			Return(nil, &ent.NotFoundError{})

		t.Log("MOCK - Update the statistics.")
		store.
			EXPECT().
			UpdateDeploymentStatistics(gomock.Any(), gomock.Eq(&ent.DeploymentStatistics{ID: 1, Count: 1})).
			DoAndReturn(func(ctx context.Context, s *ent.DeploymentStatistics) (*ent.DeploymentStatistics, error) {
				return s, nil
			})

		it := i.NewInteractor(&i.InteractorConfig{
			Store: store,
			SCM:   scm,
		})

		_, err := it.ProduceDeploymentStatisticsOfRepo(context.Background(), input.repo, input.d)
		if err != nil {
			t.Fatalf("ProduceDeploymentStatisticsOfRepo returns an error: %s", err)
		}
	})

	t.Run("Calculate changes from the lastest deployment.", func(t *testing.T) {
		input := struct {
			repo *ent.Repo
			d    *ent.Deployment
		}{
			repo: &ent.Repo{
				Namespace: "octocat",
				Name:      "HelloWorld",
			},
			d: &ent.Deployment{
				ID:  2,
				Env: "production",
				Edges: ent.DeploymentEdges{
					User: &ent.User{},
				},
			},
		}

		ctrl := gomock.NewController(t)
		store := mock.NewMockStore(ctrl)
		scm := mock.NewMockSCM(ctrl)

		t.Log("MOCK - Find the deployment_statistics by the environment.")
		store.
			EXPECT().
			FindDeploymentStatisticsOfRepoByEnv(gomock.Any(), gomock.Eq(input.repo), gomock.Eq(input.d.Env)).
			Return(&ent.DeploymentStatistics{ID: 1, Count: 1}, nil)

		t.Log("MOCK - Find the latest deployment.")
		store.
			EXPECT().
			FindPrevSuccessDeployment(gomock.Any(), gomock.Eq(input.d)).
			Return(&ent.Deployment{ID: 2}, nil)

		t.Log("MOCK - Get changed commits from the SCM.")
		scm.
			EXPECT().
			CompareCommits(gomock.Any(), gomock.AssignableToTypeOf(&ent.User{}), gomock.AssignableToTypeOf(&ent.Repo{}), gomock.Any(), gomock.Any(), gomock.Any()).
			Return([]*extent.Commit{}, []*extent.CommitFile{
				{Additions: 1, Deletions: 1, Changes: 2},
			}, nil)

		t.Log("MOCK - Update the statistics.")
		store.
			EXPECT().
			UpdateDeploymentStatistics(gomock.Any(), gomock.Eq(&ent.DeploymentStatistics{ID: 1, Count: 2, Additions: 1, Deletions: 1, Changes: 2})).
			DoAndReturn(func(ctx context.Context, s *ent.DeploymentStatistics) (*ent.DeploymentStatistics, error) {
				return s, nil
			})

		it := i.NewInteractor(&i.InteractorConfig{
			Store: store,
			SCM:   scm,
		})

		_, err := it.ProduceDeploymentStatisticsOfRepo(context.Background(), input.repo, input.d)
		if err != nil {
			t.Fatalf("ProduceDeploymentStatisticsOfRepo returns an error: %s", err)
		}
	})

}
