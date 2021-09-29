import { createSlice, PayloadAction, createAsyncThunk } from '@reduxjs/toolkit'
import { message } from "antd"

import {
    Repo,
    Config,
    Lock,
    HttpForbiddenError
} from "../models"
import {
    searchRepo,
    getConfig,
    listLocks as _listLocks,
    lock as _lock,
    unlock as _unlock
} from "../apis"

interface RepoLockState {
    display: boolean
    repo?: Repo
    config?: Config
    locks: Lock[]
}

const initialState: RepoLockState = {
    display: false,
    locks: [],
}

export const init = createAsyncThunk<Repo, {namespace: string, name: string}, { state: {repoLock: RepoLockState} }>(
    'repoLock/init', 
    async (params) => {
        const repo = await searchRepo(params.namespace, params.name)
        return repo
    },
)

export const fetchConfig = createAsyncThunk<Config, void, { state: {repoLock: RepoLockState} }>(
    "repoLock/fetchConfig", 
    async (_, { getState, rejectWithValue } ) => {
        const { repo } = getState().repoLock
        if (!repo) {
            throw new Error("The repo is undefined.")
        }

        try {
            const config = await getConfig(repo.id)
            return config
        } catch (e) {
            return rejectWithValue(e)
        }
    },
)

export const listLocks = createAsyncThunk<Lock[], void, { state: {repoLock: RepoLockState} }>(
    'repoLock/listLocks', 
    async (_, { getState, rejectWithValue }) => {
        const { repo } = getState().repoLock
        if (!repo) {
            throw new Error("The repo is undefined.")
        }

        try {
            const locks = await _listLocks(repo)
            return locks
        } catch (e) {
            return rejectWithValue(e)
        }
    },
)

export const lock = createAsyncThunk<Lock, string, { state: {repoLock: RepoLockState} }>(
    'repoLock/lock', 
    async (env, { getState, rejectWithValue }) => {
        const { repo } = getState().repoLock
        if (!repo) {
            throw new Error("The repo is undefined.")
        }

        try {
            const locks = await _lock(repo, env)
            return locks
        } catch (e) {
            if (e instanceof HttpForbiddenError) {
                message.error("Only write permission can lock.", 3)
            }
            return rejectWithValue(e)
        }
    },
)

export const unlock = createAsyncThunk<Lock, string, { state: {repoLock: RepoLockState} }>(
    'repoLock/unlock', 
    async (env, { getState, rejectWithValue }) => {
        const { repo, locks } = getState().repoLock
        if (!repo) {
            throw new Error("The repo is undefined.")
        }

        const lock = locks.find((lock) => lock.env === env)
        if (!lock) {
            throw new Error("The env is not found.")
        }

        try {
            await _unlock(repo, lock)
            return lock
        } catch (e) {
            if (e instanceof HttpForbiddenError) {
                message.error("Only write permission can unlock.", 3)
            }
            return rejectWithValue(e)
        }
    },
)

export const repoLockSlice = createSlice({
    name: "repoLock",
    initialState,
    reducers: {
        setDisplay: (state, action: PayloadAction<boolean>) => {
            state.display = action.payload
        },
    },
    extraReducers: builder => {
        builder
            .addCase(init.fulfilled, (state, action) => {
                state.repo = action.payload
            })

            .addCase(fetchConfig.fulfilled, (state, action) => {
                state.config = action.payload
            })

            .addCase(listLocks.fulfilled, (state, action) => {
                state.locks = action.payload
            })

            .addCase(lock.fulfilled, (state, action) => {
                state.locks.push(action.payload)
            })

            .addCase(unlock.fulfilled, (state, action) => {
                const idx = state.locks.findIndex((lock) => lock.id === action.payload.id)

                if (idx !== -1) {
                    state.locks.splice(idx, 1)
                }
            })
    }
})