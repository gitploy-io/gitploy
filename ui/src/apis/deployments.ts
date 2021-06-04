import { instance, headers } from './config'
import { Deployment, DeploymentType, DeploymentStatus } from '../models'

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

export const listDeployments = async (repoId: string, env: string, page: number, perPage: number) => {
    let deployments:Deployment[]

    try {
        deployments = await fetch(`${instance}/v1/repos/${repoId}/?env=${env}&page=${page}&per_page=${perPage}`, {
            headers,
            credentials: 'same-origin',
        })
            .then(response => response.json())
            .then(ds => ds.map((d: any): Deployment => {
                return {
                    id: d.id,
                    uid: d.uid,
                    type: mapDeploymentType(d.type),
                    ref: d.ref,
                    sha: d.sha,
                    env: d.env,
                    status: mapDeploymentStatus(d.status),
                    createdAt: new Date(d.created_at),
                    updatedAt: new Date(d.updatedAt),
                }
            }))
    } catch (e) {
        throw e
    }

    return deployments
}