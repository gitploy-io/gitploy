import { useEffect } from 'react'
import { shallowEqual } from 'react-redux'

import { useAppSelector, useAppDispatch } from '../redux/hooks'
import { listRepos } from '../redux/home'

import Main from './Main'
import RepoList from '../components/RepoList'
import Spin from '../components/Spin'

export default function Home(){
    const { loading, repos } = useAppSelector(state => state.home, shallowEqual)
    const dispatch = useAppDispatch()

    useEffect(() => {
        dispatch(listRepos())
    }, [dispatch])

    if (loading) {
        return (
            <Main>
                <Spin />
            </Main>
        )
    }

    return (
        <Main>
            <RepoList repos={repos}></RepoList>
        </Main>
    )
}
