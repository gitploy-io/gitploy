import { createSlice, createAsyncThunk, PayloadAction } from "@reduxjs/toolkit"
import { message } from "antd"

import { getRepo, updateRepo, deactivateRepo } from "../apis"
import { Repo, RequestStatus, HttpForbiddenError  } from "../models"

interface RepoSettingsState {
    repo?: Repo
    saveId: string,
    saving: RequestStatus
    deactivating: RequestStatus
}

const initialState: RepoSettingsState = {
    saveId: "",
    saving: RequestStatus.Idle,
    deactivating: RequestStatus.Idle,
}

export const init = createAsyncThunk<Repo, {namespace: string, name: string}, { state: {repoSettings: RepoSettingsState} }>(
    'repoSettings/init', 
    async (params) => {
        const repo = await getRepo(params.namespace, params.name)
        return repo
    },
)

export const save = createAsyncThunk<Repo, void, { state: {repoSettings: RepoSettingsState} }>(
    'repoSettings/save', 
    async (_, { getState, rejectWithValue, requestId } ) => {
        const { repo, saveId, saving } = getState().repoSettings
        if (!repo) {
            throw new Error("There is no repo.")
        }

        if (!(saving === RequestStatus.Pending || saveId === requestId)) {
            return repo
        }

        try {
            const nr = await updateRepo(repo.namespace, repo.name, {config_path: repo.configPath})
            message.success("Success to save.", 3)
            return nr
        } catch(e) {
            if (e instanceof HttpForbiddenError) {
                message.warn("Only admin permission can update.", 3)
            } 

            return rejectWithValue(e)
        }
    },
)

export const deactivate = createAsyncThunk<Repo, void, { state: {repoSettings: RepoSettingsState} }>(
    'repoSettings/deactivate', 
    async (_, { getState, rejectWithValue } ) => {
        const { repo } = getState().repoSettings
        if (!repo) throw new Error("There is no repo.")

        try {
            const nr = await deactivateRepo(repo.namespace, repo.name)
            window.location.reload()
            return nr
        } catch(e) {
            if (e instanceof HttpForbiddenError) {
                message.warn("Only admin permission can deactivate.", 3)
            } 

            return rejectWithValue(e)
        }
    },
)

export const repoSettingsSlice = createSlice({
    name: "repoSettings",
    initialState,
    reducers: {
        setConfigPath: (state, action: PayloadAction<string>) => {
            if (!state.repo) {
                return
            }

            state.repo.configPath = action.payload
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
            .addCase(deactivate.rejected, (state) => {
                state.deactivating = RequestStatus.Idle
            })
    }
})