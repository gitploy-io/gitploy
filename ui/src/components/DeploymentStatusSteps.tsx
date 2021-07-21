import { Steps, Popover, Badge } from "antd"

import { Deployment, LastDeploymentStatus } from "../models"

const { Step } = Steps

interface DeploymentStatusStepsProps {
    deployment: Deployment
}

export default function DeploymentStatusSteps(props: DeploymentStatusStepsProps) {
    if (props.deployment.statuses.length === 0) {
        return (
            <Steps size="small" progressDot>
                <Step 
                    title={props.deployment.status} 
                    status={(props.deployment.status === LastDeploymentStatus.Failure)? "error" : "process"}/>
            </Steps>
        )
    }

    return  (
        <Steps 
            current={2}
            size="small" 
            responsive>
            {props.deployment.statuses.map((status, idx) => {
                const title = (!!status.logUrl) ?
                    <a href={status.logUrl}>{status.status}</a> :
                    <span>{status.status}</span>
                return (
                    <Step 
                        key={idx}
                        style={{width: "120px"}}
                        status="finish"
                        icon={<Badge 
                                    color="purple" 
                                    style={{position: "relative", top:"-4px"}}/>
                                }
                        title={<Popover content={status.description}>
                                {title}
                            </Popover>}
                        />
                )
            })}
        </Steps>
    )
}