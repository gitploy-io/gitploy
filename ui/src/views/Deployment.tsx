import { useEffect } from "react"
import { Breadcrumb, PageHeader, Row, Col, Typography, Avatar, Button, } from "antd"
import { shallowEqual } from 'react-redux'
import { useParams } from "react-router-dom"

import { useAppSelector, useAppDispatch } from "../redux/hooks"
import { 
    init, 
    deploymentSlice, 
    fetchDeployment, 
    fetchApprovals, 
    fetchMyApproval,
    deployToSCM,
    approve,
    decline,
} from "../redux/deployment"
import { Deployment, DeploymentStatus, Approval } from "../models"

import Main from "./Main"
import ApprovalList from "../components/ApprovalList"
import ApprovalDropdown from "../components/ApprovalDropdown"
import Spin from "../components/Spin"
import DeploymentStatusBadge from "../components/DeploymentStatusBadge"

const { Text } = Typography
const { actions } = deploymentSlice

interface Params {
    namespace: string
    name: string
    number: string
}

export default function DeploymentView() {
    let { namespace, name, number } = useParams<Params>()
    const { deployment, approvals, myApproval } = useAppSelector(state => state.deployment, shallowEqual )
    const dispatch = useAppDispatch()

    useEffect(() => {
        const f = async () => {
            await dispatch(init({namespace, name}))
            await dispatch(actions.setNumber(parseInt(number, 10)))
            await dispatch(fetchDeployment())
            await dispatch(fetchApprovals())
            await dispatch(fetchMyApproval())
        }
        f()
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [dispatch])

    const onClickDeploy = () => {
        dispatch(deployToSCM())
    }

    const onClickApprove = () => {
        dispatch(approve())
    }

    const onClickDecline = () => {
        dispatch(decline())
    }

    const onBack = () => {
        window.location.href = `/${namespace}/${name}`
    }

    if (deployment === null) {
        return <Main>
            <div style={{textAlign: "center", marginTop: "100px"}}><Spin /></div>
        </Main>
    }

    // styles 
    const styleField: React.CSSProperties = { marginTop: "15px" }
    const styleFieldName: React.CSSProperties = { textAlign: "right"}

    // buttons
    const deployBtn = isDeployable(deployment, approvals)? 
        <Button type="primary" onClick={onClickDeploy}>Deploy</Button>:
        <Button type="primary" disabled>Deploy</Button>
    const approvalDropdown = (hasRequestedApproval(myApproval))?
        <ApprovalDropdown 
            key="approval" 
            onClickApprove={onClickApprove}
            onClickDecline={onClickDecline}/>:
        null

    return (
        <Main>
            <div style={{"marginTop": "20px"}}>
                <PageHeader
                    title={`Deployment #${number}`}
                    breadcrumb={
                        <Breadcrumb>
                            <Breadcrumb.Item>
                                <a href="/">Repositories</a>
                            </Breadcrumb.Item>
                            <Breadcrumb.Item>{namespace}</Breadcrumb.Item>
                            <Breadcrumb.Item>
                                <a href={`/${namespace}/${name}`}>{name}</a>
                            </Breadcrumb.Item>
                            <Breadcrumb.Item>Deployments</Breadcrumb.Item>
                            <Breadcrumb.Item>{number}</Breadcrumb.Item>
                        </Breadcrumb>}
                    extra={[
                        approvalDropdown,
                    ]}
                    onBack={onBack} />
            </div>
            <div style={{marginTop: "20px", marginBottom: "30px"}}>
                <div>
                    <Row style={styleField}>
                        <Col span="4" style={styleFieldName}>Target:&nbsp;&nbsp;</Col>
                        <Col><Text>{deployment.env}</Text></Col>
                    </Row>
                    <Row style={styleField}>
                        <Col span="4" style={styleFieldName}>Ref:&nbsp;&nbsp;</Col>
                        <Col><Text code>{deployment.ref}</Text></Col>
                    </Row>
                    <Row style={styleField}>
                        <Col span="4" style={styleFieldName}>Status:&nbsp;&nbsp;</Col>
                        <Col><DeploymentStatusBadge deployment={deployment}/></Col>
                    </Row>
                    <Row style={styleField}>
                        <Col span="4" style={styleFieldName}>Deployer:&nbsp;&nbsp;</Col>
                        <Col> 
                            {(deployment.deployer !== null)?
                                 <Text ><Avatar size="small" src={deployment.deployer.avatar} /> {deployment.deployer.login}</Text> :
                                <Avatar size="small" >U</Avatar> }
                        </Col>
                    </Row>
                    {/* Approvals */}
                    {(approvals.length !== 0) ?
                        <Row style={styleField}>
                            <Col span="4" style={styleFieldName}>Approvers:&nbsp;&nbsp;</Col>
                            <Col><ApprovalList approvals={approvals}/></Col>
                        </Row> :
                        null}
                    {(approvals.length !== 0) ?
                        <Row style={styleField}>
                            <Col span="4" style={styleFieldName}>Required Count:&nbsp;&nbsp;</Col>
                            <Col>{deployment.requiredApprovalCount}</Col>
                        </Row> :
                        null}
                    <Row style={styleField}>
                        <Col span="4" style={styleFieldName}></Col>
                        <Col>{deployBtn}</Col>
                    </Row>
                </div>
            </div>
        </Main>
    )
}

function isDeployable(deployment: Deployment, approvals: Approval[]): boolean {
    if (deployment.status !== DeploymentStatus.Waiting) {
        return false
    }

    // requiredApprovalCount have to be equal or greater than approved.
    var approved = 0
    approvals.forEach((approval) => {
        if (approval.isApproved) {
            approved++
        }
    })

    return approved >= deployment.requiredApprovalCount
}

function hasRequestedApproval(approval: Approval | null): boolean {
    return approval !== null
}