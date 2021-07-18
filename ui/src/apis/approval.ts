import { StatusCodes } from 'http-status-codes'

import { instance, headers } from './setting'
import { _fetch } from "./_base"
import { mapUser } from "./user"
import { mapDataToDeployment } from "./deployment"
import { 
    Repo,
    User, 
    Deployment, 
    Approval, 
    HttpNotFoundError,
    HttpUnprocessableEntityError
 } from '../models'

export const mapDataToApproval = (data: any): Approval => {
    let user: User | null = null
    let deployment: Deployment | null = null

    if ("user" in data.edges) {
        user = mapUser(data.edges.user)
    }

    if ("deployment" in data.edges) {
        deployment = mapDataToDeployment(data.edges.deployment)
    }

    return  {
        id: data.id,
        isApproved: data.is_approved,
        createdAt: new Date(data.created_at),
        updatedAt: new Date(data.updated_at),
        user,
        deployment
    }
}

export const listApprovals = async (id: string, number: number) => {
    const res = await _fetch(`${instance}/api/v1/repos/${id}/deployments/${number}/approvals`, {
        credentials: "same-origin",
        headers,
    })

    if (res.status === StatusCodes.NOT_FOUND) {
        throw new HttpNotFoundError("There is no requested approval.")
    }

    const approvals: Approval[] = await res.json()
        .then(data => data.map((d:any): Approval => mapDataToApproval(d)))

    return approvals
}

export const createApproval = async (repo: Repo, deployment: Deployment, approver: User) => {
    const body = {
        user_id: approver.id
    }
    const res = await _fetch(`${instance}/api/v1/repos/${repo.id}/deployments/${deployment.number}/approvals`, {
        credentials: "same-origin",
        headers,
        method: "POST",
        body: JSON.stringify(body),
    })

    if (res.status === StatusCodes.UNPROCESSABLE_ENTITY) {
        const message = await res.json().then(data => data.message)
        throw new HttpUnprocessableEntityError(message)
    }

    const approval: Approval = await res.json()
        .then(data => mapDataToApproval(data))

    return approval
}

export const deleteApproval = async (repo: Repo, approval: Approval) => {
    const res = await _fetch(`${instance}/api/v1/repos/${repo.id}/approvals/${approval.id}`, {
        credentials: "same-origin",
        headers,
        method: "DELETE",
    })

    if (res.status === StatusCodes.NOT_FOUND) {
        const message = await res.json().then(data => data.message)
        throw new HttpUnprocessableEntityError(message)
    }
}

export const getMyApproval = async (id: string, number: number) => {
    const res = await _fetch(`${instance}/api/v1/repos/${id}/deployments/${number}/approval`, {
        credentials: "same-origin",
        headers,
    })

    if (res.status === StatusCodes.NOT_FOUND) {
        throw new HttpNotFoundError("There is no requested approval.")
    }

    const approval = await res.json()
        .then(data => mapDataToApproval(data))
    return approval
}

export const setApprovalApproved = async (id: string, number: number) => {
    const body = {
        is_approved: true,
    }
    const res = await _fetch(`${instance}/api/v1/repos/${id}/deployments/${number}/approval`, {
        credentials: "same-origin",
        headers,
        method: "PATCH",
        body: JSON.stringify(body),
    })

    if (res.status === StatusCodes.NOT_FOUND) {
        throw new HttpNotFoundError("There is no requested approval.")
    }

    const approval = await res.json()
        .then(data => mapDataToApproval(data))
    return approval
}

export const setApprovalDeclined = async (id: string, number: number) => {
    const body = {
        is_approved: false,
    }
    const res = await _fetch(`${instance}/api/v1/repos/${id}/deployments/${number}/approval`, {
        credentials: "same-origin",
        headers,
        method: "PATCH",
        body: JSON.stringify(body),
    })

    if (res.status === StatusCodes.NOT_FOUND) {
        throw new HttpNotFoundError("There is no requested approval.")
    }

    const approval = await res.json()
        .then(data => mapDataToApproval(data))
    return approval
}