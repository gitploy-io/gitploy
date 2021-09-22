import { createSlice, PayloadAction, createAsyncThunk } from '@reduxjs/toolkit'

import { Repo, Deployment, Event, EventTypeEnum } from '../models'
import { searchRepo, listDeployments, getConfig } from '../apis'

export const perPage = 20;

interface RepoHomeState {
    loading: boolean
    repo?: Repo
    envs: string[]
    env: string
    deployments: Deployment[]
    page: number
}

const initialState: RepoHomeState = {
    loading: true,
    envs: [],
    env: "",
    deployments: [],
    page: 1,
}

export const init = createAsyncThunk<Repo, {namespace: string, name: string}, { state: {repoHome: RepoHomeState} }>(
    'repoHome/init', 
    async (params) => {
        const repo = await searchRepo(params.namespace, params.name)
        return repo
    },
)

export const fetchEnvs = createAsyncThunk<string[], void, { state: {repoHome: RepoHomeState} }>(
    "repoHome/fetchEnvs", 
    async (_, { getState, rejectWithValue } ) => {
        const { repo, } = getState().repoHome

        if (!repo) {
            rejectWithValue(new Error("repo doesn't exist."))
            return []
        }

        const config = await getConfig(repo.id)
        return config.envs.map(e =>  e.name)
    },
)

export const fetchDeployments = createAsyncThunk<Deployment[], void, { state: {repoHome: RepoHomeState} }>(
    'repoHome/fetchDeployments', 
    async (_, { getState, rejectWithValue } ) => {
        const { repo, env, page} = getState().repoHome

        if (!repo) {
            rejectWithValue(new Error("repo doesn't exist."))
            return []
        }

        const deployments = await listDeployments(repo.id, env, "",page, perPage)
        return deployments
    },
)

export const repoHomeSlice = createSlice({
    name: "repoHome",
    initialState,
    reducers: {
        setEnv: (state, action: PayloadAction<string>) => {
            state.env = action.payload
        },
        increasePage: (state) => {
            state.page = state.page + 1
        },
        decreasePage: (state) => {
            state.page = state.page - 1
        },
        handleDeploymentEvent: (state, action: PayloadAction<Event>) => {
            const event = action.payload

            if (event.deployment?.repo?.id !== state.repo?.id) {
                return
            }

            if (!(state.env === "" || event.deployment?.env === state.env)) {
                return
            }

            if (event.type === EventTypeEnum.Created && event.deployment) {
                state.deployments.unshift(event.deployment)
                return
            }

            state.deployments = state.deployments.map((deployment) => {
                return (event.deployment?.id ===  deployment.id)? event.deployment : deployment
            })
        },
    },
    extraReducers: builder => {
        builder
            .addCase(init.fulfilled, (state, action) => {
                state.repo = action.payload
            })

            .addCase(fetchEnvs.fulfilled, (state, action) => {
                state.envs = action.payload
            })
            
            .addCase(fetchDeployments.pending, (state) => {
                state.loading = true
            })

            .addCase(fetchDeployments.fulfilled, (state, action) => {
                state.deployments = action.payload
                state.loading = false
            })
    }
})