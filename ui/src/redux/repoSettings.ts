import { createSlice, PayloadAction, createAsyncThunk } from "@reduxjs/toolkit"
import { message } from "antd"

import { searchRepo, updateRepo, deactivateRepo } from "../apis"
import { Repo, RepoPayload, RequestStatus, HttpForbiddenError  } from "../models"

interface RepoSettingsState {
    repo: Repo | null
    saveId: string,
    saving: RequestStatus
    deactivating: RequestStatus
}

const initialState: RepoSettingsState = {
    repo: null,
    saveId: "",
    saving: RequestStatus.Idle,
    deactivating: RequestStatus.Idle,
}

export const init = createAsyncThunk<Repo, {namespace: string, name: string}, { state: {repoSettings: RepoSettingsState} }>(
    'repoSettings/init', 
    async (params, _ ) => {
        const repo = await searchRepo(params.namespace, params.name)
        return repo
    },
)

export const save = createAsyncThunk<Repo, RepoPayload, { state: {repoSettings: RepoSettingsState} }>(
    'repoSettings/save', 
    async (payload, { getState, rejectWithValue, requestId } ) => {
        const { repo, saveId, saving } = getState().repoSettings
        if (repo === null) {
            throw new Error("There is no repo.")
        }

        if (!(saving === RequestStatus.Pending || saveId === requestId)) {
            return repo
        }

        try {
            const nr = await updateRepo(repo.id, payload)
            message.success("Success to save.", 3)
            return nr
        } catch(e) {
            message.error("It has failed to save.", 3)
            return rejectWithValue(e)
        }
    },
)

export const deactivate = createAsyncThunk<Repo, void, { state: {repoSettings: RepoSettingsState} }>(
    'repoSettings/deactivate', 
    async (_, { getState, rejectWithValue } ) => {
        const { repo } = getState().repoSettings
        if (repo === null) throw new Error("There is no repo.")

        try {
            const nr = await deactivateRepo(repo.id)
            window.location.reload()
            return nr
        } catch(e) {
            if (e instanceof HttpForbiddenError) {
                message.error("Only admin permission can deactivate.", 3)
            } else {
                message.error("It has failed to save.", 3)
            }
            return rejectWithValue(e)
        }
    },
)

export const repoSettingsSlice = createSlice({
    name: "repoSettings",
    initialState,
    reducers: {},
    extraReducers: builder => {
        builder
            .addCase(init.fulfilled, (state, action) => {
                state.repo = action.payload
            })
            .addCase(save.pending, (state, action) => {
                if (state.saving === RequestStatus.Idle) {
                    state.saving = RequestStatus.Pending
                    state.saveId = action.meta.requestId
                }
            })
            .addCase(save.fulfilled, (state, action) => {
                if (state.saving === RequestStatus.Pending && state.saveId === action.meta.requestId) {
                    state.repo = action.payload
                    state.saving = RequestStatus.Idle
                }
            })
            .addCase(save.rejected, (state, action) => {
                if (state.saving === RequestStatus.Pending && state.saveId === action.meta.requestId) {
                    state.saving = RequestStatus.Idle
                }
            })
            .addCase(deactivate.pending, (state) => {
                state.deactivating = RequestStatus.Pending
            })
            .addCase(deactivate.fulfilled, (state) => {
                state.deactivating = RequestStatus.Idle
            })
            .addCase(deactivate.rejected, (state, action: PayloadAction<unknown> | PayloadAction<HttpForbiddenError>) => {
                state.deactivating = RequestStatus.Idle
            })
    }
})