import { useEffect, useState } from "react"
import { shallowEqual } from "react-redux"
import { Layout, Menu, Row, Col, Result, Button, Drawer, Avatar, Dropdown, Badge} from "antd"
import { SettingFilled } from "@ant-design/icons"

import { useAppSelector, useAppDispatch } from "../redux/hooks"
import { init, searchDeployments, searchApprovals, fetchLicense, mainSlice as slice } from "../redux/main"
import { subscribeEvents } from "../apis"

import RecentActivities from "../components/RecentActivities"
import LicenseFooter from "../components/LicenseFooter"

import Logo from "../logo.svg"

const { Header, Content, Footer } = Layout

// eslint-disable-next-line
export default function Main(props: any) { 
    const { 
        available, 
        authorized, 
        expired,
        user,
        deployments,
        approvals,
        license,
    } = useAppSelector(state => state.main, shallowEqual)
    const dispatch = useAppDispatch()

    useEffect(() => {
        dispatch(init())
        dispatch(searchDeployments())
        dispatch(searchApprovals())
        dispatch(fetchLicense())

        const sub = subscribeEvents((event) => {
            dispatch(slice.actions.handleDeploymentEvent(event))
            dispatch(slice.actions.handleApprovalEvent(event))
        })

        return () => {
            sub.close()
        }
    }, [dispatch])

    const onClickRetry = () => {
        dispatch(slice.actions.setAvailable(true))
        dispatch(slice.actions.setExpired(false))
    }

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
            subTitle="Sorry, something went wrong. Contact administrator."
            extra={[<Button key="console" type="primary" onClick={onClickRetry}>Retry</Button>]}
        />
    } else if (!authorized) {
        content = <Result
            style={{paddingTop: '120px'}}
            status="warning"
            title="Unauthorized Error"
            subTitle="Sorry, you are not authorized."
            extra={[<Button key="console" type="primary" href="/">Sign in</Button>]}
        />
    } else if (expired) {
        content = <Result
            style={{paddingTop: '120px'}}
            status="warning"
            title="License Expired"
            subTitle="Sorry, the license is expired."
            extra={[<Button key="console" type="primary" onClick={onClickRetry}>Retry</Button>]}
        />
    } else {
        content = props.children
    }

    return (
        <Layout className="layout">
            <Header>
                <Row>
                    <Col span="16">
                        <Menu theme="dark" mode="horizontal" >
                            <Menu.Item key={0}><a href="/"><img src={Logo} style={{width: 62}}/></a></Menu.Item>
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
                                        <a href="/signout">Sign out</a>
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
            <Footer style={{ textAlign: 'center' }}>
                <LicenseFooter license={license} />
                Gitploy ©2021 Created by Gitploy.IO 
            </Footer>
        </Layout>
    )
}
