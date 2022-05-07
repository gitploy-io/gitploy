import User from './User';
import Repo from './Repo';

export default interface Deployment {
  id: number;
  number: number;
  type: DeploymentType;
  ref: string;
  sha: string;
  env: string;
  status: DeploymentStatusEnum;
  uid: number;
  isRollback: boolean;
  createdAt: Date;
  updatedAt: Date;
  deployer?: User;
  repo?: Repo;
  statuses?: DeploymentStatus[];
}

export enum DeploymentType {
  Commit = 'commit',
  Branch = 'branch',
  Tag = 'tag',
}

export enum DeploymentStatusEnum {
  Waiting = 'waiting',
  Created = 'created',
  Queued = 'queued',
  Running = 'running',
  Success = 'success',
  Failure = 'failure',
  Canceled = 'canceled',
}

export interface DeploymentStatus {
  id: number;
  status: string;
  description: string;
  logUrl: string;
  createdAt: string;
  updatedAt: string;
  deploymentId: number;
  repoId: number;
  edges?: {
    deployment?: Deployment;
    repo?: Repo;
  };
}
