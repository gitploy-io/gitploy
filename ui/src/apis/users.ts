import { _fetch } from "./_base"
import { instance, headers } from "./settings"
import { User, ChatUser } from "../models"

const mapUser = (data: any): User => {
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

export const getMe = async () => {
    const user:User = await _fetch(`${instance}/api/v1/users/me`, {
        headers,
        credentials: "same-origin",
    })
        .then(response => response.json())
        .then(data => (mapUser(data)))

    return user
}