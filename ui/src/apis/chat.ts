import { _fetch } from './_base';
import { instance, headers } from './setting';
import { StatusCodes } from 'http-status-codes';

export const checkSlack = async (): Promise<boolean> => {
  const res = await _fetch(`${instance}/slack`, {
    headers,
    credentials: 'same-origin',
    method: 'HEAD',
  });
  if (res.status === StatusCodes.NOT_FOUND) {
    return false;
  }

  return true;
};
