import { Timeline, Typography } from 'antd';
import moment from 'moment';

import { Deployment } from '../../models';
import DeploymentStatusBadge from '../../components/DeploymentStatusBadge';
import UserAvatar from '../../components/UserAvatar';
import DeploymentRefCode from '../../components/DeploymentRefCode';
import { getStatusColor } from '../../components/partials';

const { Text } = Typography;

export interface ActivityHistoryProps {
  deployments: Deployment[];
}

export default function ActivityHistory(
  props: ActivityHistoryProps
): JSX.Element {
  return (
    <Timeline>
      {props.deployments.map((d, idx) => {
        return (
          <Timeline.Item key={idx} color={getStatusColor(d.status)}>
            <p>
              <Text strong>{`${d.repo?.namespace} / ${d.repo?.name}`}</Text>
              &nbsp;
              <a
                href={`/${d.repo?.namespace}/${d.repo?.name}/deployments/${d.number}`}
              >
                #{d.number}
              </a>
            </p>
            <p>
              <UserAvatar user={d.deployer} /> deployed{' '}
              <DeploymentRefCode deployment={d} /> to{' '}
              <Text strong>{d.env}</Text> on {moment(d.createdAt).format('LLL')}{' '}
              <DeploymentStatusBadge deployment={d} />
            </p>
          </Timeline.Item>
        );
      })}
    </Timeline>
  );
}
