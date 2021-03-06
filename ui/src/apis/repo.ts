import camelcaseKeys from 'camelcase-keys';
import { StatusCodes } from 'http-status-codes';

import { instance, headers } from './setting';
import { _fetch } from './_base';
import { mapDataToDeployment } from './deployment';

import {
  Repo,
  HttpForbiddenError,
  HttpUnprocessableEntityError,
} from '../models';

export const mapDataToRepo = (data: any): Repo => {
  const repo: Repo = camelcaseKeys(data);

  repo.createdAt = new Date(data.created_at);
  repo.updatedAt = new Date(data.updated_at);

  if ('deployments' in data.edges) {
    repo.deployments = data.edges.deployments.map((data) =>
      mapDataToDeployment(data)
    );
  }

  return repo;
};

export const listRepos = async (
  q: string,
  page = 1,
  perPage = 30
): Promise<Repo[]> => {
  const repos: Repo[] = await _fetch(
    `${instance}/api/v1/repos?q=${q}&sort=true&page=${page}&per_page=${perPage}`,
    {
      headers,
      credentials: 'same-origin',
    }
  )
    .then((response) => response.json())
    .then((repos) => repos.map((r: any): Repo => mapDataToRepo(r)));

  return repos;
};

export const getRepo = async (
  namespace: string,
  name: string
): Promise<Repo> => {
  const repo: Repo = await _fetch(
    `${instance}/api/v1/repos/${namespace}/${name}`,
    {
      headers,
      credentials: 'same-origin',
    }
  )
    .then((response) => response.json())
    .then((repo) => mapDataToRepo(repo));

  return repo;
};

export const updateRepo = async (
  namespace: string,
  name: string,
  payload: {
    name?: string;
    config_path?: string;
  }
): Promise<Repo> => {
  const res = await _fetch(`${instance}/api/v1/repos/${namespace}/${name}`, {
    headers,
    credentials: 'same-origin',
    method: 'PATCH',
    body: JSON.stringify(payload),
  });
  if (res.status === StatusCodes.FORBIDDEN) {
    const message = await res.json().then((data) => data.message);
    throw new HttpForbiddenError(message);
  } else if (res.status === StatusCodes.UNPROCESSABLE_ENTITY) {
    const message = await res.json().then((data) => data.message);
    throw new HttpUnprocessableEntityError(message);
  }

  const ret: Repo = await res.json().then((repo: any) => mapDataToRepo(repo));

  return ret;
};

export const activateRepo = async (
  namespace: string,
  name: string
): Promise<Repo> => {
  const body = {
    active: true,
  };
  const response = await _fetch(
    `${instance}/api/v1/repos/${namespace}/${name}`,
    {
      headers,
      credentials: 'same-origin',
      method: 'PATCH',
      body: JSON.stringify(body),
    }
  );
  if (response.status === StatusCodes.FORBIDDEN) {
    const message = await response.json().then((data) => data.message);
    throw new HttpForbiddenError(message);
  }

  const repo = await response.json().then((r: any) => mapDataToRepo(r));
  return repo;
};

export const deactivateRepo = async (
  namespace: string,
  name: string
): Promise<Repo> => {
  const body = {
    active: false,
  };
  const response = await _fetch(
    `${instance}/api/v1/repos/${namespace}/${name}`,
    {
      headers,
      credentials: 'same-origin',
      method: 'PATCH',
      body: JSON.stringify(body),
    }
  );
  if (response.status === StatusCodes.FORBIDDEN) {
    const message = await response.json().then((data) => data.message);
    throw new HttpForbiddenError(message);
  }

  const repo = await response.json().then((r: any) => mapDataToRepo(r));
  return repo;
};

export const lockRepo = async (
  namespace: string,
  name: string
): Promise<Repo> => {
  const body = {
    locked: true,
  };
  const response = await _fetch(
    `${instance}/api/v1/repos/${namespace}/${name}`,
    {
      headers,
      credentials: 'same-origin',
      method: 'PATCH',
      body: JSON.stringify(body),
    }
  );
  if (response.status === StatusCodes.FORBIDDEN) {
    const message = await response.json().then((data) => data.message);
    throw new HttpForbiddenError(message);
  }

  const repo = await response.json().then((r: any) => mapDataToRepo(r));
  return repo;
};

export const unlockRepo = async (
  namespace: string,
  name: string
): Promise<Repo> => {
  const body = {
    locked: false,
  };
  const response = await _fetch(
    `${instance}/api/v1/repos/${namespace}/${name}`,
    {
      headers,
      credentials: 'same-origin',
      method: 'PATCH',
      body: JSON.stringify(body),
    }
  );
  if (response.status === StatusCodes.FORBIDDEN) {
    const message = await response.json().then((data) => data.message);
    throw new HttpForbiddenError(message);
  }

  const repo = await response.json().then((r: any) => mapDataToRepo(r));
  return repo;
};
