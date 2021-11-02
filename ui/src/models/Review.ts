import User from "./User"
import Deployment from "./Deployment"

export interface Review {
    id: number
    status: ReviewStatusEnum
    createdAt: Date
    updatedAt: Date
    user: User | null
    deployment: Deployment | null
}

export enum ReviewStatusEnum {
    Pending = "pending",
    Approved = "approved",
    Rejected = "rejected"
}