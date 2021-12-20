import { Badge } from "antd"

import { Deployment, DeploymentStatusEnum } from "../models"

interface DeploymentStatusBadgeProps {
    deployment: Deployment
}

export default function DeploymentStatusBadge(props: DeploymentStatusBadgeProps): JSX.Element {
    const deployment = props.deployment
    return (
        <Badge color={getStatusColor(deployment.status)}text={deployment.status}/>
    )
}

// https://ant.design/components/timeline/#Timeline.Item
const getStatusColor = (status: DeploymentStatusEnum) => {
    switch (status) {
        case DeploymentStatusEnum.Waiting:
            return "gray"
        case DeploymentStatusEnum.Created:
            return "purple"
        case DeploymentStatusEnum.Queued:
            return "purple"
        case DeploymentStatusEnum.Running:
            return "purple"
        case DeploymentStatusEnum.Success:
            return "green"
        case DeploymentStatusEnum.Failure:
            return "red"
        case DeploymentStatusEnum.Canceled:
            return "gray"
        default:
            return "gray"
    }
}