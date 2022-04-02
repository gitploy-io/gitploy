import { shallowEqual } from "react-redux"
import { Tag, Descriptions, Input } from "antd"

import { useAppSelector } from "../../redux/hooks"

export default function UserDescriptions(): JSX.Element {
    const { user } = useAppSelector(state => state.settings, shallowEqual)

    return (
        <Descriptions 
            title="User Info" 
            column={2} 
            layout="vertical"
            style={{marginTop: "40px"}} 
        >
            <Descriptions.Item label="Login">
                {user?.login}
            </Descriptions.Item>
            <Descriptions.Item label="Role">
                {(user?.admin)? 
                    <Tag color="purple">Admin</Tag> 
                    : 
                    <Tag color="purple">Member</Tag>}
            </Descriptions.Item>
            <Descriptions.Item label="Token">
                <Input.Password 
                    value={user?.hash}
                    style={{width: 200, padding:0 }}
                    readOnly
                    bordered={false}
                />
            </Descriptions.Item>
        </Descriptions>
    )
}