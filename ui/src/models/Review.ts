import User from "./User"
import Deployment from "./Deployment"

export interface Review {
    id: number
    status: ReviewStatusEnum
    createdAt: Date
    updatedAt: Date
    user?: User 
    deployment?: Deployment 
}

export enum ReviewStatusEnum {
    Pending = "pending",
    Approved = "approved",
    Rejected = "rejected"
}