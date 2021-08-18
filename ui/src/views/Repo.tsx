import { useParams } from "react-router-dom";
import { Menu, Breadcrumb } from 'antd'
import { shallowEqual } from "react-redux";
import { useEffect } from "react";

import { useAppSelector, useAppDispatch } from '../redux/hooks'
import { init, activate } from '../redux/repo'

import ActivateButton from "../components/ActivateButton"
import Main from './Main'
import RepoHome from './RepoHome'
import RepoDeploy from './RepoDeploy'
import RepoRollabck from './RepoRollback'
import RepoSettings from "./RepoSettings"

interface Params {
    namespace: string
    name: string
    tab: string
}

export default function Repo(): JSX.Element {
    const { namespace, name, tab } = useParams<Params>()
    const { repo } = useAppSelector(state => state.repo, shallowEqual)
    const dispatch = useAppDispatch()

    useEffect(() => {
        const f = async () => {
            await dispatch(init({namespace, name}))
        }
        f()
        // eslint-disable-next-line
    }, [dispatch])

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
            <div >
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
                <Menu mode="horizontal" selectedKeys={[(tab)? tab : "home"]}>
                    <Menu.Item key="home">
                        <a href={`/${namespace}/${name}`}>Home</a>
                    </Menu.Item>
                    <Menu.Item key="deploy">
                        <a href={`/${namespace}/${name}/deploy`}>Deploy</a>
                    </Menu.Item>
                    <Menu.Item key="rollback">
                        <a href={`/${namespace}/${name}/rollback`}>Rollback</a>
                    </Menu.Item>
                    <Menu.Item key="settings">
                        <a href={`/${namespace}/${name}/settings`}>Settings</a>
                    </Menu.Item>
                </Menu>
            </div>
            <div style={styleActivateButton}>
                <ActivateButton onClickActivate={onClickActivate}/>
            </div>
            <div style={styleContent}>
                <div>
                    {(!tab || tab === "home") ? <RepoHome /> : null}
                </div>
                <div>
                    {(tab === "deploy") ? <RepoDeploy /> : null}
                </div> 
                <div>
                    {(tab === "rollback")? <RepoRollabck /> : null}
                </div> 
                <div>
                    {(tab === "settings")? <RepoSettings /> : null}
                </div> 
            </div>
        </Main>
    )
}
