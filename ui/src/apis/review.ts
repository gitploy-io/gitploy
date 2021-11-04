import { StatusCodes } from 'http-status-codes'

import { instance, headers } from './setting'
import { _fetch } from "./_base"
import { UserData, mapDataToUser } from "./user"
import { DeploymentData, mapDataToDeployment } from "./deployment"
import { 
    User,
    Deployment,
    Review,
    ReviewStatusEnum,
    HttpNotFoundError,
 } from '../models'

export interface ReviewData {
    id: number,
    status: string
    created_at: string
    updated_at: string
    edges: {
        user: UserData,
        deployment: DeploymentData
    }
}

// eslint-disable-next-line
export const mapDataToReview = (data: ReviewData): Review => {
    let user: User | undefined 
    let deployment: Deployment | undefined 

    if ("user" in data.edges) {
        user = mapDataToUser(data.edges.user)
    }

    if ("deployment" in data.edges) {
        deployment = mapDataToDeployment(data.edges.deployment)
    }

    return  {
        id: data.id,
        status: mapDataToReviewStatus(data.status),
        createdAt: new Date(data.created_at),
        updatedAt: new Date(data.updated_at),
        user,
        deployment
    }
}

const mapDataToReviewStatus = (status: string): ReviewStatusEnum => {
    switch (status) {
        case "pending":
            return ReviewStatusEnum.Pending
        case "approved":
            return ReviewStatusEnum.Approved
        case "rejected":
            return ReviewStatusEnum.Rejected
        default:
            return ReviewStatusEnum.Pending
    }
}

export const searchReviews = async (): Promise<Review[]> => {
    const reviews: Review[] = await _fetch(`${instance}/api/v1/search/reviews`, {
        credentials: "same-origin",
        headers,
    })
        .then(res => res.json())
        .then(data => data.map((d:any): Review => mapDataToReview(d)))

    return reviews
}

export const listReviews = async (namespace: string, name: string, number: number): Promise<Review[]> => {
    const res = await _fetch(`${instance}/api/v1/repos/${namespace}/${name}/deployments/${number}/reviews`, {
        credentials: "same-origin",
        headers,
    })

    const reviews: Review[] = await res.json()
        .then(data => data.map((d:any): Review => mapDataToReview(d)))

    return reviews
}

export const getUserReview = async (namespace: string, name: string, number: number): Promise<Review> => {
    const res = await _fetch(`${instance}/api/v1/repos/${namespace}/${name}/deployments/${number}/review`, {
        credentials: "same-origin",
        headers,
    })

    if (res.status === StatusCodes.NOT_FOUND) {
        const { message } = await res.json()
        throw new HttpNotFoundError(message)
    }

    const review = await res.json()
        .then(data => mapDataToReview(data))
    return review
}

export const approveReview = async (namespace: string, name: string, number: number): Promise<Review> => {
    const body = {
        status: "approved",
    }
    const res = await _fetch(`${instance}/api/v1/repos/${namespace}/${name}/deployments/${number}/review`, {
        credentials: "same-origin",
        headers,
        method: "PATCH",
        body: JSON.stringify(body),
    })

    if (res.status === StatusCodes.NOT_FOUND) {
        const { message } = await res.json()
        throw new HttpNotFoundError(message)
    }

    const review = await res.json()
        .then(data => mapDataToReview(data))
    return review
}

export const rejectReview = async (namespace: string, name: string, number: number): Promise<Review> => {
    const body = {
        status: "rejected",
    }
    const res = await _fetch(`${instance}/api/v1/repos/${namespace}/${name}/deployments/${number}/review`, {
        credentials: "same-origin",
        headers,
        method: "PATCH",
        body: JSON.stringify(body),
    })

    if (res.status === StatusCodes.NOT_FOUND) {
        const { message } = await res.json()
        throw new HttpNotFoundError(message)
    }

    const review = await res.json()
        .then(data => mapDataToReview(data))
    return review
}