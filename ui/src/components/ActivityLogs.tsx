import { Timeline, Typography } from 'antd'
import { SyncOutlined } from '@ant-design/icons'
import moment from "moment"

import { Deployment, DeploymentType, DeploymentStatusEnum } from "../models"
import DeploymentStatusBadge from "./DeploymentStatusBadge"
import UserAvatar from './UserAvatar'

const { Text } = Typography

interface ActivityLogsProps {
    deployments: Deployment[]
}

export default function ActivityLogs(props: ActivityLogsProps): JSX.Element {
    return <Timeline>
        {props.deployments.map((d, idx) => {
            const dot = (d.lastStatus === DeploymentStatusEnum.Running)? 
                <SyncOutlined style={{color: "purple"}} spin />: 
                null
            const ref = (d.type === DeploymentType.Commit)? 
                d.ref.substr(0, 7): 
                d.ref
            const avatar = <UserAvatar user={d.deployer} />

            return <Timeline.Item key={idx} color={getStatusColor(d.lastStatus)} dot={dot}>
                <p>
                    <Text strong>{d.env}</Text> <Text code>{ref}</Text> <a href={`/${d.repo?.namespace}/${d.repo?.name}/deployments/${d.number}`}>• View detail #{d.number}</a>
                </p>
                <p>
                    Deployed by {avatar} {moment(d.createdAt).fromNow()} <DeploymentStatusBadge deployment={d}/>
                </p>
            </Timeline.Item>
        })}
    </Timeline>
}

// https://ant.design/components/timeline/#Timeline.Item
const getStatusColor = (status: DeploymentStatusEnum) => {
    switch (status) {
        case DeploymentStatusEnum.Waiting:
            return "gray"
        case DeploymentStatusEnum.Created:
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