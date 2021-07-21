import { Steps, Popover, Badge } from "antd"

import { Deployment } from "../models"
import DeploymentStatusBadge from "./DeploymentStatusBadge"

const { Step } = Steps

interface DeploymentStatusStepsProps {
    deployment: Deployment
}

export default function DeploymentStatusSteps(props: DeploymentStatusStepsProps) {
    if (props.deployment.statuses.length === 0) {
        return (
            <DeploymentStatusBadge deployment={props.deployment}/>
        )
    }

    return  (
        <Steps 
            current={props.deployment.statuses.length - 1}
            size="small" 
            responsive>
            {props.deployment.statuses.map((status, idx) => {
                const title = (!!status.logUrl) ?
                    <a href={status.logUrl}>{status.status}</a> :
                    <span>{status.status}</span>
                return (<Step 
                        key={idx}
                        style={{width: "120px"}}
                        status="finish"
                        icon={<Badge 
                                color={guessColor(status.status)} 
                                style={{position: "relative", top:"-4px"}}
                                />}
                        title={<Popover content={status.description}>
                                {title}
                            </Popover>}
                        />)
            })}
        </Steps>
    )
}

// The deployment status has arbitrary status because it is decided by SCM.
const guessColor = (status: string) => {
    switch (status) {
        case "success":
            return "green"
        case "failure":
            return "red"
        default:
            return "purple"
    }
}