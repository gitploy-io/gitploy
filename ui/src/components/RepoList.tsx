import { Component } from "react"
import { List, Typography } from 'antd'
import moment from "moment"

import { Deployment, Repo } from '../models'
import UserAvatar from './UserAvatar'
import DeploymentStatusBadge from "./DeploymentStatusBadge"
import DeploymentRefCode from "./DeploymentRefCode"

const { Text, Paragraph } = Typography

export interface RepoListProps {
    repos: Repo[]
}

export default class RepoList extends Component<RepoListProps> {
    render(): JSX.Element {
        return (
            <List
                dataSource={this.props.repos}
                renderItem={repo => { 
                    // deployments is undeinfed if there is no deployments of the repository.
                    const deployment = (repo.deployments)? 
                        repo.deployments[0] :
                        undefined

                    return <List.Item>
                      <List.Item.Meta
                            title={<a 
                                    href={`/${repo.namespace}/${repo.name}`}
                                >
                                    {repo.namespace} / {repo.name}
                                </a>}
                            description={<Description deployment={deployment}/>}
                      />
                    </List.Item>
                }}
            />
        )
    }
}

interface DescriptionProps {
    deployment?: Deployment
}

function Description(props: DescriptionProps): JSX.Element {
    if (!props.deployment) {
        return <p></p>
    }

    return <Paragraph style={{marginTop: "10px", marginBottom: 0}}>
        <UserAvatar user={props.deployment.deployer} /> deployed <DeploymentRefCode deployment={props.deployment}/> to the <Text strong>{props.deployment.env}</Text> environment {moment(props.deployment.createdAt).fromNow()} <DeploymentStatusBadge deployment={props.deployment}/>
    </Paragraph>
}