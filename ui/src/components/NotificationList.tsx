import { List } from "antd"
import  moment from "moment"

import { Notification, NotificationType } from "../models"

interface NotificationListProps {
    notifications: Notification[]
    onClickNotificaiton(n: Notification): void
}

export default function NotificationList(props: NotificationListProps): JSX.Element {
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

            return `/${n.repoNamespace}/${n.repoName}/deployments/${n.deploymentNumber}`
    }
}

function convertToNotificationTitle(n: Notification): string {
    switch (n.type) {
        case NotificationType.Deployment:

            return `New Deployment #${n.deploymentNumber}`
    }
}

function convertToNotificationMessage(n: Notification): string {
    switch (n.type) {
        case NotificationType.Deployment:
            return `${n.repoNamespace}/${n.repoName} - Deployed to ${n.deploymentEnv} environment at ${moment(n.createdAt).fromNow()}`
    }
}