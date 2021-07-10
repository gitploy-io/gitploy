import Repo from './Repo'

export default interface Deployment {
    id: number
    number: number
    type: DeploymentType
    ref: string
    sha: string
    env: string
    status: DeploymentStatus
    uid: number
    requiredApprovalCount: number
    autoDeploy: boolean
    createdAt: Date
    updatedAt: Date
    deployer: Deployer | null
    repo: Repo | null
}

export interface Deployer {
    id: string
    login: string
    avatar: string
}

export enum DeploymentType {
    Commit = "commit",
    Branch = "branch",
    Tag = "tag"
}

export enum DeploymentStatus {
    Waiting = "waiting",
    Created = "created",
    Running = "running",
    Success = "success",
    Failure = "failure"
}