import { StatusCodes } from 'http-status-codes'
import { HttpForbiddenError } from '../models/errors'

import { instance, headers } from './setting'
import { _fetch } from "./_base"
import { Repo, RepoPayload } from '../models'

// eslint-disable-next-line
export const mapRepo = (r: any): Repo => {
    return {
        id: r.id,
        namespace: r.namespace,
        name: r.name,
        description: r.description, 
        configPath: r.config_path,
        active: r.active,
        webhookId: r.webhook_id,
        syncedAt: new Date(r.synced_at),
        createdAt: new Date(r.created_at),
        updatedAt: new Date(r.updated_at),
    }
}

export const listRepos = async (q: string, page = 1, perPage = 30): Promise<Repo[]> => {
    const repos = await _fetch(`${instance}/api/v1/repos?q=${q}&page=${page}&per_page=${perPage}`, {
        headers,
        credentials: 'same-origin',
    })
        .then(response => response.json())
        .then(repos => repos.map((r: any): Repo => (mapRepo(r))))

    return repos
}

export const searchRepo = async (namespace: string, name: string): Promise<Repo> => {
    const repo = await _fetch(`${instance}/api/v1/repos/search?namespace=${namespace}&name=${name}`, {
        headers,
        credentials: 'same-origin',
    })
        .then(response => response.json())
        .then((repo: any) => (mapRepo(repo)))

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
        .then((repo: any) => (mapRepo(repo)))

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
        .then((r:any) => mapRepo(r))
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
        .then((r:any) => mapRepo(r))
    return repo
}
