import { useState } from "react"
import { Button, Modal, Space, Input } from "antd"

import { Review, ReviewStatusEnum } from "../models"

const { TextArea } = Input

interface ReviewModalProps {
    review: Review
    onClickApprove(comment: string): void
    onClickApproveAndDeploy(comment: string): void
    onClickReject(comment: string): void
}

export default function ReviewModal(props: ReviewModalProps): JSX.Element {
    const [comment, setComment] = useState(props.review.comment)

    const onChangeComment = (e: any) => {
        setComment(e.target.value)
    }

    const [isModalVisible, setIsModalVisible] = useState(false);

    const showModal = () => {
        setIsModalVisible(true);
    }
  
    const onClickApprove = () => {
        props.onClickApprove(comment)
        setIsModalVisible(false)
    }

    const onClickApproveAndDeploy = () => {
        props.onClickApproveAndDeploy(comment)
        setIsModalVisible(false)
    }

    const onClickReject = () => {
        props.onClickReject(comment)
        setIsModalVisible(false)
    }
  
    const onClickCancel = () => {
        setIsModalVisible(false)
    }

    return (
        <>
            <Modal 
                title="Review" 
                visible={isModalVisible} 
                onCancel={onClickCancel}
                footer={
                    <Space>
                        <Button type="primary" danger onClick={onClickReject}>Reject</Button>
                        <Button type="primary" onClick={onClickApproveAndDeploy}>Approve and Deploy</Button>
                        <Button type="primary" onClick={onClickApprove}>Approve</Button>
                    </Space>
                }
            >
                <TextArea rows={3} onChange={onChangeComment} value={comment}/>
            </Modal>
            {(props.review.status === ReviewStatusEnum.Pending)? 
                <Button type="primary" onClick={showModal}>
                    Review
                </Button> :
                <Button onClick={showModal}>
                    Reviewed
                </Button> 
            }
        </>
    )

}