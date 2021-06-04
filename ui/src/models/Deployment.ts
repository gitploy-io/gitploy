export default interface Deployment {
    id: number
    uid: number
    type: DeploymentType
    ref: string
    sha: string
    env: string
    status: DeploymentStatus
    createdAt: Date
    updatedAt: Date
}

export const enum DeploymentType {
    Commit = "commit",
    Branch = "branch",
    Tag = "tag"
}

export const enum DeploymentStatus {
    Waiting = "waiting",
    Created = "created",
    Running = "running",
    Success = "success",
    Failure = "failure"
}