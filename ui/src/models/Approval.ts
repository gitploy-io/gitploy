import User from "./User"
import Deployment from "./Deployment"

export default interface Approval {
    id: number
    status: ApprovalStatus
    createdAt: Date
    updatedAt: Date
    user: User | null
    deployment: Deployment | null
}

export enum ApprovalStatus {
    Pending = "pending",
    Approved = "approved",
    Declined = "declined"
}