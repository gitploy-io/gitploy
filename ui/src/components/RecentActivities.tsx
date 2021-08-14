import { Tabs, List, Typography } from "antd"
import moment from "moment"

import { Deployment, DeploymentType, Approval } from "../models"

import DeploymentStatusBadge from "./DeploymentStatusBadge"

const { TabPane } = Tabs
const { Paragraph, Text } = Typography

interface RecentActivitiesProps {
    deployments: Deployment[]
    approvals: Approval[]
}

export default function RecentActivities(props: RecentActivitiesProps) {
    return <Tabs>
        <TabPane tab="Deployments" key={1}>
            <DeploymentList deployments={props.deployments}/>
        </TabPane>
        <TabPane tab="Approvals" key={2}>
            <ApprovalList approvals={props.approvals} />
        </TabPane>
    </Tabs>
}

interface DeploymentListProps {
    deployments: Deployment[]
}

function DeploymentList(props: DeploymentListProps) {
    return <List
        dataSource={props.deployments}
        renderItem={(d) => {
            const title = (d.repo) ? `${d.repo.namespace}/${d.repo.name} #${d.number}` : `Deployment #${d.number}`
            const link = (d.repo) ? `/${d.repo.namespace}/${d.repo.name}/deployments/${d.number}` : `#`
            const ref = (d.type === DeploymentType.Commit) ? d.ref.substr(0, 7) : d.ref

            return <List.Item>
                <List.Item.Meta
                    title={<a href={link}>{title}</a>}
                    description={<Paragraph>
                        Deployed <Text code>{ref}</Text> to the <Text code>{d.env}</Text> environment {moment(d.createdAt).fromNow()} <DeploymentStatusBadge deployment={d}/>
                    </Paragraph>}
                />
            </List.Item>
        }}
    />
}

interface ApprovalListProps {
    approvals: Approval[]
}

function ApprovalList(props: ApprovalListProps) {
    return <List
        dataSource={props.approvals}
        renderItem={(a) => {
            if (a.deployment === null) {
                throw new ReferenceError("The deployment of the approval is not found.")
            }
            const d = a.deployment
            const title = (d.repo) ? `${d.repo.namespace}/${d.repo.name} #${d.number}` : `Deployment #${d.number}`
            const link = (d.repo) ? `/${d.repo.namespace}/${d.repo.name}/deployments/${d.number}` : `#`
            const ref = (d.type === DeploymentType.Commit) ? d.ref.substr(0, 7) : d.ref

            return <List.Item>
                <List.Item.Meta
                    title={<a href={link}>{title}</a>}
                    description={<Paragraph>
                        Requested the approval to deploy <Text code>{ref}</Text> to the <Text code>{d.env}</Text> environment {moment(d.createdAt).fromNow()}
                    </Paragraph>}
                />
            </List.Item>
        }}
    />
}