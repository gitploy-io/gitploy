import { List, Avatar, Popover, Button } from "antd"
import { CheckOutlined, CloseOutlined, CommentOutlined, ClockCircleOutlined } from "@ant-design/icons"

import { User, Review, ReviewStatusEnum } from "../models"

export interface ReviewerListProps {
    reviews: Review[]
}

export default function ReviewerList(props: ReviewerListProps): JSX.Element {
    return (
        <List 
            dataSource={props.reviews}
            renderItem={(review) => {
                return (
                    <div>
                        <ReviewItem review={review} />
                    </div>
                )
            }} 
        />
    )
}

function ReviewItem(props: {review: Review}): JSX.Element {
    const status = (status: ReviewStatusEnum) => {
        switch (status) {
            case ReviewStatusEnum.Pending:
                return <ClockCircleOutlined />
            case ReviewStatusEnum.Approved:
                return <CheckOutlined style={{color: "green"}} />
            case ReviewStatusEnum.Rejected:
                return <CloseOutlined style={{color: "red"}} />
            default:
                return <ClockCircleOutlined />
        }
    }

    const avatar = (user?: User) => {
        return user?
            <span><Avatar size="small" src={user.avatar} /> {user.login}</span>:
            <span><Avatar size="small">U</Avatar> </span> 
    }

    const commentIcon = (comment: string) => {
        return comment !== ""? 
            <Popover
                title="Comment"
                trigger="click"
                content={<div style={{whiteSpace: "pre"}}>{comment}</div>}
            >
                <Button 
                    type="text" 
                    style={{padding: 0}}
                >
                    <CommentOutlined />
                </Button>
            </Popover>:
            <></>
    }

    return (
        <p>
            {status(props.review.status)} {avatar(props.review.user)} {commentIcon(props.review.comment)}
        </p>
    )
}

export function ReviewStatus(props: {reviews: Review[]}): JSX.Element {
    for (let i = 0; i < props.reviews.length; i++) {
        if (props.reviews[i].status === ReviewStatusEnum.Rejected) {
            return <span>
                <CloseOutlined style={{color: "red"}} /> Rejected
            </span>
        }
    }

    for (let i = 0; i < props.reviews.length; i++) {
        if (props.reviews[i].status === ReviewStatusEnum.Approved) {
            return <span>
                <CheckOutlined style={{color: "green"}} /> Approved
            </span>
        }
    }

    return <span><ClockCircleOutlined />&nbsp;&nbsp;Pending</span>
}
