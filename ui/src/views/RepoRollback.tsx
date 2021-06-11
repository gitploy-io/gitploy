import { useEffect } from "react";
import { useParams } from "react-router-dom";
import { PageHeader, Result, Button, message } from 'antd'
import { shallowEqual } from "react-redux";

import { useAppDispatch, useAppSelector } from "../redux/hooks"
import { repoRollbackSlice, init, fetchEnvs, fetchDeployments, rollback } from "../redux/repoRollback"

import { Deployment, RequestStatus } from '../models'
import RollbackForm from "../components/RollbackForm";

const { actions } = repoRollbackSlice

export interface Params {
    namespace: string
    name: string
}

export default function RepoHome() {
    let { namespace, name } = useParams<Params>()
    const {
        hasConfig,
        envs,
        deployments, 
        deploying } = useAppSelector(state => state.repoRollback, shallowEqual)
    const dispatch = useAppDispatch()

    useEffect(() => {
        const f = async () => {
            await dispatch(init({namespace, name}))
            await dispatch(fetchEnvs())
        }
        f()
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [dispatch])

    const onSelectEnv = (env: string) => {
        dispatch(actions.setEnv(env))
        dispatch(fetchDeployments())
    }

    const onSelectDeployment = (deployment: Deployment) => {
        dispatch(actions.setDeployment(deployment))
    }

    const onClickRollback = () => {
        dispatch(rollback())
    }

    const handleRollbackStatus = () => {
        if (deploying === RequestStatus.Failure) {
            message.error("It has failed to rollback.", 3)
            dispatch(actions.unsetDeploy())
            return 
        } else if (deploying === RequestStatus.Success) {
            message.success("It starts to rollback.", 3)
            dispatch(actions.unsetDeploy())
            return
        }
    }

    handleRollbackStatus()

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
                    deploying={deploying === RequestStatus.Pending} />
            </div>
        </div>
    )
}