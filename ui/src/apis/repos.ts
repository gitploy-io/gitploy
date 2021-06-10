import { instance, headers } from './settings'
import { Repo } from '../models'

export const listRepos = async (q: string, page: number = 1, perPage: number = 30) => {
    let repos:Repo[]

    try {
        repos = await fetch(`${instance}/api/v1/repos?q=${q}&page=${page}&per_page=${perPage}`, {
            headers,
            credentials: 'same-origin',
        })
            .then(response => response.json())
            .then(repos => repos.map((r: any): Repo => {
                return {
                    id: r.id,
                    namespace: r.namespace,
                    name: r.name,
                    description: r.description, 
                    syncedAt: new Date(r.synced_at),
                    createdAt: new Date(r.created_at),
                    updatedAt: new Date(r.updated_at),
                }
            }))
    } catch (e) {
        throw e
    }

    return repos
}

export const searchRepo = async (namespace: string, name: string) => {
    let repo:Repo

    try {
        repo = await fetch(`${instance}/api/v1/repos/search?namespace=${namespace}&name=${name}`, {
            headers,
            credentials: 'same-origin',
        })
            .then(response => response.json())
            .then((repo: any) => {
                return {
                    id: repo.id,
                    namespace: repo.namespace,
                    name: repo.name,
                    description: repo.description, 
                    syncedAt: new Date(repo.synced_at),
                    createdAt: new Date(repo.created_at),
                    updatedAt: new Date(repo.updated_at),
                }
            })
    } catch (e) {
        throw e
    }

    return repo
}