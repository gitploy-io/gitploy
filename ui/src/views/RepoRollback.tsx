import { useEffect } from "react";
import { useParams } from "react-router-dom";
import { PageHeader, Result, Button } from 'antd'
import { shallowEqual } from "react-redux";

import { useAppDispatch, useAppSelector } from "../redux/hooks"
import { 
    repoRollbackSlice, 
    fetchConfig, 
    fetchDeployments, 
    searchCandidates, 
    fetchUser,
    rollback,
} from "../redux/repoRollback"

import { User, Deployment, RequestStatus, Env } from '../models'
import RollbackForm from "../components/RollbackForm";

const { actions } = repoRollbackSlice

export interface Params {
    namespace: string
    name: string
}

export default function RepoHome(): JSX.Element {
    const { namespace, name } = useParams<Params>()
    const {
        display,
        config,
        env,
        envs,
        candidates,
        deployments, 
        deploying } = useAppSelector(state => state.repoRollback, shallowEqual)
    const dispatch = useAppDispatch()

    useEffect(() => {
        const f = async () => {
            await dispatch(actions.init({namespace, name}))
            await dispatch(fetchConfig())
            await dispatch(actions.setDisplay(true))
            await dispatch(fetchUser())
            await dispatch(searchCandidates(""))
        }
        f()
        // eslint-disable-next-line
    }, [dispatch])

    const onSelectEnv = (env: Env) => {
        dispatch(actions.setEnv(env))
        dispatch(fetchDeployments())
    }

    const onSelectDeployment = (deployment: Deployment) => {
        dispatch(actions.setDeployment(deployment))
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

    const onClickRollback = () => {
        const f = async () => {
            await dispatch(rollback())
        }
        f()
    }

    if (!display) {
        return <></>
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
                    title="Rollback"
                    subTitle="Restore to a previous deployment."/>
            </div>
            <div style={{padding: "16px 0px"}}>
                <RollbackForm 
                    envs={envs}
                    onSelectEnv={onSelectEnv}
                    deployments={deployments}
                    onSelectDeployment={onSelectDeployment}
                    onClickRollback={onClickRollback}
                    deploying={deploying === RequestStatus.Pending} 
                    approvalEnabled={(env?.approval?.enabled)? true : false}
                    candidates={candidates}
                    onSearchCandidates={onSearchCandidates}
                    onSelectCandidate={onSelectCandidate}
                    onDeselectCandidate={onDeselectCandidate} />
            </div>
        </div>
    )
}