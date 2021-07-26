import { instance, headers } from "./setting"
import { _fetch } from "./_base"
import { Notification as Noti, NotificationType } from "../models"

interface NotificationData {
    id: number
    type: string
    repo_namespace: string
    repo_name: string
    deployment_number: number
    deployment_type: string
    deployment_ref: string
    deployment_env: string
    deployment_status: string
    deployment_login: string
    approval_status: string
    approval_login: string
    notified: boolean
    checked: boolean
    created_at: string
    updated_at: string
}

const mapDataToNotification = (data: NotificationData): Noti => {
    let type: NotificationType 
    switch (data.type) {
        case "deployment":
            type = NotificationType.Deployment
            break
        case "approval_requested":
            type = NotificationType.ApprovalRequested
            break
        case "approval_responded":
            type = NotificationType.ApprovalResponded
            break
        default:
            type = NotificationType.Deployment
    }

    return { 
        id: data.id,
        type, 
        repoNamespace: data.repo_namespace,
        repoName: data.repo_name,
        deploymentNumber: data.deployment_number,
        deploymentType: data.deployment_type,
        deploymentRef: data.deployment_ref,
        deploymentEnv: data.deployment_env,
        deploymentStatus: data.deployment_status,
        deploymentLogin: data.deployment_login,
        approvalStatus: data.approval_status,
        approvalLogin: data.approval_login,
        notified: data.notified,
        checked: data.checked,
        createdAt: new Date(data.created_at),
        updatedAt: new Date(data.updated_at),
    }
}

export const subscribeNotification = (cb: (n: Noti) => void): EventSource => {
    const sse = new EventSource(`${instance}/api/v1/stream`, {
        withCredentials: true,
    })

    const eventName = "notification"
    sse.addEventListener(eventName, (e: any) => {
        const data = JSON.parse(e.data)
        cb(mapDataToNotification(data))
    })

    return sse
}

export const listNotifications = async (page = 1, perPage = 30): Promise<Noti[]> => {
    const notifications: Noti[]  = await _fetch(`${instance}/api/v1/user/notifications?page=${page}&per_page=${perPage}`, {
        headers,
        credentials: "same-origin",
    })
        .then(res => res.json())
        .then(data => data.map((n: any) => mapDataToNotification(n)))
    
    return notifications
}

export const patchNotificationChecked = async (id: number): Promise<Noti>=> {
    const body = {
        checked: true
    }
    const notification  = await _fetch(`${instance}/api/v1/user/notifications/${id}`, {
        headers,
        credentials: "same-origin",
        method: "PATCH",
        body: JSON.stringify(body)
    })
        .then(res => res.json())
        .then(data => mapDataToNotification(data))

    return notification
}
