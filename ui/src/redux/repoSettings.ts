import { createSlice, PayloadAction, createAsyncThunk } from "@reduxjs/toolkit"

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
    async (payload, { getState, requestId } ) => {
        const { repo, saveId, saving } = getState().repoSettings
        if (repo === null) {
            throw new Error("There is no repo.")
        }

        if (!(saving === RequestStatus.Pending || saveId === requestId)) {
            return repo
        }

        const nr = await updateRepo(repo.id, payload)
        return nr
    },
)

export const deactivate = createAsyncThunk<Repo, void, { state: {repoSettings: RepoSettingsState} }>(
    'repoSettings/deactivate', 
    async (_, { getState, rejectWithValue } ) => {
        const { repo } = getState().repoSettings
        if (repo === null) throw new Error("There is no repo.")

        try {
            return await deactivateRepo(repo.id)
        } catch(e) {
            return rejectWithValue(e)
        }
    },
)

export const repoSettingsSlice = createSlice({
    name: "repoSettings",
    initialState,
    reducers: {
        unsetSaving: (state) => {
            state.saving = RequestStatus.Idle
            state.saveId = ""
        },
        unsetDeactivating: (state) => {
            state.deactivating = RequestStatus.Idle
        }
    },
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
                    state.saving = RequestStatus.Success
                }
            })
            .addCase(save.rejected, (state, action) => {
                if (state.saving === RequestStatus.Pending && state.saveId === action.meta.requestId) {
                    state.saving = RequestStatus.Failure
                }
            })
            .addCase(deactivate.pending, (state) => {
                state.deactivating = RequestStatus.Idle
            })
            .addCase(deactivate.fulfilled, (state) => {
                state.deactivating = RequestStatus.Success
            })
            .addCase(deactivate.rejected, (state, action: PayloadAction<unknown> | PayloadAction<HttpForbiddenError>) => {
                if (action.payload instanceof HttpForbiddenError) {
                    state.deactivating = RequestStatus.Failure
                }

                console.log(action)
                state.deactivating = RequestStatus.Idle
            })
    }
})