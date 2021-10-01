import { StatusCodes } from 'http-status-codes'

import { instance, headers } from './setting'
import { _fetch } from "./_base"
import { DeploymentData, mapDataToDeployment } from "./deployment"

import { Repo, HttpForbiddenError, Deployment } from '../models'

export interface RepoData {
    id: string
    namespace: string
    name: string
    description: string
    config_path: string
    active: boolean
    webhook_id: number
    locked: boolean
    created_at: string
    updated_at: string
    edges: {
        deployments?: DeploymentData[]
    }
}

// eslint-disable-next-line
export const mapDataToRepo = (data: RepoData): Repo => {
    let deployments: Deployment[] | undefined

    if (typeof data.edges.deployments !== "undefined") {
        deployments = data.edges.deployments.map(data => mapDataToDeployment(data))
    }

    return {
        id: data.id,
        namespace: data.namespace,
        name: data.name,
        description: data.description, 
        configPath: data.config_path,
        active: data.active,
        webhookId: data.webhook_id,
        createdAt: new Date(data.created_at),
        updatedAt: new Date(data.updated_at),
        deployments,
    }
}

export const listRepos = async (q: string, page = 1, perPage = 30): Promise<Repo[]> => {
    const repos: Repo[] = await _fetch(`${instance}/api/v1/repos?q=${q}&sort=true&page=${page}&per_page=${perPage}`, {
        headers,
        credentials: 'same-origin',
    })
        .then(response => response.json())
        .then(repos => repos.map((r: any): Repo => (mapDataToRepo(r))))

    return repos
}

export const searchRepo = async (namespace: string, name: string): Promise<Repo> => {
    const repos: Repo[] = await _fetch(`${instance}/api/v1/repos?namespace=${namespace}&name=${name}`, {
        headers,
        credentials: 'same-origin',
    })
        .then(response => response.json())
        .then(repos => repos.map((r: any): Repo => (mapDataToRepo(r))))

    if (repos.length !== 1) {
        throw new Error(`It has failed to search the repository. The length is ${repos.length}.`)
    }

    return repos[0]
}

export const updateRepo = async (namespace: string, name: string, payload: {config_path: string}): Promise<Repo> => {
    const res = await _fetch(`${instance}/api/v1/repos/${namespace}/${name}`, {
        headers,
        credentials: 'same-origin',
        method: "PATCH",
        body: JSON.stringify(payload)
    })
    if (res.status === StatusCodes.FORBIDDEN) {
        const message = await res.json().then(data => data.message)
        throw new HttpForbiddenError(message)
    }

    const ret: Repo = await res
        .json()
        .then((repo: any) => (mapDataToRepo(repo)))

    return ret
}

export const activateRepo = async (namespace: string, name: string): Promise<Repo> => {
    const body = {
        "active": true,
    }
    const response = await _fetch(`${instance}/api/v1/repos/${namespace}/${name}`, {
        headers,
        credentials: 'same-origin',
        method: "PATCH",
        body: JSON.stringify(body)
    })
    if (response.status === StatusCodes.FORBIDDEN) {
        const message = await response.json().then(data => data.message)
        throw new HttpForbiddenError(message)
    }

    const repo = await response
        .json()
        .then((r:any) => mapDataToRepo(r))
    return repo
}

export const deactivateRepo = async (namespace: string, name: string): Promise<Repo> => {
    const body = {
        "active": false,
    }
    const response = await _fetch(`${instance}/api/v1/repos/${namespace}/${name}`, {
        headers,
        credentials: 'same-origin',
        method: "PATCH",
        body: JSON.stringify(body)
    })
    if (response.status === StatusCodes.FORBIDDEN) {
        const message = await response.json().then(data => data.message)
        throw new HttpForbiddenError(message)
    }

    const repo = await response
        .json()
        .then((r:any) => mapDataToRepo(r))
    return repo
}

export const lockRepo = async (namespace: string, name: string): Promise<Repo> => {
    const body = {
        "locked": true,
    }
    const response = await _fetch(`${instance}/api/v1/repos/${namespace}/${name}`, {
        headers,
        credentials: 'same-origin',
        method: "PATCH",
        body: JSON.stringify(body)
    })
    if (response.status === StatusCodes.FORBIDDEN) {
        const message = await response.json().then(data => data.message)
        throw new HttpForbiddenError(message)
    }

    const repo = await response
        .json()
        .then((r:any) => mapDataToRepo(r))
    return repo
}

export const unlockRepo = async (namespace: string, name: string): Promise<Repo> => {
    const body = {
        "locked": false,
    }
    const response = await _fetch(`${instance}/api/v1/repos/${namespace}/${name}`, {
        headers,
        credentials: 'same-origin',
        method: "PATCH",
        body: JSON.stringify(body)
    })
    if (response.status === StatusCodes.FORBIDDEN) {
        const message = await response.json().then(data => data.message)
        throw new HttpForbiddenError(message)
    }

    const repo = await response
        .json()
        .then((r:any) => mapDataToRepo(r))
    return repo
}