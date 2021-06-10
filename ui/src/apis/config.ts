import { StatusCodes } from 'http-status-codes'

import { instance, headers } from './settings'
import { Config, Env, HttpNotFoundError } from '../models'

export const getConfig = async (repoId: string) => {
    let conf: Config

    const response = await fetch(`${instance}/api/v1/repos/${repoId}/config`, {
        headers,
        credentials: "same-origin",
    })
    if (response.status === StatusCodes.NOT_FOUND) {
        const message = await response.json().then(data => data.message)
        throw new HttpNotFoundError(message)
    }

    conf = await response.json()
        .then(c => mapConfig(c))

    return conf
}

function mapConfig(c: any) {
    const envs: Env[] = c.envs.map((e: any) => {
        return {
            name: e.name,
            requiredContexts: e.required_contexts
        }
    })
    return {
        envs,
    } as Config
}