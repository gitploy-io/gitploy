import { createSlice, createAsyncThunk } from "@reduxjs/toolkit"

import { User } from "../models"
import { getMe, checkSlack as _checkSlack } from "../apis"

interface SettingsState {
    user: User | null
    isSlackEnabled: boolean
}

const initialState: SettingsState = {
    user: null,
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
        builder
            .addCase(checkSlack.fulfilled, (state, action) => {
                state.isSlackEnabled = action.payload
            }) 
    }
})
