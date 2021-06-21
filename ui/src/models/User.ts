export default interface User {
    id: string
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