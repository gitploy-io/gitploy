import { Avatar } from "antd"
import { CheckOutlined } from "@ant-design/icons"

import { Approval } from "../../models"

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
                const icon = (a.isApproved)?
                    <CheckOutlined style={{color: "green"}}/>:
                    <span className="ant-badge-status-dot" style={{background: "gray"}}></span>

                return (
                    <div key={idx}>
                        {icon}&nbsp;&nbsp;&nbsp;
                        {avatar}
                    </div>)
            })}
        </div>
    )
}