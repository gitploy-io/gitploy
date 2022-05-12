import camelcaseKeys from 'camelcase-keys';
import { StatusCodes } from 'http-status-codes';

import { instance, headers } from './setting';
import { _fetch } from './_base';
import {
  Deployment,
  DeploymentType,
  DeploymentStatus,
  DeploymentStatusEnum,
  Commit,
  HttpUnprocessableEntityError,
  HttpForbiddenError,
  HttpConflictError,
} from '../models';
import { UserData, mapDataToUser } from './user';
import { RepoData, mapDataToRepo } from './repo';
import { mapDataToCommit } from './commit';

export interface DeploymentData {
  id: number;
  number: number;
  type: string;
  ref: string;
  sha: string;
  env: string;
  status: string;
  uid: number;
  is_rollback: boolean;
  auto_deploy: boolean;
  created_at: string;
  updated_at: string;
  edges: {
    user?: UserData;
    repo?: RepoData;
    deployment_statuses?: DeploymentStatusData[];
  };
}

interface DeploymentStatusData {
  id: number;
  status: string;
  description: string;
  log_url: string;
  created_at: string;
  updated_at: string;
}

export const mapDataToDeployment = (data: any): Deployment => {
  const deployment: Deployment = camelcaseKeys(data);

  // Convert the time string into Date.
  deployment.createdAt = new Date(data.created_at);
  deployment.updatedAt = new Date(data.updated_at);

  if ('user' in data.edges) {
    deployment.deployer = mapDataToUser(data.edges.user);
  }

  if ('repo' in data.edges) {
    deployment.repo = mapDataToRepo(data.edges.repo);
  }

  if ('deployment_statuses' in data.edges) {
    deployment.statuses = data.edges.deployment_statuses.map((data: any) =>
      mapDataToDeploymentStatus(data)
    );
  }

  return deployment;
};

export function mapDataToDeploymentStatus(data: any): DeploymentStatus {
  const deploymentStatus = camelcaseKeys(data, { deep: true });

  if ('deployment' in data.edges) {
    deploymentStatus.edges.deployment = mapDataToDeployment(
      data.edges.deployment
    );
  }

  if ('repo' in data.edges) {
    deploymentStatus.edges.repo = mapDataToRepo(data.edges.repo);
  }

  return deploymentStatus;
}

function mapDeploymentStatusToString(status: DeploymentStatusEnum): string {
  switch (status) {
    case DeploymentStatusEnum.Waiting:
      return 'waiting';
    case DeploymentStatusEnum.Created:
      return 'created';
    case DeploymentStatusEnum.Queued:
      return 'queued';
    case DeploymentStatusEnum.Running:
      return 'running';
    case DeploymentStatusEnum.Success:
      return 'success';
    case DeploymentStatusEnum.Failure:
      return 'failure';
    case DeploymentStatusEnum.Canceled:
      return 'canceled';
    default:
      return 'waiting';
  }
}

export const searchDeployments = async (
  statuses: DeploymentStatusEnum[],
  owned: boolean,
  productionOnly: boolean,
  from?: Date,
  to?: Date,
  page = 1,
  perPage = 30
): Promise<Deployment[]> => {
  const ss: string[] = [];
  statuses.forEach((status) => {
    ss.push(mapDeploymentStatusToString(status));
  });

  const fromParam = from ? `from=${from.toISOString()}` : '';
  const toParam = to ? `&to=${to.toISOString()}` : '';

  const deployments: Deployment[] = await _fetch(
    `${instance}/api/v1/search/deployments?statuses=${ss.join(
      ','
    )}&owned=${owned}&production_only=${productionOnly}&${fromParam}&${toParam}&page=${page}&per_page=${perPage}`,
    {
      headers,
      credentials: 'same-origin',
    }
  )
    .then((response) => response.json())
    .then((ds) => ds.map((d: any): Deployment => mapDataToDeployment(d)));

  return deployments;
};

