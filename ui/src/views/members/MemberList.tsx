import { List, Switch, Button, Avatar } from 'antd';

import { User } from '../../models';

export interface MemberListProps {
  users: User[];
  onChangeSwitch(user: User, checked: boolean): void;
  onClickDelete(user: User): void;
}

export default function MemberList(props: MemberListProps): JSX.Element {
  return (
    <List
      itemLayout="horizontal"
      dataSource={props.users}
      renderItem={(user) => (
        <List.Item
          actions={[
            <Switch
              checkedChildren="Adm"
              unCheckedChildren="Mem"
              checked={user.admin}
              onChange={(checked) => {
                props.onChangeSwitch(user, checked);
              }}
            />,
            <Button
              type="primary"
              danger
              onClick={() => {
                props.onClickDelete(user);
              }}
            >
              Delete
            </Button>,
          ]}
        >
          <List.Item.Meta
            avatar={<Avatar src={user.avatar} />}
            title={user.login}
          />
        </List.Item>
      )}
    />
  );
}
