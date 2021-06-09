import { StatusCodes } from 'http-status-codes'

import { instance, headers } from './settings'
import { Branch, HttpNotFoundError } from '../models'

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

export const getBranch = async (repoId: string, name: string) => {
    const response = await fetch(`${instance}/v1/repos/${repoId}/branches/${name}`, {
        headers,
        credentials: "same-origin",
    })
    if (response.status === StatusCodes.NOT_FOUND) {
        const message = await response.json().then(data => data.message)
        throw new HttpNotFoundError(message)
    }

    const branch:Branch = await response
        .json()
        .then(b => ({
            name: b.name,
            commitSha: b.commit_sha
        }))
    
    return branch
}