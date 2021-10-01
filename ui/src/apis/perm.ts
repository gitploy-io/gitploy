import { instance, headers } from './setting'
import { _fetch } from "./_base"
import { mapDataToUser, UserData } from "./user"
import { mapDataToRepo, RepoData } from "./repo"
import { Repo, Perm } from '../models'

interface PermData{
    repo_perm: string
    synced_at: string
    created_at: string
    updated_at: string
    edges: {
        user: UserData,
        repo: RepoData,
    }
}

const mapDataToPerm = (data: PermData): Perm => {
    return {
        repoPerm: data.repo_perm,
        syncedAt: new Date(data.synced_at),
        createdAt: new Date(data.created_at),
        updatedAt: new Date(data.updated_at),
        user: mapDataToUser(data.edges.user),
        repo: mapDataToRepo(data.edges.repo),
    }
}

export const listPerms = async (namespace: string, name: string, q: string, page = 1, perPage = 30): Promise<Perm[]> => {
    const perms: Perm[] = await _fetch(`${instance}/api/v1/repos/${namespace}/${name}/perms?q=${q}&page=${page}&per_page=${perPage}`, {
        headers,
        credentials: "same-origin"
    })
        .then(res => res.json())
        .then(data => data.map((d: any) => mapDataToPerm(d)))
    
    return perms
}