import User from "./User"

export default interface Lock {
    id: number
    env: string
    createdAt: Date
    user?: User
}