export const listDeployments = async (
  namespace: string,
  name: string,
  env: string,
  status: string,
  page: number,
  perPage: number
): Promise<Deployment[]> => {
  const deployments: Deployment[] = await _fetch(
    `${instance}/api/v1/repos/${namespace}/${name}/deployments?env=${env}&status=${status}&page=${page}&per_page=${perPage}`,
    {
      headers,
      credentials: 'same-origin',
    }
  )
    .then((response) => response.json())
    .then((ds) => ds.map((d: any): Deployment => mapDataToDeployment(d)));

  return deployments;
};

export const getDeployment = async (
  namespace: string,
  name: string,
  number: number
): Promise<Deployment> => {
  const deployment = await _fetch(
    `${instance}/api/v1/repos/${namespace}/${name}/deployments/${number}`,
    {
      headers,
      credentials: 'same-origin',
    }
  )
    .then((response) => response.json())
    .then((data) => mapDataToDeployment(data));

  return deployment;
};

// eslint-disable-next-line
export const createDeployment = async (
  namespace: string,
  name: string,
  type: DeploymentType = DeploymentType.Commit,
  ref: string,
  env: string,
  // eslint-disable-next-line
  payload?: any
): Promise<Deployment> => {
  const body = JSON.stringify({
    type,
    ref,
    env,
    dynamic_payload: payload,
  });
  const response = await _fetch(
    `${instance}/api/v1/repos/${namespace}/${name}/deployments`,
    {
      headers,
      credentials: 'same-origin',
      method: 'POST',
      body: body,
    }
  );
  if (response.status === StatusCodes.FORBIDDEN) {
    const message = await response.json().then((data) => data.message);
    throw new HttpForbiddenError(message);
  } else if (response.status === StatusCodes.UNPROCESSABLE_ENTITY) {
    const message = await response.json().then((data) => data.message);
    throw new HttpUnprocessableEntityError(message);
  } else if (response.status === StatusCodes.CONFLICT) {
    const message = await response.json().then((data) => data.message);
    throw new HttpConflictError(message);
  }

  const deployment = response.json().then((d) => mapDataToDeployment(d));
  return deployment;
};

export const createRemoteDeployment = async (
  namespace: string,
  name: string,
  number: number
): Promise<Deployment> => {
  const response = await _fetch(
    `${instance}/api/v1/repos/${namespace}/${name}/deployments/${number}`,
    {
      headers,
      credentials: 'same-origin',
      method: 'PUT',
    }
  );
  if (response.status === StatusCodes.FORBIDDEN) {
    const message = await response.json().then((data) => data.message);
    throw new HttpForbiddenError(message);
  } else if (response.status === StatusCodes.UNPROCESSABLE_ENTITY) {
    const message = await response.json().then((data) => data.message);
    throw new HttpUnprocessableEntityError(message);
  }

  const deployment = response.json().then((d) => mapDataToDeployment(d));
  return deployment;
};

export const rollbackDeployment = async (
  namespace: string,
  name: string,
  number: number
): Promise<Deployment> => {
  const response = await _fetch(
    `${instance}/api/v1/repos/${namespace}/${name}/deployments/${number}/rollback`,
    {
      headers,
      credentials: 'same-origin',
      method: 'POST',
    }
  );
  if (response.status === StatusCodes.FORBIDDEN) {
    const message = await response.json().then((data) => data.message);
    throw new HttpForbiddenError(message);
  } else if (response.status === StatusCodes.UNPROCESSABLE_ENTITY) {
    const message = await response.json().then((data) => data.message);
    throw new HttpUnprocessableEntityError(message);
  } else if (response.status === StatusCodes.CONFLICT) {
    const message = await response.json().then((data) => data.message);
    throw new HttpConflictError(message);
  }

  const deployment = response.json().then((d) => mapDataToDeployment(d));
  return deployment;
};

export const listDeploymentChanges = async (
  namespace: string,
  name: string,
  number: number,
  page = 1,
  perPage = 30
): Promise<Commit[]> => {
  const res = await _fetch(
    `${instance}/api/v1/repos/${namespace}/${name}/deployments/${number}/changes?page=${page}&per_page=${perPage}`,
    {
      headers,
      credentials: 'same-origin',
    }
  );

  if (res.status === StatusCodes.NOT_FOUND) {
    return [];
  }

  const commits: Commit[] = await res
    .json()
    .then((ds) => ds.map((d: any): Commit => mapDataToCommit(d)));

  return commits;
};
