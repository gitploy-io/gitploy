export default interface Notification {
    id: number
    type: NotificationType
    repo: {
        namespace: string
        name: string
    }
    deployment: {
        number: number
        type: string
        ref: string
        env: string
        status: string
        login: string
    }
    approval: {
        status: string
        login: string
    }
    notified: boolean
    checked: boolean
    createdAt: Date
    updatedAt: Date
}

export enum NotificationType {
    Deployment = "deployment",
    ApprovalRequested = "approval_requested",
    ApprovalResponded = "approval_responded" 
}