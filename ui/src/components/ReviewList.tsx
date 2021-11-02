import { Typography } from "antd"
import { CheckOutlined, CloseOutlined } from "@ant-design/icons"

import { Review, ReviewStatusEnum } from "../models"
import UserAvatar from "./UserAvatar"

const { Text } = Typography


export interface ReviewListProps {
    reviews: Review[]
}

export default function ReviewList(props: ReviewListProps): JSX.Element {
    return (
        <div>
            <div style={{paddingLeft: "5px"}}>
                <Text strong>Reviewers</Text>
            </div>
            <div style={{marginTop: "10px", paddingLeft: "5px"}}>
                {(props.reviews.length !== 0) ?
                    <div>
                        {props.reviews.map((r, idx) => {
                            return (
                                <div key={idx}>
                                    {mapApprovalStatusToIcon(r.status)}&nbsp;<UserAvatar user={r.user}/>
                                </div>
                            )
                        })}
                    </div> :
                    <Text type="secondary"> No approvers </Text>}
            </div>
        </div>
    )
}

function mapApprovalStatusToIcon(status: ReviewStatusEnum): JSX.Element {
    switch (status) {
        case ReviewStatusEnum.Pending:
            return <span style={{color: "gray"}}>•</span>
        case ReviewStatusEnum.Approved:
            return <CheckOutlined style={{color: "green"}} />
        case ReviewStatusEnum.Rejected:
            return <CloseOutlined style={{color: "red"}} />
        default:
            return <span style={{color: "gray"}}>•</span>
    }
}