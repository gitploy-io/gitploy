import { instance, headers } from './config'
import { Repo } from '../models'

export const listRepos = async () => {
    let repos:Repo[]

    try {
        repos = await fetch(`${instance}/v1/repos`, {
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