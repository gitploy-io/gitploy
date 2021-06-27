import Deployment from "./Deployment"

export default interface Notification {
    id: number,
    type: NotificationType,
    notified: boolean,
    createdAt: Date,
    updatedAt: Date,
    deployment: Deployment | null,
}

export enum NotificationType {
    Deployment = "deployment"
}