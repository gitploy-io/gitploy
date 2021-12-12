// +build oss

package interactor

import (
	"context"

	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/pkg/e"
)

func (i *Interactor) requestReviewByLogin(ctx context.Context, d *ent.Deployment, login string) (*ent.Review, error) {
	return nil, e.NewError(e.ErrorCodeLicenseRequired, nil)
}
