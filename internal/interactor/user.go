package interactor

import (
	"context"
)

func (i *Interactor) IsAdminUser(ctx context.Context, login string) bool {
	for _, admin := range i.admins {
		if login == admin {
			return true
		}
	}

	return false
}

func (i *Interactor) IsEntryMember(ctx context.Context, login string) bool {
	if i.memberEntries == nil {
		return true
	}

	for _, m := range i.memberEntries {
		if login == m {
			return true
		}
	}

	return false
}

func (i *Interactor) IsOrgMember(ctx context.Context, orgs []string) bool {
	for _, o := range orgs {
		for _, entry := range i.memberEntries {
			if o == entry {
				return true
			}
		}
	}

	return false
}
