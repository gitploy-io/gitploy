import { useEffect } from 'react'
import { shallowEqual } from 'react-redux'
import { Input, Breadcrumb } from 'antd'

import { useAppSelector, useAppDispatch } from '../redux/hooks'
import { homeSlice, listRepos, perPage } from '../redux/home'

import Main from './Main'
import RepoList from '../components/RepoList'
import Pagination from '../components/Pagination'
import Spin from '../components/Spin'

const { Search } = Input
const { actions } = homeSlice

export default function Home(){
    const { loading, repos, page } = useAppSelector(state => state.home, shallowEqual)
    const dispatch = useAppDispatch()

    useEffect(() => {
        dispatch(listRepos())
    }, [dispatch])

    const search = (q: string) => {
        dispatch(actions.setQ(q))
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

    const isLast = repos.length < perPage

    if (loading) {
        return (
            <Main>
                <div style={{"marginTop": "20px"}}>
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
            <div style={{"marginTop": "20px"}}>
                <Breadcrumb>
                    <Breadcrumb.Item>
                        <a href="/">Repositories</a>
                    </Breadcrumb.Item>
                </Breadcrumb>
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
