import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit'
import { message } from "antd"

import { searchRepo, activateRepo } from "../apis/repo"
import { Repo, RequestStatus } from "../models"
import { HttpForbiddenError } from '../models/errors'

interface RepoState {
    key: string
    repo?: Repo
    activating: RequestStatus
}

const initialState: RepoState = {
    key: "home",
    activating: RequestStatus.Idle
}

export const init = createAsyncThunk<Repo, {namespace: string, name: string}, { state: {repo: RepoState} }>(
    'repo/init', 
    async (params) => {
        const repo = await searchRepo(params.namespace, params.name)
        return repo
    },
)

export const activate = createAsyncThunk<Repo, void, { state: {repo: RepoState} }>(
    'repo/activate', 
    async (_, { getState, rejectWithValue } ) => {
        const { repo } = getState().repo
        if (!repo) throw new Error("There is no repo.")

        try {
            const nr =  await activateRepo(repo)
            return nr
        } catch(e) {
            if (e instanceof HttpForbiddenError) {
                message.error("Only admin permission can activate.", 3)
            } else {
                message.error("It has failed to activate.", 3)
            }
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
                state.activating = RequestStatus.Idle
                state.repo = action.payload
            })
            .addCase(activate.rejected, (state) => {
                state.activating = RequestStatus.Idle
            })
    }
})