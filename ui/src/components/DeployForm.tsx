import { Form, Select, Radio, Button } from 'antd'

import { Branch, Commit, Tag, DeploymentType, StatusState } from '../models'

import CreatableSelect, {Option as Op} from './CreatableSelect'
import StatusStateIcon from './StatusStateIcon'

export type Option = Op

interface DeployFormProps {
    envs: string[]
    onSelectEnv(env: string): void
    type: DeploymentType | null
    onChangeType(type: DeploymentType): void
    branches: Branch[]
    onSelectBranch(branch: Branch): void
    onClickAddBranch(option: Option): void
    branchCheck: StatusState
    commits: Commit[]
    onSelectCommit(commit: Commit): void
    onClickAddCommit(option: Option): void
    commitCheck: StatusState
    tags: Tag[]
    onSelectTag(tag: Tag): void
    onClickAddTag(option: Option): void
    tagCheck: StatusState
    deploying: boolean
    onClickDeploy(): void
}

export default function DeployForm(props: DeployFormProps) {
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

        props.onChangeType(type)
    }

    const mapBranchToOption = (branch: Branch) => {
        return {
            label: branch.name,
            value: branch.name
        } as Option
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

    const onClickFinish = (values: any) => {
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
                    onSelect={props.onSelectEnv}
                    placeholder="Select target environment">
                    {props.envs.map((env, idx) => {
                        return <Select.Option key={idx} value={env}>{env}</Select.Option>
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
                style={(isBranchVisible(props.type)? {}: styleHide)}>
                <Form.Item
                    name="branch"
                    rules={[{required: isBranchVisible(props.type)}]}
                    noStyle>
                        <CreatableSelect 
                            options={props.branches.map(branch => mapBranchToOption(branch))}
                            onSelectOption={onSelectBranch}
                            onClickAddItem={props.onClickAddBranch}
                            placeholder="Select branch"
                            style={styleWidthForCheck}/>
                </Form.Item>
                <span style={(isBranchCheckVisible(props.type)? {}: styleHide)}>
                    &nbsp; <StatusStateIcon state={props.branchCheck} /> 
                </span>
            </Form.Item>
            <Form.Item
                label="Commit"
                {...layout}
                style={(isCommitVisible(props.type)? {}: styleHide)}>
                <Form.Item
                    name="commit"
                    rules={[{required: isCommitVisible(props.type)}]}
                    noStyle>
                        <CreatableSelect 
                            options={props.commits.map(commit => mapCommitToOption(commit))}
                            onSelectOption={onSelectCommit}
                            onClickAddItem={props.onClickAddCommit}
                            placeholder="Select commit"
                            style={styleWidthForCheck}/>
                </Form.Item>
                <span style={(isCommitCheckVisible(props.type)? {}: styleHide)}>
                    &nbsp; <StatusStateIcon state={props.commitCheck} /> 
                </span>
            </Form.Item>
            <Form.Item
                label="Tag"
                {...selectLayout}
                style={(isTagVisible(props.type)? {}: styleHide)}>
                <Form.Item
                    name="tag"
                    rules={[{required: isTagVisible(props.type)}]}
                    noStyle>
                        <CreatableSelect 
                            options={props.tags.map(tag => mapTagToOption(tag))}
                            onSelectOption={onSelectTag}
                            onClickAddItem={props.onClickAddTag}
                            placeholder="Select commit"
                            style={styleWidthForCheck}/>
                </Form.Item>
                <span style={(isTagCheckVisible(props.type)? {}: styleHide)}>
                    &nbsp; <StatusStateIcon state={props.tagCheck} /> 
                </span>
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