import { useEffect } from "react"
import { Breadcrumb, PageHeader, Row, Col, Divider, Result } from "antd"
import { shallowEqual } from 'react-redux'
import { useParams } from "react-router-dom"

import { useAppSelector, useAppDispatch } from "../redux/hooks"
import { 
    deploymentSlice as slice, 
    fetchDeployment, 
    fetchDeploymentChanges,
    deployToSCM,
    fetchReviews,
    approve,
    reject,
    fetchUserReview,
} from "../redux/deployment"
import { 
    Deployment, 
    DeploymentStatusEnum, 
    Review,
    ReviewStatusEnum,
    RequestStatus
} from "../models"
import { subscribeEvents } from "../apis"

import Main from "./Main"
import ReviewDropdown from "../components/ReviewDropdown"
import Spin from "../components/Spin"
import DeployConfirm from "../components/DeployConfirm"
import ReviewList from "../components/ReviewList"

interface Params {
    namespace: string
    name: string
    number: string
}

export default function DeploymentView(): JSX.Element {
    const { namespace, name, number } = useParams<Params>()
    const { 
        display,
        deployment, 
        changes,
        deploying,
        reviews,
        userReview,
    } = useAppSelector(state => state.deployment, shallowEqual )
    const dispatch = useAppDispatch()

    useEffect(() => {
        const f = async () => {
            await dispatch(slice.actions.init({namespace, name, number: parseInt(number, 10)}))
            await dispatch(fetchDeployment())
            await dispatch(fetchReviews())
            await dispatch(fetchUserReview())
            await dispatch(slice.actions.setDisplay(true))
            await dispatch(fetchDeploymentChanges())
        }
        f()

        const sub = subscribeEvents((event) => {
            dispatch(slice.actions.handleDeploymentEvent(event))
        })

        return () => {
            sub.close()
        }
        // eslint-disable-next-line 
    }, [dispatch])

    const hasReview = reviews.length > 0

    const onClickDeploy = () => {
        dispatch(deployToSCM())
    }

    const onClickApprove = () => {
        dispatch(approve())
    }

    const onClickReject = () => {
        dispatch(reject())
    }

    const onBack = () => {
        window.location.href = `/${namespace}/${name}`
    }

    if (!display) {
        return (
            <Main>
                <div style={{textAlign: "center", marginTop: "100px"}}>
                    <Spin />
                </div>
            </Main>
        )
    }

    if (!deployment) {
        return (
            <Main>
                <Result
                    status="warning"
                    title="The deployment is not found."
                />
            </Main>
        )
    }

    // buttons
    const approvalDropdown = (userReview)?
        <ReviewDropdown 
            key="approval" 
            onClickApprove={onClickApprove}
            onClickReject={onClickReject}/>:
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
                        {(hasReview) ? 
                            <ReviewList 
                                reviews={reviews}
                            /> :
                            null}
                        <Divider />
                    </Col>
                    <Col xs={{span: 24}} md={(hasReview)? {span: 18} : {span: 21}}>
                        <DeployConfirm 
                            isDeployable={isDeployable(deployment, reviews)}
                            deploying={RequestStatus.Pending === deploying}
                            deployment={deployment}
                            changes={changes}
                            onClickDeploy={onClickDeploy}
                        />
                    </Col>
                    <Col xs={{span: 0}} md={{span: 6}}>
                        {/* Desktop view */}
                        {(hasReview) ? 
                            <ReviewList 
                                reviews={reviews}
                            /> :
                            null}
                    </Col>
                </Row>
            </div>
        </Main>
    )
}

function isDeployable(deployment: Deployment, reviews: Review[]): boolean {
    if (deployment.status !== DeploymentStatusEnum.Waiting) {
        return false
    }

    for (let i = 0; i < reviews.length; i++) {
        if (reviews[i].status === ReviewStatusEnum.Rejected) {
            return false
        }
    }

    for (let i = 0; i < reviews.length; i++) {
        if (reviews[i].status === ReviewStatusEnum.Approved) {
            return true
        }
    }

    return false
}
