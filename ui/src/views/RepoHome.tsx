import { useParams } from "react-router-dom";
import { shallowEqual } from "react-redux";
import { PageHeader, Select, Divider } from 'antd'

import { useAppSelector, useAppDispatch } from '../redux/hooks'
import { repoHomeSlice, init, fetchDeployments, perPage } from '../redux/repoHome'

import ActivityLogs from '../components/ActivityLogs'
import Spin from '../components/Spin'
import Pagination from '../components/Pagination'
import { useEffect } from "react";

const { actions } = repoHomeSlice
const { Option } = Select

interface Params {
    namespace: string
    name: string
}

export default function RepoHome() {
    let { namespace, name } = useParams<Params>()
    const {
        loading,
        deployments,
        page
    } = useAppSelector(state => state.repoHome, shallowEqual)
    const dispatch = useAppDispatch()

    useEffect(() => {
        const f = async () => {
            await dispatch(init({namespace, name}))
            await dispatch(fetchDeployments())
        }
        f()
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [dispatch])

    const isLast = deployments.length < perPage

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
                        <Select key="1" style={{ width: 150}} defaultValue="">
                            <Option value="">All Environments</Option>
                            <Option value="dev">dev</Option>
                        </Select>,
                    ]}
                />
            </div>
            <Divider />
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