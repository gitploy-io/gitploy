import { StatusCodes } from 'http-status-codes'
import { HttpForbiddenError } from '../models/errors'

import { instance, headers } from './setting'
import { _fetch } from "./_base"
import { Repo, RepoPayload } from '../models'

export interface RepoData {
    id: string
    namespace: string
    name: string
    description: string
    config_path: string
    active: boolean
    webhook_id: number
    synced_at: string
    created_at: string
    updated_at: string
}

// eslint-disable-next-line
export const mapDataToRepo = (data: RepoData): Repo => {
    return {
        id: data.id,
        namespace: data.namespace,
        name: data.name,
        description: data.description, 
        configPath: data.config_path,
        active: data.active,
        webhookId: data.webhook_id,
        syncedAt: new Date(data.synced_at),
        createdAt: new Date(data.created_at),
        updatedAt: new Date(data.updated_at),
    }
}

export const listRepos = async (q: string, page = 1, perPage = 30): Promise<Repo[]> => {
    const repos = await _fetch(`${instance}/api/v1/repos?q=${q}&page=${page}&per_page=${perPage}`, {
        headers,
        credentials: 'same-origin',
    })
        .then(response => response.json())
        .then(repos => repos.map((r: any): Repo => (mapDataToRepo(r))))

    return repos
}

export const searchRepo = async (namespace: string, name: string): Promise<Repo> => {
    const repo = await _fetch(`${instance}/api/v1/repos/search?namespace=${namespace}&name=${name}`, {
        headers,
        credentials: 'same-origin',
    })
        .then(response => response.json())
        .then((repo: any) => (mapDataToRepo(repo)))

    return repo
}

export const updateRepo = async (repo: Repo, payload: RepoPayload): Promise<Repo> => {
    const body = {
        "config_path": payload.configPath,
        "active": repo.active
    }
    repo = await _fetch(`${instance}/api/v1/repos/${repo.id}`, {
        headers,
        credentials: 'same-origin',
        method: "PATCH",
        body: JSON.stringify(body)
    })
        .then(response => response.json())
        .then((repo: any) => (mapDataToRepo(repo)))

    return repo
}

export const activateRepo = async (repo: Repo): Promise<Repo> => {
    const body = {
        "config_path": repo.configPath,
        "active": true,
    }
    const response = await _fetch(`${instance}/api/v1/repos/${repo.id}`, {
        headers,
        credentials: 'same-origin',
        method: "PATCH",
        body: JSON.stringify(body)
    })

    if (response.status === StatusCodes.FORBIDDEN) {
        throw new HttpForbiddenError("Only admin permssion can access.")
    }

    repo = await response
        .json()
        .then((r:any) => mapDataToRepo(r))
    return repo
}

export const deactivateRepo = async (repo: Repo): Promise<Repo> => {
    const body = {
        "config_path": repo.configPath,
        "active": false,
    }
    const response = await _fetch(`${instance}/api/v1/repos/${repo.id}`, {
        headers,
        credentials: 'same-origin',
        method: "PATCH",
        body: JSON.stringify(body)
    })

    if (response.status === StatusCodes.FORBIDDEN) {
        throw new HttpForbiddenError("Only admin permssion can access.")
    }

    repo = await response
        .json()
        .then((r:any) => mapDataToRepo(r))
    return repo
}
