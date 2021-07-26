import { _fetch } from "./_base"
import { instance, headers } from "./setting"
import { User, ChatUser } from "../models"

export interface UserData {
    id: string
    login: string
    avatar: string
    admin: boolean
    created_at: string
    updated_at: string
    edges: {
        chat_user: ChatUserData
    }
}

interface ChatUserData {
    id: string
    created_at: string
    updated_at: string
}

export const mapDataToUser = (data: UserData): User => {
    let cu:ChatUser | null
    if (data.edges.chat_user) {
        const chat_user = data.edges.chat_user

        cu = {
            id: chat_user.id,
            createdAt: new Date(chat_user.created_at),
            updatedAt: new Date(chat_user.updated_at),
        }
    } else  {
        cu = null
    }

    return {
        id: data.id,
        login: data.login,
        avatar: data.avatar,
        admin: data.admin,
        createdAt: new Date(data.created_at),
        updatedAt: new Date(data.updated_at),
        chatUser: cu,
    }
}

export const getMe = async (): Promise<User> => {
    const user:User = await _fetch(`${instance}/api/v1/user`, {
        headers,
        credentials: "same-origin",
    })
        .then(response => response.json())
        .then(data => (mapDataToUser(data)))

    return user
}