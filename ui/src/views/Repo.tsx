import { useParams } from "react-router-dom";
import { Menu, Breadcrumb } from 'antd'
import { shallowEqual } from "react-redux";
import { useEffect } from "react";

import { useAppSelector, useAppDispatch } from '../redux/hooks'
import { repoSlice, init, activate } from '../redux/repo'

import ActivateButton from "../components/ActivateButton"
import Main from './Main'
import RepoHome from './RepoHome'
import RepoDeploy from './RepoDeploy'
import RepoRollabck from './RepoRollback'
import RepoSettings from "./RepoSettings"

const { actions } = repoSlice

const hide = {
    display: "none"
}

interface Params {
    namespace: string
    name: string
}

export default function Repo() {
    let { namespace, name } = useParams<Params>()
    const { key, repo } = useAppSelector(state => state.repo, shallowEqual)
    const dispatch = useAppDispatch()

    useEffect(() => {
        const f = async () => {
            await dispatch(init({namespace, name}))
        }
        f()
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [dispatch])

    const onClickMenu = (e: any) => {
        const key: string = e.key
        dispatch(actions.setKey(key))
    }

    const onClickActivate = () => {
        dispatch(activate())
    }

    const hasInit = (repo)? true : false
    const active = (repo?.active)? true : false

    const styleActivateButton: React.CSSProperties = {
        display: (hasInit && !active)? "" : "none",
        marginTop: "20px", 
        textAlign: "center"
    }

    const styleContent: React.CSSProperties = {
        display: (hasInit && active)? "" : "none",
        marginTop: "20px", 
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
            <div style={styleActivateButton}>
                <ActivateButton onClickActivate={onClickActivate}/>
            </div>
            <div style={styleContent}>
                {(key === "home")? 
                    <div><RepoHome /></div> : 
                    <div style={hide}><RepoHome /></div>}
                {(key === "deploy")?
                    <div><RepoDeploy /></div> :
                    <div style={hide}><RepoDeploy /></div>}
                {(key === "rollback")?
                    <div><RepoRollabck /></div> :
                    <div style={hide}><RepoRollabck /></div>}
                {(key === "settings")?
                    <div><RepoSettings /></div> :
                    <div style={hide}><RepoSettings /></div>}
            </div>
        </Main>
    )
}
