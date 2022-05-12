import camelcaseKeys from 'camelcase-keys';
import { instance, headers } from './setting';
import { _fetch } from './_base';
import { mapDataToUser } from './user';
import { mapDataToRepo } from './repo';
import { Perm } from '../models';

const mapDataToPerm = (data: any): Perm => {
  const perm: Perm = camelcaseKeys(data);

  perm.syncedAt = new Date(data.synced_at);
  perm.createdAt = new Date(data.created_at);
  perm.updatedAt = new Date(data.updated_at);

  // Edges
  perm.user = mapDataToUser(data.edges.user);
  perm.repo = mapDataToRepo(data.edges.repo);

  return perm;
};

export const listPerms = async (
  namespace: string,
  name: string,
  q: string,
  page = 1,
  perPage = 30
): Promise<Perm[]> => {
  const perms: Perm[] = await _fetch(
    `${instance}/api/v1/repos/${namespace}/${name}/perms?q=${q}&page=${page}&per_page=${perPage}`,
    {
      headers,
      credentials: 'same-origin',
    }
  )
    .then((res) => res.json())
    .then((data) => data.map((d: any) => mapDataToPerm(d)));

  return perms;
};
