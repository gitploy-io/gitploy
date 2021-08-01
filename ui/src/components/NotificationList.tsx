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
        case NotificationType.DeploymentCreated:
            return `New Deployment #${n.deployment.number}`
        case NotificationType.DeploymentUpdated:
            return `Deployment Updated #${n.deployment.number}`
        case NotificationType.ApprovalRequested:
            return `Approval Requested`
        case NotificationType.ApprovalResponded:
            return `Approval Responded`
    }
}

function convertToNotificationMessage(n: Notification): string {
    switch (n.type) {
        case NotificationType.DeploymentCreated:
            return `${n.deployment.login} deploys ${n.deployment.ref} to the ${n.deployment.env} environment of ${n.repo.namespace}/${n.repo.name} ${moment(n.createdAt).fromNow()}`
        case NotificationType.DeploymentUpdated:
            return `The deployment(#${n.deployment.number}) of ${n.repo.namespace}/${n.repo.name} is updated ${n.deployment.status} ${moment(n.createdAt).fromNow()}`
        case NotificationType.ApprovalRequested:
            return `${n.deployment.login} has requested the approval for the deployment(#${n.deployment.number}) of ${n.repo.namespace}/${n.repo.name} ${moment(n.createdAt).fromNow()}.`
        case NotificationType.ApprovalResponded:
            return `${n.approval.login} has responded the approval of the deployment(#${n.deployment.number}) of ${n.repo.namespace}/${n.repo.name} ${moment(n.createdAt).fromNow()}.`
    }
}