import { instance, headers } from './setting'
import { _fetch } from "./_base"
import { mapUser } from "./user"
import { mapRepo } from "./repo"
import { Repo, Perm } from '../models'

const mapDataToPerm = (data: any): Perm => {
    return {
        repoPerm: data.repo_perm,
        syncedAt: new Date(data.synced_at),
        createdAt: new Date(data.created_at),
        updatedAt: new Date(data.updated_at),
        user: mapUser(data.edges.user),
        repo: mapRepo(data.edges.repo),
    }
}

export const listPerms = async (repo: Repo, q: string, page: number = 1, perPage: number = 30) => {
    const perms: Perm[] = await _fetch(`${instance}/api/v1/repos/${repo.id}/perms?q=${q}`, {
        headers,
        credentials: "same-origin"
    })
        .then(res => res.json())
        .then(data => data.map((d: any) => mapDataToPerm(d)))
    
    return perms
}