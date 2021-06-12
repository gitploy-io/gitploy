import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit'

import { searchRepo, activateRepo } from "../apis/repos"
import { Repo, RequestStatus } from "../models"
import { HttpForbiddenError } from '../models/errors'

interface RepoState {
    key: string
    repo: Repo | null
    activating: RequestStatus
}

const initialState: RepoState = {
    key: "home",
    repo: null,
    activating: RequestStatus.Idle
}

export const init = createAsyncThunk<Repo, {namespace: string, name: string}, { state: {repo: RepoState} }>(
    'repo/init', 
    async (params, _ ) => {
        const repo = await searchRepo(params.namespace, params.name)
        return repo
    },
)

export const activate = createAsyncThunk<Repo, void, { state: {repo: RepoState} }>(
    'repo/activate', 
    async (_, { getState, rejectWithValue } ) => {
        const { repo } = getState().repo
        if (repo === null) throw new Error("There is no repo.")

        try {
            return await activateRepo(repo.id)
        } catch(e) {
            return rejectWithValue(e)
        }
    },
)

export const repoSlice = createSlice({
    name: "repo",
    initialState,
    reducers: {
        setKey: (state, action: PayloadAction<string>) => {
            state.key = action.payload
        },
        unsetActivating: (state) => {
            state.activating = RequestStatus.Idle
        }
    },
    extraReducers: builder => {
        builder
            .addCase(init.fulfilled, (state, action) => {
                const repo = action.payload
                state.repo = repo
            })
            .addCase(activate.pending, (state) => {
                state.activating = RequestStatus.Pending
            })
            .addCase(activate.fulfilled, (state, action) => {
                state.activating = RequestStatus.Success
                state.repo = action.payload
            })
            .addCase(activate.rejected, (state, action: PayloadAction<unknown> | PayloadAction<HttpForbiddenError>) => {
                if (action.payload instanceof HttpForbiddenError) {
                    state.activating = RequestStatus.Failure
                }

                console.log(action)
                state.activating = RequestStatus.Idle
            })
    }
})