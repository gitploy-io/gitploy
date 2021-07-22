import { StatusCodes } from 'http-status-codes'

import { instance, headers } from './setting'
import { _fetch } from "./_base"
import { Config, Env, HttpNotFoundError } from '../models'

export const getConfig = async (repoId: string): Promise<Config> => {
    const response = await _fetch(`${instance}/api/v1/repos/${repoId}/config`, {
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

function mapDataToConfig(data: any) {
    const envs: Env[] = data.envs.map((e: any) => {
        return {
            name: e.name,
            requiredContexts: e.required_contexts,
            approvalEnabled: e.approval? e.approval.enabled : false,
        }
    })
    return {
        envs,
    } as Config
}