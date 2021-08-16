import { createSlice, createAsyncThunk } from "@reduxjs/toolkit"

import { User, RateLimit } from "../models"
import { getMe, getRateLimit, checkSlack as _checkSlack } from "../apis"

interface SettingsState {
    user?: User
    rateLimit?: RateLimit
    isSlackEnabled: boolean
}

const initialState: SettingsState = {
    isSlackEnabled: false
}

export const fetchMe = createAsyncThunk<User, void, { state: { settings: SettingsState } }>(
    "settings/fetchMe",
    async (_, { rejectWithValue }) => {
        try {
            const user = await getMe()
            return user
        } catch (e) {
            return rejectWithValue(e)
        }
    }
)

export const fetchRateLimit = createAsyncThunk<RateLimit, void, { state: { settings: SettingsState } }>(
    "settings/fetchRateLimit",
    async (_, { rejectWithValue }) => {
        try {
            const rateLimit = await getRateLimit()
            return rateLimit
        } catch (e) {
            return rejectWithValue(e)
        }
    }
)

export const checkSlack = createAsyncThunk<boolean, void, { state: { settings: SettingsState } }>(
    "settings/checkSlack",
    async (_, { rejectWithValue }) => {
        try {
            return await _checkSlack()
        } catch (e) {
            return rejectWithValue(e)
        }
    }
)

export const settingsSlice = createSlice({
    name: "settings",
    initialState,
    reducers: {},
    extraReducers: builder => {
        builder
            .addCase(fetchMe.fulfilled, (state, action) => {
                state.user = action.payload
            }) 

            .addCase(fetchRateLimit.fulfilled, (state, action) => {
                state.rateLimit = action.payload
            })

            .addCase(checkSlack.fulfilled, (state, action) => {
                state.isSlackEnabled = action.payload
            }) 
    }
})
