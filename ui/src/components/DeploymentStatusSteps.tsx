import { Timeline, Typography } from "antd"
import moment from "moment"

import { DeploymentStatus } from "../models"

const { Paragraph, Text, Link } = Typography

interface DeploymentStatusStepsProps {
    statuses: DeploymentStatus[]
}

export default function DeploymentStatusSteps(props: DeploymentStatusStepsProps): JSX.Element {
    return  (
        <Timeline>
            {props.statuses.map((status, idx) => {
                return (
                    <Timeline.Item 
                        color={getStatusColor(status.status)}
                        style={(idx === props.statuses.length - 1)? {paddingBottom: 0} : {}}
                    >
                        <Paragraph style={{margin: 0}}>
                            <Text strong>{status.description}</Text> 
                            {(status.logUrl !== "")? <Link href={status.logUrl}> View</Link> : <></>}<br/>
                            <Text>Updated</Text> <Text code className="gitploy-code">{status.status}</Text> <Text>at {moment(status.createdAt).format('HH:mm:ss')}</Text>
                        </Paragraph>
                    </Timeline.Item>
                )
            })}
        </Timeline>
    )
}

const getStatusColor = (status: string) => {
    switch (status) {
        case "success":
            return "green"
        case "failure":
            return "red"
        default:
            return "#722ed1"
    }
}