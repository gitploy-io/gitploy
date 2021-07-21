import { StatusCodes } from 'http-status-codes'

import { instance, headers } from './setting'
import { _fetch } from "./_base"
import { Deployment, DeploymentType, LastDeploymentStatus, HttpRequestError, HttpUnprocessableEntityError } from '../models'
import { Deployer } from '../models/Deployment'
import Repo from '../models/Repo'
import { mapRepo } from './repo'

export const listDeployments = async (repoId: string, env: string, status: string, page: number, perPage: number) => {
    let deployments:Deployment[]

    deployments = await _fetch(`${instance}/api/v1/repos/${repoId}/deployments?env=${env}&status=${status}&page=${page}&per_page=${perPage}`, {
        headers,
        credentials: 'same-origin',
    })
        .then(response => response.json())
        .then(ds => ds.map((d: any): Deployment => mapDataToDeployment(d)))

    return deployments
}

export const getDeployment = async (id: string, number: number) => {
    const deployment = await _fetch(`${instance}/api/v1/repos/${id}/deployments/${number}`, {
        headers,
        credentials: 'same-origin',
    })
        .then(response => response.json())
        .then(data => mapDataToDeployment(data))

    return deployment
}

export const createDeployment = async (repoId: string, type: DeploymentType = DeploymentType.Commit, ref: string, env: string) => {
    const body = JSON.stringify({
        type,
        ref,
        env
    })
    const response = await _fetch(`${instance}/api/v1/repos/${repoId}/deployments`, {
        headers,
        credentials: 'same-origin',
        method: "POST",
        body: body,
    })
    if (response.status !== StatusCodes.CREATED) {
        throw new HttpRequestError(response.status, "It has failed to deploy.")
    } 

    const deployment = response
        .json()
        .then(d => mapDataToDeployment(d))
    return deployment
}

export const updateDeploymentStatusCreated = async (id: string, number: number) => {
    const body = JSON.stringify({
        status: "created"
    })
    const response = await _fetch(`${instance}/api/v1/repos/${id}/deployments/${number}`, {
        headers,
        credentials: 'same-origin',
        method: "PATCH",
        body: body,
    })
    if (response.status === StatusCodes.UNPROCESSABLE_ENTITY) {
        const { message } = await response.json()
        throw new HttpUnprocessableEntityError(message)
    } 

    const deployment = response
        .json()
        .then(d => mapDataToDeployment(d))
    return deployment
}

export const rollbackDeployment = async (repoId: string, number: number) => {
    const response = await _fetch(`${instance}/api/v1/repos/${repoId}/deployments/${number}/rollback`, {
        headers,
        credentials: 'same-origin',
        method: "POST",
    })
    if (response.status !== StatusCodes.CREATED) {
        throw new HttpRequestError(response.status, "It has failed to rollback.")
    }

    const deployment = response
        .json()
        .then(d => mapDataToDeployment(d))
    return deployment
}

export function mapDataToDeployment(d: any): Deployment {
    let deployer: Deployer | null = null
    let repo: Repo | null = null

    if ("user" in d.edges) {
        const ud = d.edges.user
        deployer = {
            id: ud.id,
            login: ud.login,
            avatar: ud.avatar
        }
    }

    if ("repo" in d.edges) {
        const rd = d.edges.repo
        repo = mapRepo(rd) 
    }

    return {
        id: d.id,
        number: d.number,
        type: mapDeploymentType(d.type),
        ref: d.ref,
        sha: d.sha,
        env: d.env,
        status: mapDeploymentStatus(d.status),
        uid: d.uid? d.uid : "",
        isRollback: d.is_rollback,
        isApprovalEanbled: d.is_approval_enabled,
        requiredApprovalCount: d.required_approval_count,
        autoDeploy: d.auto_deploy,
        createdAt: new Date(d.created_at),
        updatedAt: new Date(d.updatedAt),
        deployer,
        repo,
    }
}

function mapDeploymentType(t: string) {
    switch (t) {
        case "commit":
            return DeploymentType.Commit
        case "branch":
            return DeploymentType.Branch
        case "tag":
            return DeploymentType.Tag
        default:
            return DeploymentType.Commit
    }
}

function mapDeploymentStatus(s: string) {
    switch (s) {
        case "waiting":
            return LastDeploymentStatus.Waiting
        case "created":
            return LastDeploymentStatus.Created
        case "running":
            return LastDeploymentStatus.Running
        case "success":
            return LastDeploymentStatus.Success
        case "failure":
            return LastDeploymentStatus.Failure
        default:
            return LastDeploymentStatus.Waiting
    }
}
