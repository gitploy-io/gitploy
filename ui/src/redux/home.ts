import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit'

import { Repo, RequestStatus } from '../models'
import * as apis from '../apis'

export const perPage = 30

interface HomeState {
    loading: boolean 
    q: string
    repos: Repo[]
    page: number
    syncId: string
    syncing: RequestStatus
}

const initialState: HomeState = {
    loading: true,
    q: "",
    repos: [],
    page: 1,
    syncId: "",
    syncing: RequestStatus.Idle
}

export const listRepos = createAsyncThunk<Repo[], void, { state: {home: HomeState} }>(
    'home/listRepos', 
    async (_, { getState, rejectWithValue }) => {
        const {q, page } = getState().home
        try {
            const repos = await apis.listRepos(q, page, perPage)
            return repos
        } catch(e) {
            return rejectWithValue(e)
        }
    },
)

export const sync = createAsyncThunk<void, void, {state: {home: HomeState}}>(
    "home/sync",
    async (_, { getState, rejectWithValue, requestId}) => {
        const { syncId, syncing } = getState().home
        if (!(syncing === RequestStatus.Pending && syncId === requestId)) {
            return
        }

        try {
            await apis.sync()
        } catch(e) {
            return rejectWithValue(e)
        }
    }
)

export const homeSlice = createSlice({
    name: 'home',
    initialState,
    reducers: {
        setQ: (state, action: PayloadAction<string>) => {
            state.q = action.payload
        },
        increasePage: (state) => {
            state.page = state.page + 1
        },
        decreasePage: (state) => {
            state.page = state.page - 1
        }
    },
    extraReducers: builder => {
        builder
            .addCase(listRepos.fulfilled, (state, action) => {
                state.repos = action.payload
                state.loading = false
            })
            .addCase(sync.pending, (state, action) => {
                if (state.syncing === RequestStatus.Idle) {
                    state.syncId = action.meta.requestId
                    state.syncing = RequestStatus.Pending
                }
            })
            .addCase(sync.fulfilled, (state) => {
                if (state.syncing === RequestStatus.Pending) {
                    state.syncId = ""
                    state.syncing = RequestStatus.Idle
                }
            })
            .addCase(sync.rejected, (state) => {
                if (state.syncing === RequestStatus.Pending) {
                    state.syncId = ""
                    state.syncing = RequestStatus.Idle
                }
            })
    }
})
