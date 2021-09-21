import { instance } from './setting'

import { DeploymentData, mapDataToDeployment } from "./deployment"
import { ApprovalData, mapDataToApproval } from "./approval"
import { Deployment, Approval, Event, EventKindEnum, EventTypeEnum } from "../models"

interface EventData {
    id: number
    kind: string
    type: string
    deleted_entity_id: number
    edges: {
        deployment?: DeploymentData
        approval?: ApprovalData
    }
}

const mapDataToEvent = (data: EventData): Event => {
    let kind: EventKindEnum
    let type: EventTypeEnum
    let deployment: Deployment | undefined
    let approval: Approval | undefined

    switch (data.kind) {
        case "deployment":
            kind = EventKindEnum.Deployment
            break
        case "approval":
            kind = EventKindEnum.Approval
            break
        default:
            kind = EventKindEnum.Deployment
    }

    switch (data.type) {
        case "created":
            type = EventTypeEnum.Created
            break
        case "updated":
            type = EventTypeEnum.Updated
            break
        case "deleted":
            type = EventTypeEnum.Deleted
            break
        default:
            type = EventTypeEnum.Created
    }

    if (data.edges.deployment) {
        deployment = mapDataToDeployment(data.edges.deployment)
    }

    if (data.edges.approval) {
        approval = mapDataToApproval(data.edges.approval)
    }

    return {
        id: data.id,
        kind,
        type,
        deletedEntityId: data.deleted_entity_id,
        deployment,
        approval
    } 
}

export const subscribeEvents = (cb: (event: Event) => void): EventSource => {
    const sse = new EventSource(`${instance}/api/v1/stream/events`, {
        withCredentials: true,
    })

    sse.addEventListener("event", (e: any) => {
        const data = JSON.parse(e.data)
        const event = mapDataToEvent(data)

        cb(event)
    })

    return sse
}
