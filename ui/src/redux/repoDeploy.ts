import { createSlice, PayloadAction, createAsyncThunk } from '@reduxjs/toolkit'
import { StatusCodes } from 'http-status-codes'
import { message } from "antd"

import { Repo, Branch, Commit, StatusState, DeploymentType, Tag, RequestStatus, HttpNotFoundError } from '../models'
import { 
    searchRepo, 
    getConfig, 
    listBranches, 
    getBranch, 
    listCommits, 
    getCommit, 
    listStatuses, 
    listTags, 
    getTag, 
    createDeployment } from '../apis'

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
    branchCheck: StatusState
    branches: Branch[]
    commit: Commit | null
    commitCheck: StatusState
    commits: Commit[]
    tag: Tag | null
    tagCheck: StatusState
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
    branchCheck: StatusState.Null,
    branches: [],
    commit: null,
    commitCheck: StatusState.Null,
    commits: [],
    tag: null,
    tagCheck: StatusState.Null,
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
        if (repo === null) throw new Error("The repo is not set.")

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
        if (repo === null) throw new Error("The repo is not set.")

        const branches = await listBranches(repo.id, firstPage, perPage)
        return branches
    }
)

export const checkBranch = createAsyncThunk<StatusState, void, { state: {repoDeploy: RepoDeployState}}>(
    "repoDeploy/checkBranch",
    async (_, { getState }) => {
        const { repo, branch } = getState().repoDeploy
        if (repo === null || branch === null) throw new Error("The repo and branch are not set.") 

        const { state } = await listStatuses(repo.id, branch.commitSha)
        return state
    }
)

export const addBranchManually = createAsyncThunk<Branch, string, { state: {repoDeploy: RepoDeployState}}>(
    "repoDeploy/addBranchManually",
    async (name: string, { getState, rejectWithValue }) => {
        const { repo } = getState().repoDeploy
        if (repo === null) throw new Error("The repo is not set.")

        try {
            const branch = await getBranch(repo.id, name)
            return branch
        } catch(e) {
            if (e instanceof HttpNotFoundError) {
                message.error("The ref is not found. Check the ref is corrent.")
            } else {
                message.error("It has failed to add the ref.")
            }
            return rejectWithValue(e)
        }
    }
)

export const fetchCommits = createAsyncThunk<Commit[], void, { state: {repoDeploy: RepoDeployState }}>(
    "repoDeploy/fetchCommits",
    async (_, { getState }) => {
        const { repo, branch } = getState().repoDeploy
        if (repo === null) throw new Error("The repo is not set.")

        const name = (branch !== null)? branch.name : ""
        const commits = await listCommits(repo.id, name, firstPage, perPage)
        return commits
    }
)

export const checkCommit = createAsyncThunk<StatusState, void, { state: {repoDeploy: RepoDeployState}}>(
    "repoDeploy/checkCommit",
    async (_, { getState }) => {
        const { repo, commit } = getState().repoDeploy
        if (repo === null || commit === null) throw new Error("The repo and commit are not set.") 

        const { state } = await listStatuses(repo.id, commit.sha)
        return state
    }
)

export const addCommitManually = createAsyncThunk<Commit, string, { state: {repoDeploy: RepoDeployState}}>(
    "repoDeploy/addCommitManually",
    async (sha: string, { getState, rejectWithValue }) => {
        const { repo } = getState().repoDeploy
        if (repo === null) throw new Error("The repo is not set.")

        try {
            const commit = await getCommit(repo.id, sha)
            return commit
        } catch(e) {
            if (e instanceof HttpNotFoundError) {
                message.error("The ref is not found. Check the ref is corrent.")
            } else {
                message.error("It has failed to add the ref.")
            }
            return rejectWithValue(e)
        }
    }
)

export const fetchTags = createAsyncThunk<Tag[], void, { state: {repoDeploy: RepoDeployState }}>(
    "repoDeploy/fetchTags",
    async (_, { getState }) => {
        const { repo } = getState().repoDeploy
        if (repo === null) throw new Error("The repo is not set.")

        const tags = await listTags(repo.id, firstPage, perPage)
        return tags
    }
)

