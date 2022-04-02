import { useEffect } from 'react'
import { shallowEqual } from 'react-redux'
import { Input, Breadcrumb } from 'antd'
import { Helmet } from "react-helmet"

import { useAppSelector, useAppDispatch } from '../redux/hooks'
import { homeSlice, listRepos, perPage, sync, homeSlice as slice } from '../redux/home'
import { RequestStatus } from '../models'
import { subscribeEvents } from "../apis"

import Main from './main'
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

        const sub = subscribeEvents((event) => {
            dispatch(slice.actions.handleDeploymentEvent(event))
        })

        return () => {
            sub.close()
        }
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
        const f = async () => {
            await dispatch(sync())
            await dispatch(actions.setFirstPage())
            await dispatch(listRepos())
        }
        f()
    }

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
            <Helmet>
                <title>Home</title>
            </Helmet>
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
                <Pagination 
                    disabledPrev={page <= 1} 
                    disabledNext={repos.length < perPage} 
                    onClickPrev={onClickPrev} 
                    onClickNext={onClickNext} 
                />
            </div>
        </Main>
    )
}
