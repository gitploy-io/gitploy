import { Timeline, Typography, Avatar, Badge } from 'antd'
import { SyncOutlined } from '@ant-design/icons'
import moment from "moment"

import { Deployment, DeploymentType, DeploymentStatus } from '../models'

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
                        Deployed by <Avatar size="small" src={d.deployer.avatar} /> <Text strong>{d.deployer.login}</Text> {moment(d.createdAt).fromNow()}  
                        <br/>
                        {(d.repo)? 
                            <a href={`/${d.repo.namespace}/${d.repo.name}/deployments/${d.number}`}>• View detail</a>:
                            null}
                    </p>
                } else {
                    description = <p>
                        Deployed {moment(d.createdAt).fromNow()}
                        <br/>
                        {(d.repo)? 
                            <a href={`/${d.repo.namespace}/${d.repo.name}/deployments/${d.number}`}>• View detail</a>:
                            null}
                    </p>
                }

                return <Timeline.Item key={idx} color={getStatusColor(d.status)} dot={dot}>
                    <p>
                        <Text strong>{d.env}</Text>&nbsp;
                        {(d.status !== DeploymentStatus.Failure) 
                            ? <Text code>{ref}</Text>
                            : null}&nbsp;
                        <Badge color={getStatusColor(d.status)} text={d.status}/>
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