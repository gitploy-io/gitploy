import { useEffect, useState } from 'react'
import { shallowEqual } from 'react-redux'
import { Layout, Menu, Row, Col, Result, Button, Drawer, Avatar, Dropdown, Badge} from "antd"
import { BellOutlined } from "@ant-design/icons"

import { Notification as NotificationData, NotificationType } from "../models"
import { subscribeNotification } from "../apis"
import { useAppSelector, useAppDispatch } from "../redux/hooks"
import { mainSlice, init, fetchNotifications, setNotificationChecked } from "../redux/main"

import NotificationList from "../components/NotificationList"

const { actions } = mainSlice
const { Header, Content, Footer } = Layout

// eslint-disable-next-line
export default function Main(props: any) { 
    const { 
        available, 
        authorized, 
        user,
        notifications
    } = useAppSelector(state => state.main, shallowEqual)
    const dispatch = useAppDispatch()

    useEffect(() => {
        dispatch(init())
        dispatch(fetchNotifications())

        // SSE(Server Sent Event) open a connection
        // to stream a new notification.
        const sse = subscribeNotification((n) => {
            if (!n.notified) {
                notify(n)
            }
            dispatch(actions.addNotification(n))
        })

        return () => {
            sse.close();
        };
    }, [dispatch])

    // Drawer shows notifications.
    const [ isNotificationsVisible, setNotificationsVisible ] = useState(false)

    const showNotifications = () => {
        setNotificationsVisible(true)
    }

    const onCloseNotifications = () => {
        setNotificationsVisible(false)
    }

    const onClickNotificaiton = (n: NotificationData) => {
        dispatch(setNotificationChecked(n))
    }

    let content: React.ReactElement
    if (!available) {
        content = <Result
            style={{paddingTop: '120px'}}
            status="error"
            title="Server Internal Error"
            subTitle="Sorry, something went wrong."
        />
    } else if (!authorized) {
        content = <Result
            style={{paddingTop: '120px'}}
            status="warning"
            title="Unauthorized Error"
            subTitle="Sorry, you are not authorized to access."
            extra={[<Button key="console" type="primary" href="/">Sign in</Button>]}
        />
    } else {
        content = props.children
    }

    return (
        <Layout className="layout">
            <Header>
                <Row>
                    <Col span="16">
                        <Menu theme="dark" mode="horizontal" defaultActiveFirst>
                            <Menu.Item key="1">Gitploy</Menu.Item>
                            <Menu.Item key="2"><a href="/">Home</a></Menu.Item>
                        </Menu>
                    </Col>
                    <Col span="8" style={{textAlign: "right"}}>
                        {/* Notifications */}
                        <Badge size="small" count={countUnchecked(notifications)}>
                            <Button 
                                type="primary" 
                                shape="circle" 
                                icon={<BellOutlined />}
                                onClick={showNotifications}></Button>

                            </Badge>
                        <Drawer title="Notifications"
                            placement="right"
                            width={400}
                            visible={isNotificationsVisible}
                            onClose={onCloseNotifications}>
                                <NotificationList 
                                    notifications={notifications}
                                    onClickNotificaiton={onClickNotificaiton}/>
                        </Drawer>
                        &nbsp; &nbsp;

                        {/* Avatar */}
                        {(authorized) ? 
                            <Dropdown overlay={
                                <Menu style={{width: "200px"}}>
                                    <Menu.Item key="0">
                                        <a rel="noopener noreferrer" href="/settings">Settings</a>
                                    </Menu.Item>
                                </Menu> }>
                                <Avatar  src={user?.avatar}/>
                            </ Dropdown>
                            : <a href="/" style={{color: "white"}}>Sign in</a>}
                    </Col>
                </Row>
            </Header>
            <Content style={{ padding: '0 50px' }}>
                <Row>
                    <Col 
                        span={22}
                        offset={1}
                        md={{span: 14, offset: 5}} 
                        lg={{span: 10, offset: 7}}>
                        {content}
                    </Col>
                </Row>
            </Content>
            <Footer style={{ textAlign: 'center' }}>Gitploy ©2021 Created by Gitploy.io </Footer>
        </Layout>
    )
}

function notify(n: NotificationData) {
    if (!("Notification" in window)) {
        console.log("This browser does not support desktop notification")
    }

    else if (Notification.permission === "granted") {
        // To avoid duplicate notification, tag the notification with ID.
        new Notification(convertToNotificationMessage(n), {
            tag: `${n.id}`,
        });
    }
    
    else if (Notification.permission !== 'denied') {
        Notification.requestPermission(function (permission) {
          if (permission === "granted") {
            new Notification(convertToNotificationMessage(n), {
                tag: `${n.id}`
            });
          }
        });
    }
}

function convertToNotificationMessage(n: NotificationData): string {
    switch (n.type) {
        case NotificationType.Deployment:
            return `New Deployment #${n.deploymentNumber}\n ${n.repoNamespace}/${n.repoName} - ${n.deploymentLogin} has deployed to ${n.deploymentEnv} environment.`
        case NotificationType.ApprovalRequested:
            return `Approval Requested\n - ${n.repoNamespace}/${n.repoName} - ${n.deploymentLogin} has requested the approval for the deployment(#${n.deploymentNumber}).`
        case NotificationType.ApprovalResponded:
            return `Approval Responded\n - ${n.repoNamespace}/${n.repoName} - ${n.approvalLogin} has responded the approval of the deployment(#${n.deploymentNumber}).`
        default:
            return "New Event"
    }
}

function countUnchecked(notifications: NotificationData[]) {
    let cnt = 0
    notifications.forEach((n) => {
        if (!n.checked) {
            cnt++
        }
    })
    return cnt
}