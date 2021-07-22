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

    return (
        <Timeline>
            {props.deployments.map((d, idx) => {
                const dot = (d.lastStatus === LastDeploymentStatus.Running) ? <SyncOutlined spin /> : null
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

                return <Timeline.Item key={idx} color={getStatusColor(d.lastStatus)} dot={dot}>
                    <p>
                        <Text strong>{d.env}</Text>&nbsp;
                        {(d.lastStatus !== LastDeploymentStatus.Failure) 
                            ? <Text code>{ref}</Text>
                            : null}&nbsp;
                        {(d.repo)? 
                            <a href={`/${d.repo.namespace}/${d.repo.name}/deployments/${d.number}`}>• View detail #{d.number}</a>:
                            null}
                    </p>
                    {description}
                    
                </Timeline.Item>
            })}
        </Timeline>
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