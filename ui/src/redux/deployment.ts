import { message } from "antd"
import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit'

import { Repo, Deployment, Approval, RequestStatus, HttpNotFoundError, HttpUnprocessableEntityError } from "../models"
import { searchRepo, getDeployment, updateDeploymentStatusCreated, listApprovals, getApproval, setApprovalApproved, setApprovalDeclined } from "../apis"

interface DeploymentState {
    repo: Repo | null
    number: number
    deployment: Deployment | null
    deploying: RequestStatus
    deployId: string

    // approvals is requested approvals.
    approvals: Approval[]
    // myApproval exist if user have requested.
    myApproval: Approval | null
}

const initialState: DeploymentState = {
    repo: null,
    number: 0,
    deployment: null,
    deploying: RequestStatus.Idle,
    deployId: "",
    approvals: [],
    myApproval: null,
}

export const init = createAsyncThunk<Repo, {namespace: string, name: string}, { state: {deployment: DeploymentState} }>(
    'deployment/init', 
    async (params, _ ) => {
        const repo = await searchRepo(params.namespace, params.name)
        return repo
    },
)

export const fetchDeployment = createAsyncThunk<Deployment, void, { state: {deployment: DeploymentState} }>(
    'deployment/fetchDeployment', 
    async (_, { getState, rejectWithValue } ) => {
        const { repo, number } = getState().deployment
        if (repo === null) throw new Error("There is no repo.")

        try {
            const deployment = await getDeployment(repo.id, number)
            return deployment
        } catch(e) { 
            return rejectWithValue(e)
        }
    },
)

export const deployToSCM = createAsyncThunk<Deployment, void, { state: {deployment: DeploymentState} }>(
    'deployment/deployToSCM', 
    async (_, { getState, rejectWithValue, requestId } ) => {
        const { repo, number, deploying, deployId } = getState().deployment
        if (repo === null) {
            throw new Error("There is no repo.")
        }

        if (deploying !== RequestStatus.Pending || requestId !== deployId ) {
            throw new Error("The previous action is not finished.")
        }

        try {
            const deployment = await updateDeploymentStatusCreated(repo.id, number)
            return deployment
        } catch(e) { 
            if (e instanceof HttpUnprocessableEntityError) {
                message.error(`Deploy Failure: ${e.message}`)
                return rejectWithValue(e)
            }
            return rejectWithValue(e)
        }
    },
)

export const fetchApprovals = createAsyncThunk<Approval[], void, { state: {deployment: DeploymentState} }>(
    'deployment/fetchApprovals', 
    async (_, { getState, rejectWithValue } ) => {
        const { repo, number } = getState().deployment
        if (repo === null) throw new Error("There is no repo.")

        try {
            const approvals = await listApprovals(repo.id, number)
            return approvals
        } catch(e) { 
            return rejectWithValue(e)
        }
    },
)

export const fetchMyApproval = createAsyncThunk<Approval, void, { state: {deployment: DeploymentState} }>(
    'deployment/fetchMyApproval', 
    async (_, { getState, rejectWithValue } ) => {
        const { repo, number } = getState().deployment
        if (repo === null) throw new Error("There is no repo.")

        try {
            const approval = await getApproval(repo.id, number)
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
        const { repo, number } = getState().deployment
        if (repo === null) throw new Error("There is no repo.")

        try {
            const approval = await setApprovalApproved(repo.id, number)
            return approval
        } catch(e) { 
            return rejectWithValue(e)
        }
    },
)

export const decline = createAsyncThunk<Approval, void, { state: {deployment: DeploymentState} }>(
    'deployment/decline', 
    async (_, { getState, rejectWithValue } ) => {
        const { repo, number } = getState().deployment
        if (repo === null) throw new Error("There is no repo.")

        try {
            const approval = await setApprovalDeclined(repo.id, number)
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
        setNumber: (state, action: PayloadAction<number>) => {
            state.number = action.payload
        }
    },
    extraReducers: builder => {
        builder
            .addCase(init.fulfilled, (state, action) => {
                state.repo = action.payload
            })

            .addCase(fetchDeployment.fulfilled, (state, action) => {
                const deployment = action.payload
                state.deployment = deployment
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