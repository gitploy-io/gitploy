import { useParams } from "react-router-dom";
import { PageHeader, Result, Button } from "antd";
import { shallowEqual } from "react-redux";

import { useAppSelector, useAppDispatch } from "../redux/hooks"
import { DeploymentType, Branch, Commit, Tag, RequestStatus } from "../models";
import { 
    init, 
    fetchEnvs, 
    repoDeploySlice, 
    fetchBranches, 
    checkBranch,
    addBranchManually, 
    fetchCommits, 
    checkCommit,
    addCommitManually, 
    fetchTags, 
    checkTag,
    addTagManually, 
    deploy} from '../redux/repoDeploy'

import DeployForm, {Option} from '../components/DeployForm'
import { useEffect } from "react";

const { actions } = repoDeploySlice

interface Params {
    namespace: string
    name: string
}

export default function RepoDeploy() {
    let { namespace, name } = useParams<Params>()
    const { 
        hasConfig, 
        envs, 
        type, 
        branches, 
        branchCheck,
        commits, 
        commitCheck,
        tags, 
        tagCheck,
        deploying } = useAppSelector(state => state.repoDeploy, shallowEqual)
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
        dispatch(deploy())
    }

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
                    branchCheck={branchCheck}
                    commits={commits}
                    onSelectCommit={onSelectCommit}
                    onClickAddCommit={onClickAddCommit}
                    commitCheck={commitCheck}
                    tags={tags}
                    onSelectTag={onSelectTag}
                    onClickAddTag={onClickAddTag}
                    tagCheck={tagCheck}
                    deploying={deploying === RequestStatus.Pending}
                    onClickDeploy={onClickDeploy} />
            </div>
        </div>
    )
}