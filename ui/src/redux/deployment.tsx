import { message } from "antd"
import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit'

import { 
    Deployment, 
    Commit,
    Review,
    Event,
    RequestStatus, 
    HttpNotFoundError, 
    HttpForbiddenError,
    HttpUnprocessableEntityError,
} from "../models"
import { 
    getDeployment, 
    createRemoteDeployment, 
    listReviews,
    getUserReview,
    approveReview,
    rejectReview,
    listDeploymentChanges,
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
    reviews: Review[]
    userReview?: Review
}

const initialState: DeploymentState = {
    display: false,
    namespace: "",
    name: "",
    number: 0,
    changes: [],
    deploying: RequestStatus.Idle,
    deployId: "",
    reviews: [],
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
            const deployment = await createRemoteDeployment(namespace, name, number)
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

export const fetchReviews = createAsyncThunk<Review[], void, { state: {deployment: DeploymentState} }>(
    'deployment/fetchReviews', 
    async (_, { getState, rejectWithValue } ) => {
        const { namespace, name, number } = getState().deployment

        try {
            const reviews = await listReviews(namespace, name, number)
            return reviews
        } catch(e) { 
            return rejectWithValue(e)
        }
    },
)

export const approve = createAsyncThunk<Review, string, { state: {deployment: DeploymentState }}>(
    "deployment/approve",
    async (comment, { getState, rejectWithValue }) => {
        const { namespace, name, number } = getState().deployment

        try {
            const review = await approveReview(namespace, name, number, comment)
            message.info("Approve to deploy.")
            return review
        } catch(e) {
            return rejectWithValue(e)
        }
    }
)


export const reject = createAsyncThunk<Review, string, { state: {deployment: DeploymentState }}>(
    "deployment/reject",
    async (comment, { getState, rejectWithValue }) => {
        const { namespace, name, number } = getState().deployment

        try {
            const review = await rejectReview(namespace, name, number, comment)
            message.info("Reject to deploy.")
            return review
        } catch(e) {
            return rejectWithValue(e)
        }
    }
)

export const fetchUserReview = createAsyncThunk<Review, void, { state: {deployment: DeploymentState} }>(
    "deployment/fetchUserReview", 
    async (_, { getState, rejectWithValue } ) => {
        const { namespace, name, number } = getState().deployment

        try {
            const review = await getUserReview(namespace, name, number)
            return review
        } catch(e) { 
            if (e instanceof HttpNotFoundError ) {
                return rejectWithValue(e)
            }

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
            .addCase(fetchReviews.fulfilled, (state, action) => {
                state.reviews = action.payload
            })
            .addCase(approve.fulfilled, (state, action) => {
                state.userReview = action.payload
                state.reviews = state.reviews.map((review) => {
                    return (review.id === action.payload.id)? action.payload : review
                })
            })
            .addCase(reject.fulfilled, (state, action) => {
                state.userReview = action.payload
                state.reviews = state.reviews.map((review) => {
                    return (review.id === action.payload.id)? action.payload : review
                })
            })
            .addCase(fetchUserReview.fulfilled, (state, action) => {
                state.userReview = action.payload
            })
    }
})