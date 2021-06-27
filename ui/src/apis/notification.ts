import { instance } from "./setting"
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
        createdAt: new Date(data.created_at),
        updatedAt: new Date(data.updated_at),
        deployment,
    }
}
