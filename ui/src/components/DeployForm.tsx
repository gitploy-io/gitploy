import { Form, Select, Radio, Button } from 'antd'

import { Branch, Commit, Tag, DeploymentType } from '../models'
import CreatableSelect, {Option} from './CreatableSelect'

interface DeployFormProps {
    envs: string[]
    onSelectEnv(env: string): void
    type: DeploymentType | undefined
    onChangeType(type: DeploymentType): void
    branches: Branch[]
    onSelectBranch(branch: Branch): void
    onClickAddBranch(option: Option): void
    commits: Commit[]
    onSelectCommit(commit: Commit): void
    onClickAddCommit(option: Option): void
    // TODO: support tag
    // tags: Tag[]
    // onSelectTag(tag: Tag): void
    // onClickAddTag(option: Option): void
}

export default function DeployForm(props: DeployFormProps) {
    const layout = {
      labelCol: { span: 3 },
      wrapperCol: { span: 16 },
    };

    const submitLayout = {
      wrapperCol: { offset: 3, span: 16 },
    };

    const hide: React.CSSProperties = {
        display: "none"
    }

    const isBranchVisible = (type: DeploymentType | undefined) => {
        if (type === undefined) return false
        return type === DeploymentType.Commit || type === DeploymentType.Branch
    }

    const isCommitVisible = (type: DeploymentType | undefined) => {
        if (type === undefined) return false
        return type === DeploymentType.Commit 
    }

    const isTagVisible = (type: DeploymentType | undefined) => {
        if (type === undefined) return false
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
        const branch = props.branches.find(b => b.name == option.value)

        if (branch === undefined) throw new Error("The branch doesn't exist.")

        props.onSelectBranch(branch)
    }

    const mapCommitToOption = (commit: Commit) => {
        return {
            label: `${commit.sha.substr(7)} - ${commit.message}`,
            value: commit.sha,
        } as Option
    }

    const onSelectCommit = (option: Option) => {
        const commit = props.commits.find(c => c.sha == option.value)

        if (commit === undefined) throw new Error("The commit doesn't exist.")

        props.onSelectCommit(commit)
    }

    return (
        <Form
            name="deploy">
            <Form.Item
                {...layout}
                wrapperCol={{span: 8}}
                label="Environment"
                name="env">
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
                label="Type"
                name="type">
                <Radio.Group 
                    onChange={onChangeType}>
                    <Radio.Button value={DeploymentType.Commit}>Commit</Radio.Button>
                    <Radio.Button value={DeploymentType.Branch}>Branch</Radio.Button>
                    <Radio.Button value={DeploymentType.Tag}>Tag</Radio.Button>
                </Radio.Group>
            </Form.Item>
            <Form.Item
                {...layout}
                wrapperCol={{span: 8}}
                style={(isBranchVisible(props.type)? {}: hide)}
                label="Branch"
                name="branch">
                    <CreatableSelect 
                        options={props.branches.map(branch => mapBranchToOption(branch))}
                        onSelectOption={onSelectBranch}
                        onClickAddItem={props.onClickAddBranch}
                        placeholder="Select branch"/>
            </Form.Item>
            <Form.Item
                {...layout}
                style={(isCommitVisible(props.type)? {}: hide)}
                label="Commit"
                name="commit">
                    <CreatableSelect 
                        options={props.commits.map(commit => mapCommitToOption(commit))}
                        onSelectOption={onSelectCommit}
                        onClickAddItem={props.onClickAddCommit}
                        placeholder="Select commit"/>
            </Form.Item>
            <Form.Item {...submitLayout}>
                <Button type="primary" htmlType="submit">
                  Submit
                </Button>
            </Form.Item>
        </Form>
    )
}