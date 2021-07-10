import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit'

import { Repo, Deployment, DeploymentStatus, Approval, HttpNotFoundError } from "../models"
import { searchRepo, getDeployment, listApprovals, getApproval } from "../apis"

interface DeploymentState {
    repo: Repo | null
    number: number
    deployment: Deployment | null
    isDeployed: boolean

    // approvals is requested approvals.
    approvals: Approval[]
    // myApproval exist if user have requested.
    myApproval: Approval | null
    // isApproved is the state which could be deployed:
    // requiredApprovalCount is equal or greater than approved.
    isApproved: boolean
}

const initialState: DeploymentState = {
    repo: null,
    number: 0,
    deployment: null,
    isDeployed: false,
    approvals: [],
    myApproval: null,
    isApproved: false
}

export const init = createAsyncThunk<Repo, {namespace: string, name: string}, { state: {deployment: DeploymentState} }>(
    'repo/init', 
    async (params, _ ) => {
        const repo = await searchRepo(params.namespace, params.name)
        return repo
    },
)

export const fetchDeployment = createAsyncThunk<Deployment, void, { state: {deployment: DeploymentState} }>(
    'repo/fetchDeployment', 
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

export const fetchApprovals = createAsyncThunk<Approval[], void, { state: {deployment: DeploymentState} }>(
    'repo/fetchApprovals', 
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

export const fetchMyApproval = createAsyncThunk<Approval | null, void, { state: {deployment: DeploymentState} }>(
    'repo/fetchMyApproval', 
    async (_, { getState, rejectWithValue } ) => {
        const { repo, deployment } = getState().deployment
        if (repo === null) throw new Error("There is no repo.")
        if (deployment === null) throw new Error("There is no deployment.")

        try {
            const approval = await getApproval(repo.id, deployment.number)
            return approval
        } catch(e) { 
            if (e instanceof HttpNotFoundError ) {
                return null
            }

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

                if (deployment.status !== DeploymentStatus.Waiting) {
                    state.isDeployed = true
                }
            })

            .addCase(fetchApprovals.fulfilled, (state, action) => {
                state.approvals = action.payload
            })

            .addCase(fetchMyApproval.fulfilled, (state, action) => {
                state.myApproval = action.payload
            })
    }
})