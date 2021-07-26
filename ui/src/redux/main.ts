import { createSlice, Middleware, MiddlewareAPI, isRejectedWithValue, PayloadAction, createAsyncThunk } from "@reduxjs/toolkit"

import { User, Notification, HttpInternalServerError, HttpUnauthorizedError } from "../models"
import { getMe, listNotifications, patchNotificationChecked } from "../apis"

interface MainState {
    available: boolean
    authorized: boolean
    user: User | null
    notifications: Notification[]
}

const initialState: MainState = {
    available: true,
    authorized: true,
    user: null,
    notifications: []
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

export const fetchNotifications = createAsyncThunk<Notification[], void, { state: { main: MainState } }>(
    "main/fetchNotifications",
    async (_, { rejectWithValue }) => {
        try {
            const notifications = await listNotifications()
            return notifications
        } catch (e) {
            return rejectWithValue(e)
        }
    }
)

export const setNotificationChecked = createAsyncThunk<Notification, Notification, { state: { main: MainState }}>(
    "main/setNotificationChecked",
    async (n , { rejectWithValue }) => {
        try {
            const notification = await patchNotificationChecked(n.id)
            return notification
        } catch(e) {
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
        addNotification: (state, action: PayloadAction<Notification>) => {
            const ns = state.notifications
            ns.unshift(action.payload)
            state.notifications = ns
        }
    },
    extraReducers: builder => {
        builder
            .addCase(init.fulfilled, (state, action) => {
                state.user = action.payload
            }) 
            .addCase(fetchNotifications.fulfilled, (state, action) => {
                state.notifications = action.payload
            })
            .addCase(setNotificationChecked.fulfilled, (state, action) => {
                const notification = action.payload
                state.notifications = state.notifications.map((n) => {
                    return (notification.id === n.id)? notification : n
                })
            })
    }
})