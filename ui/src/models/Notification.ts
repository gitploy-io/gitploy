export default interface Notification {
    id: number,
    type: NotificationType,
    resourceId: number,
    notified: boolean,
    createdAt: Date,
    updatedAt: Date,
}

export enum NotificationType {
    Deployment = "deployment"
}