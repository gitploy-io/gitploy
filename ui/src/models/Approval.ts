import User from "./User"
import Deployment from "./Deployment"

export default interface Approval {
    id: number
    isApproved: boolean
    createdAt: Date
    updatedAt: Date
    user: User | null
    deployment: Deployment | null
}