import { instance, headers } from './config'
import { Repo } from '../models'

export const listRepos = async (q: string, page: number = 1, perPage: number = 30) => {
    let repos:Repo[]

    try {
        repos = await fetch(`${instance}/v1/repos?q=${q}&page=${page}&per_page=${perPage}`, {
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
        return [[], e]
    }

    return [repos, null]
}