import { StatusCodes } from 'http-status-codes'
import { HttpForbiddenError } from '../models/errors'

import { instance, headers } from './settings'
import { Repo, RepoPayload } from '../models'

const mapRepo = (r: any): Repo => {
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

export const listRepos = async (q: string, page: number = 1, perPage: number = 30) => {
    const repos = await fetch(`${instance}/api/v1/repos?q=${q}&page=${page}&per_page=${perPage}`, {
        headers,
        credentials: 'same-origin',
    })
        .then(response => response.json())
        .then(repos => repos.map((r: any): Repo => (mapRepo(r))))

    return repos
}

export const searchRepo = async (namespace: string, name: string) => {
    const repo = await fetch(`${instance}/api/v1/repos/search?namespace=${namespace}&name=${name}`, {
        headers,
        credentials: 'same-origin',
    })
        .then(response => response.json())
        .then((repo: any) => (mapRepo(repo)))

    return repo
}

export const updateRepo = async (id: string, payload: RepoPayload) => {
    const body = {
        "config_path": payload.configPath
    }
    const repo = await fetch(`${instance}/api/v1/repos/${id}`, {
        headers,
        credentials: 'same-origin',
        method: "PATCH",
        body: JSON.stringify(body)
    })
        .then(response => response.json())
        .then((repo: any) => (mapRepo(repo)))

    return repo
}

export const activateRepo = async (id: string) => {
    const response = await fetch(`${instance}/api/v1/repos/${id}/activate`, {
        headers,
        credentials: 'same-origin',
        method: "PATCH",
    })

    if (response.status === StatusCodes.FORBIDDEN) {
        throw new HttpForbiddenError("Only admin permssion can access.")
    }

    const repo = await response
        .json()
        .then((r:any) => mapRepo(r))
    return repo
}

export const deactivateRepo = async (id: string) => {
    const response = await fetch(`${instance}/api/v1/repos/${id}/deactivate`, {
        headers,
        credentials: 'same-origin',
        method: "PATCH",
    })

    if (response.status === StatusCodes.FORBIDDEN) {
        throw new HttpForbiddenError("Only admin permssion can access.")
    }

    const repo = await response
        .json()
        .then((r:any) => mapRepo(r))
    return repo
}
