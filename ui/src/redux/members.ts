import { createSlice, PayloadAction, createAsyncThunk } from "@reduxjs/toolkit"

import { User } from "../models"
import { listUsers, deleteUser as _deleteUser } from "../apis"
import { message } from "antd"

export const perPage = 30

interface MembersState {
    users: User[]
    q: string
    page: number
}

const initialState: MembersState = {
    users: [],
    q: "",
    page: 1,
}

export const fetchUsers = createAsyncThunk<User[], void, { state: {members: MembersState} }>(
    "members/fetchUsers", 
    async (_, { getState, rejectWithValue } ) => {
        const { q, page } = getState().members

        try {
            const users = await listUsers(q, page)
            return users
        } catch(e) {
            return rejectWithValue(e)
        }
    },
)

export const deleteUser = createAsyncThunk<string, User, { state: {members: MembersState} }>(
    "members/deleteUser", 
    async (user, { rejectWithValue } ) => {
        try {
            await _deleteUser(user.id)
            return user.id
        } catch(e) {
            message.error(`Failed to delete.`, 3000)
            return rejectWithValue(e)
        }
    },
)

export const membersSlice = createSlice({
    name: "members",
    initialState,
    reducers: {
        setQuery: (state, action: PayloadAction<string>) => {
            state.q = action.payload
        },
        increasePage: (state) => {
            state.page = state.page + 1
        },
        decreasePage: (state) => {
            if (state.page <= 1) {
                return
            }

            state.page = state.page - 1
        },
    },
    extraReducers: (builder) => {
        builder
            .addCase(fetchUsers.fulfilled, (state, action) => {
                state.users = action.payload
            })

            .addCase(deleteUser.fulfilled, (state, action) => {
                state.users = state.users.filter((u) => {
                    return u.id !== action.payload
                })
            })
    }
})