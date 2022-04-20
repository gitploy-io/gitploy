import { createSlice, PayloadAction, createAsyncThunk } from '@reduxjs/toolkit'

import { Deployment, Event, EventTypeEnum } from '../models'
import { listDeployments, getConfig } from '../apis'

export const perPage = 20;

interface RepoHomeState {
    loading: boolean
    namespace: string
    name: string
    envs: string[]
    env: string
    deployments: Deployment[]
    page: number
}

const initialState: RepoHomeState = {
    loading: true,
    namespace: "",
    name: "",
    envs: [],
    env: "",
    deployments: [],
    page: 1,
}

export const fetchEnvs = createAsyncThunk<string[], void, { state: {repoHome: RepoHomeState} }>(
    "repoHome/fetchEnvs", 
    async (_, { getState } ) => {
        const { namespace, name } = getState().repoHome

        const config = await getConfig(namespace, name)
        return config.envs.map(e =>  e.name)
    },
)

export const fetchDeployments = createAsyncThunk<Deployment[], void, { state: {repoHome: RepoHomeState} }>(
    'repoHome/fetchDeployments', 
    async (_, { getState } ) => {
        const { namespace, name, env, page} = getState().repoHome

        const deployments = await listDeployments(namespace, name, env, "",page, perPage)
        return deployments
    },
)

export const repoHomeSlice = createSlice({
    name: "repoHome",
    initialState,
    reducers: {
        init: (state, action: PayloadAction<{namespace: string, name: string}>) => {
            state.namespace = action.payload.namespace
            state.name = action.payload.name
        },
        setEnv: (state, action: PayloadAction<string>) => {
            state.env = action.payload
        },
        increasePage: (state) => {
            state.page = state.page + 1
        },
        decreasePage: (state) => {
            state.page = state.page - 1
        },
    },
    extraReducers: builder => {
        builder
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