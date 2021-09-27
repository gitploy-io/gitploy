import { createSlice, PayloadAction, createAsyncThunk } from '@reduxjs/toolkit'
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
    DeploymentStatusEnum, 
    Config,
    Env,
    RequestStatus, 
    HttpForbiddenError,
    HttpUnprocessableEntityError,
    HttpConflictError
} from "../models"

const page = 1
const perPage = 100

interface RepoRollbackState {
    display: boolean,
    repo?: Repo 
    config?: Config 
    env?: Env
    envs: Env[]
    deployment?: Deployment 
    deployments: Deployment[]
    /**
     * Approval selecter.
     * approvalEnabled - The approvers field is displayed when it is enabled.
     * approvers - selected approvers from candidates.
    */
    approvers: User[]
    candidates: User[]
    deployId: string
    deploying: RequestStatus
}

const initialState: RepoRollbackState = {
    display: false,
    envs: [],
    deployments: [],
    approvers: [],
    candidates: [],
    deployId: "",
    deploying: RequestStatus.Idle
}

export const init = createAsyncThunk<Repo, {namespace: string, name: string}, { state: {repoRollback: RepoRollbackState} }>(
    'repoRollback/init', 
    async (params) => {
        const repo = await searchRepo(params.namespace, params.name)
        return repo
    },
)

export const fetchConfig = createAsyncThunk<Config, void, { state: {repoRollback: RepoRollbackState} }>(
    "repoRollback/fetchEnvs", 
    async (_, { getState, rejectWithValue } ) => {
        const { repo } = getState().repoRollback
        if (!repo) throw new Error("The repo is not set.")

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
        if (!repo) throw new Error("The repo is not set.")
        if (!env) throw new Error("The env is not selected.")

        const deployments = await listDeployments(repo.id, env.name, DeploymentStatusEnum.Success, page, perPage)
        return deployments
    }
)

export const searchCandidates = createAsyncThunk<User[], string, { state: {repoRollback: RepoRollbackState }}>(
    "repoRollback/searchCandidates",
    async (q, { getState, rejectWithValue }) => {
        const { repo } = getState().repoRollback
        if (!repo) {
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
        const { repo, deployment, env, approvers, deployId, deploying } = getState().repoRollback
        if (!repo) {
            throw new Error("The repo is undefined.")
        }
        if (!deployment) {
            throw new Error("The deployment is undefined.")
        }
        if (!(deploying === RequestStatus.Pending && requestId === deployId )) {
            return
        }

        try {
            const rollback = await rollbackDeployment(repo.id, deployment.number)

            if (!env?.approval?.enabled) {
                const msg = <span>
                    It starts to rollback. <a href={`/${repo.namespace}/${repo.name}/deployments/${rollback.number}`}>#{rollback.number}</a>
                </span>
                message.success(msg, 3)
                return
            }

            approvers.forEach(async (approver) => {
                await createApproval(repo, rollback, approver)
            })

            const msg = <span>
                It is waiting approvals. <a href={`/${repo.namespace}/${repo.name}/deployments/${rollback.number}`}>#{rollback.number}</a>
            </span>
            message.success(msg, 3)
        } catch(e) {
            if (e instanceof HttpForbiddenError) {
                message.error("Only write permission can deploy.", 3)
            } else if (e instanceof HttpUnprocessableEntityError)  {
                const msg = <span> 
                    <span>It is unprocesable entity. Discussions <a href="https://github.com/gitploy-io/gitploy/discussions/64">#64</a></span><br/>
                    <span className="gitploy-quote">{e.message}</span>
                </span>
                message.error(msg, 3)
            } else if (e instanceof HttpConflictError) {
                message.error("It has conflicted, please retry it.", 3)
            }
            return rejectWithValue(e)
        }
    }
)

export const repoRollbackSlice = createSlice({
    name: "repoRollback",
    initialState,
    reducers: {
        setDisplay: (state, action: PayloadAction<boolean>) => {
            state.display = action.payload
        },
        setEnv: (state, action: PayloadAction<Env>) => {
            state.env = action.payload
        },
        setDeployment: (state, action: PayloadAction<Deployment>) => {
            state.deployment = action.payload
        },
        addApprover: (state, action: PayloadAction<User>) => {
            const candidate = action.payload

            // Check already exist or not.
            const approver = state.approvers.find(approver => approver.id === candidate.id)
            if (approver !== undefined) {
                return
            }

            state.approvers.push(candidate)
        },
        deleteApprover: (state, action: PayloadAction<User>) => {
            const candidate = action.payload

            const approvers = state.approvers.filter(approver => approver.id !== candidate.id)
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
                state.envs = config.envs.map(e => e)
                state.config = config
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