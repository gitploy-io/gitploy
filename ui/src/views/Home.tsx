import { Component } from "react";

import Main from './Main'
import RepoList from '../components/RepoList'
import { Repo } from '../models'

const repos:Repo[] = [{
    id: "1",
    namespace: "hanjunlee",
    name: "gitploy",
    description: "",
    syncedAt: new Date(),
    createdAt: new Date(),
    updatedAt: new Date(),
}]

interface HomeProps {
    loading: boolean
    repos: Repo[]
    listRepos(q: string, page: number, perPage: number): void
}

export default class Home extends Component<HomeProps> {
    render() {
        return (
            <Main >
                <RepoList repos={repos}></RepoList>
            </Main>
        )
    }
}