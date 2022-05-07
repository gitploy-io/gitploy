import { createSlice, PayloadAction, createAsyncThunk } from '@reduxjs/toolkit';
import { message } from 'antd';

import { Config, Lock, HttpForbiddenError, HttpNotFoundError } from '../models';
import {
  getConfig,
  listLocks as _listLocks,
  lock as _lock,
  unlock as _unlock,
  updateLock,
} from '../apis';

interface RepoLockState {
  display: boolean;
  namespace: string;
  name: string;
  config?: Config;
  locks: Lock[];
}

const initialState: RepoLockState = {
  display: false,
  namespace: '',
  name: '',
  locks: [],
};

export const fetchConfig = createAsyncThunk<
  Config,
  void,
  { state: { repoLock: RepoLockState } }
>('repoLock/fetchConfig', async (_, { getState, rejectWithValue }) => {
  const { namespace, name } = getState().repoLock;

  try {
    const config = await getConfig(namespace, name);
    return config;
  } catch (e) {
    return rejectWithValue(e);
  }
});

export const listLocks = createAsyncThunk<
  Lock[],
  void,
  { state: { repoLock: RepoLockState } }
>('repoLock/listLocks', async (_, { getState, rejectWithValue }) => {
  const { namespace, name } = getState().repoLock;

  try {
    const locks = await _listLocks(namespace, name);
    return locks;
  } catch (e) {
    return rejectWithValue(e);
  }
});

export const lock = createAsyncThunk<
  Lock,
  string,
  { state: { repoLock: RepoLockState } }
>('repoLock/lock', async (env, { getState, rejectWithValue }) => {
  const { namespace, name } = getState().repoLock;

  try {
    const locks = await _lock(namespace, name, env);
    return locks;
  } catch (e) {
    if (e instanceof HttpForbiddenError) {
      message.warn('Only write permission can lock.', 3);
    }
    return rejectWithValue(e);
  }
});

export const unlock = createAsyncThunk<
  Lock,
  string,
  { state: { repoLock: RepoLockState } }
>('repoLock/unlock', async (env, { getState, rejectWithValue }) => {
  const { namespace, name, locks } = getState().repoLock;

  const lock = locks.find((lock) => lock.env === env);
  if (!lock) {
    throw new Error('The env is not found.');
  }

  try {
    await _unlock(namespace, name, lock.id);
    return lock;
  } catch (e) {
    if (e instanceof HttpForbiddenError) {
      message.warn('Only write permission can unlock.', 3);
    }
    return rejectWithValue(e);
  }
});

export const setAutoUnlock = createAsyncThunk<
  Lock,
  { env: string; expiredAt: Date },
  { state: { repoLock: RepoLockState } }
>(
  'repoLock/setAutoUnlock',
  async ({ env, expiredAt }, { getState, rejectWithValue }) => {
    const { namespace, name, locks } = getState().repoLock;

    const lock = locks.find((lock) => lock.env === env);
    if (!lock) {
      throw new Error('The env is not found.');
    }

    try {
      const ret = await updateLock(namespace, name, lock.id, { expiredAt });
      message.info(`Setting auto-unlock.`);

      return ret;
    } catch (e) {
      if (e instanceof HttpForbiddenError) {
        message.warn('Only write permission can enable auto unlock.', 3);
      }

      if (e instanceof HttpNotFoundError) {
        message.warn('Lock is not found.');
      }
      return rejectWithValue(e);
    }
  }
);

export const repoLockSlice = createSlice({
  name: 'repoLock',
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
  },
  extraReducers: (builder) => {
    builder
      .addCase(fetchConfig.fulfilled, (state, action) => {
        state.config = action.payload;
      })
      .addCase(listLocks.fulfilled, (state, action) => {
        state.locks = action.payload;
      })
      .addCase(lock.fulfilled, (state, action) => {
        state.locks.push(action.payload);
      })
      .addCase(unlock.fulfilled, (state, action) => {
        const idx = state.locks.findIndex(
          (lock) => lock.id === action.payload.id
        );

        if (idx !== -1) {
          state.locks.splice(idx, 1);
        }
      })
      .addCase(setAutoUnlock.fulfilled, (state, action) => {
        state.locks = state.locks.map((lock) => {
          return lock.id !== action.payload.id ? lock : action.payload;
        });
      });
  },
});
