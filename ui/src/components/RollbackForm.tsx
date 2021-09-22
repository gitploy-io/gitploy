import { Form, Select, Button, Avatar } from 'antd'
import moment from 'moment'

import { User, Deployment, DeploymentType, Env } from "../models"
import ApproversSelect from "./ApproversSelect"
import DeploymentRefCode from './DeploymentRefCode'

interface RollbackFormProps {
    envs: Env[]
    onSelectEnv(env: Env): void
    deployments: Deployment[]
    onSelectDeployment(deployment: Deployment): void
    onClickRollback(): void
    deploying: boolean
    approvalEnabled: boolean
    candidates: User[]
    onSelectCandidate(candidate: User): void
    onDeselectCandidate(candidate: User): void
    onSearchCandidates(login: string): void
}

export default function RollbackForm(props: RollbackFormProps): JSX.Element {
    const layout = {
      labelCol: { span: 5},
      wrapperCol: { span: 16 },
    };

    const selectLayout = {
        ...layout,
        wrapperCol: {span: 10}
    }

    const submitLayout = {
      wrapperCol: { offset: 5, span: 16 },
    };

    const onSelectEnv = (value: string) => {
        const env = props.envs.find((e) => e.name === value)
        if (env === undefined) throw new Error("The deployment doesn't exist.")

        props.onSelectEnv(env)
    }

    const onSelectDeployment = (id: number) => {
        const deployment = props.deployments.find((d) => (d.id === id))
        if (deployment === undefined) throw new Error("The deployment doesn't exist.")

        props.onSelectDeployment(deployment)
    }
    
    const onFinish = () => {
        props.onClickRollback()
    }

    return (
        <Form
            name="rollback"
            onFinish={onFinish} >
            <Form.Item
                label="Environment"
                name="environment"
                {...selectLayout}
                rules={[{required: true}]} >
                <Select 
                    onSelect={onSelectEnv}
                    placeholder="Select target environment">
                        {props.envs.map((env, idx) => {
                            return <Select.Option key={idx} value={env.name}>{env.name}</Select.Option>
                        })}
                </Select>
            </Form.Item>
            <Form.Item
                label="Deployment"
                name="deployment"
                {...layout} 
                rules={[{required: true}]} >
                <Select 
                    onSelect={onSelectDeployment}
                    placeholder="Select the deployment">
                        {props.deployments.map((d, idx) => {
                            let option: React.ReactElement
                            const ref = (d.type === DeploymentType.Commit)? d.sha.substr(0, 7) : d.ref

                            if (d.deployer) {
                                option = <Select.Option key={idx} value={d.id}>
                                   #{d.number} - <DeploymentRefCode deployment={d}/> deployed by <Avatar size="small" src={d.deployer.avatar} /> {d.deployer.login} {moment(d.createdAt).fromNow()}
                                </Select.Option>
                            } else {
                                option = <Select.Option key={idx} value={d.id}>
                                   #{d.number} - <DeploymentRefCode deployment={d}/> deployed {moment(d.createdAt).fromNow()}
                                </Select.Option>
                            }
                            return option
                        })}
                </Select>
            </Form.Item>
            <Form.Item
                {...layout}
                style={(props.approvalEnabled)? {}: {display: "none"}}
                label="Approvers"
                name="approvers">
                    <ApproversSelect 
                        candidates={props.candidates}
                        onSearchCandidates={props.onSearchCandidates}
                        onSelectCandidate={props.onSelectCandidate}
                        onDeselectCandidate={props.onDeselectCandidate} />
            </Form.Item> 
            <Form.Item {...submitLayout}>
                <Button 
                    loading={props.deploying}
                    type="primary" 
                    htmlType="submit">
                  Submit
                </Button>
            </Form.Item>
        </Form>
    )
}