package interactor_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	i "github.com/gitploy-io/gitploy/internal/interactor"
	"github.com/gitploy-io/gitploy/internal/interactor/mock"
	"github.com/gitploy-io/gitploy/model/ent"
)

type PermMatcher struct {
	perm *ent.Perm
}

func newPermMatcher(p *ent.Perm) *PermMatcher {
	return &PermMatcher{
		perm: p,
	}
}

func (m *PermMatcher) Matches(x interface{}) bool {
	px, _ := x.(*ent.Perm)
	return px.ID == m.perm.ID
}

func (m *PermMatcher) String() string {
	return fmt.Sprintf("has same id as %d", m.perm.ID)
}

func TestPermInteractor_ResyncPerms(t *testing.T) {
	t.Run("Delete the perms.", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		store := mock.NewMockStore(ctrl)

		t.Log("Get the perms of gitploy-io organization and non-gitploy-io organization.")
		store.EXPECT().
			ListPerms(gomock.Any(), gomock.AssignableToTypeOf(&i.ListOptions{})).
			Return([]*ent.Perm{
				{
					ID: 1,
					Edges: ent.PermEdges{
						Repo: &ent.Repo{
							Namespace: "gitploy-io",
							Name:      "gitploy",
						},
					},
				},
				{
					ID: 2,
					Edges: ent.PermEdges{
						Repo: &ent.Repo{
							Namespace: "non-gitploy-io",
							Name:      "gitploy",
						},
					},
				},
			}, nil)

		store.EXPECT().
			DeletePerm(gomock.Any(), newPermMatcher(&ent.Perm{ID: 2}))

		intr := i.NewInteractor(&i.InteractorConfig{
			Store:      store,
			OrgEntries: []string{"gitploy-io"},
		})
		if err := intr.ResyncPerms(context.Background()); err != nil {
			t.Fatalf("ResyncPerms returns an error: %s", err)
		}
	})
}
