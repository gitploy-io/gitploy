import { createSlice, Middleware, MiddlewareAPI, isRejectedWithValue, PayloadAction } from "@reduxjs/toolkit"
import { HttpInternalServerError, HttpUnauthorizedError } from "../models/errors"

interface MainState {
    available: boolean
    authorized: boolean
}

const initialState: MainState = {
    available: true,
    authorized: true
}

export const apiMiddleware: Middleware = (api: MiddlewareAPI) => (
    next
) => (action) => {
    if (!isRejectedWithValue(action)) {
        next(action)
    }

    if (action.payload instanceof HttpUnauthorizedError) {
        api.dispatch(mainSlice.actions.setAuthorized(false))
    }

    if (action.payload instanceof HttpInternalServerError) {
        api.dispatch(mainSlice.actions.setAvailable(false))
    }

    next(action)
}

export const mainSlice = createSlice({
    name: "main",
    initialState,
    reducers: {
        setAvailable: (state, action: PayloadAction<boolean>) => {
            state.available = action.payload
        },
        setAuthorized: (state, action: PayloadAction<boolean>) => {
            state.authorized = action.payload
        }
    }
})