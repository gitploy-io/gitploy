import { Badge } from "antd"

import { Deployment, DeploymentStatus } from "../models"

interface DeploymentStatusBadgeProps {
    deployment: Deployment
}

export default function DeploymentStatusBadge(props: DeploymentStatusBadgeProps) {
    const deployment = props.deployment
    return (
        <Badge color={getStatusColor(deployment.status)}text={deployment.status}/>
    )
}

// https://ant.design/components/timeline/#Timeline.Item
const getStatusColor = (status: DeploymentStatus) => {
    switch (status) {
        case DeploymentStatus.Waiting:
            return "gray"
        case DeploymentStatus.Created:
            return "purple"
        case DeploymentStatus.Running:
            return "purple"
        case DeploymentStatus.Success:
            return "green"
        case DeploymentStatus.Failure:
            return "red"
        default:
            return "gray"
    }
}