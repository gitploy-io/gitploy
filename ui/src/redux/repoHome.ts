import { createSlice, PayloadAction, createAsyncThunk } from '@reduxjs/toolkit'

import { Repo, Deployment } from '../models'
import { searchRepo, listDeployments } from '../apis'

export const perPage = 20;

interface RepoHomeState {
    loading: boolean
    repo: Repo | null
    env: "",
    deployments: Deployment[]
    page: number
}

const initialState: RepoHomeState = {
    loading: true,
    repo: null,
    env: "",
    deployments: [],
    page: 1,
}

export const init = createAsyncThunk<Repo, {namespace: string, name: string}, { state: {repoHome: RepoHomeState} }>(
    'repoHome/init', 
    async (params, _ ) => {
        const repo = await searchRepo(params.namespace, params.name)
        return repo
    },
)

export const fetchDeployments = createAsyncThunk<Deployment[], void, { state: {repoHome: RepoHomeState} }>(
    'repoHome/fetchDeployments', 
    async (_, { getState, rejectWithValue } ) => {
        const { repo, env, page} = getState().repoHome

        if (repo === null) {
            rejectWithValue(new Error("repo doesn't exist."))
            return []
        }

        const deployments = await listDeployments(repo.id, env, page, perPage)
        return deployments
    },
)

export const repoHomeSlice = createSlice({
    name: "repoHome",
    initialState,
    reducers: {
        increasePage: (state) => {
            state.page = state.page + 1
        },
        decreasePage: (state) => {
            state.page = state.page - 1
        }
    },
    extraReducers: builder => {
        builder
            .addCase(init.fulfilled, (state, action) => {
                state.repo = action.payload
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