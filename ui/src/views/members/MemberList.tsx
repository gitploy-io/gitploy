import { shallowEqual } from "react-redux"
import { List, Switch, Button, Avatar } from "antd"

import { useAppSelector, useAppDispatch } from "../../redux/hooks"
import { updateUser, deleteUser } from "../../redux/members"

import { User } from "../../models"

export default function MemberList(): JSX.Element {
    const { users } = useAppSelector(state => state.members, shallowEqual)
    const dispatch = useAppDispatch()

    const onChangeSwitch = (user: User, checked: boolean) => {
        dispatch(updateUser({user, admin: checked}))
    }

    const onClickDelete = (user: User) => {
        dispatch(deleteUser(user))
    }

    return (
        <List
            itemLayout="horizontal"
            dataSource={users}
            renderItem={(user) => (
                <List.Item
                    actions={[
                        <Switch 
                            checkedChildren="Adm"  
                            unCheckedChildren="Mem"
                            checked={user.admin}
                            onChange={(checked) => {onChangeSwitch(user, checked)}}
                        />,
                        <Button
                            type="primary"
                            danger
                            onClick={() => {onClickDelete(user)}}
                        >
                            Delete
                        </Button>
                    ]}
                >
                    <List.Item.Meta 
                        avatar={<Avatar src={user.avatar} />}
                        title={user.login}
                    />
                </List.Item>
            )}
        />
    )
}