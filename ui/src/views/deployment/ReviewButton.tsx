import { useState } from "react"
import { shallowEqual } from 'react-redux'
import { Button, Modal, Space, Input } from "antd"

import { ReviewStatusEnum, DeploymentStatusEnum } from "../../models"
import { useAppSelector, useAppDispatch } from "../../redux/hooks"
import { 
    deployToSCM,
    approve,
    reject,
} from "../../redux/deployment"

export default function ReviewButton(): JSX.Element {
    const { 
        deployment, 
        userReview: review,
    } = useAppSelector(state => state.deployment, shallowEqual )
    const dispatch = useAppDispatch()

    // If no review has been assigned, an empty value is returned.
    if (!review) {
        return <></>
    }

    // Stores the review comment state and passes it as a parameter when updating.
    const [comment, setComment] = useState(review.comment)

    const onChangeComment = (e: any) => {
        setComment(e.target.value)
    }

    const [isModalVisible, setIsModalVisible] = useState(false);

    const showModal = () => {
        setIsModalVisible(true);
    }

    const onClickCancel = () => {
        setIsModalVisible(false)
    }
  
    const onClickApprove = () => {
        dispatch(approve(comment))
        setIsModalVisible(false)
    }

    const onClickApproveAndDeploy = () => {
        const f = async () => {
            await dispatch(approve(comment))
            if (deployment?.status === DeploymentStatusEnum.Waiting) {
                await dispatch(deployToSCM())
            }
        }

        f()
        setIsModalVisible(false)
    }

    const onClickReject = () => {
        dispatch(reject(comment))
        setIsModalVisible(false)
    }
  
    return (
        <>
            {(review.status === ReviewStatusEnum.Pending)? 
                <Button type="primary" onClick={showModal}>
                    Review
                </Button> 
                :
                <Button onClick={showModal}>
                    Reviewed
                </Button>}
            <Modal 
                title="Review" 
                visible={isModalVisible} 
                onCancel={onClickCancel}
                footer={(
                    <Space>
                        <Button type="primary" danger onClick={onClickReject}>Reject</Button>
                        <Button type="primary" onClick={onClickApproveAndDeploy}>Approve and Deploy</Button>
                        <Button type="primary" onClick={onClickApprove}>Approve</Button>
                    </Space>
                )}
            >
                <Input.TextArea rows={3} onChange={onChangeComment} value={comment}/>
            </Modal>
        </>
    )

}