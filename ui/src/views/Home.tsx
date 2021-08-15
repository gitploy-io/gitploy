import { useEffect } from 'react'
import { shallowEqual } from 'react-redux'
import { Input, Breadcrumb } from 'antd'

import { useAppSelector, useAppDispatch } from '../redux/hooks'
import { homeSlice, listRepos, perPage, sync } from '../redux/home'
import { RequestStatus } from '../models'

import Main from './Main'
import SyncButton from "../components/SyncButton"
import RepoList from '../components/RepoList'
import Pagination from '../components/Pagination'

const { Search } = Input
const { actions } = homeSlice

export default function Home(): JSX.Element {
    const { loading, repos, page, syncing } = useAppSelector(state => state.home, shallowEqual)
    const dispatch = useAppDispatch()

    useEffect(() => {
        dispatch(listRepos())
    }, [dispatch])

    const search = (q: string) => {
        dispatch(actions.setQ(q))
        dispatch(actions.setFirstPage())
        dispatch(listRepos())
    }

    const onClickPrev = () => {
        dispatch(actions.decreasePage())
        dispatch(listRepos())
    }

    const onClickNext = () => {
        dispatch(actions.increasePage())
        dispatch(listRepos())
    }

    const onClickSync = () => {
        dispatch(sync())
        dispatch(actions.setFirstPage())
        dispatch(listRepos())
    }

    const isLast = repos.length < perPage

    if (loading) {
        return (
            <Main>
                <div >
                    <Breadcrumb>
                        <Breadcrumb.Item>
                            <a href="/">Repositories</a>
                        </Breadcrumb.Item>
                    </Breadcrumb>
                </div>
            </Main>
        )
    }

    return (
        <Main>
            <div >
                <Breadcrumb>
                    <Breadcrumb.Item>
                        <a href="/">Repositories</a>
                    </Breadcrumb.Item>
                </Breadcrumb>
            </div>
            <div style={{textAlign: "right"}}>
                <SyncButton loading={syncing === RequestStatus.Pending} onClickSync={onClickSync}></SyncButton>
            </div>
            <div style={{"marginTop": "20px"}}>
                <Search placeholder="Search repository ..." onSearch={search} size="large" enterButton />
            </div>
            <div style={{"marginTop": "20px"}}>
                <RepoList repos={repos}></RepoList>
            </div>
            <div style={{marginTop: "20px", textAlign: "center"}}>
                <Pagination page={page} isLast={isLast} onClickPrev={onClickPrev} onClickNext={onClickNext} ></Pagination>
            </div>
        </Main>
    )
}
