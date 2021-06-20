import { StatusCodes } from 'http-status-codes'

import { instance, headers } from './settings'
import { _fetch } from "./_base"
import { Commit, Status, HttpNotFoundError, StatusState } from '../models'

export const listCommits = async (repoId: string, branch: string, page: number = 1, perPage: number = 30) => {
    const commits: Commit[] = await _fetch(`${instance}/api/v1/repos/${repoId}/commits?branch=${branch}&page=${page}&per_page=${perPage}`, {
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
    const response = await _fetch(`${instance}/api/v1/repos/${repoId}/commits/${sha}`, {
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

const mapStatusState = (state: string) => {
    if (state === "pending") {
        return StatusState.Pending
    } else if (state === "success") {
        return StatusState.Success
    } else if (state === "failure") {
        return StatusState.Failure
    }
    return StatusState.Pending
}

export const listStatuses = async (repoId: string, sha: string) => {
    const response = await _fetch(`${instance}/api/v1/repos/${repoId}/commits/${sha}/statuses`, {
        headers,
        credentials: "same-origin",
    })
    if (response.status === StatusCodes.NOT_FOUND) {
        const message = await response.json().then(data => data.message)
        throw new HttpNotFoundError(message)
    }

    const result = await response
        .json()
        .then(d => {
            let state: StatusState
            const statuses: Status[] =  d.statuses.map((s: any) => ({
                context: s.context,
                avatarUrl: s.avatar_url,
                targetUrl: s.target_url,
                state: mapStatusState(s.state)
            }))

            if (statuses.length === 0) {
                state = StatusState.Null
            } else {
                state = mapStatusState(d.state)
            }

            return {
                state,
                statuses
            }
        })
    
    return result 
}