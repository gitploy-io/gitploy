import camelcaseKeys from 'camelcase-keys';
import { StatusCodes } from 'http-status-codes';

import { instance, headers } from './setting';
import { _fetch } from './_base';
import { Config, HttpNotFoundError } from '../models';

const mapDataToConfig = (data: any): Config => {
  return camelcaseKeys(data, { deep: true });
};

export const getConfig = async (
  namespace: string,
  name: string
): Promise<Config> => {
  const response = await _fetch(
    `${instance}/api/v1/repos/${namespace}/${name}/config`,
    {
      headers,
      credentials: 'same-origin',
    }
  );
  if (response.status === StatusCodes.NOT_FOUND) {
    const message = await response.json().then((data) => data.message);
    throw new HttpNotFoundError(message);
  }

  const conf = await response.json().then((c) => mapDataToConfig(c));

  return conf;
};
