import { Form, Typography, Avatar, Button, Collapse, Timeline } from "antd"

import { Deployment, DeploymentType, Commit } from "../models"
import DeploymentStatusSteps from "./DeploymentStatusSteps"

const { Text } = Typography
const { Panel } = Collapse

interface DeployConfirmProps {
    isDeployable: boolean
    deploying: boolean
    deployment: Deployment
    changes: Commit[]
    onClickDeploy(): void
}

export default function DeployConfirm(props: DeployConfirmProps): JSX.Element {
    const layout = {
      labelCol: { span: 6},
      wrapperCol: { span: 16 },
      style: {marginBottom: 12}
    };
    const submitLayout = {
      wrapperCol: { offset: 6, span: 16 },
    };

    // Form makes it to display organized.
    return (
        <Form
            name="confirm"
        >
            <Form.Item
                {...layout}
                label="Target"
            >
                <Text>{props.deployment.env}</Text>
            </Form.Item>
            <Form.Item
                {...layout}
                label="Ref"
            >
                <Text code>{(props.deployment.type === DeploymentType.Commit)? props.deployment.ref.substr(0, 7) : props.deployment.ref}</Text>
            </Form.Item>
            <Form.Item
                {...layout}
                label="Status"
            >
                <DeploymentStatusSteps deployment={props.deployment}/>
            </Form.Item>
            <Form.Item
                {...layout}
                label="Deployer"
            >
                {(props.deployment.deployer)?
                     <Text >
                        <Avatar size="small" src={props.deployment.deployer.avatar} /> {props.deployment.deployer.login}
                    </Text> :
                    <Avatar size="small" >
                        U
                    </Avatar>}
            </Form.Item>
            {(props.deployment.isApprovalEanbled) ?
                <Form.Item
                    {...layout}
                    label="Approval"
                >
                    <Text>{props.deployment.requiredApprovalCount}</Text>
                </Form.Item>: 
                null }
            <Form.Item
                {...layout}
                label="Changes"
            >
                <Collapse ghost >
                    <Panel 
                        key={1} 
                        header="Click" 
                        // Fix the position to align with the field.
                        style={{position: "relative", top: "-5px", left: "-15px"}}
                    >
                        <CommitChanges changes={props.changes}/>
                    </Panel>
                </Collapse>
            </Form.Item>
            <Form.Item 
                {...submitLayout}
            >
                {(props.isDeployable)? 
                    <Button 
                        loading={props.deploying}
                        type="primary" 
                        onClick={props.onClickDeploy}
                        >
                      Deploy
                    </Button>:
                    <Button 
                        type="primary" 
                        disabled>
                      Deploy
                    </Button>}
            </Form.Item>
        </Form>
    )
}

interface CommitChangesProps {
    changes: Commit[]
}

function CommitChanges(props: CommitChangesProps): JSX.Element {
    if (props.changes.length === 0) {
        return <div>There are no commits.</div>
    }
    return (
        <Timeline>
            {props.changes.slice(0, 10).map((change, idx) => {
                const style: React.CSSProperties =  (idx === props.changes.length - 1) ?  {height: 0} : {}
                // Omit lines after the first feedline.
                const message = change.message.split("\n", 1)[0]

                return <Timeline.Item key={idx} color="gray" style={style}>
                    <a href={change.htmlUrl} className="gitploy-link">
                        {message.substr(0, 50)}
                    </a>
                </Timeline.Item>
            })}
        </Timeline>
    )
}