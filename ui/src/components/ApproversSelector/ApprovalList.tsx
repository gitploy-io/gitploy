import { Avatar } from "antd"
import { CheckOutlined, CloseOutlined } from "@ant-design/icons"

import { Approval, ApprovalStatus } from "../../models"

export interface ApprovalListProps {
    approvals: Approval[]
}

export default function ApprovalList(props: ApprovalListProps): JSX.Element {
    return (
        <div>
            {props.approvals.map((a, idx) => {
                const user = a.user
                const avatar = (user !== null)? 
                    <span><Avatar size="small" src={user.avatar}/> {user.login}</span> : 
                    <Avatar size="small">U</Avatar>
                const icon = mapApprovalStatusToIcon(a.status)

                return (
                    <div key={idx}>
                        {icon}&nbsp;
                        {avatar}
                    </div>)
            })}
        </div>
    )
}

function mapApprovalStatusToIcon(status: ApprovalStatus): JSX.Element {
    switch (status) {
        case ApprovalStatus.Pending:
            return <span style={{color: "gray"}}>•</span>
        case ApprovalStatus.Approved:
            return <CheckOutlined style={{color: "green"}} />
        case ApprovalStatus.Declined:
            return <CloseOutlined style={{color: "red"}} />
        default:
            return <span style={{color: "gray"}}>•</span>
    }
}