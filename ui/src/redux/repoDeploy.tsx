import { createSlice, PayloadAction, createAsyncThunk } from '@reduxjs/toolkit'
import { message } from "antd"

import { 
    User, 
    Deployment,
    Branch, 
    Commit, 
    Tag, 
    Config,
    Env,
    Status,
    DeploymentType, 
    RequestStatus, 
    HttpNotFoundError, 
    HttpForbiddenError,
    HttpConflictError,
    HttpUnprocessableEntityError
} from '../models'
import { 
    listPerms,
    getConfig, 
    listDeployments,
    listBranches, 
    getBranch, 
    listCommits, 
    getCommit, 
    listStatuses, 
    listTags, 
    getTag, 
    createDeployment,
    createApproval,
    getMe,
} from '../apis'

// fetch all at the first page.
const firstPage = 1
const perPage = 100

interface RepoDeployState {
    display: boolean
    namespace: string
    name: string
    config?: Config 
    env?: Env
    envs: Env[]
    currentDeployment?: Deployment
    type?: DeploymentType 
    branch?: Branch 
    branchStatuses: Status[]
    branches: Branch[]
    commit?: Commit 
    commitStatuses: Status[]
    commits: Commit[]
    tag?: Tag 
    tagStatuses: Status[]
    tags: Tag[]
    /**
     * Approval selecter.
     * approvalEnabled - The approvers field is displayed when it is enabled.
     * approvers - selected approvers from candidates.
    */
    approvers: User[]
    candidates: User[]
    user?: User
    deploying: RequestStatus
    deployId: string
}

const initialState: RepoDeployState = {
    display: false,
    namespace: "",
    name: "",
    envs: [],
    branchStatuses: [],
    branches: [],
    commitStatuses: [],
    commits: [],
    tagStatuses: [],
    tags: [],
    approvers: [],
    candidates: [],
    deploying: RequestStatus.Idle,
    deployId: "",
}

export const fetchConfig = createAsyncThunk<Config, void, { state: {repoDeploy: RepoDeployState} }>(
    "repoDeploy/fetchConfig", 
    async (_, { getState, rejectWithValue } ) => {
        const { namespace, name } = getState().repoDeploy

        try {
            const config = await getConfig(namespace, name)
            return config
        } catch (e) {
            return rejectWithValue(e)
        }
    },
)

export const fetchCurrentDeploymentOfEnv = createAsyncThunk<Deployment | null, Env, { state: {repoDeploy: RepoDeployState} }>(
    "repoDeploy/fetchCurrentDeployment", 
    async (env, { getState, rejectWithValue } ) => {
        const { namespace, name } = getState().repoDeploy

        try {
            const deployments = await listDeployments(namespace, name, env.name, "success", 1, 1)
            return (deployments.length > 0)? deployments[0] : null
        } catch (e) {
            return rejectWithValue(e)
        }
    },
)

export const fetchBranches = createAsyncThunk<Branch[], void, { state: {repoDeploy: RepoDeployState }}>(
    "repoDeploy/fetchBranches",
    async (_, { getState }) => {
        const { namespace, name } = getState().repoDeploy

        const branches = await listBranches(namespace, name, firstPage, perPage)
        return branches
    }
)

export const checkBranch = createAsyncThunk<Status[], void, { state: {repoDeploy: RepoDeployState}}>(
    "repoDeploy/checkBranch",
    async (_, { getState }) => {
        const { namespace, name, branch } = getState().repoDeploy
        if (!branch) {
            throw new Error("The branch is undefined.") 
        }

        const { statuses } = await listStatuses(namespace, name, branch.commitSha)
        return statuses
    }
)

export const addBranchManually = createAsyncThunk<Branch, string, { state: {repoDeploy: RepoDeployState}}>(
    "repoDeploy/addBranchManually",
    async (brnach: string, { getState, rejectWithValue }) => {
        const { namespace, name } = getState().repoDeploy

        try {
            const branch = await getBranch(namespace, name, brnach)
            return branch
        } catch(e) {
            if (e instanceof HttpNotFoundError) {
                message.warn("The branch is not found. Check the branch is correct.")
            } 

            return rejectWithValue(e)
        }
    }
)

export const fetchCommits = createAsyncThunk<Commit[], void, { state: {repoDeploy: RepoDeployState }}>(
    "repoDeploy/fetchCommits",
    async (_, { getState }) => {
        const { namespace, name, branch } = getState().repoDeploy

        const branchName = (branch)? branch.name : ""
        const commits = await listCommits(namespace, name, branchName, firstPage, perPage)
        return commits
    }
)

export const checkCommit = createAsyncThunk<Status[], void, { state: {repoDeploy: RepoDeployState}}>(
    "repoDeploy/checkCommit",
    async (_, { getState }) => {
        const { namespace, name, commit } = getState().repoDeploy

        if (!commit) {
            throw new Error("The commit is undefined.") 
        }

        const { statuses } = await listStatuses(namespace, name, commit.sha)
        return statuses
    }
)

export const addCommitManually = createAsyncThunk<Commit, string, { state: {repoDeploy: RepoDeployState}}>(
    "repoDeploy/addCommitManually",
    async (sha: string, { getState, rejectWithValue }) => {
        const { namespace, name } = getState().repoDeploy

        try {
            const commit = await getCommit(namespace, name, sha)
            return commit
        } catch(e) {
            if (e instanceof HttpNotFoundError) {
                message.warn("The ref is not found. Check the ref is correct.")
            } 

            return rejectWithValue(e)
        }
    }
)

