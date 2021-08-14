import { StatusCodes } from 'http-status-codes'

import { instance, headers } from './setting'
import { _fetch } from "./_base"
import { UserData, mapDataToUser } from "./user"
import { DeploymentData, mapDataToDeployment } from "./deployment"
import { 
    Repo,
    User, 
    Deployment, 
    Approval, 
    ApprovalStatus,
    HttpNotFoundError,
    HttpUnprocessableEntityError
 } from '../models'

interface ApprovalData {
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
export const mapDataToApproval = (data: ApprovalData): Approval => {
    let user: User | null = null
    let deployment: Deployment | null = null

    if ("user" in data.edges) {
        user = mapDataToUser(data.edges.user)
    }

    if ("deployment" in data.edges) {
        deployment = mapDataToDeployment(data.edges.deployment)
    }

    return  {
        id: data.id,
        status: mapDataToApprovalStatus(data.status),
        createdAt: new Date(data.created_at),
        updatedAt: new Date(data.updated_at),
        user,
        deployment
    }
}

const mapDataToApprovalStatus = (status: string): ApprovalStatus => {
    switch (status) {
        case "pending":
            return ApprovalStatus.Pending
        case "approved":
            return ApprovalStatus.Approved
        case "declined":
            return ApprovalStatus.Declined
        default:
            return ApprovalStatus.Pending
    }
}

function mapApprovalStatusToString(status: ApprovalStatus): string {
    switch (status) {
        case ApprovalStatus.Pending:
            return "pending"
        case ApprovalStatus.Approved:
            return "approved"
        case ApprovalStatus.Declined:
            return "declined"
        default:
            return "pending"
    }
}

export const searchApprovals = async (statuses: ApprovalStatus[], from?: Date, to?: Date, page = 1, perPage = 30): Promise<Approval[]> => {
    const ss: string[] = []
    statuses.forEach((status) => {
        ss.push(mapApprovalStatusToString(status))
    })

    const fromParam = (from)? `from=${from.toISOString()}` : ""
    const toParam = (to)? `to=${to.toISOString()}` : ""

    const approvals: Approval[] = await _fetch(`${instance}/api/v1/search/approvals?statuses=${ss.join(",")}&${fromParam}&${toParam}&page=${page}&per_page=${perPage}`, {
        credentials: "same-origin",
        headers,
    })
        .then(res => res.json())
        .then(data => data.map((d:any): Approval => mapDataToApproval(d)))

    return approvals
}

export const listApprovals = async (id: string, number: number): Promise<Approval[]> => {
    const res = await _fetch(`${instance}/api/v1/repos/${id}/deployments/${number}/approvals`, {
        credentials: "same-origin",
        headers,
    })

    if (res.status === StatusCodes.NOT_FOUND) {
        throw new HttpNotFoundError("There is no such a deployment.")
    }

    const approvals: Approval[] = await res.json()
        .then(data => data.map((d:any): Approval => mapDataToApproval(d)))

    return approvals
}

export const createApproval = async (repo: Repo, deployment: Deployment, approver: User): Promise<Approval> => {
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

export const deleteApproval = async (repo: Repo, approval: Approval): Promise<void> => {
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

export const getMyApproval = async (id: string, number: number): Promise<Approval> => {
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

export const setApprovalApproved = async (id: string, number: number): Promise<Approval> => {
    const body = {
        status: ApprovalStatus.Approved.toString(),
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

export const setApprovalDeclined = async (id: string, number: number): Promise<Approval> => {
    const body = {
        status: ApprovalStatus.Declined.toString(),
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