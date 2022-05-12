import camelcaseKeys from 'camelcase-keys';
import { instance, headers } from './setting';
import { _fetch } from './_base';

import { License } from '../models';

function mapDataToLicense(data: any): License {
  const license: License = camelcaseKeys(data);

  license.expiredAt = new Date(data.expired_at);

  return license;
}

export const getLicense = async (): Promise<License> => {
  const lic = await _fetch(`${instance}/api/v1/license`, {
    headers,
    credentials: 'same-origin',
  })
    .then((res) => res.json())
    .then((data) => mapDataToLicense(data));

  return lic;
};
