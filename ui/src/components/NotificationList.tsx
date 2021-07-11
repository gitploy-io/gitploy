import { List } from "antd"
import  moment from "moment"

import { Notification, NotificationType } from "../models"

interface NotificationListProps {
    notifications: Notification[]
    onClickNotificaiton(n: Notification): void
}

export default function NotificationList(props: NotificationListProps) {
    const uncheckedStyle: React.CSSProperties = { backgroundColor: "#efdbff" }

    return ( <List size="small"
            dataSource={props.notifications}
            renderItem={(n, idx) => {
                return (<List.Item key={idx} style={(!n.checked)? uncheckedStyle : {}}>
                        <List.Item.Meta 
                            title={
                                <a 
                                    href={convertToNotificationLink(n)}
                                    onClick={() => {props.onClickNotificaiton(n)}}>
                                    {convertToNotificationTitle(n)} 
                                </a>}
                            />
                        {convertToNotificationMessage(n)}
                    </List.Item>)
            }}>
        </List>)
}

function convertToNotificationLink(n: Notification): string {
    switch (n.type) {
        case NotificationType.Deployment:
            if (n.deployment === null) {
                throw new Error("The notification must have a deployment.")
            }

            const repo = n.repo
            const deployment = n.deployment

            return `/${repo.namespace}/${repo.name}/deployments/${deployment.number}`
    }
}

function convertToNotificationTitle(n: Notification): string {
    switch (n.type) {
        case NotificationType.Deployment:
            if (n.deployment === null) {
                throw new Error("The notification must have a deployment.")
            }

            const deployment = n.deployment

            return `New Deployment #${deployment.number}`
    }
}

function convertToNotificationMessage(n: Notification): string {
    switch (n.type) {
        case NotificationType.Deployment:
            if (n.deployment === null) {
                return `${n.repo.namespace}/${n.repo.name} - Deployed.`
            }
            return `${n.repo.namespace}/${n.repo.name} - Deployed to ${n.deployment.env} environment at ${moment(n.deployment.createdAt).fromNow()}`
    }
}