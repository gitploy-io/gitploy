import { createSlice, PayloadAction, createAsyncThunk } from '@reduxjs/toolkit'
import { StatusCodes } from 'http-status-codes'
import { message } from "antd"

import { 
    searchRepo, 
    getConfig, 
    listDeployments, 
    rollbackDeployment,
    createApproval,
    listPerms
} from "../apis"
import { 
    User,
    Repo, 
    Deployment, 
    DeploymentStatus, 
    Config,
    RequestStatus, 
    HttpNotFoundError, 
    HttpRequestError 
} from "../models"

const page = 1
const perPage = 100

interface RepoRollbackState {
    repo: Repo | null
    config: Config | null
    hasConfig: boolean
    env: string
    envs: string[]
    deployment: Deployment | null
    deployments: Deployment[]
    /**
     * Approval selecter.
     * approvalEnabled - The approvers field is displayed when it is enabled.
     * approvers - selected approvers from candidates.
    */
    approvalEnabled: boolean,
    approvers: User[]
    candidates: User[]
    deployId: string
    deploying: RequestStatus
}

const initialState: RepoRollbackState = {
    repo: null,
    config: null,
    hasConfig: true,
    env: "",
    envs: [],
    deployment: null,
    deployments: [],
    approvalEnabled: false,
    approvers: [],
    candidates: [],
    deployId: "",
    deploying: RequestStatus.Idle
}

export const init = createAsyncThunk<Repo, {namespace: string, name: string}, { state: {repoRollback: RepoRollbackState} }>(
    'repoRollback/init', 
    async (params, _ ) => {
        const repo = await searchRepo(params.namespace, params.name)
        return repo
    },
)

export const fetchConfig = createAsyncThunk<Config, void, { state: {repoRollback: RepoRollbackState} }>(
    "repoRollback/fetchEnvs", 
    async (_, { getState, rejectWithValue } ) => {
        const { repo } = getState().repoRollback
        if (repo === null) throw new Error("The repo is not set.")

        try {
            const config = await getConfig(repo.id)
            return config
        } catch (e) {
            return rejectWithValue(e)
        }
    },
)

export const fetchDeployments = createAsyncThunk<Deployment[], void, { state: {repoRollback: RepoRollbackState }}>(
    "repoRollback/fetchDeployments",
    async (_, { getState }) => {
        const { repo, env } = getState().repoRollback
        if (repo === null) throw new Error("The repo is not set.")
        if (env === "") throw new Error("The env is not selected.")

        const deployments = await listDeployments(repo.id, env, DeploymentStatus.Success, page, perPage)
        return deployments
    }
)

export const searchCandidates = createAsyncThunk<User[], string, { state: {repoRollback: RepoRollbackState }}>(
    "repoRollback/searchCandidates",
    async (q, { getState, rejectWithValue }) => {
        const { repo } = getState().repoRollback
        if (repo === null) {
            throw new Error("The repo is not set.")
        }

        try {
            const perms = await listPerms(repo, q)
            const candidates = perms.map((p) => {
                return p.user
            })
            return candidates
        } catch(e) {
            return rejectWithValue(e)
        }
    }
)

export const rollback = createAsyncThunk<void, void, { state: {repoRollback: RepoRollbackState}}> (
    "repoRollback/deploy",
    async (_ , { getState, rejectWithValue, requestId }) => {
        const { repo, deployment, approvalEnabled, approvers, deployId, deploying } = getState().repoRollback
        if (repo === null) throw new Error("The repo is not set.")
        if (deployment === null) throw new Error("The deployment is null.")

        if (!(deploying === RequestStatus.Pending && requestId === deployId )) {
            return
        }

        try {
            const rollback = await rollbackDeployment(repo.id, deployment.number)

            if (!approvalEnabled) {
                message.success("It starts to rollback.", 3)
                return
            }

            approvers.forEach(async (approver) => {
                await createApproval(repo, rollback, approver)
            })
            message.success("It starts to rollback.", 3)
        } catch(e) {
            if (e instanceof HttpRequestError && e.code === StatusCodes.CONFLICT) {
                message.error("The rollback is conflicted, please retry.", 3)
                return rejectWithValue(e)
            }
            
            message.error("It has failed to rollback.", 3)
            return rejectWithValue(e)
        }
    }
)

export const repoRollbackSlice = createSlice({
    name: "repoRollback",
    initialState,
    reducers: {
        setEnv: (state, action: PayloadAction<string>) => {
            const name = action.payload
            state.env = name

            if (state.config === null) {
                return
            }

            const env = state.config.envs.find(env => env.name === name)
            if (env !== undefined) {
                state.approvalEnabled = env.approvalEnabled
            }
        },
        setDeployment: (state, action: PayloadAction<Deployment>) => {
            state.deployment = action.payload
        },
        addApprover: (state, action: PayloadAction<string>) => {
            const userId = action.payload

            const candidate = state.candidates.find(candidate => candidate.id === userId)
            if (candidate === undefined) {
                return
            }

            // Check already exist or not.
            const approver = state.approvers.find(approver => approver.id === userId)
            if (approver !== undefined) {
                return
            }

            state.approvers.push(candidate)
        },
        deleteApprover: (state, action: PayloadAction<string>) => {
            const userId = action.payload

            const approvers = state.approvers.filter(approver => approver.id !== userId)
            state.approvers = approvers
        },
    },
    extraReducers: builder => {
        builder
            .addCase(init.fulfilled, (state, action) => {
                state.repo = action.payload
            })
            .addCase(fetchConfig.fulfilled, (state, action) => {
                const config = action.payload
                state.envs = config.envs.map(e => e.name)
                state.config = config
                state.hasConfig = true
            })
            .addCase(fetchConfig.rejected, (state, action: PayloadAction<unknown> | PayloadAction<typeof HttpNotFoundError>) => {
                if (action.payload instanceof HttpNotFoundError && action.payload.code === StatusCodes.NOT_FOUND) {
                    state.hasConfig = false
                }
            })
            .addCase(fetchDeployments.fulfilled, (state, action) => {
                state.deployments = action.payload
            })
            .addCase(searchCandidates.pending, (state) => {
                state.candidates = []
            })
            .addCase(searchCandidates.fulfilled, (state, action) => {
                state.candidates = action.payload
            })
            .addCase(rollback.pending, (state, action) => {
                if (state.deploying === RequestStatus.Idle) {
                    state.deploying = RequestStatus.Pending
                    state.deployId = action.meta.requestId
                }
            })
            .addCase(rollback.fulfilled, (state, action) => {
                if (state.deploying === RequestStatus.Pending && state.deployId === action.meta.requestId) {
                    state.deploying = RequestStatus.Idle
                }
            })
            .addCase(rollback.rejected, (state, action) => {
                if (state.deploying === RequestStatus.Pending && state.deployId === action.meta.requestId) {
                    state.deploying = RequestStatus.Idle
                }
            })
    }
})