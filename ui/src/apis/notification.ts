import { instance, headers } from "./setting"
import { _fetch } from "./_base"
import { mapRepo } from "./repo"
import { mapDataToDeployment } from "./deployment"
import { Notification as NotificationData, NotificationType, Deployment } from "../models"

export const subscribeNotification = (cb: (n: NotificationData) => void) => {
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

export const listNotifications = async (page: number, perPage: number) => {
    const notifications  = await _fetch(`${instance}/api/v1/notifications?page=${page}&per_page=${perPage}`, {
        headers,
        credentials: "same-origin",
    })
        .then(res => res.json())
        .then(data => data.map((n: any) => mapDataToNotification(n)))
    
    return notifications
}

export const patchNotificationChecked = async (id: number) => {
    const body = {
        checked: true
    }
    const notification  = await _fetch(`${instance}/api/v1/notifications/${id}`, {
        headers,
        credentials: "same-origin",
        method: "PATCH",
        body: JSON.stringify(body)
    })
        .then(res => res.json())
        .then(data => mapDataToNotification(data))

    return notification
}

function mapDataToNotification(data: any): NotificationData {
    let type: NotificationType = NotificationType.Deployment
    var deployment: Deployment | null = null

    if (data.type === "deployment") {
        type = NotificationType.Deployment
        deployment = mapDataToDeployment(data.edges.deployment)
    } 

    return { 
        id: data.id,
        type, 
        notified: data.notified,
        checked: data.checked,
        createdAt: new Date(data.created_at),
        updatedAt: new Date(data.updated_at),
        repo: mapRepo(data.edges.repo),
        deployment,
    }
}
