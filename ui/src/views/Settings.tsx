import { PageHeader, Avatar, Button, Tag } from "antd"
import moment from "moment"
import { useEffect } from "react"
import { shallowEqual } from "react-redux"

import { useAppSelector, useAppDispatch } from "../redux/hooks"
import { fetchMe, fetchRateLimit,checkSlack } from "../redux/settings"

import Main from "./Main"

export default function Settings(): JSX.Element {
    const { user, rateLimit, isSlackEnabled } = useAppSelector(state => state.settings, shallowEqual)
    const dispatch = useAppDispatch()

    useEffect(() => {
        dispatch(fetchMe())
        dispatch(fetchRateLimit())
        dispatch(checkSlack())
    }, [dispatch])

    const connected = (user?.chatUser)? true : false


    return (
        <Main>
            <div style={{"marginTop": "20px"}}>
                <PageHeader 
                    title="Settings"/>
            </div>
            <div style={{marginTop: "20px", padding: "16px 24px"}}>
                <h2>User</h2>
                <p>
                    Login: <Avatar src={user?.avatar}/> <b>{user?.login}</b> 
                </p>
                <p>
                    Role: {(user?.admin)? <Tag color="purple">Admin</Tag> : <Tag color="purple">Member</Tag>}
                </p>
            </div>
            <div style={{marginTop: "20px", padding: "16px 24px"}}>
                <h2>Rate Limit</h2>
                <p>Limit: {rateLimit?.limit}</p>
                <p>Remaining: {rateLimit?.remaining}</p>
                <p>Reset: {moment(rateLimit?.reset).fromNow()}</p>
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