import { createSlice, PayloadAction, createAsyncThunk } from '@reduxjs/toolkit'

import { Repo, Branch, Commit, DeploymentType, NotFoundError, Tag,  } from '../models'
import { searchRepo, getConfig, listBranches, listCommits, listTags } from '../apis'
import { StatusCodes } from 'http-status-codes'

interface RepoDeployState {
    deploying: boolean
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
}

const initialState: RepoDeployState = {
    deploying: false,
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
}

// fetch all at the first page.
const firstPage = 1
const perPage = 100

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
    },
    extraReducers: builder => {
        builder
            .addCase(init.fulfilled, (state, action) => {
                state.repo = action.payload
            })
            .addCase(fetchEnvs.fulfilled, (state, action) => {
                state.envs = action.payload
            })
            .addCase(fetchEnvs.rejected, (state, action: PayloadAction<any>) => {
                if (action.payload.code === StatusCodes.NOT_FOUND) {
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
    }
})
