import { useEffect, useState } from "react"
import { shallowEqual } from "react-redux"
import { Layout, Menu, Row, Col, Result, Button, Drawer, Avatar, Dropdown, Badge} from "antd"
import { SettingFilled } from "@ant-design/icons"

import { useAppSelector, useAppDispatch } from "../redux/hooks"
import { init, searchDeployments, searchApprovals, mainSlice } from "../redux/main"
import { subscribeDeploymentEvent, subscribeApprovalEvent } from "../apis"

import RecentActivities from "../components/RecentActivities"

const { Header, Content, Footer } = Layout

// eslint-disable-next-line
export default function Main(props: any) { 
    const { 
        available, 
        authorized, 
        user,
        deployments,
        approvals
    } = useAppSelector(state => state.main, shallowEqual)
    const dispatch = useAppDispatch()

    useEffect(() => {
        dispatch(init())
        dispatch(searchDeployments())
        dispatch(searchApprovals())

        const de = subscribeDeploymentEvent((d) => {
            dispatch(mainSlice.actions.handleDeploymentEvent(d))
        })

        const ae = subscribeApprovalEvent((a) => {
            dispatch(mainSlice.actions.handleApprovalEvent(a))
        })

        return () => {
            de.close()
            ae.close()
        }
    }, [dispatch])

    const [ isRecentActivitiesVisible, setRecentActivitiesVisible ] = useState(false)

    const showRecentActivities = () => {
        setRecentActivitiesVisible(true)
    }

    const onCloseRecentActivities = () => {
        setRecentActivitiesVisible(false)
    }

    // Count of recent activities.
    const countActivities = deployments.length + approvals.length

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
                            <Menu.Item key={0}>Gitploy</Menu.Item>
                            <Menu.Item key={1}><a href="/">Home</a></Menu.Item>
                            {(user?.admin)? <Menu.Item key={2}><a href="/members">Members</a></Menu.Item>: null}
                        </Menu>
                    </Col>
                    <Col span="8" style={{textAlign: "right"}}>
                        <Badge 
                            size="small" 
                            count={countActivities}
                        >
                            <Button 
                                type="primary" 
                                shape="circle" 
                                icon={<SettingFilled spin={countActivities !== 0 }/>}
                                onClick={showRecentActivities}
                            />
                        </Badge>
                        <Drawer title="Recent Activities"
                            placement="right"
                            width={400}
                            visible={isRecentActivitiesVisible}
                            onClose={onCloseRecentActivities}
                        >
                            <RecentActivities 
                                deployments={deployments}
                                approvals={approvals}
                            />
                        </Drawer>
                        &nbsp; &nbsp;

                        {/* Avatar */}
                        {(authorized) ? 
                            <Dropdown overlay={
                                <Menu style={{width: "150px"}}>
                                    <Menu.Item key="0">
                                        <a href="/settings">Settings</a>
                                    </Menu.Item>
                                    <Menu.Divider />
                                    <Menu.Item key="1">
                                        <a href="/signout">Sign Out</a>
                                    </Menu.Item>
                                </Menu> }>
                                <Avatar  src={user?.avatar}/>
                            </ Dropdown>
                            : <a href="/" style={{color: "white"}}>Sign in</a>}
                    </Col>
                </Row>
            </Header>
            <Content style={{ padding: '50px 50px' }}>
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
