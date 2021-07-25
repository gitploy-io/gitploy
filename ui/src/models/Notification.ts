export default interface Notification {
    id: number
    type: NotificationType
    repoNamespace: string
    repoName: string
    deploymentNumber: number
    deploymentType: string
    deploymentRef: string
    deploymentEnv: string
    deploymentStatus: string
    deploymentLogin: string
    approvalStatus: string
    approvalLogin: string
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