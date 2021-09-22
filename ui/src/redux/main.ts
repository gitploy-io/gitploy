import { createSlice, Middleware, MiddlewareAPI, isRejectedWithValue, PayloadAction, createAsyncThunk } from "@reduxjs/toolkit"

import { 
    User, 
    Deployment, 
    DeploymentStatusEnum, 
    Approval, 
    ApprovalStatus, 
    Event,
    EventTypeEnum,
    HttpInternalServerError, 
    HttpUnauthorizedError, 
    License,
} from "../models"
import { getMe, searchDeployments as _searchDeployments, searchApprovals as _searchApprovals, getLicense } from "../apis"
import { HttpPaymentRequiredError } from "../models/errors"

interface MainState {
    available: boolean
    authorized: boolean
    expired: boolean
    user?: User 
    deployments: Deployment[]
    approvals: Approval[]
    license?: License
}

const initialState: MainState = {
    available: true,
    authorized: true,
    expired: false,
    deployments: [],
    approvals: [],
}

const runningDeploymentStatus: DeploymentStatusEnum[] = [
    DeploymentStatusEnum.Waiting,
    DeploymentStatusEnum.Created,
    DeploymentStatusEnum.Running,
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
    } else if (action.payload instanceof HttpPaymentRequiredError) {
        api.dispatch(mainSlice.actions.setExpired(true))
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

export const fetchLicense = createAsyncThunk<License, void, { state: { main: MainState } }>(
    "main/fetchLicense",
    async (_, { rejectWithValue }) => {
        try {
            const lic = await getLicense()
            return lic
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
        setExpired: (state, action: PayloadAction<boolean>) => {
            state.expired = action.payload
        },
        handleDeploymentEvent: (state, action: PayloadAction<Event>) => {
            const user = state.user
            if (!user) {
                throw new Error("Unauthorized user.")
            }

            const event = action.payload

            // Handling the event when the owner is same.
            if (event.deployment?.deployer?.id !== user.id) {
                return
            }

            const idx = state.deployments.findIndex((deployment) => {
                return event.deployment?.id === deployment.id
            })

            if (idx !== -1 ) {
                // Remove from the list when the status is not one of 'waiting', 'created', and 'running'.
                if (!runningDeploymentStatus.includes(event.deployment.status)) {
                    state.deployments.splice(idx, 1)
                    return
                } 

                state.deployments[idx] = event.deployment
                return
            } 
            
            state.deployments.unshift(event.deployment)
        },
        handleApprovalEvent: (state, action: PayloadAction<Event>) => {
            const event = action.payload
            
            if (event.type === EventTypeEnum.Deleted) {
                const idx = state.approvals.findIndex((approval) => {
                    return event.deletedId === approval.id 
                })

                if (idx !== -1) {
                    state.approvals.splice(idx, 1)
                    return
                }
            }

            const user = state.user
            if (!user) {
                throw new Error("Unauthorized user.")
            }

            // Handling the event when the owner is same.
            if (event.approval?.user?.id !== user.id) {
                return
            }

            const idx = state.approvals.findIndex((approval) => {
                return event.approval?.id === approval.id 
            })

            if (idx !== -1) {
                if (event.type === EventTypeEnum.Updated) {
                    state.approvals.splice(idx, 1)
                    return
                }
            } 

            state.approvals.unshift(event.approval)
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

            .addCase(fetchLicense.fulfilled, (state, action) => {
                state.license = action.payload
            })
    }
})