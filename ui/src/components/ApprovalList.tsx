import { Divider, Avatar } from "antd"

import { Approval } from "../models"

interface ApprovalListProps {
    approvals: Approval[]
}

export default function ApprovalList(props: ApprovalListProps) {
    return (
        <div>
            {props.approvals.map((a, idx) => {
                const user = a.user
                const avatar = (user !== null)? 
                    <span><Avatar size="small" src={user.avatar}/>&nbsp; {user.login}</span> : 
                    <Avatar size="small">U</Avatar>
                const divider = (idx !== props.approvals.length - 1)?
                    <Divider type="vertical" />:
                    null

                return (
                    <div key={idx}>
                        {avatar}
                        {divider}
                    </div>)
            })}
        </div>
    )
}