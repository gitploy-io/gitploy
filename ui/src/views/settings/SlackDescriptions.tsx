import { Button, Descriptions } from 'antd';

import { User } from '../../models';

export interface SlackDescriptionsProps {
  user?: User;
}

export default function SlackDescriptions({
  user,
}: SlackDescriptionsProps): JSX.Element {
  const connected = user?.chatUser ? true : false;

  return (
    <Descriptions title="Slack">
      <Descriptions.Item>
        {connected ? (
          <Button href="/slack/signout" type="primary" danger>
            DISCONNECTED
          </Button>
        ) : (
          <Button href="/slack/" type="primary">
            CONNECT
          </Button>
        )}
      </Descriptions.Item>
    </Descriptions>
  );
}
