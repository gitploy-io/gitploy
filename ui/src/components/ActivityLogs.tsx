import { Timeline, Typography, Avatar } from 'antd'
import { SyncOutlined } from '@ant-design/icons'
import moment from "moment"

import { Deployment, DeploymentType, DeploymentStatus } from "../models"
import DeploymentStatusBadge from "./DeploymentStatusBadge"

const { Text } = Typography

interface ActivityLogsProps {
    deployments: Deployment[]
}

export default function ActivityLogs(props: ActivityLogsProps) {

    return (
        <Timeline>
            {props.deployments.map((d, idx) => {
                const dot = (d.status === DeploymentStatus.Running) ? <SyncOutlined spin /> : null
                const ref = (d.type === DeploymentType.Commit)? d.sha.substr(0, 7) : d.ref
                let description: React.ReactElement 

                if (d.deployer) {
                    description = <p>
                        Deployed by &nbsp;
                        <Avatar size="small" src={d.deployer.avatar} /> &nbsp;
                        <Text strong>{d.deployer.login}</Text> &nbsp;
                        {moment(d.createdAt).fromNow()} &nbsp;
                        <DeploymentStatusBadge deployment={d}/>
                    </p>
                } else {
                    // deployer is removed by admin.
                    description = <p>
                        Deployed by &nbsp;
                        <Avatar size="small">U</Avatar>&nbsp;
                        {moment(d.createdAt).fromNow()} &nbsp;
                        <DeploymentStatusBadge deployment={d}/>
                    </p>

                }

                return <Timeline.Item key={idx} color={getStatusColor(d.status)} dot={dot}>
                    <p>
                        <Text strong>{d.env}</Text>&nbsp;
                        {(d.status !== DeploymentStatus.Failure) 
                            ? <Text code>{ref}</Text>
                            : null}&nbsp;
                        {(d.repo)? 
                            <a href={`/${d.repo.namespace}/${d.repo.name}/deployments/${d.number}`}>• View detail</a>:
                            null}
                    </p>
                    {description}
                    
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