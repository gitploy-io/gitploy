import { configureStore } from '@reduxjs/toolkit'

import { mainSlice, apiMiddleware } from "./main"
import { homeSlice } from "./home"
import { repoSlice } from "./repo"
import { repoHomeSlice } from "./repoHome"
import { repoDeploySlice } from "./repoDeploy"
import { repoLockSlice } from "./repoLock"
import { repoRollbackSlice } from "./repoRollback"
import { repoSettingsSlice } from "./repoSettings"
import { settingsSlice } from "./settings"
import { deploymentSlice } from "./deployment"
import { membersSlice } from "./members"

export const store =  configureStore({
  reducer: {
    main: mainSlice.reducer,
    home: homeSlice.reducer,
    repo: repoSlice.reducer,
    repoHome: repoHomeSlice.reducer,
    repoLock: repoLockSlice.reducer,
    repoDeploy: repoDeploySlice.reducer,
    repoRollback: repoRollbackSlice.reducer,
    repoSettings: repoSettingsSlice.reducer,
    settings: settingsSlice.reducer,
    deployment: deploymentSlice.reducer,
    members: membersSlice.reducer,
  },
  middleware: (getDefaultMiddleware) => getDefaultMiddleware({
    serializableCheck: false
  })
    .concat(apiMiddleware),
  devTools: true,
})

export type RootState = ReturnType<typeof store.getState>

export type AppDispatch = typeof store.dispatch
