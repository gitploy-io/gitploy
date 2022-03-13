import { createSlice, createAsyncThunk, PayloadAction } from "@reduxjs/toolkit"

import { 
    searchDeployments as _searchDeployments, 
} from "../apis"
import { Deployment, } from "../models"

export const perPage = 30

interface ActivitiesState {
    start?: Date
    end?: Date
    loading: boolean
    deployments: Deployment[]
    page: number
}

const initialState: ActivitiesState = {
    loading: false,
    deployments: [],
    page: 1,
}

export const searchDeployments = createAsyncThunk<Deployment[], void, { state: { activities: ActivitiesState } }>(
    "activities/searchDeployments",
    async (_, { getState, rejectWithValue }) => {
        const {start, end, page} = getState().activities
        try {
            return await _searchDeployments([], false, start, end, page, perPage)
        } catch (e) {
            return rejectWithValue(e)
        }
    }
)

export const activitiesSlice = createSlice({
    name: "activities",
    initialState,
    reducers: {
        setStart: (state, action: PayloadAction<Date>) => {
            state.start = action.payload
        },
        setEnd: (state, action: PayloadAction<Date>) => {
            state.end = action.payload
        },
        increasePage: (state) => {
            state.page = state.page + 1
        },
        decreasePage: (state) => {
            state.page = state.page - 1
        },
    },
    extraReducers: builder => {
        builder
            .addCase(searchDeployments.pending, (state) => {
                state.loading = true
                state.deployments = []
            })
            .addCase(searchDeployments.fulfilled, (state, action) => {
                state.loading = false
                state.deployments = action.payload
            })
            .addCase(searchDeployments.rejected, (state) => {
                state.loading = false
            })
    }
})
