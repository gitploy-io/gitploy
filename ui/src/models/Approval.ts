import User from "./User"
import Deployment from "./Deployment"

export default interface Approval {
    isApproved: boolean
    createdAt: Date
    updatedAt: Date
    user: User | null
    deployment: Deployment | null
}