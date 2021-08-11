import { useEffect, useState } from 'react'
import { shallowEqual } from 'react-redux'
import { Layout, Menu, Row, Col, Result, Button, Drawer, Avatar, Dropdown, Badge} from "antd"
import { BellOutlined } from "@ant-design/icons"

import { useAppSelector, useAppDispatch } from "../redux/hooks"
import { init } from "../redux/main"

const { Header, Content, Footer } = Layout

// eslint-disable-next-line
export default function Main(props: any) { 
    const { 
        available, 
        authorized, 
        user,
    } = useAppSelector(state => state.main, shallowEqual)
    const dispatch = useAppDispatch()

    useEffect(() => {
        dispatch(init())

    }, [dispatch])

    const [ isNotificationsVisible, setNotificationsVisible ] = useState(false)

    const showNotifications = () => {
        setNotificationsVisible(true)
    }

    const onCloseNotifications = () => {
        setNotificationsVisible(false)
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
                        <Badge size="small" >
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
                                <p>Deployments, Approvals</p>
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
