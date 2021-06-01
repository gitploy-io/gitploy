import { Component } from "react"
import { Card } from 'react-bootstrap'

import { Repo } from '../models'

export interface RepoListProps {
    repos: Repo[]
}

export default class RepoList extends Component<RepoListProps> {
    render() {
        return (
            <div>
                {this.props.repos.map((repo, index) => {
                return <Card key={index} className="mb-3">
                        <Card.Body >
                            <p className="mb-0">{repo.namespace} / {repo.name} </p>
                        </Card.Body>
                    </Card>
                })}
            </div>
        )
    }
}