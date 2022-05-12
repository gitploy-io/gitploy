import camelcaseKeys from 'camelcase-keys';
import { StatusCodes } from 'http-status-codes';

import { instance, headers } from './setting';
import { _fetch } from './_base';
import { Branch, HttpNotFoundError } from '../models';

const mapDataToBranch = (data: any): Branch => {
  const branch = camelcaseKeys(data, { deep: true });
  return branch;
};

export const listBranches = async (
  namespace: string,
  name: string,
  page = 1,
  perPage = 30
): Promise<Branch[]> => {
  const branches: Branch[] = await _fetch(
    `${instance}/api/v1/repos/${namespace}/${name}/branches?page=${page}&per_page=${perPage}`,
    {
      headers,
      credentials: 'same-origin',
    }
  )
    .then((response) => response.json())
    .then((branches) => branches.map((b: any) => mapDataToBranch(b)));

  return branches;
};

export const getBranch = async (
  namespace: string,
  name: string,
  branch: string
): Promise<Branch> => {
  const response = await _fetch(
    `${instance}/api/v1/repos/${namespace}/${name}/branches/${branch}`,
    {
      headers,
      credentials: 'same-origin',
    }
  );
  if (response.status === StatusCodes.NOT_FOUND) {
    const message = await response.json().then((data) => data.message);
    throw new HttpNotFoundError(message);
  }

  const ret: Branch = await response
    .json()
    .then((b: any) => mapDataToBranch(b));

  return ret;
};
