import { configureStore } from '@reduxjs/toolkit'

import { homeSlice } from './home'

export const store =  configureStore({
  reducer: {
    home: homeSlice.reducer,
  },
})

export type RootState = ReturnType<typeof store.getState>

export type AppDispatch = typeof store.dispatch
