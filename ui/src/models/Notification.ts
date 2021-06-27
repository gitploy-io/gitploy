import Deployment from "./Deployment"
import Repo from "./Repo"

export default interface Notification {
    id: number,
    type: NotificationType,
    notified: boolean,
    checked: boolean,
    createdAt: Date,
    updatedAt: Date,
    repo: Repo,
    deployment: Deployment | null,
}

export enum NotificationType {
    Deployment = "deployment"
}