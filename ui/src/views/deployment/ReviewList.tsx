import { Popover, Button, Descriptions, Typography } from 'antd';
import {
  CheckOutlined,
  CloseOutlined,
  CommentOutlined,
  ClockCircleOutlined,
} from '@ant-design/icons';

import { Review, ReviewStatusEnum } from '../../models';
import UserAvatar from '../../components/UserAvatar';

const { Text } = Typography;

export interface ReviewListProps {
  reviews: Review[];
}

export default function ReviewList({ reviews }: ReviewListProps): JSX.Element {
  if (reviews.length === 0) {
    return (
      <Descriptions title="Reviewers">
        <Descriptions.Item>
          <Text type="secondary">No reviewers</Text>
        </Descriptions.Item>
      </Descriptions>
    );
  }

  return (
    <Descriptions title="Reviewers" size="small" column={1}>
      {reviews.map((review, idx) => {
        return (
          <Descriptions.Item key={idx}>
            <ReviewStatusIcon review={review} />
            &nbsp;
            <UserAvatar user={review.user} boldName={false} />
            &nbsp;
            <ReviewCommentIcon review={review} />
          </Descriptions.Item>
        );
      })}
    </Descriptions>
  );
}

function ReviewStatusIcon(props: { review: Review }): JSX.Element {
  switch (props.review.status) {
    case ReviewStatusEnum.Pending:
      return <ClockCircleOutlined />;
    case ReviewStatusEnum.Approved:
      return <CheckOutlined style={{ color: 'green' }} />;
    case ReviewStatusEnum.Rejected:
      return <CloseOutlined style={{ color: 'red' }} />;
    default:
      return <ClockCircleOutlined />;
  }
}

function ReviewCommentIcon(props: { review: Review }): JSX.Element {
  const comment = props.review.comment;

  return comment ? (
    <Popover
      title="Comment"
      trigger="click"
      content={<div style={{ whiteSpace: 'pre' }}>{comment}</div>}
    >
      <Button type="text" style={{ padding: 0 }}>
        <CommentOutlined />
      </Button>
    </Popover>
  ) : (
    <></>
  );
}
