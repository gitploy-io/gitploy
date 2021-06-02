import { useEffect } from 'react'

import { useAppSelector, useAppDispatch } from '../redux/hooks'
import { listRepos } from '../redux/home'

import Main from './Main'
import RepoList from '../components/RepoList'
import Spin from '../components/Spin'

export default function Home(){
    const loading = useAppSelector(state => state.home.loading)
    const repos = useAppSelector(state => state.home.repos)
    const dispatch = useAppDispatch()

    useEffect(() => {
        console.log("test")
        dispatch(listRepos())
    })

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
