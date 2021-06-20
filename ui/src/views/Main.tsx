import { shallowEqual } from 'react-redux';
import { Layout, Menu, Row, Col, Result, Button } from 'antd';

import { useAppSelector } from "../redux/hooks"

const { Header, Content, Footer } = Layout;

export default function Main(props: any) {
    const { available, authorized } = useAppSelector(state => state.main, shallowEqual)

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
                <Menu theme="dark" mode="horizontal" defaultActiveFirst>
                    <Menu.Item key="1">Gitploy</Menu.Item>
                    <Menu.Item key="2"><a href="/">Home</a></Menu.Item>
                </Menu>
            </Header>
            <Content style={{ padding: '0 50px' }}>
                <Row>
                    <Col span={10} offset={7}>
                        {content}
                    </Col>
                </Row>
            </Content>
            <Footer style={{ textAlign: 'center' }}>Gitploy ©2021 Created by Gitploy.io </Footer>
        </Layout>
    )
}