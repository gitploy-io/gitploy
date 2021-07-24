import { Form, Typography, Avatar, Button } from "antd"

import { Deployment } from "../models"
import DeploymentStatusSteps from "./DeploymentStatusSteps"

const { Text } = Typography

interface DeployConfirmProps {
    isDeployable: boolean
    deploying: boolean
    deployment: Deployment
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
                <Text code>{props.deployment.env}</Text>
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
                <Text>
                    {(props.deployment.deployer !== null)?
                         <Text >
                            <Avatar size="small" src={props.deployment.deployer.avatar} /> {props.deployment.deployer.login}
                        </Text> :
                        <Avatar size="small" >
                            U
                        </Avatar>}
                </Text>
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