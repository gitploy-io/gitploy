import { useParams } from "react-router-dom";
import { PageHeader, Result, Button } from "antd";
import { shallowEqual } from "react-redux";

import { useAppSelector, useAppDispatch } from "../redux/hooks"
import { init, fetchEnvs, repoDeploySlice, fetchBranches, fetchCommits, fetchTags} from '../redux/repoDeploy'
import { DeploymentType, Branch, Commit, Tag } from "../models";

import DeployForm, {Option} from '../components/DeployForm'
import { useEffect } from "react";

const { actions } = repoDeploySlice

interface Params {
    namespace: string
    name: string
}

export default function RepoDeploy() {
    let { namespace, name } = useParams<Params>()
    const { hasConfig, envs, type, branches, commits, tags } = useAppSelector(state => state.repoDeploy, shallowEqual)
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
        const branch: Branch = {
            name: option.value,
            commitSha: "",
        }
        dispatch(actions.addBranchManually(branch))
    }

    const onSelectCommit = (commit: Commit) => {
        dispatch(actions.setCommit(commit))
    }

    const onClickAddCommit = (option: Option) => {
        const commit: Commit = {
            sha: option.value,
            message: "Manually added",
            isPullRequest: false,
        }
        dispatch(actions.addCommitManually(commit))
    }

    const onSelectTag = (tag: Tag) => {
        dispatch(actions.setTag(tag))
    }

    const onClickAddTag = (option: Option) => {
        const tag: Tag = {
            name: option.value,
            commitSha: "",
        }
        dispatch(actions.addTagManually(tag))
    }

    if (!hasConfig) {
        return (
            <Result
                status="warning"
                title="There is no the configuration file."
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
            <div style={{padding: "16px 24px"}}>
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
                    onClickAddTag={onClickAddTag}/>
            </div>
        </div>
    )
}