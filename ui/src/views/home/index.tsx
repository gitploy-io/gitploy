import { useEffect } from 'react'
import { shallowEqual } from 'react-redux'
import { Helmet } from "react-helmet"
import { Input, Breadcrumb, Button } from 'antd'
import { RedoOutlined } from "@ant-design/icons"

import { useAppSelector, useAppDispatch } from '../../redux/hooks'
import { homeSlice, listRepos, perPage, sync } from '../../redux/home'
import { RequestStatus } from '../../models'

import Main from '../main'
import RepoList, { RepoListProps } from './RepoList'
import Pagination from '../../components/Pagination'
import Spin from '../../components/Spin'

const { Search } = Input
const { actions } = homeSlice

// Binding the state to the deployment page.
export default ():JSX.Element => {
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
        const f = async () => {
            await dispatch(sync())
            await dispatch(actions.setFirstPage())
            await dispatch(listRepos())
        }
        f()
    }

    return (
        <Main>
            <Home
                loading={loading}
                syncing={syncing}
                page={page}
                repos={repos}
                search={search}
                onClickSync={onClickSync}
                onClickPrev={onClickPrev}
                onClickNext={onClickNext}
            />
        </Main>
    )
}

interface HomeProps extends RepoListProps {
    loading: boolean
    syncing: RequestStatus
    page: number
    search(q: string): void
    onClickSync(): void
    onClickPrev(): void
    onClickNext(): void
}

function Home({
    loading,
    page,
    syncing,
    repos,
    search,
    onClickSync,
    onClickPrev,
    onClickNext,
}: HomeProps): JSX.Element {

    return (
        <>
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
                <Button
                    loading={syncing === RequestStatus.Pending}
                    icon={<RedoOutlined />}
                    onClick={onClickSync}
                >
                    Sync
                </Button>
            </div>
            <div style={{"marginTop": "20px"}}>
                <Search placeholder="Search repository ..." onSearch={search} size="large" enterButton />
            </div>
            <div style={{"marginTop": "20px"}}>
                {(loading)? 
                    <div style={{textAlign: "center"}}>
                        <Spin />
                    </div>
                    :
                    <RepoList repos={repos} />}
            </div>
            <div style={{marginTop: "20px", textAlign: "center"}}>
                <Pagination 
                    disabledPrev={page <= 1} 
                    disabledNext={repos.length < perPage} 
                    onClickPrev={onClickPrev} 
                    onClickNext={onClickNext} 
                />
            </div>
        </>
    )
}