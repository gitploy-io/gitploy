import User from "./User"

export default interface Lock {
    id: number
    env: string
    expiredAt?: Date
    createdAt: Date
    user?: User
}