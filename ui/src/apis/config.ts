import { StatusCodes } from 'http-status-codes'

import { instance, headers } from './settings'
import { Config, Env } from '../models'

export const getConfig = async (repoId: string) => {
    let conf: Config

    try {
        const response = await fetch(`${instance}/v1/repos/${repoId}/config`, {
            headers,
            credentials: "same-origin",
        })

        if (response.status === StatusCodes.NOT_FOUND) {
            conf = { envs: [] }
            return conf
        }

        conf = await response.json()
            .then(c => {
                const envs: Env[] = c.envs.map((e: any) => {
                    return {
                        name: e.name,
                        requiredContexts: e.required_contexts
                    }
                })
                return {
                    envs: envs
                }
            })
    } catch (e) {
        throw e
    }

    return conf
}