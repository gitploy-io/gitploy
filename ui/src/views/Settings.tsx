import { useEffect } from "react"
import { shallowEqual } from "react-redux"
import { Helmet } from "react-helmet"
import { Button, Tag, Descriptions, Input } from "antd"

import { useAppSelector, useAppDispatch } from "../redux/hooks"
import { fetchMe, checkSlack } from "../redux/settings"

import Main from "./main"

export default function Settings(): JSX.Element {
    const { user, isSlackEnabled } = useAppSelector(state => state.settings, shallowEqual)
    const dispatch = useAppDispatch()

    useEffect(() => {
        dispatch(fetchMe())
        dispatch(checkSlack())
    }, [dispatch])

    const connected = (user?.chatUser)? true : false


    return (
        <Main>
            <Helmet>
                <title>Settings</title>
            </Helmet>
            <h1>Settings</h1>
            <Descriptions title="User Info" column={2} style={{marginTop: "40px"}} layout="vertical">
                <Descriptions.Item label="Login">{user?.login}</Descriptions.Item>
                <Descriptions.Item label="Role">
                    {(user?.admin)? 
                        <Tag color="purple">Admin</Tag> 
                        : 
                        <Tag color="purple">Member</Tag>}
                </Descriptions.Item>
                <Descriptions.Item label="Token">
                    <Input.Password 
                        value={user?.hash}
                        style={{width: 200, padding:0 }}
                        readOnly
                        bordered={false}
                    />
                </Descriptions.Item>
            </Descriptions>
            {(isSlackEnabled)?
                <div style={{marginTop: "40px", marginBottom: "20px"}}>
                    <Descriptions title="Slack">
                        <Descriptions.Item>
                            {(connected)? 
                                <Button href="/slack/signout" type="primary" danger>DISCONNECTED</Button>:
                                <Button href="/slack/" type="primary">CONNECT</Button>}
                        </Descriptions.Item>
                    </Descriptions>
                </div>:
                null}
        </Main>
    )
}