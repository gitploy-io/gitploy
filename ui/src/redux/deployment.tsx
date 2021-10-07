import { message } from "antd"
import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit'

import { 
    User,
    Deployment, 
    Commit,
    Approval, 
    Event,
    RequestStatus, 
    HttpNotFoundError, 
    HttpForbiddenError,
    HttpUnprocessableEntityError,
} from "../models"
import { 
    getDeployment, 
    updateDeploymentStatusCreated, 
    listPerms,
    listApprovals, 
    createApproval as _createApproval,
    deleteApproval as _deleteApproval,
    getMyApproval, 
    setApprovalApproved, 
    setApprovalDeclined, 
    listDeploymentChanges
} from "../apis"

interface DeploymentState {
    display: boolean
    namespace: string
    name: string
    number: number
    deployment?: Deployment
    changes: Commit[]
    deploying: RequestStatus
    deployId: string

    // approvals is requested approvals.
    approvals: Approval[]
    candidates: User[]
    // myApproval exist if user have requested.
    myApproval?: Approval
}

const initialState: DeploymentState = {
    display: false,
    namespace: "",
    name: "",
    number: 0,
    changes: [],
    deploying: RequestStatus.Idle,
    deployId: "",
    approvals: [],
    candidates: [],
}

export const fetchDeployment = createAsyncThunk<Deployment, void, { state: {deployment: DeploymentState} }>(
    'deployment/fetchDeployment', 
    async (_, { getState, rejectWithValue } ) => {
        const { namespace, name, number } = getState().deployment

        try {
            const deployment = await getDeployment(namespace, name, number)
            return deployment
        } catch(e) { 
            return rejectWithValue(e)
        }
    },
)

export const fetchDeploymentChanges = createAsyncThunk<Commit[], void, { state: {deployment: DeploymentState} }>(
    'deployment/fetchDeploymentChanges', 
    async (_, { getState, rejectWithValue } ) => {
        const { namespace, name, number } = getState().deployment

        try {
            const commits = await listDeploymentChanges(namespace, name, number)
            return commits
        } catch(e) { 
            return rejectWithValue(e)
        }
    },
)

export const deployToSCM = createAsyncThunk<Deployment, void, { state: {deployment: DeploymentState} }>(
    'deployment/deployToSCM', 
    async (_, { getState, rejectWithValue, requestId } ) => {
        const { namespace, name, number, deploying, deployId } = getState().deployment

        if (deploying !== RequestStatus.Pending || requestId !== deployId ) {
            throw new Error("The previous action is not finished.")
        }

        try {
            const deployment = await updateDeploymentStatusCreated(namespace, name, number)
            message.info("It starts to deploy.", 3)

            return deployment
        } catch(e) { 
            if (e instanceof HttpForbiddenError) {
                message.warn("Only write permission can deploy.", 3)
            } else if (e instanceof HttpUnprocessableEntityError)  {
                const msg = <span> 
                    <span>It is unprocesable entity. Discussions <a href="https://github.com/gitploy-io/gitploy/discussions/64">#64</a></span><br/>
                    <span className="gitploy-quote">{e.message}</span>
                </span>
                message.error(msg, 3)
            } 

            return rejectWithValue(e)
        }
    },
)

export const fetchApprovals = createAsyncThunk<Approval[], void, { state: {deployment: DeploymentState} }>(
    'deployment/fetchApprovals', 
    async (_, { getState, rejectWithValue } ) => {
        const { namespace, name, number } = getState().deployment

        try {
            const approvals = await listApprovals(namespace, name, number)
            return approvals
        } catch(e) { 
            return rejectWithValue(e)
        }
    },
)

export const searchCandidates = createAsyncThunk<User[], string, { state: {deployment: DeploymentState }}>(
    "deployment/fetchCandidates",
    async (q, { getState, rejectWithValue }) => {
        const { namespace, name } = getState().deployment

        try {
            const perms = await listPerms(namespace, name, q)
            const candidates = perms.map((p) => {
                return p.user
            })
            return candidates
        } catch(e) {
            return rejectWithValue(e)
        }
    }
)

export const createApproval = createAsyncThunk<Approval, User, { state: {deployment: DeploymentState }}>(
    "deployment/createApprover",
    async (candidate, { getState, rejectWithValue }) => {
        const { namespace, name, number } = getState().deployment

        try {
            const approval = await _createApproval(namespace, name, number, candidate.id)
            return approval
        } catch(e) {
            return rejectWithValue(e)
        }
    }
)


