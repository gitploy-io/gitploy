import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit'

import { Repo, RequestStatus, Deployment } from '../models'
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
        setFirstPage: (state) => {
            state.page = 1
        },
        increasePage: (state) => {
            state.page = state.page + 1
        },
        decreasePage: (state) => {
            state.page = state.page - 1
        },
        handleDeploymentEvent: (state, action: PayloadAction<Deployment>) => {
            const event = action.payload
            
            state.repos = state.repos.map((repo) => {
                if (event.repo?.id !== repo.id) {
                    return repo
                }

                if (!repo.deployments) {
                    repo.deployments = []
                }

                // Unshift a new deployment when the event is create.
                if (event.createdAt.getTime() === event.updatedAt.getTime()) {
                    repo.deployments.unshift(event)
                    return repo
                }

                repo.deployments = repo.deployments.map((deployment) => {
                    return (deployment.id === event.id)? event : deployment
                })
                return repo
            })
        },
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
