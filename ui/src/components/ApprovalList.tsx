import { Divider, Avatar } from "antd"
import { CheckCircleTwoTone, MinusCircleOutlined } from "@ant-design/icons"

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
                    <span><Avatar size="small" src={user.avatar}/> {user.login}</span> : 
                    <Avatar size="small">U</Avatar>
                const divider = (idx !== props.approvals.length - 1)?
                    <Divider type="vertical" />:
                    null
                const icon = (a.isApproved)?
                    <CheckCircleTwoTone twoToneColor="#52c41a" />:
                    <MinusCircleOutlined />

                return (
                    <div key={idx}>
                        {avatar}&nbsp;
                        {icon}
                        {divider}
                    </div>)
            })}
        </div>
    )
}