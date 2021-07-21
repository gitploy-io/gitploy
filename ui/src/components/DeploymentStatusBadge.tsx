import { Badge } from "antd"

import { Deployment, LastDeploymentStatus } from "../models"

interface DeploymentStatusBadgeProps {
    deployment: Deployment
}

export default function DeploymentStatusBadge(props: DeploymentStatusBadgeProps) {
    const deployment = props.deployment
    return (
        <Badge color={getStatusColor(deployment.lastStatus)}text={deployment.lastStatus}/>
    )
}

// https://ant.design/components/timeline/#Timeline.Item
const getStatusColor = (status: LastDeploymentStatus) => {
    switch (status) {
        case LastDeploymentStatus.Waiting:
            return "gray"
        case LastDeploymentStatus.Created:
            return "purple"
        case LastDeploymentStatus.Running:
            return "purple"
        case LastDeploymentStatus.Success:
            return "green"
        case LastDeploymentStatus.Failure:
            return "red"
        default:
            return "gray"
    }
}