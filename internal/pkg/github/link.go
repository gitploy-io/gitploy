package github

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gitploy-io/gitploy/model/ent"
	"github.com/gitploy-io/gitploy/pkg/e"
)

func (g *Github) GetConfigRedirectURL(ctx context.Context, u *ent.User, r *ent.Repo) (string, error) {
	remote, res, err := g.Client(ctx, u.Token).
		Repositories.
		Get(ctx, r.Namespace, r.Name)
	if res.StatusCode == http.StatusForbidden {
		return "", e.NewError(e.ErrorPermissionRequired, err)
	} else if res.StatusCode == http.StatusNotFound {
		return "", e.NewError(e.ErrorCodeEntityNotFound, err)
	} else if err != nil {
		return "", e.NewError(e.ErrorCodeInternalError, err)
	}

	// The latest version file on the main branch.
	// https://docs.github.com/en/repositories/working-with-files/using-files/getting-permanent-links-to-files
	url := fmt.Sprintf("%s/blob/%s/%s", *remote.HTMLURL, *remote.DefaultBranch, r.ConfigPath)
	return url, nil
}

func (g *Github) GetNewFileRedirectURL(ctx context.Context, u *ent.User, r *ent.Repo) (string, error) {
	remote, res, err := g.Client(ctx, u.Token).
		Repositories.
		Get(ctx, r.Namespace, r.Name)
	if res.StatusCode == http.StatusForbidden {
		return "", e.NewError(e.ErrorPermissionRequired, err)
	} else if res.StatusCode == http.StatusNotFound {
		return "", e.NewError(e.ErrorCodeEntityNotFound, err)
	} else if err != nil {
		return "", e.NewError(e.ErrorCodeInternalError, err)
	}

	// Redirect to the URL to create a configuration file.
	// https://docs.github.com/en/enterprise-server@3.0/repositories/working-with-files/managing-files/creating-new-files
	url := fmt.Sprintf("%s/new/%s/%s", *remote.HTMLURL, *remote.DefaultBranch, r.ConfigPath)
	return url, nil
}
