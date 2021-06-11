import { createSlice, PayloadAction, createAsyncThunk } from '@reduxjs/toolkit'
import { StatusCodes } from 'http-status-codes'

import { searchRepo, getConfig, listDeployments, createDeployment } from "../apis"
import { Repo, Deployment, DeploymentStatus, RequestStatus, DeploymentType, HttpNotFoundError } from "../models"

const page = 1
const perPage = 100

interface RepoRollbackState {
    repo: Repo | null
    hasConfig: boolean
    env: string
    envs: string[]
    deployment: Deployment | null
    deployments: Deployment[]
    deployId: string
    deploying: RequestStatus
}

const initialState: RepoRollbackState = {
    repo: null,
    hasConfig: true,
    env: "",
    envs: [],
    deployment: null,
    deployments: [],
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

export const fetchEnvs = createAsyncThunk<string[], void, { state: {repoRollback: RepoRollbackState} }>(
    "repoRollback/fetchEnvs", 
    async (_, { getState, rejectWithValue } ) => {
        const { repo } = getState().repoRollback
        if (repo === null) throw new Error("The repo is not set.")

        try {
            const config = await getConfig(repo.id)
            return config.envs.map(e =>  e.name)
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

export const rollback = createAsyncThunk<void, void, { state: {repoRollback: RepoRollbackState}}> (
    "repoRollback/deploy",
    async (_ , { getState, rejectWithValue, requestId }) => {
        const { repo, env, deployment, deployId, deploying } = getState().repoRollback
        if (repo === null) throw new Error("The repo is not set.")
        if (deployment === null) throw new Error("The deployment is null.")

        if (!(deploying === RequestStatus.Pending && requestId === deployId )) {
            return
        }

        try {
            if (deployment.type === DeploymentType.Branch || deployment.type === DeploymentType.Commit) {
                await createDeployment(repo.id, DeploymentType.Commit, deployment.sha, env)
            } else if (deployment.type === DeploymentType.Tag) {
                await createDeployment(repo.id, DeploymentType.Tag, deployment.ref, env)
            } 
            return
        } catch(e) {
            return rejectWithValue(e)
        }
    }
)

export const repoRollbackSlice = createSlice({
    name: "repoRollback",
    initialState,
    reducers: {
        setEnv: (state, action: PayloadAction<string>) => {
            state.env = action.payload
        },
        setDeployment: (state, action: PayloadAction<Deployment>) => {
            state.deployment = action.payload
        },
        unsetDeploy: (state) => {
            state.deploying = RequestStatus.Idle
            state.deployId = ""
        }
    },
    extraReducers: builder => {
        builder
            .addCase(init.fulfilled, (state, action) => {
                state.repo = action.payload
            })
            .addCase(fetchEnvs.fulfilled, (state, action) => {
                state.envs = action.payload
            })
            .addCase(fetchEnvs.rejected, (state, action: PayloadAction<unknown> | PayloadAction<typeof HttpNotFoundError>) => {
                if (action.payload instanceof HttpNotFoundError && action.payload.code === StatusCodes.NOT_FOUND) {
                    state.hasConfig = false
                }
            })
            .addCase(fetchDeployments.fulfilled, (state, action) => {
                state.deployments = action.payload
            })
            .addCase(rollback.pending, (state, action) => {
                if (state.deploying === RequestStatus.Idle) {
                    state.deploying = RequestStatus.Pending
                    state.deployId = action.meta.requestId
                }
            })
            .addCase(rollback.fulfilled, (state, action) => {
                if (state.deploying === RequestStatus.Pending && state.deployId === action.meta.requestId) {
                    state.deploying = RequestStatus.Success
                }
            })
            .addCase(rollback.rejected, (state, action) => {
                if (state.deploying === RequestStatus.Pending && state.deployId === action.meta.requestId) {
                    state.deploying = RequestStatus.Failure
                }
            })
    }
})