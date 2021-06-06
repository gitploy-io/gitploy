import { createSlice, PayloadAction, createAsyncThunk } from '@reduxjs/toolkit'

import { Repo, Branch, Commit, DeploymentType, Tag, RequestStatus, HttpNotFoundError } from '../models'
import { searchRepo, getConfig, listBranches, listCommits, listTags, createDeployment } from '../apis'
import { StatusCodes } from 'http-status-codes'

// fetch all at the first page.
const firstPage = 1
const perPage = 100

interface RepoDeployState {
    repo: Repo | null
    hasConfig: boolean
    env: string
    envs: string[]
    type: DeploymentType | null
    branch: Branch | null
    branches: Branch[]
    commit: Commit | null
    commits: Commit[]
    tag: Tag | null
    tags: Tag[]
    deploying: RequestStatus
    deployId: string
}

const initialState: RepoDeployState = {
    repo: null,
    hasConfig: true,
    env: "",
    envs: [],
    type: null,
    branch: null,
    branches: [],
    commit: null,
    commits: [],
    tag: null,
    tags: [],
    deploying: RequestStatus.Idle,
    deployId: "",
}

export const init = createAsyncThunk<Repo, {namespace: string, name: string}, { state: {repoDeploy: RepoDeployState} }>(
    'repoDeploy/init', 
    async (params, _ ) => {
        const repo = await searchRepo(params.namespace, params.name)
        return repo
    },
)

export const fetchEnvs = createAsyncThunk<string[], void, { state: {repoDeploy: RepoDeployState} }>(
    "repoDeploy/fetchEnvs", 
    async (_, { getState, rejectWithValue } ) => {
        const { repo } = getState().repoDeploy
        if (repo === null) throw new Error("The repo was not set.")

        try {
            const config = await getConfig(repo.id)
            return config.envs.map(e =>  e.name)
        } catch (e) {
            return rejectWithValue(e)
        }
    },
)

export const fetchBranches = createAsyncThunk<Branch[], void, { state: {repoDeploy: RepoDeployState }}>(
    "repoDeploy/fetchBranches",
    async (_, { getState }) => {
        const { repo } = getState().repoDeploy
        if (repo === null) throw new Error("The repo was not set.")

        const branches = await listBranches(repo.id, firstPage, perPage)
        return branches
    }
)

export const fetchCommits = createAsyncThunk<Commit[], void, { state: {repoDeploy: RepoDeployState }}>(
    "repoDeploy/fetchCommits",
    async (_, { getState }) => {
        const { repo, branch } = getState().repoDeploy
        if (repo === null) throw new Error("The repo was not set.")

        const name = (branch !== null)? branch.name : ""
        const commits = await listCommits(repo.id, name, firstPage, perPage)
        return commits
    }
)

export const fetchTags = createAsyncThunk<Tag[], void, { state: {repoDeploy: RepoDeployState }}>(
    "repoDeploy/fetchTags",
    async (_, { getState }) => {
        const { repo } = getState().repoDeploy
        if (repo === null) throw new Error("The repo was not set.")

        const tags = await listTags(repo.id, firstPage, perPage)
        return tags
    }
)

// TODO: support dynamic addition for commit, branch, tag by async.

export const deploy = createAsyncThunk<void, void, { state: {repoDeploy: RepoDeployState}}> (
    "repoDeploy/deploy",
    async (_ , { getState, rejectWithValue, requestId }) => {
        const { repo, env, type, branch, commit, tag, deploying, deployId } = getState().repoDeploy
        if (repo === null) {
            throw new Error("The repo was not set.")
        }
        if (deploying !== RequestStatus.Pending || requestId !== deployId ) {
            return
        }

        try {
            if (type === DeploymentType.Commit && commit !== null) {
                createDeployment(repo.id, type, commit.sha, env)
            } else if (type === DeploymentType.Branch && branch !== null) {
                createDeployment(repo.id, type, branch.name, env)
            } else if (type === DeploymentType.Tag && tag !== null) {
                createDeployment(repo.id, type, tag.name, env)
            } else {
                throw new Error("failed")
            }
        } catch(e) {
            return rejectWithValue(e)
        }
    }
)

export const repoDeploySlice = createSlice({
    name: "repoDeploy",
    initialState,
    reducers: {
        setEnv: (state, action: PayloadAction<string>) => {
            state.env = action.payload
        },
        setType: (state, action: PayloadAction<DeploymentType>) => {
            state.type = action.payload
        },
        setBranch: (state, action: PayloadAction<Branch>) => {
            state.branch = action.payload
        },
        addBranchManually: (state, action: PayloadAction<Branch>) => {
            state.branches.unshift(action.payload)
        },
        setCommit: (state, action: PayloadAction<Commit>) => {
            state.commit = action.payload
        },
        addCommitManually: (state, action: PayloadAction<Commit>) => {
            state.commits.unshift(action.payload)
        },
        setTag: (state, action: PayloadAction<Tag>) => {
            state.tag = action.payload
        },
        addTagManually: (state, action: PayloadAction<Tag>) => {
            state.tags.unshift(action.payload)
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
            .addCase(fetchBranches.fulfilled, (state, action) => {
                state.branches = action.payload
            })
            .addCase(fetchCommits.fulfilled, (state, action) => {
                state.commits = action.payload
            })
            .addCase(fetchTags.fulfilled, (state, action) => {
                state.tags = action.payload
            })
            .addCase(deploy.pending, (state, action) => {
                if (state.deploying === RequestStatus.Idle) {
                    state.deploying = RequestStatus.Pending
                    state.deployId = action.meta.requestId
                }
            })
            .addCase(deploy.fulfilled, (state, action) => {
                if (state.deploying === RequestStatus.Pending && state.deployId === action.meta.requestId) {
                    state.deploying = RequestStatus.Success
                }
            })
            .addCase(deploy.rejected, (state, action) => {
                if (state.deploying === RequestStatus.Pending && state.deployId === action.meta.requestId) {
                    state.deploying = RequestStatus.Failure
                }
            })
    }
})
