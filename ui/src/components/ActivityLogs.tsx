import { Timeline, Typography, Avatar } from 'antd'
import { SyncOutlined } from '@ant-design/icons'
import moment from "moment"

import { Deployment, DeploymentType, LastDeploymentStatus } from "../models"
import DeploymentStatusBadge from "./DeploymentStatusBadge"

const { Text } = Typography

interface ActivityLogsProps {
    deployments: Deployment[]
}

export default function ActivityLogs(props: ActivityLogsProps): JSX.Element {
    return <Timeline>
        {props.deployments.map((d, idx) => {
            const dot = (d.lastStatus === LastDeploymentStatus.Running)? 
                <SyncOutlined style={{color: "purple"}} spin />: 
                null
            const ref = (d.type === DeploymentType.Commit)? 
                d.ref.substr(0, 7): 
                d.ref
            const avatar = (d.deployer)? 
                <span><Avatar size="small" src={d.deployer.avatar} /> <Text strong>{d.deployer.login}</Text></span> :
                <span><Avatar size="small">U</Avatar> </span> 

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
        case LastDeploymentStatus.Canceled:
            return "gray"
        default:
            return "gray"
    }
}