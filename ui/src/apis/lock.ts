import { StatusCodes } from "http-status-codes"

import { instance, headers } from "./setting"
import { _fetch } from "./_base"
import { UserData, mapDataToUser } from "./user"
import { Repo, Lock, User, HttpForbiddenError, HttpUnprocessableEntityError } from "../models"

interface LockData {
    id: number
    env: string
    created_at: string
    edges: {
        user?: UserData
    }
}

const mapDataToLock = (data: LockData): Lock => {
    let user: User | undefined

    if (data.edges.user) {
        user = mapDataToUser(data.edges.user)
    }

    return {
        id: data.id,
        env: data.env,
        createdAt: new Date(data.created_at),
        user,
    }
}

export const listLocks = async (namespace: string, name: string): Promise<Lock[]> => {
    const locks = await _fetch(`${instance}/api/v1/repos/${namespace}/${name}/locks`, {
        headers,
        credentials: 'same-origin',
    })
        .then(res => res.json())
        .then(datas => datas.map((d: any): Lock => mapDataToLock(d)))

    return locks
}

export const lock = async (namespace: string, name: string, env: string): Promise<Lock> => {
    const res = await _fetch(`${instance}/api/v1/repos/${namespace}/${name}/locks`, {
        headers,
        credentials: 'same-origin',
        method: "POST",
        body: JSON.stringify({env})
    })

    if (res.status === StatusCodes.FORBIDDEN) {
        const {message} = await res.json()
        throw new HttpForbiddenError(message)
    } else if (res.status === StatusCodes.UNPROCESSABLE_ENTITY) {
        const {message} = await res.json()
        throw new HttpUnprocessableEntityError(message)
    }

    const lock = await res
        .json()
        .then(data => mapDataToLock(data))
    return lock
}

export const unlock = async (namespace: string, name: string, id: number): Promise<void> => {
    const res = await _fetch(`${instance}/api/v1/repos/${namespace}/${name}/locks/${id}`, {
        headers,
        credentials: 'same-origin',
        method: "DELETE",
    })

    if (res.status === StatusCodes.FORBIDDEN) {
        const {message} = await res.json()
        throw new HttpForbiddenError(message)
    }
}