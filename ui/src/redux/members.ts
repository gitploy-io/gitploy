import { createSlice, PayloadAction, createAsyncThunk } from "@reduxjs/toolkit"
import { message } from "antd"

import { User, HttpForbiddenError } from "../models"
import { listUsers, updateUser as _updateUser, deleteUser as _deleteUser } from "../apis"

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
            if (e instanceof HttpForbiddenError) {
                message.warn("Only admin can access.", 3)
            }

            return rejectWithValue(e)
        }
    },
)

export const updateUser = createAsyncThunk<User, {user: User, admin: boolean}, { state: {members: MembersState} }>(
    "members/updateUser", 
    async ({user, admin}, { rejectWithValue } ) => {
        try {
            const u = await _updateUser(user.id, {admin})
            return u
        } catch(e) {
            if (e instanceof HttpForbiddenError) {
                message.warn("Only admin can access.", 3)
            }

            return rejectWithValue(e)
        }
    },
)

export const deleteUser = createAsyncThunk<number, User, { state: {members: MembersState} }>(
    "members/deleteUser", 
    async (user, { rejectWithValue } ) => {
        try {
            await _deleteUser(user.id)
            return user.id
        } catch(e) {
            if (e instanceof HttpForbiddenError) {
                message.warn("Only admin can access.", 3)
            }

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

            .addCase(updateUser.fulfilled, (state, action) => {
                const idx = state.users.findIndex((u) => {
                    return u.id === action.payload.id
                })
                if (idx === -1) {
                    console.log("The updated user is not found.")
                    return
                }

                state.users[idx] = action.payload
            })

            .addCase(deleteUser.fulfilled, (state, action) => {
                state.users = state.users.filter((u) => {
                    return u.id !== action.payload
                })
            })
    }
})