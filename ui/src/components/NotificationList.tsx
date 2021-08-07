import { List, Typography } from "antd"
import  moment from "moment"

import { Notification, NotificationType } from "../models"

const { Paragraph, Text } = Typography

interface NotificationListProps {
    notifications: Notification[]
    onClickNotificaiton(n: Notification): void
}

export default function NotificationList(props: NotificationListProps): JSX.Element {
    const uncheckedStyle: React.CSSProperties = { backgroundColor: "#efdbff" }

    return ( <List 
            size="small"
            itemLayout="vertical"
            dataSource={props.notifications}
            renderItem={(n, idx) => {
                return (
                    <List.Item 
                        key={idx} 
                        style={(!n.checked)? uncheckedStyle : {}}
                    >
                        <List.Item.Meta 
                            style={{margin: 0}}
                            title={<a 
                                        href={convertToNotificationLink(n)}
                                        onClick={() => {props.onClickNotificaiton(n)}}
                                    >
                                        {convertToNotificationTitle(n)} 
                                    </a>}
                            description={convertToNotificationMessage(n)}
                        />
                    </List.Item>
                )
            } }>
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

function convertToNotificationMessage(n: Notification): JSX.Element {
    const ref = (n.deployment.type === "commit")? n.deployment.ref.substr(0, 7) : n.deployment.ref
    const style: React.CSSProperties = {margin: 0}

    switch (n.type) {
        case NotificationType.DeploymentCreated:
            return <Paragraph style={style}>
                <Text strong>{n.deployment.login}</Text> deploys <Text code>{ref}</Text> to the <Text code>{n.deployment.env}</Text> environment of <Text code>{n.repo.namespace}/{n.repo.name}</Text> {moment(n.createdAt).fromNow()}
            </Paragraph>
        case NotificationType.DeploymentUpdated:
            return <Paragraph style={style}>
                The deployment(#{n.deployment.number}) of <Text code>{n.repo.namespace}/{n.repo.name}</Text> is updated {n.deployment.status} {moment(n.createdAt).fromNow()}
            </Paragraph>
        case NotificationType.ApprovalRequested:
            return <Paragraph style={style}>
                <Text strong>{n.deployment.login}</Text> has requested the approval for the deployment(#{n.deployment.number}) of <Text code>{n.repo.namespace}/{n.repo.name}</Text> {moment(n.createdAt).fromNow()}
            </Paragraph>
        case NotificationType.ApprovalResponded:
            return <Paragraph style={style}>
                <Text strong>{n.deployment.login}</Text> has <Text strong>{n.approval.status}</Text> the approval for the deployment(#{n.deployment.number}) of <Text code>{n.repo.namespace}/{n.repo.name}</Text> {moment(n.createdAt).fromNow()}
            </Paragraph>
    }
}