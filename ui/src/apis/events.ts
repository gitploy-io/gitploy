import { instance } from './setting'

import { DeploymentData, mapDataToDeployment } from "./deployment"
import { ReviewData, mapDataToReview } from "./review"
import { Deployment, Review, Event, EventKindEnum, EventTypeEnum } from "../models"

interface EventData {
    id: number
    kind: string
    type: string
    deleted_id: number
    edges: {
        deployment?: DeploymentData
        review?: ReviewData
    }
}

const mapDataToEvent = (data: EventData): Event => {
    let kind: EventKindEnum
    let type: EventTypeEnum
    let deployment: Deployment | undefined
    let review: Review | undefined

    switch (data.kind) {
        case "deployment":
            kind = EventKindEnum.Deployment
            break
        case "review":
            kind = EventKindEnum.Review
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

    if (data.edges.review) {
        review = mapDataToReview(data.edges.review)
    }

    return {
        id: data.id,
        kind,
        type,
        deletedId: data.deleted_id,
        deployment,
        review
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