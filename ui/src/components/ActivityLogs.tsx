import { Timeline } from 'antd'
import { SyncOutlined } from '@ant-design/icons';

import { Deployment, DeploymentStatus } from '../models'

interface ActivityLogsProps {
    deployments: Deployment[]
}

export default function ActivityLogs(props: ActivityLogsProps) {

    return (
        <Timeline>
            {props.deployments.map((d, idx) => {
                return <Timeline.Item key={idx} color={getStatusColor(d.status)} dot={(d.status === DeploymentStatus.Running) ? <SyncOutlined spin /> : null}>
                    {d.env} {d.ref}
                </Timeline.Item>
            })}
        </Timeline>
    )
}

// https://ant.design/components/timeline/#Timeline.Item
const getStatusColor = (status: DeploymentStatus) => {
    switch (status) {
        case DeploymentStatus.Waiting:
            return "gray"
        case DeploymentStatus.Created:
            return "blue"
        case DeploymentStatus.Running:
            return "blue"
        case DeploymentStatus.Success:
            return "green"
        case DeploymentStatus.Failure:
            return "red"
        default:
            return "gray"
    }
}