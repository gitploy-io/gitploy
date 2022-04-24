import { createSlice, createAsyncThunk } from "@reduxjs/toolkit"
import { message } from "antd"

import { getRepo, updateRepo, deactivateRepo } from "../apis"
import { Repo, RequestStatus, HttpForbiddenError, HttpUnprocessableEntityError  } from "../models"

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

export const save = createAsyncThunk<
    Repo, 
    {
        name: string,
        config_path: string
    },
    { state: {repoSettings: RepoSettingsState} }
>(
    'repoSettings/save', 
    async (values, { getState, rejectWithValue, requestId } ) => {
        const { repo, saveId, saving } = getState().repoSettings
        if (!repo) {
            throw new Error("There is no repo.")
        }

        if (!(saving === RequestStatus.Pending || saveId === requestId)) {
            return repo
        }

        try {
            const nr = await updateRepo(repo.namespace, repo.name, values)
            message.success("Success to save.", 3)
            return nr
        } catch(e) {
            if (e instanceof HttpForbiddenError) {
                message.warn("Only admin permission can update.", 3)
            } else if (e instanceof HttpUnprocessableEntityError) {
                message.error(<> 
                    <span>It is unprocesable entity.</span><br/>
                    <span className="gitploy-quote">{e.message}</span>
                </>, 3)
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
            .addCase(deactivate.rejected, (state) => {
                state.deactivating = RequestStatus.Idle
            })
    }
})