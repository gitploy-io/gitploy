import { Component } from 'react';
import { Layout, Menu } from 'antd';

const { Header, Content, Footer } = Layout;

export default class Main extends Component {
    render() {
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
                    <div>
                        {this.props.children}
                    </div>
                </Content>
                <Footer style={{ textAlign: 'center' }}>Gitploy ©2021 Created by Gitploy.io </Footer>
            </Layout>
        )
    }
}