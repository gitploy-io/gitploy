import { instance, headers } from './settings'
import { Branch } from '../models'

export const listBranches = async (repoId: string, page: number = 1, perPage: number = 30) => {
    const branches: Branch[] = await fetch(`${instance}/v1/repos/${repoId}/branches?page=${page}&per_page=${perPage}`, {
        headers,
        credentials: "same-origin",
    })
        .then(response => response.json())
        .then(branches => branches.map((b: any): Branch => {
            return {
                name: b.name,
                commitSha: b.commit_sha
            } 
        }))
    
    return branches
}