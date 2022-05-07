import { Breadcrumb } from 'antd';

export interface HeaderBreadcrumbProps {
  namespace: string;
  name: string;
  number: string;
}

export default function HeaderBreadcrumb({
  namespace,
  name,
  number,
}: HeaderBreadcrumbProps): JSX.Element {
  return (
    <Breadcrumb>
      <Breadcrumb.Item>
        <a href="/">Repositories</a>
      </Breadcrumb.Item>
      <Breadcrumb.Item>{namespace}</Breadcrumb.Item>
      <Breadcrumb.Item>
        <a href={`/${namespace}/${name}`}>{name}</a>
      </Breadcrumb.Item>
      <Breadcrumb.Item>Deployments</Breadcrumb.Item>
      <Breadcrumb.Item>{number}</Breadcrumb.Item>
    </Breadcrumb>
  );
}
