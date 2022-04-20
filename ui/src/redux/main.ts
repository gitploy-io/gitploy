import { createSlice, Middleware, MiddlewareAPI, isRejectedWithValue, PayloadAction, createAsyncThunk } from "@reduxjs/toolkit"
import { message } from "antd"

import { 
    User, 
    Deployment, 
    DeploymentStatusEnum, 
    Review,
    HttpInternalServerError, 
    HttpUnauthorizedError, 
    HttpPaymentRequiredError,
    License,
    ReviewStatusEnum,
} from "../models"
import { 
    getMe, 
    searchDeployments as _searchDeployments, 
    searchReviews as _searchReviews,
    getLicense 
} from "../apis"
import { getShortRef } from "../libs"

interface MainState {
    available: boolean
    authorized: boolean
    expired: boolean
    user?: User 
    deployments: Deployment[]
    reviews: Review[]
    license?: License
}

const initialState: MainState = {
    available: true,
    authorized: true,
    expired: false,
    deployments: [],
    reviews: [],
}

export const apiMiddleware: Middleware = (api: MiddlewareAPI) => (
    next
) => (action) => {
    if (!isRejectedWithValue(action)) {
        next(action)
        return
    }

    if (action.payload instanceof HttpUnauthorizedError) {
        api.dispatch(mainSlice.actions.setAuthorized(false))
    } else if (action.payload instanceof HttpInternalServerError) {
        api.dispatch(mainSlice.actions.setAvailable(false))
    } else if (action.payload instanceof HttpPaymentRequiredError) {
        // Only display the message to prevent damaging the user expirence.
        if (process.env.REACT_APP_GITPLOY_OSS?.toUpperCase() === "TRUE") {
            message.warn("Sorry, it is limited to the community edition.")
        } else {
            api.dispatch(mainSlice.actions.setExpired(true))
        }
    }

    next(action)
}

export const init = createAsyncThunk<User, void, { state: { main: MainState } }>(
    "main/init",
    async (_, { rejectWithValue }) => {
        try {
            const user = await getMe()
            return user
        } catch (e) {
            return rejectWithValue(e)
        }
    }
)

/**
 * Search all processing deployments that the user can access.
 */
export const searchDeployments = createAsyncThunk<Deployment[], void, { state: { main: MainState } }>(
    "main/searchDeployments",
    async (_, { rejectWithValue }) => {
        try {
            const deployments = await _searchDeployments([
                DeploymentStatusEnum.Waiting, 
                DeploymentStatusEnum.Created, 
                DeploymentStatusEnum.Queued,
                DeploymentStatusEnum.Running,
            ], false, false)
            return deployments
        } catch (e) {
            return rejectWithValue(e)
        }
    }
)

/**
 * Search all reviews has requested.
 */
export const searchReviews = createAsyncThunk<Review[], void, { state: { main: MainState } }>(
    "main/searchReviews",
    async (_, { rejectWithValue }) => {
        try {
            const reviews = await _searchReviews()
            return reviews
        } catch (e) {
            return rejectWithValue(e)
        }
    }
)

export const fetchLicense = createAsyncThunk<License, void, { state: { main: MainState } }>(
    "main/fetchLicense",
    async (_, { rejectWithValue }) => {
        try {
            const lic = await getLicense()
            return lic
        } catch (e) {
            return rejectWithValue(e)
        }
    }
)

const notify = (title: string, options?: NotificationOptions) => {
    if (!("Notification" in window)) {
        console.log("This browser doesn't support the notification.")
        return
    }

    if (Notification.permission === "default") {
        Notification.requestPermission()
    }

    new Notification(title, options)
}

/**
 * The browser notifies only the user who triggers the deployment.
 */
export const notifyDeploymentEvent = createAsyncThunk<void, Deployment, { state: { main: MainState } }>(
    "main/notifyDeploymentEvent",
    async (deployment, { getState }) => {
        const { user } = getState().main

        if (deployment.status === DeploymentStatusEnum.Created) {
            notify(`New Deployment #${deployment.number}`, {
                icon: "/logo192.png",
                body: `Start to deploy ${getShortRef(deployment)} to the ${deployment.env} environment of ${deployment.repo?.namespace}/${deployment.repo?.name}.`,
                tag: String(deployment.id),
            })
            return
        }

        notify(`Deployment Updated #${deployment.number}`, {
            icon: "/logo192.png",
            body: `The deployment ${deployment.number} of ${deployment.repo?.namespace}/${deployment.repo?.name} is updated ${deployment.status}.`,
            tag: String(deployment.id),
        })
    }
)

/**
 * The browser notifies the requester when the review is responded to, 
 * but it should notify the reviewer when the review is requested.
 */
export const notifyReviewmentEvent = createAsyncThunk<void, Review, { state: { main: MainState } }>(
    "main/notifyReviewmentEvent",
    async (review, { getState }) => {
        const { user } = getState().main
        if (review.status === ReviewStatusEnum.Pending) {
            notify(`Review Requested`, {
                icon: "/logo192.png",
                body: `${review.deployment?.deployer?.login} requested the review for the deployment ${review.deployment?.number} of ${review.deployment?.repo?.namespace}/${review.deployment?.repo?.name}`,
                tag: String(review.id),
            })
            return
        }

        notify(`Review Responded`, {
            icon: "/logo192.png",
            body: `${review.user?.login} ${review.status} the deployment ${review.deployment?.number} of ${review.deployment?.repo?.namespace}/${review.deployment?.repo?.name}`,
            tag: String(review.id),
        })
        return
    }
)


export const mainSlice = createSlice({
    name: "main",
    initialState,
    reducers: {
        setAvailable: (state, action: PayloadAction<boolean>) => {
            state.available = action.payload
        },
        setAuthorized: (state, action: PayloadAction<boolean>) => {
            state.authorized = action.payload
        },
        setExpired: (state, action: PayloadAction<boolean>) => {
            state.expired = action.payload
        },
        /**
         * Update the status of the deployment with an event.
         */
        handleDeploymentEvent: (state, { payload: deployment }: PayloadAction<Deployment>) => {
            if (deployment.status === DeploymentStatusEnum.Created) {
                state.deployments.unshift(deployment)
                return
            }

            state.deployments = state.deployments.filter((item) => {
                return !(item.status === DeploymentStatusEnum.Success 
                    || item.status === DeploymentStatusEnum.Failure)
            })

            state.deployments = state.deployments.map((item) => {
                return (item.id === deployment.id)? deployment : item
            })
        },
        /**
         * Reviews are removed from the state.
         */
        handleReviewEvent: (state, { payload: review }: PayloadAction<Review>) => {
            state.reviews = state.reviews.filter((item) => {
                return item.id !== review.id
            })
        },
    },
    extraReducers: builder => {
        builder
            .addCase(init.fulfilled, (state, action) => {
                state.user = action.payload
            }) 

            .addCase(searchDeployments.fulfilled, (state, action) => {
                state.deployments = action.payload
            })

            .addCase(searchReviews.fulfilled, (state, action) => {
                state.reviews = action.payload
            })

            .addCase(fetchLicense.fulfilled, (state, action) => {
                state.license = action.payload
            })
    }
})