import { useEffect } from 'react';
import { shallowEqual } from 'react-redux';
import { Layout, Menu, Row, Col, Result, Button, Avatar, Dropdown} from 'antd';

import { Notification as NotificationData, NotificationType } from "../models"
import { subscribeNotification } from "../apis"
import { useAppSelector, useAppDispatch } from "../redux/hooks"
import { init } from "../redux/main"

const { Header, Content, Footer } = Layout;

export default function Main(props: any) {
    const { available, authorized, user } = useAppSelector(state => state.main, shallowEqual)
    const dispatch = useAppDispatch()

    useEffect(() => {
        dispatch(init())
        const sse = subscribeNotification((n) => {
            if (!n.notified) {
                notify(n)
            }
        })

        return () => {
            console.log("close the stream.")
            sse.close();
        };
    }, [dispatch])

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


    const userMenu = <Menu style={{width: "200px"}}>
        <Menu.Item key="0">
            <a rel="noopener noreferrer" href="/notifications">Notifications</a>
        </Menu.Item>
        <Menu.Divider />
        <Menu.Item key="1">
            <a rel="noopener noreferrer" href="/settings">Settings</a>
        </Menu.Item>
    </Menu>

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
                        {(authorized) ? 
                            <Dropdown overlay={userMenu}>
                                <Avatar src={user?.avatar}/>
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
            return `New Deployment - #${n.resourceId}`
    }
}
