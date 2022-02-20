import { useEffect } from "react"
import { shallowEqual } from "react-redux"
import { Helmet } from "react-helmet"
import { Avatar, Button, Tag, Descriptions } from "antd"
import moment from "moment"

import { useAppSelector, useAppDispatch } from "../redux/hooks"
import { fetchMe, fetchRateLimit, checkSlack } from "../redux/settings"

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
            <Helmet>
                <title>Settings</title>
            </Helmet>
            <h1>Settings</h1>
            <Descriptions title="User Info" column={1} style={{marginTop: "40px"}}>
                <Descriptions.Item label="Login">
                    <b>{user?.login}</b>
                </Descriptions.Item>
                <Descriptions.Item label="Role">
                    {(user?.admin)? 
                        <Tag color="purple">Admin</Tag> 
                        : 
                        <Tag color="purple">Member</Tag>}
                </Descriptions.Item>
            </Descriptions>
            <Descriptions title="Rate Limit" style={{marginTop: "40px"}} column={1}>
                <Descriptions.Item label="Limit">{rateLimit?.limit}</Descriptions.Item>
                <Descriptions.Item label="Remaining">{rateLimit?.remaining}</Descriptions.Item>
                <Descriptions.Item label="Reset">{moment(rateLimit?.reset).fromNow()}</Descriptions.Item>
            </Descriptions>
            {(isSlackEnabled)?
                <div style={{marginTop: "40px", marginBottom: "20px"}}>
                    <h2>Slack</h2>
                    {(connected)? 
                        <Button href="/slack/signout" type="primary" danger>DISCONNECTED</Button>:
                        <Button href="/slack/" type="primary">CONNECT</Button>}
                </div>:
                null}
        </Main>
    )
}