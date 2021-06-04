import { useParams } from "react-router-dom";
import { Menu, Breadcrumb } from 'antd'

import { useAppSelector, useAppDispatch } from '../redux/hooks'
import { repoSlice } from '../redux/repo'

import Main from './Main'
import RepoHome from './RepoHome'

const { actions } = repoSlice

interface Params {
    namespace: string
    name: string
}

export default function Repo() {
    let { namespace, name } = useParams<Params>()
    const key = useAppSelector(state => state.repo.key)
    const dispatch = useAppDispatch()

    const onClickMenu = (e: any) => {
        const key: string = e.key
        dispatch(actions.setKey(key))
    }

    return (
        <Main>
            <div style={{"marginTop": "20px"}}>
                <Breadcrumb>
                    <Breadcrumb.Item>
                        <a href="/">Repositories</a>
                    </Breadcrumb.Item>
                    <Breadcrumb.Item>{namespace}</Breadcrumb.Item>
                    <Breadcrumb.Item>
                        <a href={`/${namespace}/${name}`}>{name}</a>
                    </Breadcrumb.Item>
                </Breadcrumb>
            </div>
            <div style={{"marginTop": "20px"}}>
                <Menu mode="horizontal" onClick={onClickMenu} selectedKeys={[key]}>
                    <Menu.Item key="home">
                        Home
                    </Menu.Item>
                    <Menu.Item key="deploy">
                        Deploy
                    </Menu.Item>
                    <Menu.Item key="rollback">
                        Rollback
                    </Menu.Item>
                    <Menu.Item key="settings">
                        Settings
                    </Menu.Item>
                </Menu>
            </div>
            <div style={{"marginTop": "20px"}}>
                {(key === "home")? <RepoHome /> : null}
            </div>
        </Main>
    )
}