export const fetchTags = createAsyncThunk<Tag[], void, { state: {repoDeploy: RepoDeployState }}>(
    "repoDeploy/fetchTags",
    async (_, { getState }) => {
        const { namespace, name } = getState().repoDeploy

        const tags = await listTags(namespace, name, firstPage, perPage)
        return tags
    }
)

export const checkTag = createAsyncThunk<Status[], void, { state: {repoDeploy: RepoDeployState}}>(
    "repoDeploy/checkTag",
    async (_, { getState }) => {
        const { namespace, name, tag } = getState().repoDeploy

        if (!tag) {
            throw new Error("The tag is undefined.") 
        }

        const { statuses } = await listStatuses(namespace, name, tag.commitSha)
        return statuses
    }
)

export const addTagManually = createAsyncThunk<Tag, string, { state: {repoDeploy: RepoDeployState}}>(
    "repoDeploy/addTagManually",
    async (tagName: string, { getState, rejectWithValue }) => {
        const { namespace, name } = getState().repoDeploy

        try {
            const tag = await getTag(namespace, name, tagName)
            return tag
        } catch(e) {
            if (e instanceof HttpNotFoundError) {
                message.warn("The tag is not found. Check the tag is correct.")
            } 

            return rejectWithValue(e)
        }
    }
)

export const searchCandidates = createAsyncThunk<User[], string, { state: {repoDeploy: RepoDeployState }}>(
    "repoDeploy/searchCandidates",
    async (q, { getState, rejectWithValue }) => {
        const { namespace, name } = getState().repoDeploy

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

export const fetchUser = createAsyncThunk<User, void, { state: {repoDeploy: RepoDeployState }}>(
    "repoDeploy/fetchUser",
    async (_, { rejectWithValue }) => {
        try {
            const user = await getMe()
            return user
        } catch(e) {
            return rejectWithValue(e)
        }
    }
)

export const deploy = createAsyncThunk<void, void, { state: {repoDeploy: RepoDeployState}}> (
    "repoDeploy/deploy",
    async (_ , { getState, rejectWithValue, requestId }) => {
        const { namespace, name, env, type, branch, commit, tag, approvers, deploying, deployId } = getState().repoDeploy
        if (!env) {
            throw new Error("The env is undefined.")
        }

        if (deploying !== RequestStatus.Pending || requestId !== deployId ) {
            return
        }

        try {
            let deployment: Deployment
            if (type === DeploymentType.Commit && commit) {
                deployment = await createDeployment(namespace, name, type, commit.sha, env.name)
            } else if (type === DeploymentType.Branch && branch) {
                deployment = await createDeployment(namespace, name, type, branch.name, env.name)
            } else if (type === DeploymentType.Tag && tag) {
                deployment = await createDeployment(namespace, name, type, tag.name, env.name)
            } else {
                throw new Error("The type should be one of them: commit, branch, and tag.")
            }

            if (!env.approval?.enabled) {
                const msg = <span>
                    It starts to deploy. <a href={`/${namespace}/${name}/deployments/${deployment.number}`}>#{deployment.number}</a>
                </span>
                message.success(msg, 3)
                return
            }

            approvers.forEach(async (approver) => {
                await createApproval(namespace, name, deployment.number, approver.id)
            })

            const msg = <span>
                It is waiting approvals. <a href={`/${namespace}/${name}/deployments/${deployment.number}`}>#{deployment.number}</a>
            </span>
            message.success(msg, 3)
        } catch(e) {
            if (e instanceof HttpForbiddenError) {
                message.warn("Only write permission can deploy.", 3)
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

export const repoDeploySlice = createSlice({
    name: "repoDeploy",
    initialState,
    reducers: {
        init: (state, action: PayloadAction<{namespace: string, name: string}>) => {
            state.namespace = action.payload.namespace
            state.name = action.payload.name
        },
        setDisplay: (state, action: PayloadAction<boolean>) => {
            state.display = action.payload
        },
        setEnv: (state, action: PayloadAction<Env>) => {
            state.env =  action.payload
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
            .addCase(fetchConfig.fulfilled, (state, action) => {
                const config = action.payload
                state.envs = config.envs.map(e => e)
                state.config = config
            })
            .addCase(fetchCurrentDeploymentOfEnv.fulfilled, (state, action) => {
                if (action.payload) {
                    state.currentDeployment = action.payload
                }
            })
            .addCase(fetchCurrentDeploymentOfEnv.rejected, (state) => {
                state.currentDeployment = undefined
            })
            .addCase(fetchBranches.fulfilled, (state, action) => {
                state.branches = action.payload
            })
            .addCase(checkBranch.fulfilled, (state, action) => {
                state.branchStatuses = action.payload
            })
            .addCase(addBranchManually.fulfilled, (state, action) => {
                state.branches.unshift(action.payload)
            })
            .addCase(fetchCommits.fulfilled, (state, action) => {
                state.commits = action.payload
            })
            .addCase(checkCommit.fulfilled, (state, action) => {
                state.commitStatuses = action.payload
            })
            .addCase(addCommitManually.fulfilled, (state, action) => {
                state.commits.unshift(action.payload)
            })
            .addCase(fetchTags.fulfilled, (state, action) => {
                state.tags = action.payload
            })
            .addCase(checkTag.fulfilled, (state, action) => {
                state.tagStatuses = action.payload
            })
            .addCase(addTagManually.fulfilled, (state, action) => {
                state.tags.unshift(action.payload)
            })
            .addCase(searchCandidates.pending, (state) => {
                state.candidates = []
            })
            .addCase(searchCandidates.fulfilled, (state, action) => {
                state.candidates = action.payload.filter(candidate => (candidate.id !== state.user?.id))
            })
            .addCase(fetchUser.fulfilled, (state, action) => {
                state.user = action.payload
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
