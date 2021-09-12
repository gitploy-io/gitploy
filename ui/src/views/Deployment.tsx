import { useEffect } from "react"
import { Breadcrumb, PageHeader, Row, Col, Divider } from "antd"
import { shallowEqual } from 'react-redux'
import { useParams } from "react-router-dom"

import { useAppSelector, useAppDispatch } from "../redux/hooks"
import { 
    init, 
    deploymentSlice as slice, 
    fetchDeployment, 
    fetchDeploymentChanges,
    fetchApprovals, 
    fetchMyApproval,
    deployToSCM,
    searchCandidates,
    createApproval,
    deleteApproval,
    approve,
    decline,
} from "../redux/deployment"
import { 
    User, 
    Deployment, 
    DeploymentStatusEnum, 
    Approval,
    ApprovalStatus,
    RequestStatus
} from "../models"
import { subscribeDeploymentEvent } from "../apis"

import Main from "./Main"
import ApproversSelector from "../components/ApproversSelector"
import ApprovalDropdown from "../components/ApprovalDropdown"
import Spin from "../components/Spin"
import DeployConfirm from "../components/DeployConfirm"

interface Params {
    namespace: string
    name: string
    number: string
}

export default function DeploymentView(): JSX.Element {
    const { namespace, name, number } = useParams<Params>()
    const { 
        repo,
        deployment, 
        changes,
        deploying,
        approvals, 
        candidates, 
        myApproval 
    } = useAppSelector(state => state.deployment, shallowEqual )
    const dispatch = useAppDispatch()

    useEffect(() => {
        const f = async () => {
            await dispatch(init({namespace, name}))
            await dispatch(slice.actions.setNumber(parseInt(number, 10)))
            await dispatch(fetchDeployment())
            await dispatch(fetchDeploymentChanges())
            await dispatch(fetchApprovals())
            await dispatch(fetchMyApproval())
        }
        f()

        const de = subscribeDeploymentEvent((d) => {
            dispatch(slice.actions.handleDeploymentEvent(d))
        })

        return () => {
            de.close()
        }
        // eslint-disable-next-line 
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

    if (!deployment) {
        return <Main>
            <div style={{textAlign: "center", marginTop: "100px"}}><Spin /></div>
        </Main>
    }

    // buttons
    const approvalDropdown = (myApproval)?
        <ApprovalDropdown 
            key="approval" 
            onClickApprove={onClickApprove}
            onClickDecline={onClickDecline}/>:
        null

    return (
        <Main>
            <div>
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
                    onBack={onBack} 
                />
            </div>
            <div style={{marginTop: "20px", marginBottom: "30px"}}>
                <Row>
                    <Col xs={{span: 24}} md={{span: 0}}>
                        {/* Mobile view */}
                        {(deployment.isApprovalEanbled) ? 
                            <ApproversSelector 
                                approvers={approvers}
                                candidates={candidates}
                                approvals={approvals}
                                onSearchCandidates={onSearchCandidates}
                                onSelectCandidate={onSelectCandidate}
                            /> :
                            null}
                        <Divider />
                    </Col>
                    <Col xs={{span: 24}} md={{span: 18}}>
                        <DeployConfirm 
                            isDeployable={(repo?.locked)? false : isDeployable(deployment, approvals)}
                            deploying={RequestStatus.Pending === deploying}
                            deployment={deployment}
                            changes={changes}
                            onClickDeploy={onClickDeploy}
                        />
                    </Col>
                    <Col xs={{span: 0}} md={{span: 6}}>
                        {/* Desktop view */}
                        {(deployment.isApprovalEanbled) ? 
                            <ApproversSelector 
                                approvers={approvers}
                                candidates={candidates}
                                approvals={approvals}
                                onSearchCandidates={onSearchCandidates}
                                onSelectCandidate={onSelectCandidate}
                            /> :
                            null}
                    </Col>
                </Row>
            </div>
        </Main>
    )
}

function isDeployable(deployment: Deployment, approvals: Approval[]): boolean {
    if (deployment.status !== DeploymentStatusEnum.Waiting) {
        return false
    }

    // requiredApprovalCount have to be equal or greater than approved.
    let approved = 0
    approvals.forEach((approval) => {
        if (approval.status === ApprovalStatus.Approved) {
            approved++
        }
    })

    return approved >= deployment.requiredApprovalCount
}
