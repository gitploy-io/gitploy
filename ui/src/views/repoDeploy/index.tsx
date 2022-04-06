import { useEffect, useState } from "react";
import { shallowEqual } from "react-redux";
import { useParams } from "react-router-dom";
import { PageHeader, Result, Button } from "antd";

import { useAppSelector, useAppDispatch } from "../../redux/hooks"
import { DeploymentType, Branch, Commit, Tag, RequestStatus, Env } from "../../models";
import { 
    fetchConfig, 
    repoDeploySlice, 
    fetchCurrentDeploymentOfEnv,
    fetchBranches, 
    checkBranch,
    addBranchManually, 
    fetchCommits, 
    checkCommit,
    addCommitManually, 
    fetchTags, 
    checkTag,
    addTagManually, 
    searchCandidates,
    fetchUser,
    deploy
} from "../../redux/repoDeploy"
import DeployForm, { DeployFormProps,  Option} from "./DeployForm"
import DynamicPayloadModal, { DynamicPayloadModalProps } from "./DynamicPayloadModal"

const { actions } = repoDeploySlice

export default (): JSX.Element => {
    const { namespace, name } = useParams<{
        namespace: string
        name: string
    }>()

    const { 
        display,
        config, 
        envs, 
        env,
        currentDeployment,
        branches, 
        branchStatuses,
        commits, 
        commitStatuses,
        tags, 
        tagStatuses,
        deploying } = useAppSelector(state => state.repoDeploy, shallowEqual)

    const dispatch = useAppDispatch()

    useEffect(() => {
        const f = async () => {
            await dispatch(actions.init({namespace, name}))
            await dispatch(fetchConfig())
            await dispatch(actions.setDisplay(true))
            await dispatch(fetchBranches())
            await dispatch(fetchTags())
            await dispatch(fetchUser())
            await dispatch(searchCandidates(""))
        }
        f()
        // eslint-disable-next-line 
    }, [dispatch])

    const onSelectEnv = (env: Env) => {
        dispatch(actions.setEnv(env))
        dispatch(fetchCurrentDeploymentOfEnv(env))
    }

    const onChangeType = (type: DeploymentType) => {
        dispatch(actions.setType(type))
    }

    const onSelectBranch = (branch: Branch) => {
        dispatch(actions.setBranch(branch))
        dispatch(checkBranch())
        dispatch(fetchCommits())
    }

    const onClickAddBranch = (option: Option) => {
        dispatch(addBranchManually(option.value))
    }

    const onSelectCommit = (commit: Commit) => {
        dispatch(actions.setCommit(commit))
        dispatch(checkCommit())
    }

    const onClickAddCommit = (option: Option) => {
        dispatch(addCommitManually(option.value))
    }

    const onSelectTag = (tag: Tag) => {
        dispatch(actions.setTag(tag))
        dispatch(checkTag())
    }

    const onClickAddTag = (option: Option) => {
        dispatch(addTagManually(option.value))
    }

    const onClickDeploy = () => {
        dispatch(deploy(null))
    }

    const onClickDeployWithPayload = (values: any) => {
        dispatch(deploy(values))
    }

    if (!display) {
        return <div />
    } 

    if (!config) {
        return (
            <Result
                status="warning"
                title="There is no configuration file."
                extra={[
                    <Button type="primary" key="console" target="_blank" href="https://www.gitploy.io/docs/concepts/deploy.yml">
                      Read Document
                    </Button>,
                    <Button type="link" key="link" target="_blank" href={`/link/${namespace}/${name}/config/new`}>
                      New Configuration
                    </Button>
                ]}
            />
        )
    }
    return (
        <RepoDeploy 
            envs={envs}
            onSelectEnv={onSelectEnv}
            onChangeType={onChangeType}
            currentDeployment={currentDeployment}
            branches={branches}
            onSelectBranch={onSelectBranch}
            onClickAddBranch={onClickAddBranch}
            branchStatuses={branchStatuses}
            commits={commits}
            onSelectCommit={onSelectCommit}
            onClickAddCommit={onClickAddCommit}
            commitStatuses={commitStatuses}
            tags={tags}
            onSelectTag={onSelectTag}
            onClickAddTag={onClickAddTag}
            tagStatuses={tagStatuses}
            deploying={deploying === RequestStatus.Pending}
            onClickDeploy={onClickDeploy}
            env={env}
            onClickOk={onClickDeployWithPayload}
        />
    )
}

interface RepoDeployProps extends 
    DeployFormProps,
    Omit<DynamicPayloadModalProps, "visible" | "env" | "onClickCancel"> {
    env?: Env
}

function RepoDeploy({
    // Properities for the DeployForm component.
    envs, 
    onSelectEnv,
    onChangeType,
    currentDeployment,
    branches,
    onSelectBranch,
    onClickAddBranch,
    branchStatuses,
    commits,
    onSelectCommit,
    onClickAddCommit,
    commitStatuses,
    tags,
    onSelectTag,
    onClickAddTag,
    tagStatuses,
    deploying,
    onClickDeploy,
    // Properties for the DynamicPayloadModal component.
    env,
    onClickOk,
}: RepoDeployProps): JSX.Element {

    const [payloadModalVisible, setPayloadModalVisible] = useState(false);

    const _onClickDeploy = () => {
        if (env?.dynamicPayload?.enabled) {
            setPayloadModalVisible(true)
        } else {
            onClickDeploy()
        }
    }

    const _onClickOk = (values: any) => {
        onClickOk(values)
        setPayloadModalVisible(false)
    }

    const onClickCancel = () => {
        setPayloadModalVisible(false)
    }

    return (
        <div>
            <div>
                <PageHeader
                    title="Deploy"
                />
            </div>
            <div style={{padding: "16px 0px"}}>
                <DeployForm 
                    envs={envs}
                    onSelectEnv={onSelectEnv}
                    onChangeType={onChangeType}
                    currentDeployment={currentDeployment}
                    branches={branches}
                    onSelectBranch={onSelectBranch}
                    onClickAddBranch={onClickAddBranch}
                    branchStatuses={branchStatuses}
                    commits={commits}
                    onSelectCommit={onSelectCommit}
                    onClickAddCommit={onClickAddCommit}
                    commitStatuses={commitStatuses}
                    tags={tags}
                    onSelectTag={onSelectTag}
                    onClickAddTag={onClickAddTag}
                    tagStatuses={tagStatuses}
                    deploying={deploying}
                    onClickDeploy={_onClickDeploy} 
                />
                {(env)? 
                    <DynamicPayloadModal
                        visible={payloadModalVisible}
                        env={env}
                        onClickOk={_onClickOk}
                        onClickCancel={onClickCancel}
                    />
                    :
                    <></>
                }
            </div>
        </div>
    )
}