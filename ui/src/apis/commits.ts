import { StatusCodes } from 'http-status-codes'

import { instance, headers } from './settings'
import { Commit, HttpNotFoundError } from '../models'

export const listCommits = async (repoId: string, branch: string, page: number = 1, perPage: number = 30) => {
    const commits: Commit[] = await fetch(`${instance}/v1/repos/${repoId}/commits?branch=${branch}&page=${page}&per_page=${perPage}`, {
        headers,
        credentials: "same-origin",
    })
        .then(response => response.json())
        .then(commits => commits.map((c: any): Commit => {
            return {
                sha: c.sha,
                message: c.message,
                isPullRequest: c.is_pull_request,
            } 
        }))
    
    return commits
}

export const getCommit = async (repoId: string, sha: string) => {
    const response = await fetch(`${instance}/v1/repos/${repoId}/commits/${sha}`, {
        headers,
        credentials: "same-origin",
    })
    if (response.status === StatusCodes.NOT_FOUND) {
        const message = await response.json().then(data => data.message)
        throw new HttpNotFoundError(message)
    }

    const commit: Commit = await response
        .json()
        .then(c => ({
            sha: c.sha,
            message: c.message,
            isPullRequest: c.is_pull_request
        }))
    
    return commit
}