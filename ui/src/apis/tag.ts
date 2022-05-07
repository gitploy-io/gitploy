import { StatusCodes } from 'http-status-codes';

import { instance, headers } from './setting';
import { _fetch } from './_base';
import { Tag, HttpNotFoundError } from '../models';

export const listTags = async (
  namespace: string,
  name: string,
  page = 1,
  perPage = 30
): Promise<Tag[]> => {
  const tags: Tag[] = await _fetch(
    `${instance}/api/v1/repos/${namespace}/${name}/tags?page=${page}&per_page=${perPage}`,
    {
      headers,
      credentials: 'same-origin',
    }
  )
    .then((response) => response.json())
    .then((tags) =>
      tags.map((t: any): Tag => {
        return {
          name: t.name,
          commitSha: t.commit_sha,
        };
      })
    );

  return tags;
};

export const getTag = async (
  namespace: string,
  name: string,
  tag: string
): Promise<Tag> => {
  const response = await _fetch(
    `${instance}/api/v1/repos/${namespace}/${name}/tags/${tag}`,
    {
      headers,
      credentials: 'same-origin',
    }
  );
  if (response.status === StatusCodes.NOT_FOUND) {
    const message = await response.json().then((data) => data.message);
    throw new HttpNotFoundError(message);
  }

  const ret: Tag = await response.json().then((t) => ({
    name: t.name,
    commitSha: t.commit_sha,
  }));
  return ret;
};
