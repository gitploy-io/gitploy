import { StatusCodes } from 'http-status-codes'

import { instance, headers } from './settings'
import { _fetch } from "./_base"
import { Deployment, DeploymentType, DeploymentStatus, HttpRequestError } from '../models'
import { Deployer } from '../models/Deployment'

const mapDeploymentType = (t: string) => {
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

const mapDeploymentStatus = (s: string) => {
    switch (s) {
        case "waiting":
            return DeploymentStatus.Waiting
        case "created":
            return DeploymentStatus.Created
        case "running":
            return DeploymentStatus.Running
        case "success":
            return DeploymentStatus.Success
        case "failure":
            return DeploymentStatus.Failure
        default:
            return DeploymentStatus.Waiting
    }
}

export const listDeployments = async (repoId: string, env: string, status: string, page: number, perPage: number) => {
    let deployments:Deployment[]

    deployments = await _fetch(`${instance}/api/v1/repos/${repoId}/deployments?env=${env}&status=${status}&page=${page}&per_page=${perPage}`, {
        headers,
        credentials: 'same-origin',
    })
        .then(response => response.json())
        .then(ds => ds.map((d: any): Deployment => {
            const u = d.edges.user
            const user: Deployer = {
                id: u.id,
                login: u.login,
                avatar: u.avatar
            }
            return {
                id: d.id,
                uid: d.uid? d.uid : "",
                type: mapDeploymentType(d.type),
                ref: d.ref,
                sha: d.sha? d.sha : "",
                env: d.env,
                status: mapDeploymentStatus(d.status),
                createdAt: new Date(d.created_at),
                updatedAt: new Date(d.updatedAt),
                deployer: user,
            }
        }))

    return deployments
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
}