import { Timeline, Typography } from 'antd'
import { SyncOutlined } from '@ant-design/icons'
import moment from "moment"

import { Deployment, DeploymentStatusEnum } from "../../models"
import DeploymentStatusBadge from "../../components/DeploymentStatusBadge"
import UserAvatar from '../../components/UserAvatar'
import DeploymentRefCode from '../../components/DeploymentRefCode'
import { getStatusColor } from "../../components/partials"

const { Text } = Typography

export interface ActivityLogsProps {
    deployments: Deployment[]
}

export default function ActivityLogs({ deployments }: ActivityLogsProps): JSX.Element {
    return (
        <Timeline>
            {deployments.map((d, idx) => {
                const dot = (d.status === DeploymentStatusEnum.Running)? 
                    <SyncOutlined style={{color: "purple"}} spin />
                    : 
                    null
                const avatar = <UserAvatar user={d.deployer} />

                return (
                    <Timeline.Item key={idx} color={getStatusColor(d.status)} dot={dot}>
                        <p>
                            <Text strong>{d.env}</Text> <DeploymentRefCode deployment={d}/> <a href={`/${d.repo?.namespace}/${d.repo?.name}/deployments/${d.number}`}>• View detail #{d.number}</a>
                        </p>
                        <p>
                            Deployed by {avatar} {moment(d.createdAt).fromNow()} <DeploymentStatusBadge deployment={d}/>
                        </p>
                    </Timeline.Item>
                )
            })}
        </Timeline>
    )
}
