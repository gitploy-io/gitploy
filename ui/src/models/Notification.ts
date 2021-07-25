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
    notified: boolean
    checked: boolean
    createdAt: Date
    updatedAt: Date
}

export enum NotificationType {
    Deployment = "deployment"
}