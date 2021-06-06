import { configureStore, getDefaultMiddleware } from '@reduxjs/toolkit'

import { homeSlice } from './home'
import { repoSlice } from './repo'
import { repoHomeSlice } from './repoHome'
import { repoDeploySlice } from './repoDeploy'

export const store =  configureStore({
  reducer: {
    home: homeSlice.reducer,
    repo: repoSlice.reducer,
    repoHome: repoHomeSlice.reducer,
    repoDeploy: repoDeploySlice.reducer,
  },
  middleware: getDefaultMiddleware({
    serializableCheck: false
  }),
  devTools: true,
})

export type RootState = ReturnType<typeof store.getState>

export type AppDispatch = typeof store.dispatch