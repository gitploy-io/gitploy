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
    return `/${n.repo.namespace}/${n.repo.name}/deployments/${n.deployment.number}`
}

function convertToNotificationTitle(n: Notification): string {
    switch (n.type) {
        case NotificationType.Deployment:
            return `New Deployment #${n.deployment.number}`
        case NotificationType.ApprovalRequested:
            return `Approval Requested`
        case NotificationType.ApprovalResponded:
            return `Approval Responded`
    }
}

function convertToNotificationMessage(n: Notification): string {
    switch (n.type) {
        case NotificationType.Deployment:
            return `${n.repo.namespace}/${n.repo.name} - ${n.deployment.login} has deployed to ${n.deployment.env} environment ${moment(n.createdAt).fromNow()}`
        case NotificationType.ApprovalRequested:
            return `${n.repo.namespace}/${n.repo.name} - ${n.deployment.login} has requested the approval for the deployment(#${n.deployment.number}) ${moment(n.createdAt).fromNow()}.`
        case NotificationType.ApprovalResponded:
            return `${n.repo.namespace}/${n.repo.name} - ${n.approval.login} has responded the approval of the deployment(#${n.deployment.number}) ${moment(n.createdAt).fromNow()}.`
    }
}