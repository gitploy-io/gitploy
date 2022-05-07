import { Avatar, Typography } from 'antd';

import { User } from '../models';

const { Text } = Typography;

interface UserAvatarProps {
  boldName?: boolean;
  user?: User;
}

export default function UserAvatar(props: UserAvatarProps): JSX.Element {
  const boldName = props.boldName === undefined ? true : props.boldName;
  return props.user ? (
    <span>
      <Avatar size="small" src={props.user.avatar} />
      &nbsp;
      <Text strong={boldName}>{props.user.login}</Text>
    </span>
  ) : (
    <span>
      <Avatar size="small">U</Avatar>
    </span>
  );
}
