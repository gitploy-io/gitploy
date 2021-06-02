import { createSlice, createAsyncThunk } from '@reduxjs/toolkit'

import { Repo } from '../models'
import * as apis from '../apis'

const perPage = 30

interface HomeSate {
    loading: boolean 
    q: string
    repos: Repo[]
    page: number
}

const initialState: HomeSate = {
    loading: true,
    q: "",
    repos: [],
    page: 1,
}

export const listRepos = createAsyncThunk<Repo[], void, { state: {home: HomeSate} }>(
    'home/listRepos', 
    async (_, { getState }) => {
        const {q, page } = getState().home
        const repos = await apis.listRepos(q, page, perPage)
        return repos
    },
)

export const homeSlice = createSlice({
    name: 'home',
    initialState,
    reducers: {},
    extraReducers: builder => {
        builder
            .addCase(listRepos.fulfilled, (state, action) => {
                console.log("test")
                state.repos = action.payload
                state.loading = false
            })
    }
})
