import { useEffect } from "react";
import { useParams } from "react-router-dom";
import { shallowEqual } from "react-redux";
import { PageHeader, Select } from 'antd'

import { useAppSelector, useAppDispatch } from '../redux/hooks'
import { repoHomeSlice, init, fetchEnvs, fetchDeployments, perPage } from '../redux/repoHome'

import ActivityLogs from '../components/ActivityLogs'
import Spin from '../components/Spin'
import Pagination from '../components/Pagination'

const { actions } = repoHomeSlice
const { Option } = Select

interface Params {
    namespace: string
    name: string
}

export default function RepoHome(): JSX.Element {
    const { namespace, name } = useParams<Params>()
    const {
        loading,
        deployments,
        envs,
        page
    } = useAppSelector(state => state.repoHome, shallowEqual)
    const dispatch = useAppDispatch()

    useEffect(() => {
        const f = async () => {
            await dispatch(init({namespace, name}))
            await dispatch(fetchEnvs())
            await dispatch(fetchDeployments())
        }
        f()
        // eslint-disable-next-line
    }, [dispatch])

    const isLast = deployments.length < perPage

    const onChangeEnv = (env: string) => {
        dispatch(actions.setEnv(env))
        dispatch(fetchDeployments())
    }

    const onClickPrev = () => {
        dispatch(actions.decreasePage())
        dispatch(fetchDeployments())
    }

    const onClickNext = () => {
        dispatch(actions.increasePage())
        dispatch(fetchDeployments())
    }

    return (
        <div>
            <div>
                <PageHeader
                    title="Activity Log"
                    extra={[
                        <Select key="1" style={{ width: 150}} defaultValue="" onChange={onChangeEnv}>
                            <Option value="">All Environments</Option>
                            {envs.map((env, idx) => {
                                return <Option key={idx} value={env}>{env}</Option>
                            })}
                        </Select>,
                    ]}
                />
            </div>
            <div style={{marginTop: "30px", padding: "16px 24px"}}>
                {(loading)? 
                    <div style={{textAlign: "center"}}><Spin /></div> :
                    <ActivityLogs deployments={deployments}/>
                }
            </div>
            <div style={{marginTop: "20px", textAlign: "center"}}>
                <Pagination page={page} isLast={isLast} onClickPrev={onClickPrev} onClickNext={onClickNext} ></Pagination>
            </div>
        </div>
    )
}