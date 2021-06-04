import { useParams } from "react-router-dom";
import { Menu, Breadcrumb } from 'antd'

import Main from './Main'

interface Params {
    namespace: string
    name: string
}

export default function Repo() {
    let { namespace, name } = useParams<Params>()

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
                <Menu mode="horizontal" selectedKeys={["home"]}>
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
        </Main>
    )
}