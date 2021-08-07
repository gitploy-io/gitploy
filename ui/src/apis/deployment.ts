import { StatusCodes } from "http-status-codes"

import { instance, headers } from "./setting"
import { _fetch } from "./_base"
import { 
    User,
    Repo,
    Deployment, 
    DeploymentType, 
    DeploymentStatus,
    LastDeploymentStatus, 
    HttpRequestError, 
    HttpUnprocessableEntityError 
} from "../models"
import { UserData, mapDataToUser } from "./user"
import { RepoData, mapDataToRepo } from "./repo"

export interface DeploymentData {
    id: number
    number: number
    type: string
    ref: string
    sha: string
    env: string
    status: string
    uid: number
    is_rollback: boolean
    is_approval_enabled: boolean
    required_approval_count: number
    auto_deploy: boolean
    created_at: string
    updated_at: string
    edges: {
        user: UserData,
        repo: RepoData,
        deployment_statuses: DeploymentStatusData[]
    }
}

interface DeploymentStatusData {
    id: number
    status: string
    description: string
    log_url: string
    created_at: string
    updated_at: string
}

export const mapDataToDeployment = (data: DeploymentData): Deployment => {
    let deployer: User | null = null
    let repo: Repo | null = null
    let statuses: DeploymentStatus[] = []

    if ("user" in data.edges) {
        deployer = mapDataToUser(data.edges.user)
    }

    if ("repo" in data.edges) {
        repo = mapDataToRepo(data.edges.repo) 
    }

    if ("deployment_statuses" in data.edges) {
        statuses =  data.edges.deployment_statuses
            .map((data: any) => mapDataToDeploymentStatus(data))
    }

    return {
        id: data.id,
        number: data.number,
        type: mapDeploymentType(data.type),
        ref: data.ref,
        sha: data.sha,
        env: data.env,
        lastStatus: mapLastDeploymentStatus(data.status),
        uid: data.uid,
        isRollback: data.is_rollback,
        isApprovalEanbled: data.is_approval_enabled,
        requiredApprovalCount: data.required_approval_count,
        createdAt: new Date(data.created_at),
        updatedAt: new Date(data.updated_at),
        deployer,
        repo,
        statuses,
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

function mapLastDeploymentStatus(s: string) {
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
        case "canceled":
            return LastDeploymentStatus.Canceled
        default:
            return LastDeploymentStatus.Waiting
    }
}

function mapDataToDeploymentStatus(data: any): DeploymentStatus {
    return {
        id: data.id,
        status: data.status,
        description: data.description,
        logUrl: data.log_url,
        createdAt: data.created_at,
        updatedAt: data.updated_at,
    }
}

export const listDeployments = async (repoId: string, env: string, status: string, page: number, perPage: number): Promise<Deployment[]> => {
    const deployments: Deployment[] = await _fetch(`${instance}/api/v1/repos/${repoId}/deployments?env=${env}&status=${status}&page=${page}&per_page=${perPage}`, {
        headers,
        credentials: 'same-origin',
    })
        .then(response => response.json())
        .then(ds => ds.map((d: any): Deployment => mapDataToDeployment(d)))

    return deployments
}

export const getDeployment = async (id: string, number: number): Promise<Deployment> => {
    const deployment = await _fetch(`${instance}/api/v1/repos/${id}/deployments/${number}`, {
        headers,
        credentials: 'same-origin',
    })
        .then(response => response.json())
        .then(data => mapDataToDeployment(data))

    return deployment
}

export const createDeployment = async (repoId: string, type: DeploymentType = DeploymentType.Commit, ref: string, env: string): Promise<Deployment> => {
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

export const updateDeploymentStatusCreated = async (id: string, number: number): Promise<Deployment> => {
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

export const rollbackDeployment = async (repoId: string, number: number): Promise<Deployment> => {
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
