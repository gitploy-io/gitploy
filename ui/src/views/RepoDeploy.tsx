import { useEffect } from "react";
import { shallowEqual } from "react-redux";
import { useParams } from "react-router-dom";
import { PageHeader, Result, Button } from "antd";

import { useAppSelector, useAppDispatch } from "../redux/hooks"
import { User,DeploymentType, Branch, Commit, Tag, RequestStatus, Env } from "../models";
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
    deploy} from "../redux/repoDeploy"

import DeployForm, {Option} from "../components/DeployForm"

const { actions } = repoDeploySlice

interface Params {
    namespace: string
    name: string
}

export default function RepoDeploy(): JSX.Element {
    const { namespace, name } = useParams<Params>()
    const { 
        display,
        config, 
        env,
        envs, 
        currentDeployment,
        branches, 
        branchStatuses,
        commits, 
        commitStatuses,
        tags, 
        tagStatuses,
        candidates,
        deploying } = useAppSelector(state => state.repoDeploy, shallowEqual)
    const dispatch = useAppDispatch()

    useEffect(() => {
        const f = async () => {
            await dispatch(actions.init({namespace, name}))
            await dispatch(fetchConfig())
            await dispatch(actions.setDisplay(true))
            await dispatch(fetchBranches())
            await dispatch(fetchTags())
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

    const onSearchCandidates = (login: string) => {
        dispatch(searchCandidates(login))
    }

    const onSelectCandidate = (candidate: User) => {
        dispatch(actions.addApprover(candidate))
    }

    const onDeselectCandidate = (candidate: User) => {
        dispatch(actions.deleteApprover(candidate))
    }

    const onClickDeploy = () => {
        const f = async () => {
            await dispatch(deploy())
        }
        f()
    }

    if (!display) {
        return <div />
    } 

    if (!config) {
        return (
            <Result
                status="warning"
                title="There is no configuration file."
                extra={
                    <Button type="primary" key="console" href="https://docs.gitploy.io/concepts/deploy.yml">
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
                    approvalEnabled={(env?.approval?.enabled)? true : false}
                    candidates={candidates}
                    onSearchCandidates={onSearchCandidates}
                    onSelectCandidate={onSelectCandidate}
                    onDeselectCandidate={onDeselectCandidate}
                />
            </div>
        </div>
    )
}