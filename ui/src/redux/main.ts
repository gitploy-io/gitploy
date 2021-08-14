import { createSlice, Middleware, MiddlewareAPI, isRejectedWithValue, PayloadAction, createAsyncThunk } from "@reduxjs/toolkit"

import { User, Deployment, LastDeploymentStatus, Approval, ApprovalStatus, HttpInternalServerError, HttpUnauthorizedError,  } from "../models"
import { getMe, searchDeployments as _searchDeployments, searchApprovals as _searchApprovals } from "../apis"

interface MainState {
    available: boolean
    authorized: boolean
    user: User | null
    deployments: Deployment[]
    approvals: Approval[]
}

const initialState: MainState = {
    available: true,
    authorized: true,
    user: null,
    deployments: [],
    approvals: []
}

export const apiMiddleware: Middleware = (api: MiddlewareAPI) => (
    next
) => (action) => {
    if (!isRejectedWithValue(action)) {
        next(action)
        return
    }

    if (action.payload instanceof HttpUnauthorizedError) {
        api.dispatch(mainSlice.actions.setAuthorized(false))
    } else if (action.payload instanceof HttpInternalServerError) {
        api.dispatch(mainSlice.actions.setAvailable(false))
    }

    next(action)
}

export const init = createAsyncThunk<User, void, { state: { main: MainState } }>(
    "main/init",
    async (_, { rejectWithValue }) => {
        try {
            const user = await getMe()
            return user
        } catch (e) {
            return rejectWithValue(e)
        }
    }
)

export const searchDeployments = createAsyncThunk<Deployment[], void, { state: { main: MainState } }>(
    "main/searchDeployments",
    async (_, { rejectWithValue }) => {
        try {
            const deployments = await _searchDeployments([
                LastDeploymentStatus.Waiting,
                LastDeploymentStatus.Created,
                LastDeploymentStatus.Running,
            ], false)
            return deployments
        } catch (e) {
            return rejectWithValue(e)
        }
    }
)

export const searchApprovals = createAsyncThunk<Approval[], void, { state: { main: MainState } }>(
    "main/searchApprovals",
    async (_, { rejectWithValue }) => {
        try {
            const approvals = await _searchApprovals([
                ApprovalStatus.Pending,
            ])
            return approvals
        } catch (e) {
            return rejectWithValue(e)
        }
    }
)

export const mainSlice = createSlice({
    name: "main",
    initialState,
    reducers: {
        setAvailable: (state, action: PayloadAction<boolean>) => {
            state.available = action.payload
        },
        setAuthorized: (state, action: PayloadAction<boolean>) => {
            state.authorized = action.payload
        },
    },
    extraReducers: builder => {
        builder
            .addCase(init.fulfilled, (state, action) => {
                state.user = action.payload
            }) 

            .addCase(searchDeployments.fulfilled, (state, action) => {
                state.deployments = action.payload
            })

            .addCase(searchApprovals.fulfilled, (state, action) => {
                state.approvals = action.payload
            })
    }
})