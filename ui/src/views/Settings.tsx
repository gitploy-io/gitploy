import { PageHeader, Avatar, Space, Button } from "antd"
import { useEffect } from "react"
import { shallowEqual } from "react-redux"

import { useAppSelector, useAppDispatch } from "../redux/hooks"
import { fetchMe, checkSlack } from "../redux/settings"

import Main from "./Main"

export default function Settings() {
    const { user, isSlackEnabled } = useAppSelector(state => state.settings, shallowEqual)
    const dispatch = useAppDispatch()

    useEffect(() => {
        const f = async () => {
            await dispatch(fetchMe())
            await dispatch(checkSlack())
        }
        f()
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [dispatch])

    let connected = (user?.chatUser)? true : false


    return (
        <Main>
            <div style={{"marginTop": "20px"}}>
                <PageHeader 
                    title="Settings"/>
            </div>
            <div style={{marginTop: "20px", padding: "16px 24px"}}>
                <h2>User</h2>
                <Space><b>{user?.login}</b><Avatar src={user?.avatar}/></Space>
            </div>
            {(isSlackEnabled)?
                <div style={{marginTop: "20px", marginBottom: "20px", padding: "16px 24px"}}>
                    <h2>Slack</h2>
                    {(connected)? 
                        <Button href="#" type="primary" danger>DISCONNECTED</Button>:
                        <Button href="/slack/" type="primary">CONNECT</Button>}
                </div>:
                null}
        </Main>
    )
}