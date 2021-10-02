import { Steps, Popover } from "antd"
import moment from "moment"

import { Deployment } from "../models"
import DeploymentStatusBadge from "./DeploymentStatusBadge"

const { Step } = Steps

interface DeploymentStatusStepsProps {
    deployment: Deployment
}

export default function DeploymentStatusSteps(props: DeploymentStatusStepsProps): JSX.Element {
    if (typeof props.deployment.statuses === "undefined" 
        || props.deployment.statuses.length === 0) {
        return (
            <DeploymentStatusBadge deployment={props.deployment}/>
        )
    }

    return  (
        <Steps 
            current={props.deployment.statuses.length - 1}
            size="small" 
            responsive
        >
            {props.deployment.statuses.map((status, idx) => {
                const title = (status.logUrl)?
                    <a href={status.logUrl}>{status.status}</a> :
                    <span>{status.status}</span>

                const description = (status.description)?
                    `${status.description.replace(/\.$/,' ')} at ${moment(status.createdAt).format('HH:mm:ss')}` :
                    moment(status.createdAt).format('HH:mm:ss')

                return (<Step 
                            key={idx}
                            style={{width: "100px"}}
                            status="finish"
                            icon={<span>•</span>}
                            title={<Popover content={description}>{title}</Popover>}
                        />)
            })}
        </Steps>
    )
}