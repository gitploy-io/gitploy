import camelcaseKeys from 'camelcase-keys';
import { StatusCodes } from 'http-status-codes';

import { instance, headers } from './setting';
import { _fetch } from './_base';
import { mapDataToUser } from './user';
import { mapDataToDeployment } from './deployment';
import { Review, HttpNotFoundError } from '../models';

export const mapDataToReview = (data: any): Review => {
  const review: Review = camelcaseKeys(data);

  review.createdAt = new Date(data.created_at);
  review.updatedAt = new Date(data.updated_at);

  if ('user' in data.edges) {
    review.user = mapDataToUser(data.edges.user);
  }

  if ('deployment' in data.edges) {
    review.deployment = mapDataToDeployment(data.edges.deployment);
  }

  return review;
};

export const searchReviews = async (): Promise<Review[]> => {
  const reviews: Review[] = await _fetch(`${instance}/api/v1/search/reviews`, {
    credentials: 'same-origin',
    headers,
  })
    .then((res) => res.json())
    .then((data) => data.map((d: any): Review => mapDataToReview(d)));

  return reviews;
};

export const listReviews = async (
  namespace: string,
  name: string,
  number: number
): Promise<Review[]> => {
  const res = await _fetch(
    `${instance}/api/v1/repos/${namespace}/${name}/deployments/${number}/reviews`,
    {
      credentials: 'same-origin',
      headers,
    }
  );

  const reviews: Review[] = await res
    .json()
    .then((data) => data.map((d: any): Review => mapDataToReview(d)));

  return reviews;
};

export const getUserReview = async (
  namespace: string,
  name: string,
  number: number
): Promise<Review> => {
  const res = await _fetch(
    `${instance}/api/v1/repos/${namespace}/${name}/deployments/${number}/review`,
    {
      credentials: 'same-origin',
      headers,
    }
  );

  if (res.status === StatusCodes.NOT_FOUND) {
    const { message } = await res.json();
    throw new HttpNotFoundError(message);
  }

  const review = await res.json().then((data) => mapDataToReview(data));
  return review;
};

export const approveReview = async (
  namespace: string,
  name: string,
  number: number,
  comment?: string
): Promise<Review> => {
  const body = {
    status: 'approved',
    comment,
  };
  const res = await _fetch(
    `${instance}/api/v1/repos/${namespace}/${name}/deployments/${number}/review`,
    {
      credentials: 'same-origin',
      headers,
      method: 'PATCH',
      body: JSON.stringify(body),
    }
  );

  if (res.status === StatusCodes.NOT_FOUND) {
    const { message } = await res.json();
    throw new HttpNotFoundError(message);
  }

  const review = await res.json().then((data) => mapDataToReview(data));
  return review;
};

export const rejectReview = async (
  namespace: string,
  name: string,
  number: number,
  comment?: string
): Promise<Review> => {
  const body = {
    status: 'rejected',
    comment,
  };
  const res = await _fetch(
    `${instance}/api/v1/repos/${namespace}/${name}/deployments/${number}/review`,
    {
      credentials: 'same-origin',
      headers,
      method: 'PATCH',
      body: JSON.stringify(body),
    }
  );

  if (res.status === StatusCodes.NOT_FOUND) {
    const { message } = await res.json();
    throw new HttpNotFoundError(message);
  }

  const review = await res.json().then((data) => mapDataToReview(data));
  return review;
};
