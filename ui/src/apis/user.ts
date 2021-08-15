import { StatusCodes } from 'http-status-codes'

import { _fetch } from "./_base"
import { instance, headers } from "./setting"
import { User, ChatUser, RateLimit, HttpRequestError, HttpForbiddenError } from "../models"

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

interface RateLimitData {
    limit: number
    remaining: number
    reset: string
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

export const mapDataToRateLimit = (data: RateLimitData): RateLimit => {
    return {
        limit: data.limit,
        remaining: data.remaining,
        reset: new Date(data.reset)
    }
}

export const listUsers = async (q: string, page = 1, perPage = 30): Promise<User[]> => {
    const res = await _fetch(`${instance}/api/v1/users?q=${q}&page=${page}&per_page=${perPage}&`, {
        headers,
        credentials: "same-origin",
    })

    if (res.status === StatusCodes.FORBIDDEN) {
        throw new HttpForbiddenError("Only admin can access.")
    }

    const users = await res
        .json()
        .then((data: UserData[]) => (data.map(d => mapDataToUser(d))))

    return users
}

export const deleteUser = async (id: string): Promise<void> => {
    const res = await _fetch(`${instance}/api/v1/users/${id}`, {
        headers,
        credentials: "same-origin",
        method: "DELETE"
    })

    if (res.status !== StatusCodes.OK) {
        const message = await res.json().then(data => data.message)
        throw new HttpRequestError(res.status, message)
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

export const getRateLimit = async (): Promise<RateLimit> => {
    const rateLimit:RateLimit = await _fetch(`${instance}/api/v1/user/rate-limit`, {
        headers,
        credentials: "same-origin",
    })
        .then(response => response.json())
        .then(data => (mapDataToRateLimit(data)))

    return rateLimit
}