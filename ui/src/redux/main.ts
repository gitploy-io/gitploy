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

const runningDeploymentStatus: LastDeploymentStatus[] = [
    LastDeploymentStatus.Waiting,
    LastDeploymentStatus.Created,
    LastDeploymentStatus.Running,
]

const pendingApprovalStatus: ApprovalStatus[] = [
    ApprovalStatus.Pending
]

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
            const deployments = await _searchDeployments(runningDeploymentStatus, false)
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
            const approvals = await _searchApprovals(pendingApprovalStatus)
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
        handleDeploymentEvent: (state, action: PayloadAction<Deployment>) => {
            const user = state.user
            if (user === null) {
                throw new Error("Unauthorized user.")
            }

            // It updates the status of the deployment if the deployment is found,
            // But if the deployment is not found it pushes to the list.
            const deployment = action.payload
            if (deployment.deployer?.id !== user.id) {
                return
            }

            const idx = state.deployments.findIndex((d) => {
                return d.id === deployment.id
            })

            if (idx !== -1 ) {
                if (!runningDeploymentStatus.includes(deployment.lastStatus)) {
                    state.deployments.splice(idx, 1)
                    return
                } 

                state.deployments[idx] = deployment
                return
            } 
            
            state.deployments.push(deployment)
        },
        handleApprovalEvent: (state, action: PayloadAction<Approval>) => {
            const user = state.user
            if (user === null) {
                throw new Error("Unauthorized user.")
            }

            // It updates the status of the approval if the approval is found,
            // But if the approval is not found it pushes to the list.
            const approval = action.payload
            if (approval.user?.id !== user.id) {
                return
            }

            const idx = state.approvals.findIndex((a) => {
                return a.id === approval.id
            })

            if (idx !== -1) {
                if (!pendingApprovalStatus.includes(approval.status)) {
                    console.log("remove")
                    state.approvals.splice(idx, 1)
                    return
                }

                state.approvals[idx] = approval
                return
            } 

            state.approvals.push(approval)
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