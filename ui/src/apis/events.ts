import { instance } from './setting'

import { mapDataToDeployment } from "./deployment"
import { mapDataToApproval } from "./approval"
import { Deployment, Approval } from "../models"

export const subscribeDeploymentEvent = (cb: (e: Deployment) => void): EventSource => {
    const eventName = "deployment"

    const sse = new EventSource(`${instance}/api/v1/stream/events`, {
        withCredentials: true,
    })

    sse.addEventListener(eventName, (e: any) => {
        const data = JSON.parse(e.data)
        cb(mapDataToDeployment(data))
    })

    return sse
}

export const subscribeApprovalEvent = (cb: (e: Approval) => void): EventSource => {
    const eventName = "approval"

    const sse = new EventSource(`${instance}/api/v1/stream/events`, {
        withCredentials: true,
    })

    sse.addEventListener(eventName, (e: any) => {
        const data = JSON.parse(e.data)
        cb(mapDataToApproval(data))
    })

    return sse
}