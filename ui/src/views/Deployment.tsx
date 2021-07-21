import { useEffect } from "react"
import { Breadcrumb, PageHeader, Row, Col, Typography, Avatar, Button } from "antd"
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
    searchCandidates,
    createApproval,
    deleteApproval,
    approve,
    decline,
} from "../redux/deployment"
import { User, Deployment, LastDeploymentStatus, Approval } from "../models"

import Main from "./Main"
import ApprovalList from "../components/ApprovalList"
import ApproversSearch from "../components/ApproversSearch"
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
    const { deployment, approvals, candidates, myApproval } = useAppSelector(state => state.deployment, shallowEqual )
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

    const approvers: User[] = []
    approvals.forEach(approval => {
        if (approval.user !== null) {
            approvers.push(approval.user)
        }
    })

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

    const onSearchCandidates = (login: string) => {
        dispatch(searchCandidates(login))
    }

    const onSelectCandidate = (id: string) => {
        const approval = approvals.find(a => a.user?.id === id) 

        if (approval !== undefined) {
            dispatch(deleteApproval(approval))
            return
        }

        const candidate = candidates.find(c => c.id === id)
        if (candidate === undefined) {
            throw new Error("The candidate is not found")
        }

        dispatch(createApproval(candidate))
    }

    if (deployment === null) {
        return <Main>
            <div style={{textAlign: "center", marginTop: "100px"}}><Spin /></div>
        </Main>
    }

    // styles 
    const styleField: React.CSSProperties = { marginTop: "18px" }
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
            {/* TODO: support mobile view */}
            <div style={{marginTop: "20px", marginBottom: "30px"}}>
                <Row>
                    <Col span="18">
                        <Row >
                            <Col span="6" style={styleFieldName}>Target:&nbsp;&nbsp;</Col>
                            <Col><Text>{deployment.env}</Text></Col>
                        </Row>
                        <Row style={styleField}>
                            <Col span="6" style={styleFieldName}>Ref:&nbsp;&nbsp;</Col>
                            <Col><Text code>{deployment.ref}</Text></Col>
                        </Row>
                        <Row style={styleField}>
                            <Col span="6" style={styleFieldName}>Status:&nbsp;&nbsp;</Col>
                            <Col><DeploymentStatusBadge deployment={deployment}/></Col>
                        </Row>
                        <Row style={styleField}>
                            <Col span="6" style={styleFieldName}>Deployer:&nbsp;&nbsp;</Col>
                            <Col> 
                                {(deployment.deployer !== null)?
                                     <Text ><Avatar size="small" src={deployment.deployer.avatar} /> {deployment.deployer.login}</Text> :
                                    <Avatar size="small" >U</Avatar> }
                            </Col>
                        </Row>
                        {(deployment.isApprovalEanbled) ?
                            <Row style={styleField}>
                                <Col span="6" style={styleFieldName}>Required Approval:&nbsp;&nbsp;</Col>
                                <Col>{deployment.requiredApprovalCount}</Col>
                            </Row> :
                            null}
                        <Row style={styleField}>
                            <Col span="6" style={styleFieldName}></Col>
                            <Col>{deployBtn}</Col>
                        </Row>
                    </Col>
                    <Col span="6">
                        {(deployment.isApprovalEanbled) ? 
                            <div>
                                <div style={{paddingLeft: "5px"}}>
                                    <Text strong>Approvers</Text>
                                </div>
                                <div style={{marginTop: "5px"}}>
                                    <ApproversSearch
                                        style={{width: "100%"}}
                                        value="Select Approvers"
                                        approvers={approvers}
                                        candidates={candidates}
                                        onSearchCandidates={onSearchCandidates}
                                        onSelectCandidate={onSelectCandidate} />
                                </div>
                                <div style={{marginTop: "10px", paddingLeft: "5px"}}>
                                    {(approvals.length !== 0) ?
                                        <ApprovalList approvals={approvals}/>:
                                        <Text type="secondary"> No approvers </Text>}
                                </div>
                            </div> : null}
                    </Col>
                </Row>
            </div>
        </Main>
    )
}

function isDeployable(deployment: Deployment, approvals: Approval[]): boolean {
    if (deployment.status !== LastDeploymentStatus.Waiting) {
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