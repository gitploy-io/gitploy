import { Breadcrumb, PageHeader, Row, Col, Typography, Avatar, Badge } from "antd"
import { shallowEqual } from 'react-redux'
import { useParams } from "react-router-dom"

import { useAppSelector, useAppDispatch } from "../redux/hooks"
import { init, deploymentSlice, fetchDeployment, fetchApprovals, fetchMyApproval } from "../redux/deployment"
import { DeploymentStatus } from "../models"

import Main from "./Main"
import ApprovalList from "../components/ApprovalList"
import Spin from "../components/Spin"
import { useEffect } from "react"

const { Text } = Typography
const { actions } = deploymentSlice

interface Params {
    namespace: string
    name: string
    number: string
}

export default function Deployment() {
    let { namespace, name, number } = useParams<Params>()
    const { deployment, approvals } = useAppSelector(state => state.deployment, shallowEqual )
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

    const onBack = () => {
        window.location.href = `/${namespace}/${name}`
    }

    if (deployment === null) {
        return <Main>
            <div style={{textAlign: "center", marginTop: "100px"}}><Spin /></div>
        </Main>
    }

    // styles 
    const styleField: React.CSSProperties = { marginTop: "10px" }
    const styleFieldName: React.CSSProperties = { textAlign: "right"}

    const deployer = deployment.deployer

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
                        <Col>
                            {/* TODO: add a new component - DeploymentStatusBadge */}
                            <Badge color={getStatusColor(deployment.status)}text={deployment.status}/>
                        </Col>
                    </Row>
                    <Row style={styleField}>
                        <Col span="4" style={styleFieldName}>Deployer:&nbsp;&nbsp;</Col>
                        <Col> 
                            {(deployer !== null)?
                                 <Text ><Avatar size="small" src={deployer.avatar} /> {deployer.login}</Text> :
                                <Avatar size="small" >U</Avatar> }
                        </Col>
                    </Row>
                    {(approvals.length !== 0) ?
                    <Row style={styleField}>
                            <Col span="4" style={styleFieldName}>Approvers:&nbsp;&nbsp;</Col>
                            <Col><ApprovalList approvals={approvals}/></Col>
                        </Row> :
                        null}
                </div>
            </div>
        </Main>
    )
}

// https://ant.design/components/timeline/#Timeline.Item
const getStatusColor = (status: DeploymentStatus) => {
    switch (status) {
        case DeploymentStatus.Waiting:
            return "gray"
        case DeploymentStatus.Created:
            return "purple"
        case DeploymentStatus.Running:
            return "purple"
        case DeploymentStatus.Success:
            return "green"
        case DeploymentStatus.Failure:
            return "red"
        default:
            return "gray"
    }
}