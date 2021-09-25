import { StatusCodes } from "http-status-codes"

import { instance, headers } from "./setting"
import { _fetch } from "./_base"
import { UserData, mapDataToUser } from "./user"
import { Repo, Lock, User, HttpUnprocessableEntityError } from "../models"

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

export const listLocks = async (repo: Repo): Promise<Lock[]> => {
    const locks = await _fetch(`${instance}/api/v1/repos/${repo.id}/locks`, {
        headers,
        credentials: 'same-origin',
    })
        .then(res => res.json())
        .then(datas => datas.map((d: any): Lock => mapDataToLock(d)))

    return locks
}

export const lock = async (repo: Repo, env: string): Promise<Lock> => {
    const res = await _fetch(`${instance}/api/v1/repos/${repo.id}/locks`, {
        headers,
        credentials: 'same-origin',
        method: "POST",
        body: JSON.stringify({env})
    })

    if (res.status === StatusCodes.UNPROCESSABLE_ENTITY) {
        const {message} = await res.json()
        throw new HttpUnprocessableEntityError(message)
    }

    const lock = await res
        .json()
        .then(data => mapDataToLock(data))
    return lock
}

export const unlock = async (repo: Repo, lock: Lock): Promise<void> => {
    await _fetch(`${instance}/api/v1/repos/${repo.id}/locks/${lock.id}`, {
        headers,
        credentials: 'same-origin',
        method: "DELETE",
    })
}