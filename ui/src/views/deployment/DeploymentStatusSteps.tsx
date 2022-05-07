import { Timeline, Typography } from 'antd';
import { ClockCircleOutlined } from '@ant-design/icons';
import moment from 'moment';

import { DeploymentStatus } from '../../models';

const { Text, Link } = Typography;

export interface DeploymentStatusStepsProps {
  statuses: DeploymentStatus[];
}

export default function DeploymentStatusSteps(
  props: DeploymentStatusStepsProps
): JSX.Element {
  return (
    <Timeline>
      {props.statuses.map((status, idx) => {
        return (
          <Timeline.Item key={idx} color={getStatusColor(status.status)}>
            <ClockCircleOutlined />{' '}
            {moment(status.createdAt).format('YYYY-MM-DD HH:mm:ss')}
            <br />
            <b>{status.description}</b>&nbsp;&nbsp;
            {status.logUrl !== '' ? (
              <Link href={status.logUrl} target="_blank">
                View Detail
              </Link>
            ) : (
              <></>
            )}
            <br />
            Updated{' '}
            <Text className="gitploy-code" code>
              {status.status}
            </Text>{' '}
            {moment(status.createdAt).fromNow()}
          </Timeline.Item>
        );
      })}
    </Timeline>
  );
}

const getStatusColor = (status: string) => {
  switch (status) {
    case 'success':
      return 'green';
    case 'failure':
      return 'red';
    default:
      return 'purple';
  }
};
