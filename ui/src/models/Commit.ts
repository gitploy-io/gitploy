export default interface Commit {
    sha: string
    message: string
    isPullRequest: boolean
}

export interface Status {
    context: string
    avatarUrl: string
    targetUrl: string
    state: StatusState
}

export enum StatusState {
    Null = "null",
    Pending = "pending",
    Success = "success",
    Failure = "failure",
}