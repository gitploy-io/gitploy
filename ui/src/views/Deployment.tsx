import { useEffect } from "react"
import { Breadcrumb, PageHeader, Result } from "antd"
import { shallowEqual } from 'react-redux'
import { useParams } from "react-router-dom"
import { Helmet } from "react-helmet"

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
import ReviewModal from "../components/ReviewModal"
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
            dispatch(slice.actions.handleReviewEvent(event))
        })

        return () => {
            sub.close()
        }
        // eslint-disable-next-line 
    }, [dispatch])

    const onClickDeploy = () => {
        dispatch(deployToSCM())
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

    const onClickApprove = (comment: string) => {
       dispatch(approve(comment))
    }

    const onClickReject = (comment: string) => {
        dispatch(reject(comment))
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

    const reviewBtn = (userReview)?
        <ReviewModal 
            key={0}
            review={userReview}
            onClickApproveAndDeploy={onClickApproveAndDeploy}
            onClickApprove={onClickApprove}
            onClickReject={onClickReject}
        />:
        <></>

    return (
        <Main>
            <Helmet>
                <title>Deployment #{number} - {namespace}/{name}</title>
            </Helmet>
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
                    extra={reviewBtn}
                    onBack={onBack} 
                />
            </div>
            <div style={{marginTop: "20px", marginBottom: "30px"}}>
                <DeployConfirm 
                    isDeployable={isDeployable(deployment, reviews)}
                    deploying={RequestStatus.Pending === deploying}
                    deployment={deployment}
                    changes={changes}
                    reviews={reviews}
                    onClickDeploy={onClickDeploy}
                />
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
