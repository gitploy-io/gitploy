import { useParams } from "react-router-dom";
import { PageHeader, Result, Button, message } from "antd";
import { shallowEqual } from "react-redux";

import { useAppSelector, useAppDispatch } from "../redux/hooks"
import { init, fetchEnvs, repoDeploySlice, fetchBranches, addBranchManually, fetchCommits, addCommitManually, fetchTags, addTagManually, deploy} from '../redux/repoDeploy'
import { DeploymentType, Branch, Commit, Tag, RequestStatus } from "../models";

import DeployForm, {Option} from '../components/DeployForm'
import { useEffect } from "react";

const { actions } = repoDeploySlice

interface Params {
    namespace: string
    name: string
}

export default function RepoDeploy() {
    let { namespace, name } = useParams<Params>()
    const { hasConfig, envs, type, branches, commits, tags, adding, deploying } = useAppSelector(state => state.repoDeploy, shallowEqual)
    const dispatch = useAppDispatch()

    useEffect(() => {
        const f = async () => {
            await dispatch(init({namespace, name}))
            await dispatch(fetchEnvs())
            await dispatch(fetchBranches())
            await dispatch(fetchTags())
        }
        f()
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [dispatch])

    const onSelectEnv = (env: string) => {
        dispatch(actions.setEnv(env))
    }

    const onChangeType = (type: DeploymentType) => {
        dispatch(actions.setType(type))
    }

    const onSelectBranch = (branch: Branch) => {
        dispatch(actions.setBranch(branch))
        dispatch(fetchCommits())
    }

    const onClickAddBranch = (option: Option) => {
        dispatch(addBranchManually(option.value))
    }

    const onSelectCommit = (commit: Commit) => {
        dispatch(actions.setCommit(commit))
    }

    const onClickAddCommit = (option: Option) => {
        dispatch(addCommitManually(option.value))
    }

    const onSelectTag = (tag: Tag) => {
        dispatch(actions.setTag(tag))
    }

    const onClickAddTag = (option: Option) => {
        dispatch(addTagManually(option.value))
    }

    const handleAddManuallyStatus = () => {
        if (adding === RequestStatus.Failure) {
            message.error("It has failed to add the item. Check Ref is correct.")
            dispatch(actions.unsetAddManually())
        }
    }


    const onClickDeploy = () => {
        dispatch(deploy())
    }

    const handleDeployStatus = () => {
        if (deploying === RequestStatus.Failure) {
            message.error("It has failed to deploy.")
            dispatch(actions.unsetDeploy())
            return 
        } else if (deploying === RequestStatus.Success) {
            message.success("It starts to deploy.", 3)
            dispatch(actions.unsetDeploy())
            return
        }
    }

    handleAddManuallyStatus()
    handleDeployStatus()

    if (!hasConfig) {
        return (
            <Result
                status="warning"
                title="There isn't the configuration file."
                extra={
                    <Button type="primary" key="console">
                      Read Document
                    </Button>
                }
            />
        )
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
                    type={type}
                    onChangeType={onChangeType}
                    branches={branches}
                    onSelectBranch={onSelectBranch}
                    onClickAddBranch={onClickAddBranch}
                    commits={commits}
                    onSelectCommit={onSelectCommit}
                    onClickAddCommit={onClickAddCommit}
                    tags={tags}
                    onSelectTag={onSelectTag}
                    onClickAddTag={onClickAddTag}
                    deploying={deploying === RequestStatus.Pending}
                    onClickDeploy={onClickDeploy} />
            </div>
        </div>
    )
}