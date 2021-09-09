import { useParams } from "react-router-dom";
import { PageHeader, Result, Button } from "antd";
import { shallowEqual } from "react-redux";

import { useAppSelector, useAppDispatch } from "../redux/hooks"
import { User,DeploymentType, Branch, Commit, Tag, RequestStatus, Env } from "../models";
import { 
    init, 
    fetchConfig, 
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
    searchCandidates,
    deploy} from "../redux/repoDeploy"

import DeployForm, {Option} from "../components/DeployForm"
import { useEffect } from "react";

const { actions } = repoDeploySlice

interface Params {
    namespace: string
    name: string
}

export default function RepoDeploy(): JSX.Element {
    const { namespace, name } = useParams<Params>()
    const { 
        display,
        repo,
        config, 
        env,
        envs, 
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
            await dispatch(init({namespace, name}))
            await dispatch(fetchConfig())
            await dispatch(actions.setDisplay(true))
            await dispatch(fetchBranches())
            await dispatch(fetchTags())
        }
        f()
        // eslint-disable-next-line 
    }, [dispatch])

    const onSelectEnv = (env: Env) => {
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

    if (!display || !repo) {
        return <div />
    } 

    if (repo.locked) {
        return (
            <Result
                status="warning"
                title="The repository is locked."
                subTitle="Sorry, you can't deploy until the repository is unlocked."
            />
        )
    }
    
    if (repo && !config) {
        return (
            <Result
                status="warning"
                title="There is no configuration file."
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
                    onChangeType={onChangeType}
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