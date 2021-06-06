import { instance, headers } from './settings'
import { Commit } from '../models'

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