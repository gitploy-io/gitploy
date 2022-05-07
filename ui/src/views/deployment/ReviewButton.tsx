import { useState } from 'react';
import { Button, Modal, Space, Input } from 'antd';

import { Review, ReviewStatusEnum } from '../../models';

export interface ReviewButtonProps {
  review?: Review;
  onClickApprove(comment: string): void;
  onClickApproveAndDeploy(comment: string): void;
  onClickReject(comment: string): void;
}

export default function ReviewButton(props: ReviewButtonProps): JSX.Element {
  const { review } = props;

  // If no review has been assigned, an empty value is returned.
  if (!review) {
    return <></>;
  }

  // Stores the review comment state and passes it as a parameter when updating.
  const [comment, setComment] = useState(review.comment);

  const onChangeComment = (e: any) => {
    setComment(e.target.value);
  };

  const [isModalVisible, setIsModalVisible] = useState(false);

  const showModal = () => {
    setIsModalVisible(true);
  };

  const onClickCancel = () => {
    setIsModalVisible(false);
  };

  const onClickApprove = () => {
    props.onClickApprove(comment);
    setIsModalVisible(false);
  };

  const onClickApproveAndDeploy = () => {
    props.onClickApproveAndDeploy(comment);
    setIsModalVisible(false);
  };

  const onClickReject = () => {
    props.onClickReject(comment);
    setIsModalVisible(false);
  };

  return (
    <>
      {review.status === ReviewStatusEnum.Pending ? (
        <Button type="primary" onClick={showModal}>
          Review
        </Button>
      ) : (
        <Button onClick={showModal}>Reviewed</Button>
      )}
      <Modal
        title="Review"
        visible={isModalVisible}
        onCancel={onClickCancel}
        footer={
          <Space>
            <Button type="primary" danger onClick={onClickReject}>
              Reject
            </Button>
            <Button type="primary" onClick={onClickApproveAndDeploy}>
              Approve and Deploy
            </Button>
            <Button type="primary" onClick={onClickApprove}>
              Approve
            </Button>
          </Space>
        }
      >
        <Input.TextArea rows={3} onChange={onChangeComment} value={comment} />
      </Modal>
    </>
  );
}
