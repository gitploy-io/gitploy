import { Avatar, Typography } from "antd"

import { User } from "../models";

const { Text } = Typography

interface UserAvatarProps {
    user?: User
}

export default function UserAvatar(props: UserAvatarProps): JSX.Element {
    return (
        props.user?
                <span><Avatar size="small" src={props.user.avatar} /> <Text strong>{props.user.login}</Text></span> :
                <span><Avatar size="small">U</Avatar> </span> 
    )
}