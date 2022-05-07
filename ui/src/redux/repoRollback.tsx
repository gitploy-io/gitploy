import { createSlice, PayloadAction, createAsyncThunk } from '@reduxjs/toolkit';
import { message } from 'antd';

import {
  getConfig,
  listDeployments,
  rollbackDeployment,
  listPerms,
  getMe,
} from '../apis';
import {
  User,
  Deployment,
  DeploymentStatusEnum,
  Config,
  Env,
  RequestStatus,
  HttpForbiddenError,
  HttpUnprocessableEntityError,
  HttpConflictError,
} from '../models';

const page = 1;
const perPage = 100;

interface RepoRollbackState {
  display: boolean;
  namespace: string;
  name: string;
  config?: Config;
  env?: Env;
  envs: Env[];
  deployment?: Deployment;
  deployments: Deployment[];
  deployId: string;
  deploying: RequestStatus;
}

const initialState: RepoRollbackState = {
  display: false,
  namespace: '',
  name: '',
  envs: [],
  deployments: [],
  deployId: '',
  deploying: RequestStatus.Idle,
};

export const fetchConfig = createAsyncThunk<
  Config,
  void,
  { state: { repoRollback: RepoRollbackState } }
>('repoRollback/fetchEnvs', async (_, { getState, rejectWithValue }) => {
  const { namespace, name } = getState().repoRollback;

  try {
    const config = await getConfig(namespace, name);
    return config;
  } catch (e) {
    return rejectWithValue(e);
  }
});

export const fetchDeployments = createAsyncThunk<
  Deployment[],
  void,
  { state: { repoRollback: RepoRollbackState } }
>('repoRollback/fetchDeployments', async (_, { getState }) => {
  const { namespace, name, env } = getState().repoRollback;

  if (!env) {
    throw new Error('The env is not selected.');
  }

  // Return the deployment history except the latest one.
  const deployments = await listDeployments(
    namespace,
    name,
    env.name,
    DeploymentStatusEnum.Success,
    page,
    perPage
  );
  return deployments.slice(1);
});

export const searchCandidates = createAsyncThunk<
  User[],
  string,
  { state: { repoRollback: RepoRollbackState } }
>('repoRollback/searchCandidates', async (q, { getState, rejectWithValue }) => {
  const { namespace, name } = getState().repoRollback;

  try {
    const perms = await listPerms(namespace, name, q);
    const candidates = perms.map((p) => {
      return p.user;
    });
    return candidates;
  } catch (e) {
    return rejectWithValue(e);
  }
});

export const fetchUser = createAsyncThunk<
  User,
  void,
  { state: { repoRollback: RepoRollbackState } }
>('repoRollback/fetchUser', async (_, { rejectWithValue }) => {
  try {
    const user = await getMe();
    return user;
  } catch (e) {
    return rejectWithValue(e);
  }
});

export const rollback = createAsyncThunk<
  void,
  void,
  { state: { repoRollback: RepoRollbackState } }
>(
  'repoRollback/deploy',
  async (_, { getState, rejectWithValue, requestId }) => {
    const { namespace, name, deployment, env, deployId, deploying } =
      getState().repoRollback;
    if (!deployment) {
      throw new Error('The deployment is undefined.');
    }
    if (!(deploying === RequestStatus.Pending && requestId === deployId)) {
      return;
    }

    try {
      const rollback = await rollbackDeployment(
        namespace,
        name,
        deployment.number
      );

      if (!env?.review?.enabled) {
        const msg = (
          <span>
            Starts to rollback.{' '}
            <a href={`/${namespace}/${name}/deployments/${rollback.number}`}>
              #{rollback.number}
            </a>
          </span>
        );
        message.success(msg, 3);
        return;
      }

      const msg = (
        <span>
          Request a review to reviewers{' '}
          <a href={`/${namespace}/${name}/deployments/${rollback.number}`}>
            #{rollback.number}
          </a>
        </span>
      );
      message.success(msg, 3);
    } catch (e) {
      if (e instanceof HttpForbiddenError) {
        message.warn('Only write permission can deploy.', 3);
      } else if (e instanceof HttpUnprocessableEntityError) {
        const msg = (
          <span>
            <span>It is unprocesable entity.</span>
            <br />
            <span className="gitploy-quote">{e.message}</span>
          </span>
        );
        message.error(msg, 3);
      } else if (e instanceof HttpConflictError) {
        message.error('It has conflicted, please retry it.', 3);
      }
      return rejectWithValue(e);
    }
  }
);

export const repoRollbackSlice = createSlice({
  name: 'repoRollback',
  initialState,
  reducers: {
    init: (
      state,
      action: PayloadAction<{ namespace: string; name: string }>
    ) => {
      state.namespace = action.payload.namespace;
      state.name = action.payload.name;
    },
    setDisplay: (state, action: PayloadAction<boolean>) => {
      state.display = action.payload;
    },
    setEnv: (state, action: PayloadAction<Env>) => {
      state.env = action.payload;
    },
    setDeployment: (state, action: PayloadAction<Deployment>) => {
      state.deployment = action.payload;
    },
  },
  extraReducers: (builder) => {
    builder
      .addCase(fetchConfig.fulfilled, (state, action) => {
        const config = action.payload;
        state.envs = config.envs.map((e) => e);
        state.config = config;
      })
      .addCase(fetchDeployments.fulfilled, (state, action) => {
        state.deployments = action.payload;
      })
      .addCase(rollback.pending, (state, action) => {
        if (state.deploying === RequestStatus.Idle) {
          state.deploying = RequestStatus.Pending;
          state.deployId = action.meta.requestId;
        }
      })
      .addCase(rollback.fulfilled, (state, action) => {
        if (
          state.deploying === RequestStatus.Pending &&
          state.deployId === action.meta.requestId
        ) {
          state.deploying = RequestStatus.Idle;
        }
      })
      .addCase(rollback.rejected, (state, action) => {
        if (
          state.deploying === RequestStatus.Pending &&
          state.deployId === action.meta.requestId
        ) {
          state.deploying = RequestStatus.Idle;
        }
      });
  },
});
