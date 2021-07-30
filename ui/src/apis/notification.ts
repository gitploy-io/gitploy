import { instance, headers } from "./setting"
import { _fetch } from "./_base"
import { Notification as Noti, NotificationType } from "../models"

interface NotificationData {
    id: number
    type: string
    repo: {
        namespace: string
        name: string
    }
    deployment: {
        number: number
        type: string
        ref: string
        env: string
        status: string
        login: string
    }
    approval: {
        status: string
        login: string
    }
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
        repo: {
            namespace: data.repo.namespace,
            name: data.repo.name
        },
        deployment: {
            number: data.deployment.number,
            type: data.deployment.type,
            ref: data.deployment.ref,
            env: data.deployment.env,
            status: data.deployment.status,
            login: data.deployment.login,
        },
        approval: {
            status: data.approval.status,
            login: data.approval.login,
        },
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
