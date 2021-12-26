export default interface Commit {
    sha: string
    message: string
    isPullRequest: boolean
    htmlUrl: string
    author?: Author
}

export interface Author {
    login: string
    avatarUrl: string
    date: Date
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
    Cancelled = "cancelled",
    Skipped = "skipped"
}