import { createSlice, PayloadAction } from '@reduxjs/toolkit'

interface RepoState {
    key: string
}

const initialState: RepoState = {
    key: "home",
}

export const repoSlice = createSlice({
    name: "repo",
    initialState,
    reducers: {
        setKey: (state, action: PayloadAction<string>) => {
            state.key = action.payload
        }
    }
})