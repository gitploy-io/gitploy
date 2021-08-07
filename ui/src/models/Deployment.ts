import User from "./User"
import Repo from "./Repo"

export default interface Deployment {
    id: number
    number: number
    type: DeploymentType
    ref: string
    sha: string
    env: string
    lastStatus: LastDeploymentStatus
    uid: number
    isRollback: boolean 
    isApprovalEanbled: boolean
    requiredApprovalCount: number
    createdAt: Date
    updatedAt: Date
    deployer: User | null
    repo: Repo | null
    statuses: DeploymentStatus[]
}

export enum DeploymentType {
    Commit = "commit",
    Branch = "branch",
    Tag = "tag"
}

export enum LastDeploymentStatus {
    Waiting = "waiting",
    Created = "created",
    Running = "running",
    Success = "success",
    Failure = "failure",
    Canceled = "canceled"
}

export interface DeploymentStatus {
    id: number
    status: string
    description: string
    logUrl: string
    createdAt: string
    updatedAt: string
}