import { Form, Typography, Avatar, Button, Collapse, Timeline } from "antd"
import moment from "moment"
import { useState } from "react"

import { Deployment, Commit } from "../models"
import DeploymentRefCode from "./DeploymentRefCode"
import DeploymentStatusSteps from "./DeploymentStatusSteps"

const { Paragraph, Text } = Typography
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
                <DeploymentRefCode deployment={props.deployment}/>
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
            <Form.Item
                {...layout}
                label="Deployed At"
            >
                <Text>{moment(props.deployment.createdAt).format("YYYY-MM-DD HH:mm:ss")}</Text>
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
                return <Timeline.Item key={idx} color="gray">
                    <CommitChange commit={change} />
                </Timeline.Item>
            })}
        </Timeline>
    )
}

interface CommitChangeProps {
    commit: Commit
}

function CommitChange(props: CommitChangeProps): JSX.Element {
    const [message, ...description] = props.commit.message.split(/(\r\n|\n|\r)/g)

    const [hide, setHide] = useState(true)

    const onClickHide = () => {
        setHide(!hide)
    }

    return (
        <span >
            <a href={props.commit.htmlUrl} className="gitploy-link"><strong>{message.substring(0, 50)}</strong></a>
            {(description.length) ? 
                <Button size="small" type="text" onClick={onClickHide}>
                    <Text className="gitploy-code" code>...</Text>
                </Button> :
                null}
            {/* Display the description of the commit. */}
            {(!hide) ?
                <Paragraph style={{margin: 0}}>
                    <pre style={{marginBottom: 0, fontSize: 12}}>
                        {description.join("").trim()}
                    </pre>
                </Paragraph> :
                null}
            <br />
            {(props.commit?.author) ?
                <span >
                    <Text >&nbsp;<Avatar size="small" src={props.commit.author.avatarUrl} /> {props.commit.author.login}</Text> <Text >committed {moment(props.commit.author?.date).fromNow()}</Text>
                </span> :
                null} 
        </span>
    )
}