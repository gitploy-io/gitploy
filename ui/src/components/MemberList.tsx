import { List, Switch, Button, Avatar } from "antd"

import { User } from "../models"

interface MemberListProps {
    users: User[]
    onChangeSwitch(user: User, checked: boolean): void
    onClickDelete(user: User): void
}

export default function MemberList(props: MemberListProps): JSX.Element {
    return <List
        itemLayout="horizontal"
        dataSource={props.users}
        renderItem={(u) => {
            return <List.Item
                actions={[
                    <Switch 
                        checkedChildren="Adm"  
                        unCheckedChildren="Mem"
                        checked={u.admin}
                        onChange={(checked) => {props.onChangeSwitch(u, checked)}}
                    />,
                    <Button
                        type="primary"
                        danger
                        onClick={() => {props.onClickDelete(u)}}
                    >Delete</Button>
                ]}
            >
                <List.Item.Meta 
                    avatar={<Avatar src={u.avatar} />}
                    title={u.login}
                />
            </List.Item>
        }}
    >

    </List>
}