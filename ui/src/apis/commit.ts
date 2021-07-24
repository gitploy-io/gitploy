import { StatusCodes } from 'http-status-codes'

import { instance, headers } from './setting'
import { _fetch } from "./_base"
import { Commit, Status, HttpNotFoundError, StatusState } from '../models'

interface CommitData {
    sha: string
    message: string
    is_pull_request: boolean
}

const mapDataToCommit = (data: CommitData): Commit => {
    return {
        sha: data.sha,
        message: data.message,
        isPullRequest: data.is_pull_request
    }
}

export const listCommits = async (repoId: string, branch: string, page = 1, perPage = 30): Promise<Commit[]> => {
    const commits: Commit[] = await _fetch(`${instance}/api/v1/repos/${repoId}/commits?branch=${branch}&page=${page}&per_page=${perPage}`, {
        headers,
        credentials: "same-origin",
    })
        .then(response => response.json())
        .then(commits => commits.map((c: CommitData) => mapDataToCommit(c)))
    
    return commits
}

export const getCommit = async (repoId: string, sha: string): Promise<Commit> => {
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
        .then((c: CommitData) => mapDataToCommit(c))
    
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

export const listStatuses = async (repoId: string, sha: string): Promise<{state: StatusState, statuses: Status[]}> => {
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