import { List, Typography } from 'antd';
import moment from 'moment';

import { Repo, Deployment } from '../../models';
import UserAvatar from '../../components/UserAvatar';
import DeploymentStatusBadge from '../../components/DeploymentStatusBadge';
import DeploymentRefCode from '../../components/DeploymentRefCode';

const { Text, Paragraph } = Typography;

export interface RepoListProps {
  repos: Repo[];
}

export default function RepoList({ repos }: RepoListProps): JSX.Element {
  return (
    <List
      dataSource={repos}
      renderItem={(repo) => {
        // deployments is undeinfed if there is no deployments of the repository.
        const deployment = repo.deployments ? repo.deployments[0] : null;

        return (
          <List.Item>
            <List.Item.Meta
              title={
                <a href={`/${repo.namespace}/${repo.name}`}>
                  {repo.namespace} / {repo.name}
                </a>
              }
              description={<Description deployment={deployment} />}
            />
          </List.Item>
        );
      }}
    />
  );
}

interface DescriptionProps {
  deployment: Deployment | null;
}

function Description(props: DescriptionProps): JSX.Element {
  if (!props.deployment) {
    return <></>;
  }

  return (
    <Paragraph style={{ marginTop: '10px', marginBottom: 0 }}>
      <UserAvatar user={props.deployment.deployer} /> deployed{' '}
      <DeploymentRefCode deployment={props.deployment} /> to the{' '}
      <Text strong>{props.deployment.env}</Text> environment{' '}
      {moment(props.deployment.createdAt).fromNow()}{' '}
      <DeploymentStatusBadge deployment={props.deployment} />
    </Paragraph>
  );
}
