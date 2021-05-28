package repos

import "github.com/hanjunlee/gitploy/ent"

func mapReposToRepoDatas(rs []*ent.Repo) []*repoData {
	rds := []*repoData{}

	for _, r := range rs {
		rds = append(rds, mapRepoToRepoData(r))
	}

	return rds
}

func mapRepoToRepoData(r *ent.Repo) *repoData {
	return &repoData{
		Repo:     r,
		FullName: r.Namespace + "/" + r.Name,
	}
}