export const deleteApproval = createAsyncThunk<Approval, Approval, { state: {deployment: DeploymentState }}>(
    "deployment/deleteApprover",
    async (approval, { getState, rejectWithValue }) => {
        const { namespace, name } = getState().deployment

        try {
            await _deleteApproval(namespace, name, approval.id)
            return approval
        } catch(e) {
            return rejectWithValue(e)
        }
    }
)

export const fetchMyApproval = createAsyncThunk<Approval, void, { state: {deployment: DeploymentState} }>(
    'deployment/fetchMyApproval', 
    async (_, { getState, rejectWithValue } ) => {
        const { namespace, name, number } = getState().deployment

        try {
            const approval = await getMyApproval(namespace, name, number)
            return approval
        } catch(e) { 
            if (e instanceof HttpNotFoundError ) {
                return rejectWithValue(e)
            }

            return rejectWithValue(e)
        }
    },
)

export const approve = createAsyncThunk<Approval, void, { state: {deployment: DeploymentState} }>(
    'deployment/approve', 
    async (_, { getState, rejectWithValue } ) => {
        const { namespace, name, number } = getState().deployment

        try {
            const approval = await setApprovalApproved(namespace, name, number)
            return approval
        } catch(e) { 
            return rejectWithValue(e)
        }
    },
)

export const decline = createAsyncThunk<Approval, void, { state: {deployment: DeploymentState} }>(
    'deployment/decline', 
    async (_, { getState, rejectWithValue } ) => {
        const { namespace, name, number } = getState().deployment

        try {
            const approval = await setApprovalDeclined(namespace, name, number)
            return approval
        } catch(e) { 
            return rejectWithValue(e)
        }
    },
)

export const deploymentSlice = createSlice({
    name: "deployment",
    initialState,
    reducers: {
        init: (state, action: PayloadAction<{namespace: string, name: string, number: number}>) => {
            state.namespace = action.payload.namespace
            state.name = action.payload.name
            state.number = action.payload.number
        },
        setDisplay: (state, action: PayloadAction<boolean>) => {
            state.display = action.payload
        },
        handleDeploymentEvent: (state, action: PayloadAction<Event>) => {
            const event = action.payload

            if (event.deployment?.id !== state.deployment?.id) {
                return
            }

            state.deployment = event.deployment
        }
    },
    extraReducers: builder => {
        builder
            .addCase(fetchDeployment.fulfilled, (state, action) => {
                state.deployment = action.payload
            })
            .addCase(fetchDeploymentChanges.fulfilled, (state, action) => {
                state.changes = action.payload
            })
            .addCase(deployToSCM.pending, (state, action) => {
                if (state.deploying === RequestStatus.Idle) {
                    state.deploying = RequestStatus.Pending
                    state.deployId = action.meta.requestId
                }
            })
            .addCase(deployToSCM.fulfilled, (state, action) => {
                state.deployment = action.payload
                state.deploying = RequestStatus.Idle
            })
            .addCase(deployToSCM.rejected, (state) => {
                state.deploying = RequestStatus.Idle
            })
            .addCase(fetchApprovals.fulfilled, (state, action) => {
                state.approvals = action.payload
            })
            .addCase(searchCandidates.pending, (state) => {
                state.candidates = []
            })
            .addCase(searchCandidates.fulfilled, (state, action) => {
                state.candidates = action.payload.filter(candidate => candidate.id !== state.deployment?.deployer?.id)
            })
            .addCase(createApproval.fulfilled, (state, action) => {
                state.approvals.push(action.payload)
            })
            .addCase(deleteApproval.fulfilled, (state, action) => {
                const approval = action.payload
                state.approvals = state.approvals.filter(a => a.id !== approval.id)
            })
            .addCase(fetchMyApproval.fulfilled, (state, action) => {
                state.myApproval = action.payload
            })
            .addCase(approve.fulfilled, (state, action) => {
                const myApproval = action.payload
                state.myApproval = myApproval
                state.approvals = state.approvals.map((approval) => {
                    if (approval.id === myApproval.id) {
                        return myApproval
                    } 
                    return approval
                })
            })
            .addCase(decline.fulfilled, (state, action) => {
                const myApproval = action.payload
                state.myApproval = myApproval
                state.approvals = state.approvals.map((approval) => {
                    if (approval.id === myApproval.id) {
                        return myApproval
                    } 
                    return approval
                })
            })
    }
})