import { StatusCodes } from 'http-status-codes'

import { instance, headers } from './setting'
import { _fetch } from "./_base"
import { Config, Env, HttpNotFoundError } from '../models'

interface ConfigData {
    envs: EnvData[]
}

interface EnvData {
    name: string
    required_contexts?: string[]
    approval?: {
        enabled: boolean
        required_count: number
    }
    review?: {
        enabled: boolean
        reviewers: string[]
    }
}

const mapDataToConfig = (data: ConfigData): Config => {
    const envs: Env[] = data.envs.map((e: EnvData) => {
        const { review } = e

        return {
            name: e.name,
            requiredContexts: e.required_contexts,
            review,
        }
    })

    return {
        envs,
    }
}

export const getConfig = async (namespace: string, name: string): Promise<Config> => {
    const response = await _fetch(`${instance}/api/v1/repos/${namespace}/${name}/config`, {
        headers,
        credentials: "same-origin",
    })
    if (response.status === StatusCodes.NOT_FOUND) {
        const message = await response.json().then(data => data.message)
        throw new HttpNotFoundError(message)
    }

    const conf = await response.json()
        .then(c => mapDataToConfig(c))

    return conf
}
