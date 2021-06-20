import { Layout, Menu, Row, Col } from 'antd';

const { Header, Content, Footer } = Layout;

export default function Main(props: any) {
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
                            {props.children}
                        </Col>
                    </Row>
                </Content>
                <Footer style={{ textAlign: 'center' }}>Gitploy ©2021 Created by Gitploy.io </Footer>
            </Layout>
        )
}