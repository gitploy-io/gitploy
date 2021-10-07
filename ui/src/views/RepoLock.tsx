import { useEffect } from "react";
import { shallowEqual } from "react-redux";
import { useParams } from "react-router-dom";
import { PageHeader, Button } from 'antd'
import { Result } from "antd"

import { useAppSelector, useAppDispatch } from "../redux/hooks"
import { fetchConfig, listLocks, lock, unlock, repoLockSlice as slice} from "../redux/repoLock"
import LockList from '../components/LockList'

interface Params {
    namespace: string
    name: string
}

export default function RepoLock(): JSX.Element {
    const { namespace, name } = useParams<Params>()
    const { display, config, locks } = useAppSelector(state => state.repoLock, shallowEqual)
    const dispatch = useAppDispatch()

    useEffect(() => {
        const f = async () => {
            await dispatch(slice.actions.init({namespace, name}))
            await dispatch(fetchConfig())
            await dispatch(listLocks())
            await dispatch(slice.actions.setDisplay(true))
        }
        f()
        // eslint-disable-next-line 
    }, [dispatch])

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

    const onClickLock = (env: string) => {
        dispatch(lock(env))
    }

    const onClickUnlock = (env: string) => {
        dispatch(unlock(env))
    }

    return <div>
        <div>
            <PageHeader title="Lock" subTitle="Lock the environment."/>
        </div>
        <div style={{padding: "16px 24px"}}>
            <LockList
                envs={(config)? config.envs:[]}
                locks={locks}
                onClickLock={onClickLock}
                onClickUnlock={onClickUnlock}
            />
        </div>
    </div>
}