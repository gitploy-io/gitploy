import { createSlice, Middleware, MiddlewareAPI, isRejectedWithValue, PayloadAction, createAsyncThunk } from "@reduxjs/toolkit"
import { message } from "antd"

import { 
    User, 
    Deployment, 
    DeploymentStatusEnum, 
    Review,
    Event,
    EventKindEnum,
    EventTypeEnum,
    HttpInternalServerError, 
    HttpUnauthorizedError, 
    License,
} from "../models"
import { 
    getMe, 
    searchDeployments as _searchDeployments, 
    searchReviews as _searchReviews,
    getLicense 
} from "../apis"
import { HttpPaymentRequiredError } from "../models/errors"

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

const runningDeploymentStatus: DeploymentStatusEnum[] = [
    DeploymentStatusEnum.Waiting,
    DeploymentStatusEnum.Created,
    DeploymentStatusEnum.Running,
]

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

export const searchDeployments = createAsyncThunk<Deployment[], void, { state: { main: MainState } }>(
    "main/searchDeployments",
    async (_, { rejectWithValue }) => {
        try {
            const deployments = await _searchDeployments(runningDeploymentStatus, false)
            return deployments
        } catch (e) {
            return rejectWithValue(e)
        }
    }
)

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
        handleDeploymentEvent: (state, action: PayloadAction<Event>) => {
            const user = state.user
            if (!user) {
                throw new Error("Unauthorized user.")
            }

            const event = action.payload

            // Handling the event when the owner is same.
            if (event.deployment?.deployer?.id !== user.id) {
                return
            }

            const idx = state.deployments.findIndex((deployment) => {
                return event.deployment?.id === deployment.id
            })

            if (idx !== -1 ) {
                // Remove from the list when the status is not one of 'waiting', 'created', and 'running'.
                if (!runningDeploymentStatus.includes(event.deployment.status)) {
                    state.deployments.splice(idx, 1)
                    return
                } 

                state.deployments[idx] = event.deployment
                return
            } 
            
            state.deployments.unshift(event.deployment)
        },
        handleReviewEvent: (state, action: PayloadAction<Event>) => {
            const event = action.payload

            if (action.payload.kind !== EventKindEnum.Review) {
                return
            }
            
            const user = state.user
            if (!user) {
                throw new Error("Unauthorized user.")
            }

            // Handling the event when the user own the event.
            if (event.review?.user?.id !== user.id) {
                return
            }

            if (event.type === EventTypeEnum.Created) {
                state.reviews.unshift(event.review)
                return
            }

            const idx = state.reviews.findIndex((review) => {
                return event.review?.id === review.id 
            })

            if (idx !== -1) {
                state.reviews.splice(idx, 1)
                return
            } 
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