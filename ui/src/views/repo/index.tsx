import { useParams } from "react-router-dom";
import { Menu, Breadcrumb, Result, } from 'antd'
import { shallowEqual } from "react-redux";
import { useEffect } from "react";
import { Helmet } from "react-helmet"

import { useAppSelector, useAppDispatch } from '../../redux/hooks'
import { init, activate, repoSlice as slice } from '../../redux/repo'

import ActivateButton from "../../components/ActivateButton"
import Main from '../main'
import RepoHome from '../repoHome'
import RepoLock from "../repoLock"
import RepoDeploy from '../repoDeploy'
import RepoRollabck from '../repoRollback'
import RepoSettings from "../repoSettings"

interface Params {
    namespace: string
    name: string
    tab: string
}

export default (): JSX.Element => {
    const { namespace, name, tab } = useParams<Params>()
    const { display, repo } = useAppSelector(state => state.repo, shallowEqual)
    const dispatch = useAppDispatch()

    useEffect(() => {
        const f = async () => {
            await dispatch(init({namespace, name}))
            await dispatch(slice.actions.setDisplay(true))
        }
        f()
        // eslint-disable-next-line
    }, [dispatch])

    const onClickActivate = () => {
        dispatch(activate())
    }

    if (!display) {
        return (
            <Main>
                <div />
            </Main>
        )
    } else if (display && !repo) {
        return (
            <Main>
                <Result
                    style={{paddingTop: '120px'}}
                    status="warning"
                    title="The page is not found."
                    subTitle="Please check the URL."
                />
            </Main>
        )
    }

    return (
        <Repo 
            namespace={namespace}
            name={name}
            tab={tab}
            active={(repo?.active)? true : false}
            onClickActivate={onClickActivate}
        />
    )
}

interface RepoProps {
    namespace: string
    name: string
    tab: string
    active: boolean
    onClickActivate(): void
}

function Repo({
    namespace,
    name,
    tab,
    active,
    onClickActivate
}: RepoProps): JSX.Element {
    const styleActivateButton: React.CSSProperties = {
        display: (!active)? "" : "none",
        marginTop: "20px", 
        textAlign: "center"
    }

    const styleContent: React.CSSProperties = {
        display: (active)? "" : "none",
        marginTop: "20px", 
    }

    return (
        <Main>
            <Helmet>
                <title>{namespace}/{name}</title>
            </Helmet>
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
                    <Menu.Item key="lock">
                        <a href={`/${namespace}/${name}/lock`}>Lock</a>
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
                    {(tab === "lock")? <RepoLock /> : null}
                </div> 
                <div>
                    {(tab === "settings")? <RepoSettings /> : null}
                </div> 
            </div>
        </Main>
    )
}
