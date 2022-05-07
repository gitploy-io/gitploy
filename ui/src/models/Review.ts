import User from './User';
import Deployment from './Deployment';

export interface Review {
  id: number;
  status: ReviewStatusEnum;
  comment: string;
  createdAt: Date;
  updatedAt: Date;
  user?: User;
  deployment?: Deployment;
}

export enum ReviewStatusEnum {
  Pending = 'pending',
  Approved = 'approved',
  Rejected = 'rejected',
}
