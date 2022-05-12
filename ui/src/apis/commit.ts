import camelcaseKeys from 'camelcase-keys';
import { StatusCodes } from 'http-status-codes';

import { instance, headers } from './setting';
import { _fetch } from './_base';
import { Commit, Status, HttpNotFoundError, StatusState } from '../models';

export const mapDataToCommit = (data: any): Commit => {
  const commit: Commit = camelcaseKeys(data, { deep: true });

  if (commit.author) {
    // Convert the type of date field.
    commit.author.date = new Date(data.author.date);
  }

  return commit;
};

const mapDataToStatus = (data: any): Status => {
  const status = camelcaseKeys(data, { deep: true });
  return status;
};

const mapStatusState = (state: string): StatusState => {
  switch (state) {
    case 'pending':
      return StatusState.Pending;
    case 'success':
      return StatusState.Success;
    case 'failure':
      return StatusState.Failure;
    case 'cancelled':
      return StatusState.Cancelled;
    case 'skipped':
      return StatusState.Skipped;
    default:
      return StatusState.Pending;
  }
};

export const listCommits = async (
  namespace: string,
  name: string,
  branch: string,
  page = 1,
  perPage = 30
): Promise<Commit[]> => {
  const commits: Commit[] = await _fetch(
    `${instance}/api/v1/repos/${namespace}/${name}/commits?branch=${branch}&page=${page}&per_page=${perPage}`,
    {
      headers,
      credentials: 'same-origin',
    }
  )
    .then((response) => response.json())
    .then((commits) => commits.map((c: any) => mapDataToCommit(c)));

  return commits;
};

export const getCommit = async (
  namespace: string,
  name: string,
  sha: string
): Promise<Commit> => {
  const response = await _fetch(
    `${instance}/api/v1/repos/${namespace}/${name}/commits/${sha}`,
    {
      headers,
      credentials: 'same-origin',
    }
  );
  if (response.status === StatusCodes.NOT_FOUND) {
    const message = await response.json().then((data) => data.message);
    throw new HttpNotFoundError(message);
  }

  const commit: Commit = await response
    .json()
    .then((c: any) => mapDataToCommit(c));

  return commit;
};

export const listStatuses = async (
  namespace: string,
  name: string,
  sha: string
): Promise<{ state: StatusState; statuses: Status[] }> => {
  const response = await _fetch(
    `${instance}/api/v1/repos/${namespace}/${name}/commits/${sha}/statuses`,
    {
      headers,
      credentials: 'same-origin',
    }
  );
  if (response.status === StatusCodes.NOT_FOUND) {
    const message = await response.json().then((data) => data.message);
    throw new HttpNotFoundError(message);
  }

  const result = await response.json().then((d) => {
    let state: StatusState;
    const statuses: Status[] = d.statuses.map((status: any) =>
      mapDataToStatus(status)
    );

    if (statuses.length === 0) {
      state = StatusState.Null;
    } else {
      state = mapStatusState(d.state);
    }

    return {
      state,
      statuses,
    };
  });

  return result;
};