export const checkTag = createAsyncThunk<StatusState, void, { state: {repoDeploy: RepoDeployState}}>(
    "repoDeploy/checkTag",
    async (_, { getState }) => {
        const { repo, tag } = getState().repoDeploy
        if (repo === null || tag === null) throw new Error("The repo and tag are not set.") 

        const { state } = await listStatuses(repo.id, tag.commitSha)
        return state
    }
)

export const addTagManually = createAsyncThunk<Tag, string, { state: {repoDeploy: RepoDeployState}}>(
    "repoDeploy/addTagManually",
    async (name: string, { getState, rejectWithValue }) => {
        const { repo } = getState().repoDeploy
        if (repo === null) throw new Error("The repo is not set.")

        try {
            const tag = await getTag(repo.id, name)
            return tag
        } catch(e) {
            if (e instanceof HttpNotFoundError) {
                message.error("The ref is not found. Check the ref is corrent.")
            } else {
                message.error("It has failed to add the ref.")
            }
            return rejectWithValue(e)
        }
    }
)

export const deploy = createAsyncThunk<void, void, { state: {repoDeploy: RepoDeployState}}> (
    "repoDeploy/deploy",
    async (_ , { getState, rejectWithValue, requestId }) => {
        const { repo, env, type, branch, commit, tag, deploying, deployId } = getState().repoDeploy
        if (repo === null) {
            throw new Error("The repo is not set.")
        }
        if (deploying !== RequestStatus.Pending || requestId !== deployId ) {
            return
        }

        try {
            if (type === DeploymentType.Commit && commit !== null) {
                await createDeployment(repo.id, type, commit.sha, env)
            } else if (type === DeploymentType.Branch && branch !== null) {
                await createDeployment(repo.id, type, branch.name, env)
            } else if (type === DeploymentType.Tag && tag !== null) {
                await createDeployment(repo.id, type, tag.name, env)
            } 

            message.success("It starts to deploy.", 3)
            return
        } catch(e) {
            message.error("It has failed to deploy.", 3)
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
        setCommit: (state, action: PayloadAction<Commit>) => {
            state.commit = action.payload
        },
        setTag: (state, action: PayloadAction<Tag>) => {
            state.tag = action.payload
        },
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
            .addCase(checkBranch.fulfilled, (state, action) => {
                state.branchCheck = action.payload
            })
            .addCase(addBranchManually.fulfilled, (state, action) => {
                state.branches.unshift(action.payload)
            })
            .addCase(fetchCommits.fulfilled, (state, action) => {
                state.commits = action.payload
            })
            .addCase(checkCommit.fulfilled, (state, action) => {
                state.commitCheck = action.payload
            })
            .addCase(addCommitManually.fulfilled, (state, action) => {
                state.commits.unshift(action.payload)
            })
            .addCase(fetchTags.fulfilled, (state, action) => {
                state.tags = action.payload
            })
            .addCase(checkTag.fulfilled, (state, action) => {
                state.tagCheck = action.payload
            })
            .addCase(addTagManually.fulfilled, (state, action) => {
                state.tags.unshift(action.payload)
            })
            .addCase(deploy.pending, (state, action) => {
                if (state.deploying === RequestStatus.Idle) {
                    state.deploying = RequestStatus.Pending
                    state.deployId = action.meta.requestId
                }
            })
            .addCase(deploy.fulfilled, (state, action) => {
                if (state.deploying === RequestStatus.Pending && state.deployId === action.meta.requestId) {
                    state.deploying = RequestStatus.Idle
                }
            })
            .addCase(deploy.rejected, (state, action) => {
                if (state.deploying === RequestStatus.Pending && state.deployId === action.meta.requestId) {
                    state.deploying = RequestStatus.Idle
                }
            })
    }
})
