import { instance } from "./setting"

import { Notification as NotificationData, NotificationType } from "../models"

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

function mapDataToNotification(data: any): NotificationData {
    let type: NotificationType

    if (data.type === "deployment") {
        type = NotificationType.Deployment

    } else {
        type = NotificationType.Deployment
    }

    return { 
        id: data.id,
        type, 
        resourceId: data.resource_id,
        notified: data.notified,
        createdAt: new Date(data.created_at),
        updatedAt: new Date(data.updated_at),
    }
}
