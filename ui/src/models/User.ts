export default interface User {
    id: number
    login: string
    avatar: string
    admin: boolean
    createdAt: Date
    updatedAt: Date
    chatUser: ChatUser | null
}

export interface ChatUser {
    id: string
    createdAt: Date
    updatedAt: Date
}

export interface RateLimit {
    limit: number
    remaining: number
    reset: Date
}