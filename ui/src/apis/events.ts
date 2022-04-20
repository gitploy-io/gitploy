import { instance } from './setting'

import { mapDataToDeployment } from "./deployment"
import {  mapDataToReview } from "./review"
import { Deployment, Review  } from "../models"


export const subscribeDeploymentEvents = (cb: (deployment: Deployment) => void): EventSource => {
    const sse = new EventSource(`${instance}/api/v1/stream/events`, {
        withCredentials: true,
    })

    sse.addEventListener("deployment", (e: any) => {
        const data = JSON.parse(e.data)
        const deployment = mapDataToDeployment(data)

        cb(deployment)
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
