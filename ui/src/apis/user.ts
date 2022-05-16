import camelcaseKeys from 'camelcase-keys';
import { StatusCodes } from 'http-status-codes';

import { _fetch } from './_base';
import { instance, headers } from './setting';
import { User, RateLimit, HttpForbiddenError } from '../models';

export interface UserData {
  id: number;
  login: string;
  avatar: string;
  admin: boolean;
  hash?: string;
  created_at: string;
  updated_at: string;
  edges: {
    chat_user: ChatUserData;
  };
}

interface ChatUserData {
  id: string;
  created_at: string;
  updated_at: string;
}

interface RateLimitData {
  limit: number;
  remaining: number;
  reset: string;
}

export const mapDataToUser = (data: any): User => {
  const user: User = camelcaseKeys(data);

  user.createdAt = new Date(data.created_at);
  user.updatedAt = new Date(data.updated_at);

  if ('chat_user' in data.edges) {
    const { chat_user } = data.edges;

    user.chatUser = {
      id: chat_user.id,
      createdAt: new Date(chat_user.created_at),
      updatedAt: new Date(chat_user.updated_at),
    };
  }

  return user;
};

export const mapDataToRateLimit = (data: any): RateLimit => {
  const rateLimit = camelcaseKeys(data);

  rateLimit.reset = new Date(data.reset);

  return rateLimit;
};

/**
 * Only admin can access to users API.
 */
export const listUsers = async (
  q: string,
  page = 1,
  perPage = 30
): Promise<User[]> => {
  const res = await _fetch(
    `${instance}/api/v1/users?q=${q}&page=${page}&per_page=${perPage}`,
    {
      headers,
      credentials: 'same-origin',
    }
  );
  if (res.status === StatusCodes.FORBIDDEN) {
    const message = await res.json().then((data) => data.message);
    throw new HttpForbiddenError(message);
  }

  const users = await res
    .json()
    .then((data: UserData[]) => data.map((d) => mapDataToUser(d)));

  return users;
};

export const updateUser = async (
  id: number,
  payload: { admin: boolean }
): Promise<User> => {
  const res = await _fetch(`${instance}/api/v1/users/${id}`, {
    headers,
    credentials: 'same-origin',
    method: 'PATCH',
    body: JSON.stringify(payload),
  });
  if (res.status === StatusCodes.FORBIDDEN) {
    const message = await res.json().then((data) => data.message);
    throw new HttpForbiddenError(message);
  }

  const user: User = await res
    .json()
    .then((data: UserData) => mapDataToUser(data));

  return user;
};

export const deleteUser = async (id: number): Promise<void> => {
  const res = await _fetch(`${instance}/api/v1/users/${id}`, {
    headers,
    credentials: 'same-origin',
    method: 'DELETE',
  });
  if (res.status === StatusCodes.FORBIDDEN) {
    const message = await res.json().then((data) => data.message);
    throw new HttpForbiddenError(message);
  }
};

export const getMe = async (): Promise<User> => {
  const user: User = await _fetch(`${instance}/api/v1/user`, {
    headers,
    credentials: 'same-origin',
  })
    .then((response) => response.json())
    .then((data) => mapDataToUser(data));

  return user;
};

export const getRateLimit = async (): Promise<RateLimit> => {
  const rateLimit: RateLimit = await _fetch(
    `${instance}/api/v1/user/rate-limit`,
    {
      headers,
      credentials: 'same-origin',
    }
  )
    .then((response) => response.json())
    .then((data) => mapDataToRateLimit(data));

  return rateLimit;
};
