import { useEffect } from "react"
import { shallowEqual } from 'react-redux'
import { useParams } from "react-router-dom"
import { Helmet } from "react-helmet"
import { Button, PageHeader, Result, Row, Col } from "antd"

import { useAppSelector, useAppDispatch } from "../../redux/hooks"
import { 
    deploymentSlice as slice, 
    fetchDeployment, 
    fetchDeploymentChanges,
    deployToSCM,
    fetchReviews,
    fetchUserReview,
    approve,
    reject,
} from "../../redux/deployment"
import { 
    Deployment, 
    DeploymentStatusEnum, 
    Review,
    ReviewStatusEnum,
    RequestStatus
} from "../../models"
import { subscribeEvents } from "../../apis"

import Main from "../main"
import HeaderBreadcrumb, { HeaderBreadcrumbProps } from "./HeaderBreadcrumb"
import ReviewButton, { ReviewButtonProps } from "./ReviewButton"
import ReviewerList, { ReviewListProps } from "./ReviewList"
import DeploymentDescriptor from "./DeploymentDescriptor"
import Spin from "../../components/Spin"
import DeploymentStatusSteps from "../../components/DeploymentStatusSteps"

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
        userReview: review,
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
            dispatch(slice.actions.handleReviewEvent(event))
        })

        return () => {
            sub.close()
        }
        // eslint-disable-next-line 
    }, [dispatch])

    const onClickApprove = (comment: string) => {
        dispatch(approve(comment))
    }

    const onClickApproveAndDeploy = (comment: string) => {
        const f = async () => {
            await dispatch(approve(comment))
            if (deployment?.status === DeploymentStatusEnum.Waiting) {
                await dispatch(deployToSCM())
            }
        }
        f()
    }

    const onClickReject = (comment: string) => {
        dispatch(reject(comment))
    }

    const onClickDeploy = () => {
        dispatch(deployToSCM())
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

    return (
        <Main>
            <Helmet>
                <title>Deployment #{number} - {namespace}/{name}</title>
            </Helmet>
            <div>
                <PageHeader
                    title={`Deployment #${number}`}
                    breadcrumb={
                        <HeaderBreadcrumb 
                            namespace={namespace} 
                            name={name} 
                            number={number} 
                        />
                    }
                    extra={
                        <ReviewButton 
                            review={review} 
                            onClickApprove={onClickApprove}
                            onClickApproveAndDeploy={onClickApproveAndDeploy}
                            onClickReject={onClickReject}
                        />
                    }
                    onBack={onBack} 
                />
            </div>
            <Row>
                <Col  span={23} offset={1} lg={{span: 13, offset: 1}}>
                    <DeploymentDescriptor 
                        changes={changes} 
                        deployment={deployment}
                    />
                </Col>
                <Col span={23} offset={1}  lg={{span: 6, offset: 2}}>
                    <ReviewerList 
                        reviews={reviews}
                    /> 
                </Col>
            </Row>
            <Row style={{marginTop: 40}}>
                <Col offset={1} span={22} md={{offset: 2}}>
                    {deployment.statuses?
                        <DeploymentStatusSteps statuses={deployment.statuses}/>
                        :
                        <></>}
                </Col>
            </Row>
            <Row style={{marginTop: 20}}>
                <Col offset={16}>
                    {isDeployable(deployment, reviews)?
                        <Button 
                            loading={RequestStatus.Pending === deploying} 
                            type="primary" 
                            onClick={onClickDeploy}
                        >
                            Deploy
                        </Button>
                        :
                        <Button 
                            type="primary" 
                            disabled
                        >
                          Deploy
                        </Button>}
                </Col>
            </Row>
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
