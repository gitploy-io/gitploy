import { Component } from "react"
import { List } from 'antd'

import { Repo } from '../models'

export interface RepoListProps {
    repos: Repo[]
}

export default class RepoList extends Component<RepoListProps> {
    render() {
        return (
            <List
                itemLayout="horizontal"
                dataSource={this.props.repos}
                renderItem={repo => (
                  <List.Item>
                    <List.Item.Meta
                      title={<a href="https://ant.design">{repo.namespace} / {repo.name}</a>}
                    />
                  </List.Item>
                )}
            />
        )
    }
}