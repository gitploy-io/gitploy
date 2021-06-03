import { Layout, Menu, Row, Col } from 'antd';

const { Header, Content, Footer } = Layout;

export default function Main(props: any) {
        return (
            <Layout className="layout">
                <Header>
                    <div className="logo" />
                    <Menu theme="dark" mode="horizontal" defaultSelectedKeys={['2']}>
                        <Menu.Item key="1">nav 1</Menu.Item>
                        <Menu.Item key="2">nav 2</Menu.Item>
                        <Menu.Item key="3">nav 3</Menu.Item>
                    </Menu>
                </Header>
                <Content style={{ padding: '0 50px' }}>
                    <Row>
                        <Col span={10} offset={7}>
                            {props.children}
                        </Col>
                    </Row>
                </Content>
                <Footer style={{ textAlign: 'center' }}>Gitploy ©2021 Created by Gitploy.io </Footer>
            </Layout>
        )
}