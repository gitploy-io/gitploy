import { useState } from "react"
import { Form, Select, Radio, Button } from "antd"

import { Branch, Commit, Tag, DeploymentType, Status, User, Env } from "../models"

import CreatableSelect, {Option as Op} from "./CreatableSelect"
import StatusStateIcon from "./StatusStateIcon"
import ApproversSelect from "./ApproversSelect"

export type Option = Op

interface DeployFormProps {
    envs: Env[]
    onSelectEnv(env: Env): void
    onChangeType(type: DeploymentType): void
    branches: Branch[]
    onSelectBranch(branch: Branch): void
    onClickAddBranch(option: Option): void
    branchStatuses: Status[]
    commits: Commit[]
    onSelectCommit(commit: Commit): void
    onClickAddCommit(option: Option): void
    commitStatuses: Status[]
    tags: Tag[]
    onSelectTag(tag: Tag): void
    onClickAddTag(option: Option): void
    tagStatuses: Status[]
    deploying: boolean
    onClickDeploy(): void
    // properties for approvers.
    approvalEnabled: boolean
    candidates: User[]
    onSelectCandidate(candidate: User): void
    onDeselectCandidate(candidate: User): void
    onSearchCandidates(login: string): void
}

export default function DeployForm(props: DeployFormProps): JSX.Element {
    const [ deploymentType, setDeploymentType ] = useState<DeploymentType | null>(null) 

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

    const styleHide: React.CSSProperties = {
        display: "none"
    }

    const styleWidthForCheck: React.CSSProperties = {
        width: "90%"
    }

    const isBranchVisible = (type: DeploymentType | null) => {
        if (type === null) return false
        return type === DeploymentType.Commit || type === DeploymentType.Branch
    }

    const isBranchCheckVisible = (type: DeploymentType | null) => {
        if (type === null) return false
        return type === DeploymentType.Branch
    }

    const isCommitVisible = (type: DeploymentType | null) => {
        if (type === null) return false
        return type === DeploymentType.Commit 
    }

    const isCommitCheckVisible = (type: DeploymentType | null) => {
        if (type === null) return false
        return type === DeploymentType.Commit 
    }

    const isTagVisible = (type: DeploymentType | null) => {
        if (type === null) return false
        return type === DeploymentType.Tag 
    }

    const isTagCheckVisible = (type: DeploymentType | null) => {
        if (type === null) return false
        return type === DeploymentType.Tag
    }

    const onChangeType = (e: any) => {
        let type: DeploymentType

        switch (e.target.value) {
            case "commit":
                type = DeploymentType.Commit
                break
            case "branch":
                type = DeploymentType.Branch
                break
            case "tag":
                type = DeploymentType.Tag
                break
            default:
                type = DeploymentType.Commit
        }

        setDeploymentType(type)
        props.onChangeType(type)
    }

    const mapBranchToOption = (branch: Branch) => {
        return {
            label: branch.name,
            value: branch.name
        } as Option
    }

    const onSelectEnv = (value: string) => {
        const env = props.envs.find(env => env.name === value)
        if (env === undefined) throw new Error("The env doesn't exist.")

        props.onSelectEnv(env)
    }

    const onSelectBranch = (option: Option) => {
        const branch = props.branches.find(b => b.name === option.value)
        if (branch === undefined) throw new Error("The branch doesn't exist.")

        props.onSelectBranch(branch)
    }

    const mapCommitToOption = (commit: Commit) => {
        return {
            label: `${commit.sha.substr(0, 7)} - ${commit.message}`,
            value: commit.sha,
        } as Option
    }

    const onSelectCommit = (option: Option) => {
        const commit = props.commits.find(c => c.sha === option.value)
        if (commit === undefined) throw new Error("The commit doesn't exist.")

        props.onSelectCommit(commit)
    }

    const mapTagToOption = (tag: Tag) => {
        return {
            label: tag.name,
            value: tag.name
        } as Option
    }

    const onSelectTag = (option: Option) => {
        const tag = props.tags.find(t => t.name === option.value)
        if (tag === undefined) throw new Error("The tag doesn't exist.")

        props.onSelectTag(tag)
    }

    const onClickFinish = () => {
        props.onClickDeploy()
    }

    return (
        <Form
            onFinish={onClickFinish}
            name="deploy">
            <Form.Item
                {...selectLayout}
                rules={[{required: true}]}
                label="Environment"
                name="environment">
                <Select 
                    onSelect={onSelectEnv}
                    placeholder="Select target environment">
                    {props.envs.map((env, idx) => {
                        return <Select.Option key={idx} value={env.name}>{env.name}</Select.Option>
                    })}
                </Select>
            </Form.Item>
            <Form.Item
                {...layout}
                rules={[{required: true}]}
                label="Type"
                name="type">
                <Radio.Group 
                    onChange={onChangeType}>
                    <Radio.Button value={DeploymentType.Commit}>Commit</Radio.Button>
                    <Radio.Button value={DeploymentType.Branch}>Branch</Radio.Button>
                    <Radio.Button value={DeploymentType.Tag}>Tag</Radio.Button>
                </Radio.Group>
            </Form.Item>
            {/* https://ant.design/components/form/#components-form-demo-complex-form-control */}
            <Form.Item
                label="Branch"
                {...selectLayout}
                style={(isBranchVisible(deploymentType)? {}: styleHide)}>
                <Form.Item
                    name="branch"
                    rules={[{required: isBranchVisible(deploymentType)}]}
                    noStyle>
                        <CreatableSelect 
                            options={props.branches.map(branch => mapBranchToOption(branch))}
                            onSelectOption={onSelectBranch}
                            onClickAddItem={props.onClickAddBranch}
                            placeholder="Select branch"
                            style={styleWidthForCheck}/>
                </Form.Item>
                <span style={(isBranchCheckVisible(deploymentType)? {}: styleHide)}>
                    &nbsp; <StatusStateIcon statuses={props.branchStatuses} /> 
                </span>
            </Form.Item>
            <Form.Item
                label="Commit"
                {...layout}
                style={(isCommitVisible(deploymentType)? {}: styleHide)}>
                <Form.Item
                    name="commit"
                    rules={[{required: isCommitVisible(deploymentType)}]}
                    noStyle>
                        <CreatableSelect 
                            options={props.commits.map(commit => mapCommitToOption(commit))}
                            onSelectOption={onSelectCommit}
                            onClickAddItem={props.onClickAddCommit}
                            placeholder="Select commit"
                            style={styleWidthForCheck}/>
                </Form.Item>
                <span style={(isCommitCheckVisible(deploymentType)? {}: styleHide)}>
                    &nbsp; <StatusStateIcon statuses={props.commitStatuses} /> 
                </span>
            </Form.Item>
            <Form.Item
                label="Tag"
                {...selectLayout}
                style={(isTagVisible(deploymentType)? {}: styleHide)}>
                <Form.Item
                    name="tag"
                    rules={[{required: isTagVisible(deploymentType)}]}
                    noStyle>
                        <CreatableSelect 
                            options={props.tags.map(tag => mapTagToOption(tag))}
                            onSelectOption={onSelectTag}
                            onClickAddItem={props.onClickAddTag}
                            placeholder="Select commit"
                            style={styleWidthForCheck}/>
                </Form.Item>
                <span style={(isTagCheckVisible(deploymentType)? {}: styleHide)}>
                    &nbsp; <StatusStateIcon statuses={props.tagStatuses} /> 
                </span>
            </Form.Item>
            <Form.Item
                {...layout}
                style={(props.approvalEnabled)? {}: styleHide}
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