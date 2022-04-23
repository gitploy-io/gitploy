import { instance } from './setting'

import { mapDataToDeploymentStatus } from "./deployment"
import { mapDataToReview } from "./review"
import { DeploymentStatus, Review  } from "../models"


export const subscribeDeploymentStatusEvents = (cb: (status: DeploymentStatus) => void): EventSource => {
    const sse = new EventSource(`${instance}/api/v1/stream/events`, {
        withCredentials: true,
    })

    sse.addEventListener("deployment_status", (e: any) => {
        const data = JSON.parse(e.data)
        const status = mapDataToDeploymentStatus(data)

        cb(status)
    })

    return sse
}

export const subscribeReviewEvents = (cb: (review: Review) => void): EventSource => {
    const sse = new EventSource(`${instance}/api/v1/stream/events`, {
        withCredentials: true,
    })

    sse.addEventListener("review", (e: any) => {
        const data = JSON.parse(e.data)
        const review = mapDataToReview(data)

        cb(review)
    })

    return sse
}
